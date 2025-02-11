package base

import (
	"context"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/violet-eva-01/spark-connect/internal/generated"
	"github.com/violet-eva-01/spark-connect/spark/sql/types"
	"github.com/violet-eva-01/spark-connect/spark/sql/utils"
)

type SparkConnectRPCClient generated.SparkConnectServiceClient

// SparkConnectClient is the interface for executing a plan in Spark.
//
// This interface does not deal with the public Spark API abstractions but roughly deals on the
// RPC API level and the necessary translation of Arrow to Row objects.
type SparkConnectClient interface {
	ExecutePlan(ctx context.Context, plan *generated.Plan) (ExecuteResponseStream, error)
	ExecuteCommand(ctx context.Context, plan *generated.Plan) (arrow.Table, *types.StructType, map[string]any, error)
	AnalyzePlan(ctx context.Context, plan *generated.Plan) (*generated.AnalyzePlanResponse, error)
	Explain(ctx context.Context, plan *generated.Plan, explainMode utils.ExplainMode) (*generated.AnalyzePlanResponse, error)
	Persist(ctx context.Context, plan *generated.Plan, storageLevel utils.StorageLevel) error
	Unpersist(ctx context.Context, plan *generated.Plan) error
	GetStorageLevel(ctx context.Context, plan *generated.Plan) (*utils.StorageLevel, error)
	SparkVersion(ctx context.Context) (string, error)
	DDLParse(ctx context.Context, sql string) (*types.StructType, error)
	SameSemantics(ctx context.Context, plan1 *generated.Plan, plan2 *generated.Plan) (bool, error)
	SemanticHash(ctx context.Context, plan *generated.Plan) (int32, error)
	Config(ctx context.Context, configRequest *generated.ConfigRequest_Operation) (*generated.ConfigResponse, error)
}

type ExecuteResponseStream interface {
	ToTable() (*types.StructType, arrow.Table, error)
	Properties() map[string]any
}
