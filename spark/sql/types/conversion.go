package types

import (
	"errors"

	"github.com/violet-eva-01/spark-connect/internal/generatedCustom"
	"github.com/violet-eva-01/spark-connect/spark/sparkerrors"
)

func ConvertProtoDataTypeToStructType(input *generatedCustom.DataType) (*StructType, error) {
	dataTypeStruct := input.GetStruct()
	if dataTypeStruct == nil {
		return nil, sparkerrors.WithType(errors.New("dataType.GetStruct() is nil"), sparkerrors.InvalidInputError)
	}
	return &StructType{
		Fields: ConvertProtoStructFields(dataTypeStruct.Fields),
	}, nil
}

func ConvertProtoStructFields(input []*generatedCustom.DataType_StructField) []StructField {
	result := make([]StructField, len(input))
	for i, f := range input {
		result[i] = ConvertProtoStructField(f)
	}
	return result
}

func ConvertProtoStructField(field *generatedCustom.DataType_StructField) StructField {
	return StructField{
		Name:     field.Name,
		DataType: ConvertProtoDataTypeToDataType(field.DataType),
		Nullable: field.Nullable,
		Metadata: field.Metadata,
	}
}

// ConvertProtoDataTypeToDataType converts protobuf data type to Spark connect sql data type
func ConvertProtoDataTypeToDataType(input *generatedCustom.DataType) DataType {
	switch v := input.GetKind().(type) {
	case *generatedCustom.DataType_Boolean_:
		return BooleanType{}
	case *generatedCustom.DataType_Byte_:
		return ByteType{}
	case *generatedCustom.DataType_Short_:
		return ShortType{}
	case *generatedCustom.DataType_Integer_:
		return IntegerType{}
	case *generatedCustom.DataType_Long_:
		return LongType{}
	case *generatedCustom.DataType_Float_:
		return FloatType{}
	case *generatedCustom.DataType_Double_:
		return DoubleType{}
	case *generatedCustom.DataType_Decimal_:
		return DecimalType{}
	case *generatedCustom.DataType_String_:
		return StringType{}
	case *generatedCustom.DataType_Binary_:
		return BinaryType{}
	case *generatedCustom.DataType_Timestamp_:
		return TimestampType{}
	case *generatedCustom.DataType_TimestampNtz:
		return TimestampNtzType{}
	case *generatedCustom.DataType_Date_:
		return DateType{}
	default:
		return UnsupportedType{
			TypeInfo: v,
		}
	}
}
