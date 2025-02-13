// Package sql @author: Violet-Eva @date  : 2025/2/13 @notes :
package sql

import (
	"context"
	"fmt"
	proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"
	"github.com/violet-eva-01/spark-connect/spark/sparkerrors"
	"github.com/violet-eva-01/spark-connect/spark/sql/utils"
)

type DataFrameStreamQuery interface {
	Status(ctx context.Context, status bool) (map[string]any, error)
	LastProgress(ctx context.Context, lastProgress bool) (map[string]any, error)
	RecentProgress(ctx context.Context, recentProgress bool) (map[string]any, error)
	Stop(ctx context.Context, stop bool) (map[string]any, error)
	ProcessAllAvailable(ctx context.Context, processAllAvailable bool) (map[string]any, error)
	Exception(ctx context.Context, exception bool) (map[string]any, error)
	AwaitTermination(ctx context.Context, timeoutMs int64) (map[string]any, error)
	Explain(ctx context.Context, explain bool, explainMode utils.ExplainMode) (string, error)
}

func NewDataFrameStreamQuery(sparkExecutor *sparkSessionImpl, id, runId string) DataFrameStreamQuery {
	return &dataFrameStreamQueryImpl{
		sparkExecutor: sparkExecutor,
		id:            id,
		runId:         runId,
	}
}

type dataFrameStreamQueryImpl struct {
	sparkExecutor *sparkSessionImpl
	id            string
	runId         string
}

func (w *dataFrameStreamQueryImpl) Status(ctx context.Context, status bool) (map[string]any, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_Status{
							Status: status,
						},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	_, _, err = responseClient.ToTable()
	if err != nil {
		return nil, err
	}
	properties := responseClient.Properties()
	return properties, err
}

func (w *dataFrameStreamQueryImpl) LastProgress(ctx context.Context, lastProgress bool) (map[string]any, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_LastProgress{
							LastProgress: lastProgress,
						},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	_, _, err = responseClient.ToTable()
	if err != nil {
		return nil, err
	}
	properties := responseClient.Properties()
	return properties, err
}

func (w *dataFrameStreamQueryImpl) RecentProgress(ctx context.Context, recentProgress bool) (map[string]any, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_RecentProgress{
							RecentProgress: recentProgress,
						},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	_, _, err = responseClient.ToTable()
	if err != nil {
		return nil, err
	}
	properties := responseClient.Properties()
	return properties, err
}

func (w *dataFrameStreamQueryImpl) Stop(ctx context.Context, stop bool) (map[string]any, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_Stop{
							Stop: stop,
						},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	_, _, err = responseClient.ToTable()
	if err != nil {
		return nil, err
	}
	properties := responseClient.Properties()
	return properties, err
}

func (w *dataFrameStreamQueryImpl) ProcessAllAvailable(ctx context.Context, processAllAvailable bool) (map[string]any, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_ProcessAllAvailable{
							ProcessAllAvailable: processAllAvailable,
						},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	_, _, err = responseClient.ToTable()
	if err != nil {
		return nil, err
	}
	properties := responseClient.Properties()
	return properties, err
}

func (w *dataFrameStreamQueryImpl) Explain(ctx context.Context, explain bool, explainMode utils.ExplainMode) (string, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_Explain{
							Explain: &proto.StreamingQueryCommand_ExplainCommand{
								Extended: explain,
							},
						},
					},
				},
			},
		},
	}
	response, err := w.sparkExecutor.client.Explain(ctx, plan, explainMode)
	if err != nil {
		return "", sparkerrors.WithType(fmt.Errorf("failed to execute plan: %w", err), sparkerrors.ExecutionError)
	}
	return response.GetExplain().GetExplainString(), err
}

func (w *dataFrameStreamQueryImpl) Exception(ctx context.Context, exception bool) (map[string]any, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_Exception{
							Exception: exception,
						},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	_, _, err = responseClient.ToTable()
	if err != nil {
		return nil, err
	}
	properties := responseClient.Properties()
	return properties, err
}

func (w *dataFrameStreamQueryImpl) AwaitTermination(ctx context.Context, timeoutMs int64) (map[string]any, error) {
	var timeout *int64
	timeout = &timeoutMs
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryCommand{
					StreamingQueryCommand: &proto.StreamingQueryCommand{
						QueryId: &proto.StreamingQueryInstanceId{
							Id:    w.id,
							RunId: w.runId,
						},
						Command: &proto.StreamingQueryCommand_AwaitTermination{
							AwaitTermination: &proto.StreamingQueryCommand_AwaitTerminationCommand{
								TimeoutMs: timeout,
							},
						},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	_, _, err = responseClient.ToTable()
	if err != nil {
		return nil, err
	}
	properties := responseClient.Properties()
	return properties, err
}
