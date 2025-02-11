package sql

import (
	"context"
	"github.com/violet-eva-01/spark-connect/spark/sql/types"
)

type DataFrameNaFunctions interface {
	Drop(ctx context.Context, cols ...string) (DataFrame, error)
	DropAll(ctx context.Context, cols ...string) (DataFrame, error)
	DropWithThreshold(ctx context.Context, threshold int32, cols ...string) (DataFrame, error)
	Fill(ctx context.Context, value types.PrimitiveTypeLiteral, cols ...string) (DataFrame, error)
	FillWithValues(ctx context.Context, values map[string]types.PrimitiveTypeLiteral) (DataFrame, error)
	Replace(ctx context.Context, toReplace []types.PrimitiveTypeLiteral,
		values []types.PrimitiveTypeLiteral, cols ...string) (DataFrame, error)
}

type dataFrameNaFunctionsImpl struct {
	dataFrame DataFrame
}

func (d *dataFrameNaFunctionsImpl) Drop(ctx context.Context, cols ...string) (DataFrame, error) {
	return d.dataFrame.DropNa(ctx, cols...)
}

func (d *dataFrameNaFunctionsImpl) DropAll(ctx context.Context, cols ...string) (DataFrame, error) {
	return d.dataFrame.DropNaAll(ctx, cols...)
}

func (d *dataFrameNaFunctionsImpl) DropWithThreshold(ctx context.Context, threshold int32, cols ...string) (DataFrame, error) {
	return d.dataFrame.DropNaWithThreshold(ctx, threshold, cols...)
}

func (d *dataFrameNaFunctionsImpl) Fill(ctx context.Context, value types.PrimitiveTypeLiteral, cols ...string) (DataFrame, error) {
	return d.dataFrame.FillNa(ctx, value, cols...)
}

func (d *dataFrameNaFunctionsImpl) FillWithValues(ctx context.Context,
	values map[string]types.PrimitiveTypeLiteral,
) (DataFrame, error) {
	return d.dataFrame.FillNaWithValues(ctx, values)
}

func (d *dataFrameNaFunctionsImpl) Replace(ctx context.Context,
	toReplace []types.PrimitiveTypeLiteral, values []types.PrimitiveTypeLiteral, cols ...string,
) (DataFrame, error) {
	return d.dataFrame.Replace(ctx, toReplace, values, cols...)
}
