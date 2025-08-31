package sql

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"reflect"
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
	CreateDataFrameFromStruct(ctx context.Context, data any) (DataFrame, error)
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

func NewSparkSQL(ip string, port int, args map[string]string, ctxL ...context.Context) (SparkSession, error) {
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

func getStructType(v interface{}) (*types.StructType, error) {
	var (
		fields []types.StructField
	)
	vf := reflect.ValueOf(v)
	tf := reflect.TypeOf(v)
	sparkTags := types.ExtractSparkTags(v)
	if tf.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct type, got %T", v)
	}

	for i := 0; i < vf.NumField(); i++ {
		var filed types.StructField

		if sparkTags[tf.Field(i).Name].Column != "" {
			filed.Name = sparkTags[tf.Field(i).Name].Column
		} else {
			filed.Name = tf.Field(i).Name
		}

		if tag := sparkTags[tf.Field(i).Name]; tag.Type != "" {
			switch tag.Type {
			case "bool":
				filed.DataType = types.BOOLEAN
			case "[]bool":
				filed.DataType = types.BOOLEAN_SLICE
			case "int8":
				filed.DataType = types.BYTE
			case "[]int8":
				filed.DataType = types.BYTE_SLICE
			case "int16":
				filed.DataType = types.SHORT
			case "[]int16":
				filed.DataType = types.SHORT_SLICE
			case "int", "int32":
				filed.DataType = types.INTEGER
			case "[]int", "[]int32":
				filed.DataType = types.INTEGER_SLICE
			case "int64":
				filed.DataType = types.LONG
			case "[]int64":
				filed.DataType = types.LONG_SLICE
			case "float32", "float":
				filed.DataType = types.FLOAT
			case "[]float32", "[]float":
				filed.DataType = types.FLOAT_SLICE
			case "float64", "double":
				filed.DataType = types.DOUBLE
			case "[]float64", "[]double":
				filed.DataType = types.DOUBLE_SLICE
			case "string":
				filed.DataType = types.STRING
			case "[]string":
				filed.DataType = types.STRING_SLICE
			case "date":
				filed.DataType = types.DATE
			case "[]date":
				filed.DataType = types.DATE_SLICE
			case "timestamp":
				filed.DataType = types.TIMESTAMP
			case "[]timestamp":
				filed.DataType = types.TIMESTAMP_SLICE
			case "timestamp_ms":
				filed.DataType = types.TIMESTAMP_MTZ
			case "[]timestamp_ms":
				filed.DataType = types.TIMESTAMP_MTZ_SLICE
			case "timestamp_us":
				filed.DataType = types.TIMESTAMP_UTZ
			case "[]timestamp_us":
				filed.DataType = types.TIMESTAMP_UTZ_SLICE
			case "timestamp_ns":
				filed.DataType = types.TIMESTAMP_NTZ
			case "[]timestamp_ns":
				filed.DataType = types.TIMESTAMP_NTZ_SLICE
			default:
				panic(fmt.Errorf("unsupported data type: %s", tag.Type))
			}
		} else {
			switch vt := vf.Field(i).Interface().(type) {
			case bool:
				filed.DataType = types.BOOLEAN
			case []bool:
				filed.DataType = types.BOOLEAN_SLICE
			case int8:
				filed.DataType = types.BYTE
			case []int8:
				filed.DataType = types.BYTE_SLICE
			case int16:
				filed.DataType = types.SHORT
			case []int16:
				filed.DataType = types.SHORT_SLICE
			case int, int32:
				filed.DataType = types.INTEGER
			case []int, []int32:
				filed.DataType = types.INTEGER_SLICE
			case int64:
				filed.DataType = types.LONG
			case []int64:
				filed.DataType = types.LONG_SLICE
			case float32:
				filed.DataType = types.FLOAT
			case []float32:
				filed.DataType = types.FLOAT_SLICE
			case float64:
				filed.DataType = types.DOUBLE
			case []float64:
				filed.DataType = types.DOUBLE_SLICE
			case string:
				filed.DataType = types.STRING
			case []string:
				filed.DataType = types.STRING_SLICE
			case time.Time:
				filed.DataType = types.TIMESTAMP_UTZ
			case []time.Time:
				filed.DataType = types.TIMESTAMP_UTZ_SLICE
			default:
				panic(fmt.Errorf("unsupported data type: %s", vt))
			}
		}

		if sparkTags[tf.Field(i).Name].Format != "" {
			filed.Format = sparkTags[tf.Field(i).Name].Format
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

func (s *sparkSessionImpl) CreateDataFrameFromStruct(ctx context.Context, data any) (DataFrame, error) {
	rows := dataConversion(data)
	if len(rows) == 0 {
		return nil, fmt.Errorf("no data")
	}
	structType, err := getStructType(rows[0])
	if err != nil {
		return nil, err
	}
	sliceAny := getData(structType, rows...)
	return s.createDataFrame(ctx, sliceAny, structType)
}

func (s *sparkSessionImpl) createDataFrame(ctx context.Context, data [][]any, schema *types.StructType) (DataFrame, error) {
	pool := memory.NewGoAllocator()
	arrowSchema := arrow.NewSchema(schema.ToArrowType().Fields(), nil)
	rb := array.NewRecordBuilder(pool, arrowSchema)
	defer rb.Release()
	for _, row := range data {
		for i, field := range schema.Fields {
			switch field.DataType {
			case types.MAP:
			default:
				if err := fieldBuilder(i, field, rb, row); err != nil {
					return nil, err
				}
			}
		}
	}
	rec := rb.NewRecord()
	defer rec.Release()
	tbl := array.NewTableFromRecords(arrowSchema, []arrow.Record{rec})
	defer tbl.Release()
	return s.CreateDataFrameFromArrow(ctx, tbl)
}

var tz = time.UTC

func InitTZ(timeZone *time.Location) {
	tz = timeZone
}

func fieldBuilder(index int, field types.StructField, rb *array.RecordBuilder, row []any) error {
	if row[index] == nil {
		rb.Field(index).AppendNull()
		return nil
	} else if row[index] == "" {
		switch field.DataType {
		case types.TIMESTAMP, types.TIMESTAMP_MTZ, types.TIMESTAMP_UTZ, types.TIMESTAMP_NTZ, types.DATE:
			rb.Field(index).AppendNull()
			return nil
		default:
			rb.Field(index).AppendEmptyValue()
			return nil
		}
	} else if row[index] == "[]" {
		rb.Field(index).AppendEmptyValue()
		return nil
	}
	switch field.DataType {
	case types.BOOLEAN:
		rowData, ok := row[index].(bool)
		if !ok {
			parseBool, err := strconv.ParseBool(fmt.Sprintf("%v", row[index]))
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to bool failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = parseBool
		}
		rb.Field(index).(*array.BooleanBuilder).Append(rowData)
	case types.BOOLEAN_SLICE:
		rowData, ok := row[index].([]bool)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []bool failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, rd := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.BooleanBuilder).Append(rd)
		}
	case types.BYTE:
		rowData, ok := row[index].(int8)
		if !ok {
			parseByte, err := strconv.ParseInt(fmt.Sprintf("%v", row[index]), 10, 8)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to int8 failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = int8(parseByte)
		}
		rb.Field(index).(*array.Int8Builder).Append(rowData)
	case types.BYTE_SLICE:
		rowData, ok := row[index].([]int8)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []int8 failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, rd := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Int8Builder).Append(rd)
		}
	case types.SHORT:
		rowData, ok := row[index].(int16)
		if !ok {
			parseShort, err := strconv.ParseInt(fmt.Sprintf("%v", row[index]), 10, 16)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to int16 failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = int16(parseShort)
		}
		rb.Field(index).(*array.Int16Builder).Append(rowData)
	case types.SHORT_SLICE:
		rowData, ok := row[index].([]int16)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []int16 failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, rd := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Int16Builder).Append(rd)
		}
	case types.INTEGER:
		rowData, ok := row[index].(int32)
		if !ok {
			parseInt, err := strconv.ParseInt(fmt.Sprintf("%v", row[index]), 10, 32)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to int32 failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = int32(parseInt)
		}
		rb.Field(index).(*array.Int32Builder).Append(rowData)
	case types.INTEGER_SLICE:
		switch row[index].(type) {
		case []int:
			rowData, ok := row[index].([]int)
			if !ok {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to []int32 failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
			}
			rb.Field(index).(*array.ListBuilder).Append(true)
			for _, rd := range rowData {
				rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Int32Builder).Append(int32(rd))
			}
		case []int32:
			rowData, ok := row[index].([]int32)
			if !ok {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to []int32 failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
			}
			rb.Field(index).(*array.ListBuilder).Append(true)
			for _, b := range rowData {
				rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Int32Builder).Append(b)
			}
		}
	case types.LONG:
		rowData, ok := row[index].(int64)
		if !ok {
			parseLong, err := strconv.ParseInt(fmt.Sprintf("%v", row[index]), 10, 64)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to int64 failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = parseLong
		}
		rb.Field(index).(*array.Int64Builder).Append(rowData)
	case types.LONG_SLICE:
		rowData, ok := row[index].([]int64)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []int64 failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, b := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Int64Builder).Append(b)
		}
	case types.FLOAT:
		rowData, ok := row[index].(float32)
		if !ok {
			parseFloat, err := strconv.ParseFloat(fmt.Sprintf("%v", row[index]), 32)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to float32 failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = float32(parseFloat)
		}
		rb.Field(index).(*array.Float32Builder).Append(rowData)
	case types.FLOAT_SLICE:
		rowData, ok := row[index].([]float32)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []float32 failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, b := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Float32Builder).Append(b)
		}
	case types.DOUBLE:
		rowData, ok := row[index].(float64)
		if !ok {
			parseDouble, err := strconv.ParseFloat(fmt.Sprintf("%v", row[index]), 64)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to float64 failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = parseDouble
		}
		rb.Field(index).(*array.Float64Builder).Append(rowData)
	case types.DOUBLE_SLICE:
		rowData, ok := row[index].([]float64)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []float64 failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, b := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Float64Builder).Append(b)
		}
	case types.STRING:
		rowData, ok := row[index].(string)
		if !ok {
			rowData = fmt.Sprintf("%v", row[index])
		}
		rb.Field(index).(*array.StringBuilder).Append(rowData)
	case types.STRING_SLICE:
		rowData, ok := row[index].([]string)
		if !ok {
			if field.Format == "" {
				rowData = strings.Fields(fmt.Sprintf("%v", row[index]))
			} else {
				rowData = strings.Split(fmt.Sprintf("%v", row[index]), field.Format)
			}
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, str := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.StringBuilder).Append(str)
		}
	case types.DATE:
		rowData, ok := row[index].(time.Time)
		if !ok {
			format := "2006-01-02 15:04:05"
			if field.Format != "" {
				format = field.Format
			}
			// 2006-01-02 格式数据跨时区转换可能会导致时间转换异常，推荐使用 2006-01-02 15:04:05 格式转换数据。
			parse, err := time.ParseInLocation(format, fmt.Sprintf("%s", row[index]), tz)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to time.Time failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
			rowData = parse
		} else {
			if fmt.Sprintf("%s", time.Time{}) == fmt.Sprintf("%s", row[index]) {
				rb.Field(index).(*array.Date32Builder).AppendNull()
				return nil
			}
		}
		rb.Field(index).(*array.Date32Builder).Append(arrow.Date32FromTime(rowData))
	case types.DATE_SLICE:
		rowData, ok := row[index].([]time.Time)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []time.Time failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, b := range rowData {
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.Date32Builder).Append(arrow.Date32FromTime(b))
		}
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
		ts, ok := row[index].(time.Time)
		if !ok {
			format := "2006-01-02 15:04:05"
			if field.Format != "" {
				format = field.Format
			}
			ts, err = time.ParseInLocation(format, fmt.Sprintf("%s", row[index]), tz)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to time.Time failed, err is %s", field.Name, row[index], err), sparkerrors.CreateDataFrameError)
			}
		} else {
			if fmt.Sprintf("%s", time.Time{}) == fmt.Sprintf("%s", row[index]) {
				rb.Field(index).(*array.TimestampBuilder).AppendNull()
				return nil
			}
		}
		rowData, err = arrow.TimestampFromTime(ts, tu)
		if err != nil {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to arrow.Timestamp(%s, UTC) failed, err is %s", field.Name, row[index], tu.String(), err), sparkerrors.CreateDataFrameError)
		}
		rb.Field(index).(*array.TimestampBuilder).Append(rowData)
	case types.TIMESTAMP_SLICE, types.TIMESTAMP_MTZ_SLICE, types.TIMESTAMP_UTZ_SLICE, types.TIMESTAMP_NTZ_SLICE:
		var (
			tu      arrow.TimeUnit
			rowData arrow.Timestamp
			err     error
		)
		switch field.DataType {
		case types.TIMESTAMP_SLICE:
			tu = arrow.Second
		case types.TIMESTAMP_MTZ_SLICE:
			tu = arrow.Millisecond
		case types.TIMESTAMP_UTZ_SLICE:
			tu = arrow.Microsecond
		case types.TIMESTAMP_NTZ_SLICE:
			tu = arrow.Nanosecond
		}
		ts, ok := row[index].([]time.Time)
		if !ok {
			return sparkerrors.WithType(fmt.Errorf(
				"field [%s] cast value [%s] to []time.Time failed", field.Name, row[index]), sparkerrors.CreateDataFrameError)
		} else {
			if fmt.Sprintf("%s", []time.Time{}) == fmt.Sprintf("%s", row[index]) {
				rb.Field(index).(*array.ListBuilder).AppendEmptyValue()
				return nil
			}
		}
		rb.Field(index).(*array.ListBuilder).Append(true)
		for _, b := range ts {
			rowData, err = arrow.TimestampFromTime(b, tu)
			if err != nil {
				return sparkerrors.WithType(fmt.Errorf(
					"field [%s] cast value [%s] to []arrow.Timestamp(%s, UTC) failed,err is %s", field.Name, row[index], tu.String(), err), sparkerrors.CreateDataFrameError)
			}
			rb.Field(index).(*array.ListBuilder).ValueBuilder().(*array.TimestampBuilder).Append(rowData)
		}

	default:
		return sparkerrors.WithType(fmt.Errorf(
			"unsupported data type: %s", field.DataType), sparkerrors.NotImplementedError)
	}
	return nil
}
