package types

import (
	"fmt"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
)

type DataType interface {
	TypeName() string
	IsNumeric() bool
	ToArrowType() arrow.DataType
}

type BooleanType struct{}

func (t BooleanType) TypeName() string {
	return getDataTypeName(t)
}

func (t BooleanType) IsNumeric() bool {
	return false
}

func (t BooleanType) ToArrowType() arrow.DataType {
	return arrow.FixedWidthTypes.Boolean
}

type BooleanSliceType struct{}

func (t BooleanSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t BooleanSliceType) IsNumeric() bool {
	return false
}

func (t BooleanSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.FixedWidthTypes.Boolean)
}

type ByteType struct{}

func (t ByteType) IsNumeric() bool {
	return true
}

func (t ByteType) ToArrowType() arrow.DataType {
	return arrow.PrimitiveTypes.Int8
}

func (t ByteType) TypeName() string {
	return getDataTypeName(t)
}

type ByteSliceType struct{}

func (t ByteSliceType) IsNumeric() bool {
	return true
}

func (t ByteSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.PrimitiveTypes.Int8)
}

func (t ByteSliceType) TypeName() string {
	return getDataTypeName(t)
}

type ShortType struct{}

func (t ShortType) TypeName() string {
	return getDataTypeName(t)
}

func (t ShortType) IsNumeric() bool {
	return true
}

func (t ShortType) ToArrowType() arrow.DataType {
	return arrow.PrimitiveTypes.Int16
}

type ShortSliceType struct{}

func (t ShortSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t ShortSliceType) IsNumeric() bool {
	return true
}

func (t ShortSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.PrimitiveTypes.Int16)
}

type IntegerType struct{}

func (t IntegerType) TypeName() string {
	return getDataTypeName(t)
}

func (t IntegerType) IsNumeric() bool {
	return true
}

func (t IntegerType) ToArrowType() arrow.DataType {
	return arrow.PrimitiveTypes.Int32
}

type IntegerSliceType struct{}

func (t IntegerSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t IntegerSliceType) IsNumeric() bool {
	return true
}

func (t IntegerSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.PrimitiveTypes.Int32)
}

type LongType struct{}

func (t LongType) TypeName() string {
	return getDataTypeName(t)
}

func (t LongType) IsNumeric() bool {
	return true
}

func (t LongType) ToArrowType() arrow.DataType {
	return arrow.PrimitiveTypes.Int64
}

type LongSliceType struct{}

func (t LongSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t LongSliceType) IsNumeric() bool {
	return true
}

func (t LongSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.PrimitiveTypes.Int64)
}

type FloatType struct{}

func (t FloatType) TypeName() string {
	return getDataTypeName(t)
}

func (t FloatType) IsNumeric() bool {
	return true
}

func (t FloatType) ToArrowType() arrow.DataType {
	return arrow.PrimitiveTypes.Float32
}

type FloatSliceType struct{}

func (t FloatSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t FloatSliceType) IsNumeric() bool {
	return true
}

func (t FloatSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.PrimitiveTypes.Float32)
}

type DoubleType struct{}

func (t DoubleType) TypeName() string {
	return getDataTypeName(t)
}

func (t DoubleType) IsNumeric() bool {
	return true
}

func (t DoubleType) ToArrowType() arrow.DataType {
	return arrow.PrimitiveTypes.Float64
}

type DoubleSliceType struct{}

func (t DoubleSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t DoubleSliceType) IsNumeric() bool {
	return true
}

func (t DoubleSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.PrimitiveTypes.Float64)
}

type DecimalType struct {
	Precision int32
	Scale     int32
}

func (t DecimalType) TypeName() string {
	return getDataTypeName(t)
}

func (t DecimalType) IsNumeric() bool {
	return true
}

func (t DecimalType) ToArrowType() arrow.DataType {
	return &arrow.Decimal128Type{
		Precision: t.Precision,
		Scale:     t.Scale,
	}
}

type DecimalSliceType struct {
	Precision int32
	Scale     int32
}

func (t DecimalSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t DecimalSliceType) IsNumeric() bool {
	return true
}

func (t DecimalSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(&arrow.Decimal128Type{
		Precision: t.Precision,
		Scale:     t.Scale,
	})
}

type StringType struct{}

func (t StringType) TypeName() string {
	return getDataTypeName(t)
}

func (t StringType) IsNumeric() bool {
	return false
}

func (t StringType) ToArrowType() arrow.DataType {
	return arrow.BinaryTypes.String
}

type StringSliceType struct{}

func (t StringSliceType) TypeName() string { return getDataTypeName(t) }

func (t StringSliceType) IsNumeric() bool { return false }

func (t StringSliceType) ToArrowType() arrow.DataType { return arrow.ListOf(arrow.BinaryTypes.String) }

type BinaryType struct{}

func (t BinaryType) TypeName() string {
	return getDataTypeName(t)
}

