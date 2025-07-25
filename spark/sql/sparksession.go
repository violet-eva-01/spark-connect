package sql

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/ipc"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/google/uuid"
	proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"
	"github.com/violet-eva-01/spark-connect/spark/client"
	"github.com/violet-eva-01/spark-connect/spark/client/base"
	"github.com/violet-eva-01/spark-connect/spark/client/channel"
	"github.com/violet-eva-01/spark-connect/spark/client/options"
	"github.com/violet-eva-01/spark-connect/spark/sparkerrors"
	"github.com/violet-eva-01/spark-connect/spark/sql/types"
	"google.golang.org/grpc/metadata"
)

type SparkSession interface {
	Read() DataFrameReader
	ReadStream() DataFrameReaderStream
	Sql(ctx context.Context, query string) (DataFrame, error)
	Stop() error
	Table(name string) (DataFrame, error)
	StreamManager() DataFrameStreamManager
	StreamQuery(id, runId string) DataFrameStreamQuery
	CreateDataFrameFromArrow(ctx context.Context, data arrow.Table) (DataFrame, error)
	CreateDataFrame(ctx context.Context, data [][]any, schema *types.StructType) (DataFrame, error)
	Config() client.RuntimeConfig
	CreateDataFrameFromStruct(ctx context.Context, data any, isRename bool) (DataFrame, error)
}

// NewSessionBuilder creates a new session builder for starting a new spark session
func NewSessionBuilder() *SparkSessionBuilder {
	return &SparkSessionBuilder{}
}

type SparkSessionBuilder struct {
	connectionString string
	channelBuilder   channel.Builder
}

// Remote sets the connection string for remote connection
func (s *SparkSessionBuilder) Remote(connectionString string) *SparkSessionBuilder {
	s.connectionString = connectionString
	return s
}

func (s *SparkSessionBuilder) WithChannelBuilder(cb channel.Builder) *SparkSessionBuilder {
	s.channelBuilder = cb
	return s
}

func (s *SparkSessionBuilder) Build(ctx context.Context) (SparkSession, error) {
	if s.channelBuilder == nil {
		cb, err := channel.NewBuilder(s.connectionString)
		if err != nil {
			return nil, sparkerrors.WithType(fmt.Errorf(
				"failed to connect to remote %s: %w", s.connectionString, err), sparkerrors.ConnectionError)
		}
		s.channelBuilder = cb
	}
	conn, err := s.channelBuilder.Build(ctx)
	if err != nil {
		return nil, sparkerrors.WithType(fmt.Errorf("failed to connect to remote %s: %w",
			s.connectionString, err), sparkerrors.ConnectionError)
	}

	// Add metadata to the request.
	meta := metadata.MD{}
	for k, v := range s.channelBuilder.Headers() {
		meta[k] = append(meta[k], v)
	}

	sessionId := uuid.NewString()

	// Update the options according to the configuration.
	opts := options.NewSparkClientOptions(options.DefaultSparkClientOptions.ReattachExecution)
	opts.UserAgent = s.channelBuilder.UserAgent()
	opts.UserId = s.channelBuilder.User()

	return &sparkSessionImpl{
		sessionId: sessionId,
		client:    client.NewSparkExecutor(conn, meta, sessionId, opts),
		conn:      conn,
	}, nil
}

type sparkSessionImpl struct {
	sessionId string
	client    base.SparkConnectClient
	conn      *grpc.ClientConn
}

func (s *sparkSessionImpl) Config() client.RuntimeConfig {
	return client.NewRuntimeConfig(&s.client)
}

func (s *sparkSessionImpl) Read() DataFrameReader {
	return NewDataframeReader(s)
}

func (s *sparkSessionImpl) ReadStream() DataFrameReaderStream { return newDataFrameReaderStream(s) }

func (s *sparkSessionImpl) StreamQuery(id, runId string) DataFrameStreamQuery {
	return NewDataFrameStreamQuery(s, id, runId)
}

func (s *sparkSessionImpl) StreamManager() DataFrameStreamManager {
	return NewDataFrameStreamManager(s)
}

