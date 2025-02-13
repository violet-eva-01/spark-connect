package sql

import (
	"context"
	proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"
)

// DataFrameWriterStream supports writing data frame to storage.
type DataFrameWriterStream interface {
	// Mode specifies saving mode for the data, e.g. Append, Overwrite, ErrorIfExists.
	Mode(saveMode string) DataFrameWriterStream
	// Format specifies data format (data source type) for the underlying data, e.g. parquet.
	Format(source string) DataFrameWriterStream
	// Trigger The trigger settings of a streaming query define the timing of streaming data processing
	Trigger(triggerType string, value interface{}) DataFrameWriterStream

	QueryName(queryName string) DataFrameWriterStream

	Foreach() DataFrameWriterStream

	ForeachBatch() DataFrameWriterStream
	// Option set params
	Option(key, value string) DataFrameWriterStream

	Start(ctx context.Context) (map[string]any, error)
}

func newDataFrameWriterStream(sparkExecutor *sparkSessionImpl, relation *proto.Relation) DataFrameWriterStream {
	return &dataFrameWriterStreamImpl{
		sparkExecutor: sparkExecutor,
		write: &proto.WriteStreamOperationStart{
			Input: relation,
		},
	}
}

// dataFrameWriterImpl is an implementation of DataFrameWriter interface.
type dataFrameWriterStreamImpl struct {
	sparkExecutor *sparkSessionImpl
	write         *proto.WriteStreamOperationStart
}

func (w *dataFrameWriterStreamImpl) Foreach() DataFrameWriterStream {
	return w
}

func (w *dataFrameWriterStreamImpl) ForeachBatch() DataFrameWriterStream {
	return w
}

func (w *dataFrameWriterStreamImpl) QueryName(queryName string) DataFrameWriterStream {
	w.write.QueryName = queryName
	return w
}

func (w *dataFrameWriterStreamImpl) Mode(saveMode string) DataFrameWriterStream {
	w.write.OutputMode = saveMode
	return w
}

func (w *dataFrameWriterStreamImpl) Format(format string) DataFrameWriterStream {
	w.write.Format = format
	return w
}

func (w *dataFrameWriterStreamImpl) Option(key, value string) DataFrameWriterStream {
	if w.write.Options == nil {
		w.write.Options = make(map[string]string)
	}
	w.write.Options[key] = value
	return w
}

func (w *dataFrameWriterStreamImpl) Start(ctx context.Context) (map[string]any, error) {
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_WriteStreamOperationStart{
					WriteStreamOperationStart: w.write,
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

func (w *dataFrameWriterStreamImpl) Trigger(triggerType string, value any) DataFrameWriterStream {
	switch triggerType {
	case "ProcessingTimeInterval":
		w.write.Trigger = &proto.WriteStreamOperationStart_ProcessingTimeInterval{ProcessingTimeInterval: value.(string)}
	case "AvailableNow":
		w.write.Trigger = &proto.WriteStreamOperationStart_AvailableNow{AvailableNow: value.(bool)}
	case "Once":
		w.write.Trigger = &proto.WriteStreamOperationStart_Once{Once: value.(bool)}
	case "ContinuousCheckpointInterval":
		w.write.Trigger = &proto.WriteStreamOperationStart_ContinuousCheckpointInterval{ContinuousCheckpointInterval: value.(string)}
	}
	return w
}
