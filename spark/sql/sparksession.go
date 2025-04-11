package sql

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/violet-eva-01/spark-connect/spark/sql/types"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/ipc"
	"github.com/violet-eva-01/spark-connect/spark/client/base"

	"github.com/violet-eva-01/spark-connect/spark/client/options"

	"github.com/google/uuid"
	proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"
	"github.com/violet-eva-01/spark-connect/spark/client"
	"github.com/violet-eva-01/spark-connect/spark/client/channel"
	"github.com/violet-eva-01/spark-connect/spark/sparkerrors"
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
	CreateDataFrameFromMap(ctx context.Context, v interface{}, isTag bool, isRename bool, data ...map[string]interface{}) (DataFrame, error)
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
	rr := array.NewTableReader(data, int64(data.NumRows()))
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
				rb.Field(i).(*array.Float32Builder).Append(float32(row[i].(float32)))
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

// convStructTags
// @Description: get tag name & tag elem type or get elem name & tag name
// @param data
// @param tagName
// @param isGetType  true , get tag name & tag elem type . false , get elem name & tag name.
// @param splitKey
// @return map[string]string
func convStructTags(data any, tagName string, isGetType bool, splitKey ...string) map[string]string {

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

// structToStructType
// @Description: isTag is false , get struct elem name assign to structField name. isTag is true  , get json tag name assign to structField name.
// @param v
// @param isTag
// @return *types.StructType
// @return error
func structToStructType(v interface{}, isRename bool) (*types.StructType, error) {
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
		sparkTags = convStructTags(v, "spark", false)
	}
	for i := 0; i < vf.NumField(); i++ {
		var filed types.StructField
		if isRename {
			filed.Name = sparkTags[tf.Field(i).Name]
			if filed.Name == "" {
				filed.Name = tf.Field(i).Name
			}
		} else {
			filed.Name = tf.Field(i).Name
		}
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
			tag := tf.Field(i).Tag.Get("type")
			if tag == "timestamp" {
				filed.DataType = types.TIMESTAMP
			} else {
				filed.DataType = types.DATE
			}
		default:
			panic(fmt.Errorf("unsupported data type: %s", vt))
		}
		filed.Metadata = nil
		filed.Nullable = true
		fields = append(fields, filed)
	}
	return &types.StructType{
		Fields: fields,
	}, nil
}

// sAToTSA
// @Description: any slice  -> any 2D slicing
// @param structType
// @param data
// @return [][]interface{}
func sAToTSA(structType *types.StructType, data ...any) [][]interface{} {
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

// anyToSliceAny
// @Description: any -> any slice
// @param data
// @return []interface{}
func anyToSliceAny(data any) []interface{} {
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

// CreateDataFrameFromStruct
// @Description:
// @param ctx
// @param data
// @param isRename == true ,Rename the dataframe based on the spark tag , Insufficient tags are supplemented by elem name
// @return sql.DataFrame
// @return error
func (s *sparkSessionImpl) CreateDataFrameFromStruct(ctx context.Context, data any, isRename bool) (DataFrame, error) {
	rows := anyToSliceAny(data)
	if len(rows) == 0 {
		return nil, fmt.Errorf("no data")
	}
	structType, err := structToStructType(rows[0], isRename)
	if err != nil {
		return nil, err
	}
	sliceAny := sAToTSA(structType, rows...)
	return s.createDataFrame(ctx, sliceAny, structType)
}

func convStructDoubleTags(data any, tagName1, tagName2 string, splitKey ...[2]string) map[string]string {
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
			var (
				tag1Name string
				tag2Name string
			)
			field := valueOf.Type().Field(i)
			tag := field.Tag
			tag1Value := tag.Get(tagName1)
			if tag1Value != "" {
				splitValue := strings.Split(tag1Value, ",")
				for _, Value := range splitValue {
					if strings.HasPrefix(Value, splitKey[0][0]) {
						tag1Name = strings.TrimPrefix(Value, splitKey[0][0])
					}
				}
			}
			tag2Value := tag.Get(tagName2)
			if tag2Value != "" {
				splitValue := strings.Split(tag1Value, ",")
				for _, Value := range splitValue {
					if strings.HasPrefix(Value, splitKey[0][0]) {
						tag2Name = strings.TrimPrefix(Value, splitKey[0][0])
					}
				}
			}
			output[tag1Name] = tag2Name
		}
	} else {
		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Type().Field(i)
			tag := field.Tag
			tag1Value := tag.Get(tagName1)
			if tag1Value == "" {
				tag1Value = field.Name
			}
			tag2Value := tag.Get(tagName2)
			output[tag1Value] = tag2Value
		}
	}
	return output
}

