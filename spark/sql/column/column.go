package column

import (
	"context"

	"github.com/violet-eva-01/spark-connect/spark/sql/types"

	proto "github.com/violet-eva-01/spark-connect/internal/generated"
)

// Convertible is the interface for all things that can be converted into a protobuf expression.
type Convertible interface {
	ToProto(ctx context.Context) (*proto.Expression, error)
}

type Column struct {
	expr expression
}

func (c Column) ToProto(ctx context.Context) (*proto.Expression, error) {
	return c.expr.ToProto(ctx)
}

func (c Column) Lt(other Column) Column {
	return NewColumn(NewUnresolvedFunction("<", []expression{c.expr, other.expr}, false))
}

func (c Column) Le(other Column) Column {
	return NewColumn(NewUnresolvedFunction("<=", []expression{c.expr, other.expr}, false))
}

func (c Column) Gt(other Column) Column {
	return NewColumn(NewUnresolvedFunction(">", []expression{c.expr, other.expr}, false))
}

func (c Column) Ge(other Column) Column {
	return NewColumn(NewUnresolvedFunction(">=", []expression{c.expr, other.expr}, false))
}

func (c Column) Eq(other Column) Column {
	return NewColumn(NewUnresolvedFunction("==", []expression{c.expr, other.expr}, false))
}

func (c Column) Neq(other Column) Column {
	cmp := NewUnresolvedFunction("==", []expression{c.expr, other.expr}, false)
	return NewColumn(NewUnresolvedFunction("not", []expression{cmp}, false))
}

func (c Column) Mul(other Column) Column {
	return NewColumn(NewUnresolvedFunction("*", []expression{c.expr, other.expr}, false))
}

func (c Column) Div(other Column) Column {
	return NewColumn(NewUnresolvedFunction("/", []expression{c.expr, other.expr}, false))
}

func (c Column) Desc() Column {
	return NewColumn(&sortExpression{
		child:        c.expr,
		direction:    proto.Expression_SortOrder_SORT_DIRECTION_DESCENDING,
		nullOrdering: proto.Expression_SortOrder_SORT_NULLS_LAST,
	})
}

func (c Column) GetItem(key types.LiteralType) Column {
	return NewColumn(NewUnresolvedExtractValue("getItem", c.expr, NewLiteral(key)))
}

func (c Column) Asc() Column {
	return NewColumn(&sortExpression{
		child:        c.expr,
		direction:    proto.Expression_SortOrder_SORT_DIRECTION_ASCENDING,
		nullOrdering: proto.Expression_SortOrder_SORT_NULLS_FIRST,
	})
}

func (c Column) Alias(alias string) Column {
	return NewColumn(NewColumnAlias(alias, c.expr))
}

func NewColumn(expr expression) Column {
	return Column{
		expr: expr,
	}
}

type SchemaDataFrame interface {
	PlanId() int64
	Schema(ctx context.Context) (*types.StructType, error)
}

func OfDF(df SchemaDataFrame, colName string) Column {
	return NewColumn(&delayedColumnReference{colName, df})
}

func OfDFWithRegex(df SchemaDataFrame, colRegex string) Column {
	planId := df.PlanId()
	return NewColumn(&unresolvedRegex{colRegex, &planId})
}

type Alias struct {
	Name string
	Col  Convertible
}

func (a Alias) ToProto(ctx context.Context) (*proto.Expression, error) {
	col, err := a.Col.ToProto(ctx)
	if err != nil {
		return nil, err
	}
	return &proto.Expression{
		ExprType: &proto.Expression_Alias_{
			Alias: &proto.Expression_Alias{
				Expr: col,
				Name: []string{a.Name},
			},
		},
	}, nil
}

func WithAlias(name string, col Convertible) Alias {
	return Alias{
		Name: name,
		Col:  col,
	}
}