// Sql executes a sql query and returns the result as a DataFrame
func (s *sparkSessionImpl) Sql(ctx context.Context, query string) (DataFrame, error) {
	// Due to the nature of Spark, we have to first submit the SQL query immediately as a command
	// to make sure that all side effects have been executed properly. If no side effects are present,
	// then simply prepare this as a SQL relation.

	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_SqlCommand{
					SqlCommand: &proto.SqlCommand{
						Sql: query,
					},
				},
			},
		},
	}
	// We need an execute command here.
	_, _, properties, err := s.client.ExecuteCommand(ctx, plan)
	if err != nil {
		return nil, sparkerrors.WithType(fmt.Errorf("failed to execute sql: %s: %w", query, err), sparkerrors.ExecutionError)
	}

	val, ok := properties["sql_command_result"]
	if !ok {
		plan := &proto.Relation{
			Common: &proto.RelationCommon{
				PlanId: newPlanId(),
			},
			RelType: &proto.Relation_Sql{
				Sql: &proto.SQL{
					Query: query,
				},
			},
		}
		return NewDataFrame(s, plan), nil
	} else {
		rel := val.(*proto.Relation)
		rel.Common = &proto.RelationCommon{
			PlanId: newPlanId(),
		}
		return NewDataFrame(s, rel), nil
	}
}

func (s *sparkSessionImpl) Stop() error {
	if s.conn != nil {
		err := s.conn.Close()
		return err
	} else {
		return nil
	}
}

func (s *sparkSessionImpl) GetSessionId() string {
	return s.sessionId
}

func (s *sparkSessionImpl) Table(name string) (DataFrame, error) {
	return s.Read().Table(name)
}

func (s *sparkSessionImpl) CreateDataFrameFromArrow(ctx context.Context, data arrow.Table) (DataFrame, error) {
	// Generate the schema.
	// schema := types.ArrowSchemaToProto(data.Schema())
	// schemaString := ""
	// TODO (PySpark does a lot of casting here to convert the schema that does not exist yet.

	// Convert the Arrow Table into a byte array of arrow IPC messages.
	buf := new(bytes.Buffer)
	w := ipc.NewWriter(buf, ipc.WithSchema(data.Schema()))
	defer w.Close()

	// Create a RecordReader from the table
	rr := array.NewTableReader(data, data.NumRows())
	defer rr.Release()

	// Read the records from the table and write them to the buffer
	for rr.Next() {
		record := rr.Record()
		if err := w.Write(record); err != nil {
			return nil, sparkerrors.WithType(fmt.Errorf("failed to write record: %w", err), sparkerrors.WriteError)
		}
	}

	// Create a local relation object
	plan := &proto.Relation{
		Common: &proto.RelationCommon{
			PlanId: newPlanId(),
		},
		RelType: &proto.Relation_LocalRelation{
			LocalRelation: &proto.LocalRelation{
				// Schema: &schemaString,
				Data: buf.Bytes(),
			},
		},
	}

	// Capture the column names from the schema:
	columnNames := make([]string, data.NumCols())
	for i, field := range data.Schema().Fields() {
		columnNames[i] = field.Name
	}

	dfPlan := &proto.Relation{
		Common: &proto.RelationCommon{
			PlanId: newPlanId(),
		},
		RelType: &proto.Relation_ToDf{
			ToDf: &proto.ToDF{
				Input:       plan,
				ColumnNames: columnNames,
			},
		},
	}
	return NewDataFrame(s, dfPlan), nil
}

