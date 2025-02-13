// Package sql @author: Violet-Eva @date  : 2025/2/12 @notes :
package sql

import (
	"context"
	proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"
)

type DataFrameStreamManager interface {
}

func NewDataFrameStreamManager(sparkExecutor *sparkSessionImpl) DataFrameStreamManager {
	return &dataFrameStreamManagerImpl{
		sparkExecutor: sparkExecutor,
	}
}

type dataFrameStreamManagerImpl struct {
	sparkExecutor *sparkSessionImpl
	relation      *proto.Relation
	saveMode      string
	formatSource  string
	options       map[string]string
	write         *proto.WriteOperation
}

func (w *dataFrameWriterImpl) SaveStreamManager(ctx context.Context, queryName string) error {

	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_StreamingQueryManagerCommand{
					StreamingQueryManagerCommand: &proto.StreamingQueryManagerCommand{
						Command: &proto.StreamingQueryManagerCommand_Active{},
					},
				},
			},
		},
	}
	responseClient, err := w.sparkExecutor.client.ExecutePlan(ctx, plan)
	if err != nil {
		return err
	}
	_, _, err = responseClient.ToTable()
	return err
}
