package mocks

import (
	"context"
	"errors"

	"github.com/violet-eva-01/spark-connect/spark/sql/utils"

	"github.com/violet-eva-01/spark-connect/spark/client/base"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/violet-eva-01/spark-connect/internal/generatedCustom"
	"github.com/violet-eva-01/spark-connect/spark/sql/types"
)

type TestExecutor struct {
	Client   base.ExecuteResponseStream
	response *generatedCustom.AnalyzePlanResponse
	Err      error
}

func (t *TestExecutor) ExecutePlan(ctx context.Context, plan *generatedCustom.Plan) (base.ExecuteResponseStream, error) {
	if t.Err != nil {
		return nil, t.Err
	}
	return t.Client, nil
}

func (t *TestExecutor) AnalyzePlan(ctx context.Context, plan *generatedCustom.Plan) (*generatedCustom.AnalyzePlanResponse, error) {
	return t.response, nil
}

func (t *TestExecutor) Explain(ctx context.Context, plan *generatedCustom.Plan,
	explainMode utils.ExplainMode,
) (*generatedCustom.AnalyzePlanResponse, error) {
	return nil, errors.New("not implemented")
}

func (t *TestExecutor) ExecuteCommand(ctx context.Context, plan *generatedCustom.Plan) (arrow.Table, *types.StructType, map[string]interface{}, error) {
	if t.Err != nil {
		return nil, nil, nil, t.Err
	}
	return nil, nil, nil, nil
}

func (t *TestExecutor) Persist(ctx context.Context, plan *generatedCustom.Plan, storageLevel utils.StorageLevel) error {
	return errors.New("not implemented")
}

func (t *TestExecutor) Unpersist(ctx context.Context, plan *generatedCustom.Plan) error {
	return errors.New("not implemented")
}

func (t *TestExecutor) GetStorageLevel(ctx context.Context, plan *generatedCustom.Plan) (*utils.StorageLevel, error) {
	return nil, errors.New("not implemented")
}

func (t *TestExecutor) SparkVersion(ctx context.Context) (string, error) {
	return "", errors.New("not implemented")
}

func (t *TestExecutor) DDLParse(ctx context.Context, sql string) (*types.StructType, error) {
	return nil, errors.New("not implemented")
}

func (t *TestExecutor) SameSemantics(ctx context.Context, plan1 *generatedCustom.Plan, plan2 *generatedCustom.Plan) (bool, error) {
	return false, errors.New("not implemented")
}

func (t *TestExecutor) SemanticHash(ctx context.Context, plan *generatedCustom.Plan) (int32, error) {
	return 0, errors.New("not implemented")
}

func (t *TestExecutor) Config(ctx context.Context, configRequest *generatedCustom.ConfigRequest_Operation) (*generatedCustom.ConfigResponse, error) {
	return nil, errors.New("not implemented")
}