func (s *sparkSessionImpl) CreateDataFrame(ctx context.Context, data [][]any, schema *types.StructType) (DataFrame, error) {
	pool := memory.NewGoAllocator()
	// Convert the data into an Arrow Table
	arrowSchema := arrow.NewSchema(schema.ToArrowType().Fields(), nil)
	rb := array.NewRecordBuilder(pool, arrowSchema)
	defer rb.Release()
	// Iterate over all fields and add the values:
	for _, row := range data {
		for i, field := range schema.Fields {
			if row[i] == nil {
				rb.Field(i).AppendNull()
				continue
			}
			switch field.DataType {
			case types.BOOLEAN:
				rb.Field(i).(*array.BooleanBuilder).Append(row[i].(bool))
			case types.BYTE:
				rb.Field(i).(*array.Int8Builder).Append(int8(row[i].(int)))
			case types.SHORT:
				rb.Field(i).(*array.Int16Builder).Append(int16(row[i].(int)))
			case types.INTEGER:
				rb.Field(i).(*array.Int32Builder).Append(int32(row[i].(int)))
			case types.LONG:
				rb.Field(i).(*array.Int64Builder).Append(int64(row[i].(int)))
			case types.FLOAT:
				rb.Field(i).(*array.Float32Builder).Append(row[i].(float32))
			case types.DOUBLE:
				rb.Field(i).(*array.Float64Builder).Append(row[i].(float64))
			case types.STRING:
				rb.Field(i).(*array.StringBuilder).Append(row[i].(string))
			case types.DATE:
				rb.Field(i).(*array.Date32Builder).Append(
					arrow.Date32FromTime(row[i].(time.Time)))
			case types.TIMESTAMP:
				ts, err := arrow.TimestampFromTime(row[i].(time.Time), arrow.Millisecond)
				if err != nil {
					return nil, err
				}
				rb.Field(i).(*array.TimestampBuilder).Append(ts)
			default:
				return nil, sparkerrors.WithType(fmt.Errorf(
					"unsupported data type: %s", field.DataType), sparkerrors.NotImplementedError)
			}
		}
	}
	rec := rb.NewRecord()
	defer rec.Release()
	tbl := array.NewTableFromRecords(arrowSchema, []arrow.Record{rec})
	defer tbl.Release()
	return s.CreateDataFrameFromArrow(ctx, tbl)
}

func SparkConnServer(ip string, port int, args map[string]string, ctxL ...context.Context) (SparkSession, error) {
	var (
		param    string
		remote   = fmt.Sprintf("sc://%s:%d", ip, port)
		builder  *channel.BaseBuilder
		sparkSQL SparkSession
		err      error
		ctx      context.Context
	)

	if len(ctxL) > 0 {
		ctx = ctxL[0]
	} else {
		ctx = context.Background()
	}

	if args != nil && len(args) > 0 {
		param = "/"
		for k, v := range args {
			if k == "user" {
				param += fmt.Sprintf(";%s=%s", "user_id", v)
			}
			param += fmt.Sprintf(";%s=%s", k, v)
		}
		remote += param
		builder, err = channel.NewBuilder(remote)
		if err != nil {
			return nil, err
		}
		sparkSQL, err = NewSessionBuilder().WithChannelBuilder(builder).Build(ctx)
	} else {
		sparkSQL, err = NewSessionBuilder().Remote(remote).Build(ctx)
	}
	if err != nil {
		return nil, err
	}
	_, err = sparkSQL.Sql(ctx, "select 1")
	if err != nil {
		return nil, err
	}
	return sparkSQL, nil
}

func NewSparkSQL(ip string, port int, args map[string]string, retryTime int, retryInterval time.Duration, ctxL ...context.Context) (conn SparkSession, err error) {
	for i := 0; i < retryTime; i++ {
		conn, err = SparkConnServer(ip, port, args, ctxL...)
		if err != nil {
			if i != retryTime-1 {
				time.Sleep(retryInterval * time.Second)
				continue
			} else {
				return nil, fmt.Errorf("connect spark connect server failed ,err is %s", err)
			}
		}
	}
	return
}

