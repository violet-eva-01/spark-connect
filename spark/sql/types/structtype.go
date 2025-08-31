package types

import "github.com/apache/arrow-go/v18/arrow"

// StructField represents a field in a StructType.
type StructField struct {
	Name     string
	DataType DataType
	Format   string
	Nullable bool // default should be true
	Metadata *string
}

func (t *StructField) ToArrowType() arrow.Field {
	return arrow.Field{
		Name:     t.Name,
		Type:     t.DataType.ToArrowType(),
		Nullable: t.Nullable,
	}
}

// StructType represents a struct type.
type StructType struct {
	Fields []StructField
}

func (t *StructType) TypeName() string {
	return "structtype"
}

func (t *StructType) IsNumeric() bool {
	return false
}

func (t *StructType) ToArrowType() *arrow.StructType {
	fields := make([]arrow.Field, len(t.Fields))
	for i, f := range t.Fields {
		fields[i] = f.ToArrowType()
	}
	return arrow.StructOf(fields...)
}

func StructOf(fields ...StructField) *StructType {
	return &StructType{Fields: fields}
}

func NewStructField(name string, dataType DataType) StructField {
	return StructField{
		Name:     name,
		DataType: dataType,
		Nullable: true,
	}
}
