package functions

import (
	"github.com/violet-eva-01/spark-connect/spark/sql/column"
	"github.com/violet-eva-01/spark-connect/spark/sql/types"
)

func Expr(expr string) column.Column {
	return column.NewColumn(column.NewSQLExpression(expr))
}

func Col(name string) column.Column {
	return column.NewColumn(column.NewColumnReference(name))
}

func Lit(value types.LiteralType) column.Column {
	return column.NewColumn(column.NewLiteral(value))
}

func Int8Lit(value int8) column.Column {
	return Lit(types.Int8(value))
}

func Int16Lit(value int16) column.Column {
	return Lit(types.Int16(value))
}

func Int32Lit(value int32) column.Column {
	return Lit(types.Int32(value))
}

func Int64Lit(value int64) column.Column {
	return Lit(types.Int64(value))
}

func Float32Lit(value float32) column.Column {
	return Lit(types.Float32(value))
}

func Float64Lit(value float64) column.Column {
	return Lit(types.Float64(value))
}

func StringLit(value string) column.Column {
	return Lit(types.String(value))
}

func BoolLit(value bool) column.Column {
	return Lit(types.Boolean(value))
}

func BinaryLit(value []byte) column.Column {
	return Lit(types.Binary(value))
}

func IntLit(value int) column.Column {
	return Lit(types.Int(value))
}