func (t BinaryType) IsNumeric() bool {
	return false
}

func (t BinaryType) ToArrowType() arrow.DataType {
	return arrow.BinaryTypes.Binary
}

type BinarySliceType struct{}

func (t BinarySliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t BinarySliceType) IsNumeric() bool {
	return false
}

func (t BinarySliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.BinaryTypes.Binary)
}

type TimestampType struct{}

func (t TimestampType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampType) IsNumeric() bool {
	return false
}

func (t TimestampType) ToArrowType() arrow.DataType {
	return arrow.FixedWidthTypes.Timestamp_s
}

type TimestampSliceType struct{}

func (t TimestampSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampSliceType) IsNumeric() bool {
	return false
}

func (t TimestampSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.FixedWidthTypes.Timestamp_s)
}

type TimestampMtzType struct{}

func (t TimestampMtzType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampMtzType) IsNumeric() bool {
	return false
}

func (t TimestampMtzType) ToArrowType() arrow.DataType {
	return arrow.FixedWidthTypes.Timestamp_ms
}

type TimestampMtzSliceType struct{}

func (t TimestampMtzSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampMtzSliceType) IsNumeric() bool {
	return false
}

func (t TimestampMtzSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.FixedWidthTypes.Timestamp_ms)
}

type TimestampUtzType struct{}

func (t TimestampUtzType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampUtzType) IsNumeric() bool {
	return false
}

func (t TimestampUtzType) ToArrowType() arrow.DataType {
	return arrow.FixedWidthTypes.Timestamp_us
}

type TimestampUtzSliceType struct{}

func (t TimestampUtzSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampUtzSliceType) IsNumeric() bool {
	return false
}

func (t TimestampUtzSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.FixedWidthTypes.Timestamp_us)
}

type TimestampNtzType struct{}

func (t TimestampNtzType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampNtzType) IsNumeric() bool {
	return false
}

func (t TimestampNtzType) ToArrowType() arrow.DataType {
	return arrow.FixedWidthTypes.Timestamp_ns
}

type TimestampNtzSliceType struct{}

func (t TimestampNtzSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t TimestampNtzSliceType) IsNumeric() bool {
	return false
}

func (t TimestampNtzSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.FixedWidthTypes.Timestamp_ns)
}

type DateType struct{}

func (t DateType) TypeName() string {
	return getDataTypeName(t)
}

func (t DateType) IsNumeric() bool {
	return false
}

func (t DateType) ToArrowType() arrow.DataType {
	return arrow.FixedWidthTypes.Date32
}

type DateSliceType struct{}

func (t DateSliceType) TypeName() string {
	return getDataTypeName(t)
}

func (t DateSliceType) IsNumeric() bool {
	return false
}

func (t DateSliceType) ToArrowType() arrow.DataType {
	return arrow.ListOf(arrow.FixedWidthTypes.Date32)
}

type UnsupportedType struct {
	TypeInfo any
}

func (t UnsupportedType) TypeName() string {
	return getDataTypeName(t)
}

func (t UnsupportedType) IsNumeric() bool {
	return false
}

func (t UnsupportedType) ToArrowType() arrow.DataType {
	return nil
}

func getDataTypeName(dataType DataType) string {
	typeName := fmt.Sprintf("%T", dataType)
	nonQualifiedTypeName := strings.Split(typeName, ".")[1]
	return strings.TrimSuffix(nonQualifiedTypeName, "Type")
}

var (
	BOOLEAN             = BooleanType{}
	BYTE                = ByteType{}
	SHORT               = ShortType{}
	INTEGER             = IntegerType{}
	LONG                = LongType{}
	FLOAT               = FloatType{}
	DOUBLE              = DoubleType{}
	DATE                = DateType{}
	TIMESTAMP           = TimestampType{}
	TIMESTAMP_MTZ       = TimestampMtzType{}
	TIMESTAMP_UTZ       = TimestampUtzType{}
	TIMESTAMP_NTZ       = TimestampNtzType{}
	STRING              = StringType{}
	BOOLEAN_SLICE       = BooleanSliceType{}
	BYTE_SLICE          = ByteSliceType{}
	SHORT_SLICE         = ShortSliceType{}
	INTEGER_SLICE       = IntegerSliceType{}
	LONG_SLICE          = LongSliceType{}
	FLOAT_SLICE         = FloatSliceType{}
	DOUBLE_SLICE        = DoubleSliceType{}
	DATE_SLICE          = DateSliceType{}
	TIMESTAMP_SLICE     = TimestampSliceType{}
	TIMESTAMP_MTZ_SLICE = TimestampMtzSliceType{}
	TIMESTAMP_UTZ_SLICE = TimestampUtzSliceType{}
	TIMESTAMP_NTZ_SLICE = TimestampNtzSliceType{}
	STRING_SLICE        = StringSliceType{}
)