func getStructTagValues(data any, tagName string, isGetType bool, splitKey ...string) map[string]string {

	valueOf := reflect.ValueOf(data)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}
	if valueOf.Kind() != reflect.Struct {
		return nil
	}

	output := make(map[string]string, valueOf.NumField())
	if len(splitKey) > 0 {
		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Type().Field(i)
			tag := field.Tag
			tagValue := tag.Get(tagName)
			var fieldType string
			if isGetType {
				fieldType = field.Type.String()
			} else {
				fieldType = field.Name
			}
			if tagValue != "" {
				splitValue := strings.Split(tagValue, ",")
				for _, Value := range splitValue {
					if strings.HasPrefix(Value, splitKey[0]) {
						columnName := strings.TrimPrefix(Value, splitKey[0])
						if isGetType {
							output[columnName] = fieldType
						} else {
							output[fieldType] = columnName
						}
					}
				}
			}
		}
	} else {
		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Type().Field(i)
			tag := field.Tag
			tagValue := tag.Get(tagName)
			var fieldType string
			if isGetType {
				fieldType = field.Type.String()
			} else {
				fieldType = field.Name
			}
			if isGetType {
				// tag name : elem type
				output[tagValue] = fieldType
			} else {
				// elem name : tag name
				output[fieldType] = tagValue
			}
		}
	}

	return output
}

func getStructType(v interface{}, isRename bool) (*types.StructType, error) {
	var (
		fields    []types.StructField
		sparkTags map[string]string
	)
	vf := reflect.ValueOf(v)
	tf := reflect.TypeOf(v)
	if tf.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct type, got %T", v)
	}
	if isRename {
		sparkTags = getStructTagValues(v, "spark", false)
	}
	for i := 0; i < vf.NumField(); i++ {
		var filed types.StructField
		if isRename {
			filed.Name = sparkTags[tf.Field(i).Name]
		}
		if filed.Name == "" {
			filed.Name = tf.Field(i).Name
		}
		if sparkTypeTag := tf.Field(i).Tag.Get("sparkType"); sparkTypeTag != "" {
			tagValue := strings.Split(sparkTypeTag, ",")
			sparkType := tagValue[0]
			switch sparkType {
			case "bool":
				filed.DataType = types.BOOLEAN
			case "int":
				switch runtime.GOARCH {
				case "386", "arm":
					filed.DataType = types.INTEGER
				default:
					filed.DataType = types.LONG
				}
			case "int8":
				filed.DataType = types.BYTE
			case "int16":
				filed.DataType = types.SHORT
			case "int32":
				filed.DataType = types.INTEGER
			case "int64":
				filed.DataType = types.LONG
			case "float32", "float":
				filed.DataType = types.FLOAT
			case "float64", "double":
				filed.DataType = types.DOUBLE
			case "string":
				filed.DataType = types.STRING
			case "date":
				filed.DataType = types.DATE
				if len(tagValue) > 1 {
					filed.TimeFormat = tagValue[1]
				}
			case "timestamp":
				filed.DataType = types.TIMESTAMP
				if len(tagValue) > 1 {
					filed.TimeFormat = tagValue[1]
				}
			case "timestamp_ms":
				filed.DataType = types.TIMESTAMP_MTZ
				if len(tagValue) > 1 {
					filed.TimeFormat = tagValue[1]
				}
			case "timestamp_us":
				filed.DataType = types.TIMESTAMP_UTZ
				if len(tagValue) > 1 {
					filed.TimeFormat = tagValue[1]
				}
			case "timestamp_ns":
				filed.DataType = types.TIMESTAMP_NTZ
				if len(tagValue) > 1 {
					filed.TimeFormat = tagValue[1]
				}
			}
		} else {
			switch vt := vf.Field(i).Interface().(type) {
			case int:
				switch runtime.GOARCH {
				case "386", "arm":
					filed.DataType = types.INTEGER
				default:
					filed.DataType = types.LONG
				}
			case bool:
				filed.DataType = types.BOOLEAN
			case int8:
				filed.DataType = types.BYTE
			case int16:
				filed.DataType = types.SHORT
			case int32:
				filed.DataType = types.INTEGER
			case int64:
				filed.DataType = types.LONG
			case float32:
				filed.DataType = types.FLOAT
			case float64:
				filed.DataType = types.DOUBLE
			case string:
				filed.DataType = types.STRING
			case time.Time:
				filed.DataType = types.TIMESTAMP_UTZ
			default:
				panic(fmt.Errorf("unsupported data type: %s", vt))
			}
		}
		filed.Metadata = nil
		filed.Nullable = true
		fields = append(fields, filed)
	}
	return &types.StructType{
		Fields: fields,
	}, nil
}

