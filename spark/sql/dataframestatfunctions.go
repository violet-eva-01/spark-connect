package sql

import "context"

type DataFrameStatFunctions interface {
	ApproxQuantile(ctx context.Context, probabilities []float64, relativeError float64, cols ...string) ([][]float64, error)
	Cov(ctx context.Context, col1, col2 string) (float64, error)
	Corr(ctx context.Context, col1, col2 string) (float64, error)
	CorrWithMethod(ctx context.Context, col1, col2 string, method string) (float64, error)
	CrossTab(ctx context.Context, col1, col2 string) DataFrame
	FreqItems(ctx context.Context, cols ...string) DataFrame
	FreqItemsWithSupport(ctx context.Context, support float64, cols ...string) DataFrame
	Sample(ctx context.Context, fraction float64) (DataFrame, error)
	SampleWithReplacement(ctx context.Context, withReplacement bool, fraction float64) (DataFrame, error)
	SampleWithSeed(ctx context.Context, fraction float64, seed int64) (DataFrame, error)
	SampleWithReplacementAndSeed(ctx context.Context, withReplacement bool, fraction float64, seed int64) (DataFrame, error)
}

type dataFrameStatFunctionsImpl struct {
	df DataFrame
}

func (d *dataFrameStatFunctionsImpl) Sample(ctx context.Context, fraction float64) (DataFrame, error) {
	return d.df.Sample(ctx, fraction)
}

func (d *dataFrameStatFunctionsImpl) SampleWithReplacement(ctx context.Context,
	withReplacement bool, fraction float64,
) (DataFrame, error) {
	return d.df.SampleWithReplacement(ctx, withReplacement, fraction)
}

func (d *dataFrameStatFunctionsImpl) SampleWithSeed(ctx context.Context, fraction float64, seed int64) (DataFrame, error) {
	return d.df.SampleWithSeed(ctx, fraction, seed)
}

func (d *dataFrameStatFunctionsImpl) SampleWithReplacementAndSeed(ctx context.Context,
	withReplacement bool, fraction float64, seed int64,
) (DataFrame, error) {
	return d.df.SampleWithReplacementAndSeed(ctx, withReplacement, fraction, seed)
}

func (d *dataFrameStatFunctionsImpl) ApproxQuantile(ctx context.Context, probabilities []float64,
	relativeError float64, cols ...string,
) ([][]float64, error) {
	return d.df.ApproxQuantile(ctx, probabilities, relativeError, cols...)
}

func (d *dataFrameStatFunctionsImpl) Cov(ctx context.Context, col1, col2 string) (float64, error) {
	return d.df.Cov(ctx, col1, col2)
}

func (d *dataFrameStatFunctionsImpl) Corr(ctx context.Context, col1, col2 string) (float64, error) {
	return d.df.Corr(ctx, col1, col2)
}

func (d *dataFrameStatFunctionsImpl) CorrWithMethod(ctx context.Context, col1, col2 string, method string) (float64, error) {
	return d.df.CorrWithMethod(ctx, col1, col2, method)
}

func (d *dataFrameStatFunctionsImpl) CrossTab(ctx context.Context, col1, col2 string) DataFrame {
	return d.df.CrossTab(ctx, col1, col2)
}

func (d *dataFrameStatFunctionsImpl) FreqItems(ctx context.Context, cols ...string) DataFrame {
	return d.df.FreqItems(ctx, cols...)
}

func (d *dataFrameStatFunctionsImpl) FreqItemsWithSupport(ctx context.Context, support float64, cols ...string) DataFrame {
	return d.df.FreqItemsWithSupport(ctx, support, cols...)
}