func mapTurnOver(input map[string]string) map[string]string {
	var output = make(map[string]string, len(input))
	for key, value := range input {
		output[value] = key
	}
	return output
}

func mapToSliceAny(structType *types.StructType, v interface{}, isTag bool, isRename bool, data ...map[string]interface{}) ([][]interface{}, error) {
	var (
		mappingTags map[string]string
	)
	length := len(structType.Fields)
	if isTag && isRename {
		mappingTags = convStructDoubleTags(v, "spark", "json")
	} else if isRename {
		mappingTags = convStructTags(v, "spark", false)
		mappingTags = mapTurnOver(mappingTags)
	} else if isTag {
		mappingTags = convStructTags(v, "json", false)
	}
	var rows [][]interface{}
	for _, row := range data {
		var record []interface{}
		for i := 0; i < length; i++ {
			var rec interface{}
			if len(mappingTags) > 0 {
				rec = row[mappingTags[structType.Fields[i].Name]]
				if rec == nil && !isTag {
					rec = row[structType.Fields[i].Name]
				}
			} else {
				rec = row[structType.Fields[i].Name]
			}
			record = append(record, rec)
		}
		rows = append(rows, record)
	}
	return rows, nil
}

// CreateDataFrameFromMap
// @Description: []map[string]interface{} -> sql.DataFrame , map
// @param ctx
// @param v
// @param isTag true ,Assign values to dataframes based on JSON tags. false, Assign values to dataframes based on elem name. This field is because it is not possible to ignore the case match to the value, and the field is added.
// @param isRename  true ,Rename the dataframe based on the spark tag , Insufficient tags are supplemented by elem name
// @param data
// @return sql.DataFrame
// @return error
func (s *sparkSessionImpl) CreateDataFrameFromMap(ctx context.Context, v interface{}, isTag bool, isRename bool, data ...map[string]interface{}) (DataFrame, error) {
	structType, err := structToStructType(v, isRename)
	if err != nil {
		return nil, err
	}
	sliceAny, err := mapToSliceAny(structType, v, isTag, isRename, data...)
	if err != nil {
		return nil, err
	}
	return s.createDataFrame(ctx, sliceAny, structType)
}

func (s *sparkSessionImpl) createDataFrame(ctx context.Context, data [][]any, schema *types.StructType) (DataFrame, error) {
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
				rb.Field(i).(*array.Int8Builder).Append(row[i].(int8))
			case types.SHORT:
				rb.Field(i).(*array.Int16Builder).Append(row[i].(int16))
			case types.INTEGER:
				switch row[i].(type) {
				case int:
					rb.Field(i).(*array.Int32Builder).Append(int32(row[i].(int)))
				default:
					rb.Field(i).(*array.Int32Builder).Append(row[i].(int32))
				}
			case types.LONG:
				switch row[i].(type) {
				case int:
					rb.Field(i).(*array.Int64Builder).Append(int64(row[i].(int)))
				default:
					rb.Field(i).(*array.Int64Builder).Append(row[i].(int64))
				}
			case types.FLOAT:
				rb.Field(i).(*array.Float32Builder).Append(row[i].(float32))
			case types.DOUBLE:
				rb.Field(i).(*array.Float64Builder).Append(row[i].(float64))
			case types.STRING:
				rb.Field(i).(*array.StringBuilder).Append(row[i].(string))
			case types.DATE:
				rb.Field(i).(*array.Date32Builder).Append(arrow.Date32FromTime(row[i].(time.Time)))
			// case filed , err is execution error: [Internal] [UNSUPPORTED_ARROWTYPE] Unsupported arrow type Timestamp(MILLISECOND, UTC).
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