func getData(structType *types.StructType, data ...any) [][]interface{} {
	length := len(structType.Fields)
	var rows [][]interface{}
	for _, row := range data {
		var record []interface{}
		vf := reflect.ValueOf(row)
		for i := 0; i < length; i++ {
			rec := vf.Field(i).Interface()
			record = append(record, rec)
		}
		rows = append(rows, record)
	}
	return rows
}

func dataConversion(data any) []interface{} {
	vf := reflect.ValueOf(data)
	if vf.Kind() != reflect.Slice && vf.Kind() == reflect.Struct {
		return []interface{}{data}
	}
	rows := make([]any, vf.Len())
	for i := 0; i < vf.Len(); i++ {
		rows[i] = vf.Index(i).Interface()
	}
	return rows
}

func (s *sparkSessionImpl) CreateDataFrameFromStruct(ctx context.Context, data any, isRename bool) (DataFrame, error) {
	rows := dataConversion(data)
	if len(rows) == 0 {
		return nil, fmt.Errorf("no data")
	}
	structType, err := getStructType(rows[0], isRename)
	if err != nil {
		return nil, err
	}
	sliceAny := getData(structType, rows...)
	return s.createDataFrame(ctx, sliceAny, structType)
}

var TZ = time.Local

func (s *sparkSessionImpl) createDataFrame(ctx context.Context, data [][]any, schema *types.StructType) (DataFrame, error) {
	pool := memory.NewGoAllocator()
	arrowSchema := arrow.NewSchema(schema.ToArrowType().Fields(), nil)
	rb := array.NewRecordBuilder(pool, arrowSchema)
	defer rb.Release()
	for _, row := range data {
		for i, field := range schema.Fields {
			if row[i] == nil {
				rb.Field(i).AppendNull()
				continue
			} else if row[i] == "" {
				switch field.DataType {
				case types.TIMESTAMP, types.TIMESTAMP_MTZ, types.TIMESTAMP_UTZ, types.TIMESTAMP_NTZ, types.DATE:
					rb.Field(i).AppendNull()
					continue
				default:
					rb.Field(i).AppendEmptyValue()
					continue
				}
			}
			switch field.DataType {
			case types.BOOLEAN:
				rowData, ok := row[i].(bool)
				if !ok {
					parseBool, err := strconv.ParseBool(fmt.Sprintf("%v", row[i]))
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to bool failed, err is %s", row[i], err), sparkerrors.CreateDataFrameError)
					}
					rowData = parseBool
				}
				rb.Field(i).(*array.BooleanBuilder).Append(rowData)
			case types.BYTE:
				rowData, ok := row[i].(int8)
				if !ok {
					parseByte, err := strconv.ParseInt(fmt.Sprintf("%v", row[i]), 10, 8)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to int8 failed, err is %s", row[i], err), sparkerrors.CreateDataFrameError)
					}
					rowData = int8(parseByte)
				}
				rb.Field(i).(*array.Int8Builder).Append(rowData)
			case types.SHORT:
				rowData, ok := row[i].(int16)
				if !ok {
					parseShort, err := strconv.ParseInt(fmt.Sprintf("%v", row[i]), 10, 16)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to int16 failed, err is %s", row[i], err), sparkerrors.CreateDataFrameError)
					}
					rowData = int16(parseShort)
				}
				rb.Field(i).(*array.Int16Builder).Append(rowData)
			case types.INTEGER:
				rowData, ok := row[i].(int32)
				if !ok {
					parseInt, err := strconv.ParseInt(fmt.Sprintf("%v", row[i]), 10, 32)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to int32 failed, err is %s", row[i], err), sparkerrors.CreateDataFrameError)
					}
					rowData = int32(parseInt)
				}
				rb.Field(i).(*array.Int32Builder).Append(rowData)
			case types.LONG:
				rowData, ok := row[i].(int64)
				if !ok {
					parseLong, err := strconv.ParseInt(fmt.Sprintf("%v", row[i]), 10, 64)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to int64 failed, err is %s", row[i], err), sparkerrors.CreateDataFrameError)
					}
					rowData = parseLong
				}
				rb.Field(i).(*array.Int64Builder).Append(rowData)
			case types.FLOAT:
				rowData, ok := row[i].(float32)
				if !ok {
					parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%v", row[i]), 32)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to float32 failed, err is %s", row[i], err), sparkerrors.CreateDataFrameError)
					}
					rowData = float32(parseFloat)
				}
				rb.Field(i).(*array.Float32Builder).Append(rowData)
			case types.DOUBLE:
				rowData, ok := row[i].(float64)
				if !ok {
					parseDouble, err := strconv.ParseFloat(fmt.Sprintf("%v", row[i]), 64)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to float64/double failed, err is %s", row[i], err), sparkerrors.CreateDataFrameError)
					}
					rowData = parseDouble
				}
				rb.Field(i).(*array.Float64Builder).Append(rowData)
			case types.STRING:
				rowData, ok := row[i].(string)
				if !ok {
					rowData = fmt.Sprintf("%v", row[i])
				}
				rb.Field(i).(*array.StringBuilder).Append(rowData)
			case types.DATE:
				rowData, ok := row[i].(time.Time)
				if !ok {
					parse, err := time.ParseInLocation(field.TimeFormat, fmt.Sprintf("%s", row[i]), TZ)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to date failed, format is [%s], err is %s", row[i], field.TimeFormat, err), sparkerrors.CreateDataFrameError)
					}
					rowData = parse
				} else {
					if fmt.Sprintf("%s", time.Time{}) == fmt.Sprintf("%s", row[i]) {
						rb.Field(i).(*array.Date32Builder).AppendNull()
						continue
					}
				}
				rb.Field(i).(*array.Date32Builder).Append(arrow.Date32FromTime(rowData))
			case types.TIMESTAMP, types.TIMESTAMP_MTZ, types.TIMESTAMP_UTZ, types.TIMESTAMP_NTZ:
				var (
					tu      arrow.TimeUnit
					rowData arrow.Timestamp
					err     error
				)
				switch field.DataType {
				case types.TIMESTAMP:
					tu = arrow.Second
				case types.TIMESTAMP_MTZ:
					tu = arrow.Millisecond
				case types.TIMESTAMP_UTZ:
					tu = arrow.Microsecond
				case types.TIMESTAMP_NTZ:
					tu = arrow.Nanosecond
				}
				ts, ok := row[i].(time.Time)
				if !ok {
					ts, err = time.ParseInLocation(field.TimeFormat, fmt.Sprintf("%s", row[i]), TZ)
					if err != nil {
						return nil, sparkerrors.WithType(fmt.Errorf(
							"cast value [%s] to time.Time failed, format is [%s], err is %s", row[i], field.TimeFormat, err), sparkerrors.CreateDataFrameError)
					}
				} else {
					if fmt.Sprintf("%s", time.Time{}) == fmt.Sprintf("%s", row[i]) {
						rb.Field(i).(*array.TimestampBuilder).AppendNull()
						continue
					}
				}
				rowData, err = arrow.TimestampFromTime(ts, tu)
				if err != nil {
					return nil, sparkerrors.WithType(fmt.Errorf(
						"cast value [%s] to arrow.Timestamp(%s, UTC) failed,err is %s", row[i], tu.String(), err), sparkerrors.CreateDataFrameError)
				}
				rb.Field(i).(*array.TimestampBuilder).Append(rowData)
			default:
				return nil, sparkerrors.WithType(fmt.Errorf(
					"unsupported data type: %s", field.DataType), sparkerrors.NotImplementedError)
			}
		}
	}
	rec := rb.NewRecord()
	defer rec.Release()
	tbl := array.NewTableFromRecords(arrowSchema, []arrow.Record{rec})
	defer tbl.Release()
	return s.CreateDataFrameFromArrow(ctx, tbl)
}
