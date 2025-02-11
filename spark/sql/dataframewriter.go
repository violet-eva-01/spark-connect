package sql

import (
	"context"
	"fmt"
	"strings"

	proto "github.com/violet-eva-01/spark-connect/internal/generated"
	"github.com/violet-eva-01/spark-connect/spark/sparkerrors"
)

// DataFrameWriter supports writing data frame to storage.
type DataFrameWriter interface {
	// Mode specifies saving mode for the data, e.g. Append, Overwrite, ErrorIfExists.
	Mode(saveMode string) DataFrameWriter
	// Format specifies data format (data source type) for the underlying data, e.g. parquet.
	Format(source string) DataFrameWriter
	// Save writes data frame to the given path.
	Save(ctx context.Context, path string) error
}

func newDataFrameWriter(sparkExecutor *sparkSessionImpl, relation *proto.Relation) DataFrameWriter {
	return &dataFrameWriterImpl{
		sparkExecutor: sparkExecutor,
		relation:      relation,
	}
}

// dataFrameWriterImpl is an implementation of DataFrameWriter interface.
type dataFrameWriterImpl struct {
	sparkExecutor *sparkSessionImpl
	relation      *proto.Relation
	saveMode      string
	formatSource  string
	options       map[string]string
}

func (w *dataFrameWriterImpl) Mode(saveMode string) DataFrameWriter {
	w.saveMode = saveMode
	return w
}

func (w *dataFrameWriterImpl) Format(source string) DataFrameWriter {
	w.formatSource = source
	return w
}

func (w *dataFrameWriterImpl) Option(key, value string) DataFrameWriter {
	if w.options == nil {
		w.options = make(map[string]string)
	}
	w.options[key] = value
	return w
}

func (w *dataFrameWriterImpl) Save(ctx context.Context, path string) error {
	saveMode, err := getSaveMode(w.saveMode)
	if err != nil {
		return err
	}
	var source *string
	if w.formatSource != "" {
		source = &w.formatSource
	}
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_WriteOperation{
					WriteOperation: &proto.WriteOperation{
						Input:   w.relation,
						Mode:    saveMode,
						Source:  source,
						Options: w.options,
						SaveType: &proto.WriteOperation_Path{
							Path: path,
						},
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

func getSaveMode(mode string) (proto.WriteOperation_SaveMode, error) {
	if mode == "" {
		return proto.WriteOperation_SAVE_MODE_UNSPECIFIED, nil
	} else if strings.EqualFold(mode, "Append") {
		return proto.WriteOperation_SAVE_MODE_APPEND, nil
	} else if strings.EqualFold(mode, "Overwrite") {
		return proto.WriteOperation_SAVE_MODE_OVERWRITE, nil
	} else if strings.EqualFold(mode, "ErrorIfExists") {
		return proto.WriteOperation_SAVE_MODE_ERROR_IF_EXISTS, nil
	} else if strings.EqualFold(mode, "Ignore") {
		return proto.WriteOperation_SAVE_MODE_IGNORE, nil
	} else {
		return 0, sparkerrors.WithType(fmt.Errorf("unsupported save mode: %s", mode), sparkerrors.InvalidInputError)
	}
}
