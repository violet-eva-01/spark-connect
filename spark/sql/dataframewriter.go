package sql

import (
	"context"
	"fmt"
	"strings"

	proto "github.com/violet-eva-01/spark-connect/internal/generatedCustom"
	"github.com/violet-eva-01/spark-connect/spark/sparkerrors"
)

// DataFrameWriter supports writing data frame to storage.
type DataFrameWriter interface {
	// Mode specifies saving mode for the data, e.g. Append, Overwrite, ErrorIfExists.
	Mode(saveMode string) DataFrameWriter
	// Format specifies data format (data source type) for the underlying data, e.g. parquet.
	Format(source string) DataFrameWriter
	// SaveAsPath writes data frame to the given path.
	SaveAsPath(ctx context.Context, path string) error
	// SaveAsTable writes data frame to the table.
	SaveAsTable(ctx context.Context, table string) error
	SaveAsTableTable(ctx context.Context) error
	// SaveAsTableV2 writes data frame to the table.
	SaveAsTableV2(ctx context.Context, table string) error
	SaveAsTableV2Table(ctx context.Context) error
	SaveAsStream(ctx context.Context) error
	SaveAsStreamTable(ctx context.Context, table string) error
	// Option set params
	Option(key, value string) DataFrameWriter
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

func (w *dataFrameWriterImpl) SaveAsPath(ctx context.Context, path string) error {
	saveMode, err := getSavePathMode(w.saveMode)
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

func (w *dataFrameWriterImpl) SaveAsTable(ctx context.Context, table string) error {
	saveMode, err := getSavePathMode(w.saveMode)
	if err != nil {
		return err
	}
	saveTableMode, err := getSaveTableMode(w.saveMode)
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
						Source:  source,
						Options: w.options,
						Mode:    saveMode,
						SaveType: &proto.WriteOperation_Table{
							Table: &proto.WriteOperation_SaveTable{
								TableName:  table,
								SaveMethod: saveTableMode,
							},
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

func (w *dataFrameWriterImpl) SaveAsTableTable(ctx context.Context) error {
	saveMode, err := getSavePathMode(w.saveMode)
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
						Source:  source,
						Options: w.options,
						Mode:    saveMode,
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

func getSavePathMode(mode string) (proto.WriteOperation_SaveMode, error) {
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
		return 0, sparkerrors.WithType(fmt.Errorf("unsupported save path mode: %s", mode), sparkerrors.InvalidInputError)
	}
}

func getSaveTableMode(mode string) (proto.WriteOperation_SaveTable_TableSaveMethod, error) {
	if mode == "" {
		return proto.WriteOperation_SaveTable_TABLE_SAVE_METHOD_UNSPECIFIED, nil
	} else if strings.EqualFold(mode, "Append") {
		return proto.WriteOperation_SaveTable_TABLE_SAVE_METHOD_INSERT_INTO, nil
	} else if strings.EqualFold(mode, "Overwrite") {
		return proto.WriteOperation_SaveTable_TABLE_SAVE_METHOD_SAVE_AS_TABLE, nil
	} else {
		return 0, sparkerrors.WithType(fmt.Errorf("unsupported save table mode: %s", mode), sparkerrors.InvalidInputError)
	}
}

func (w *dataFrameWriterImpl) SaveAsTableV2(ctx context.Context, table string) error {
	saveTableMode, err := getSaveTableModeV2(w.saveMode)
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
				CommandType: &proto.Command_WriteOperationV2{
					WriteOperationV2: &proto.WriteOperationV2{
						Input:     w.relation,
						Options:   w.options,
						TableName: table,
						Provider:  source,
						Mode:      saveTableMode,
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

func (w *dataFrameWriterImpl) SaveAsTableV2Table(ctx context.Context) error {
	saveTableMode, err := getSaveTableModeV2(w.saveMode)
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
				CommandType: &proto.Command_WriteOperationV2{
					WriteOperationV2: &proto.WriteOperationV2{
						Input:    w.relation,
						Options:  w.options,
						Provider: source,
						Mode:     saveTableMode,
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

func getSaveTableModeV2(mode string) (proto.WriteOperationV2_Mode, error) {
	if mode == "" {
		return proto.WriteOperationV2_MODE_UNSPECIFIED, nil
	} else if strings.EqualFold(mode, "Create") {
		return proto.WriteOperationV2_MODE_CREATE, nil
	} else if strings.EqualFold(mode, "Overwrite") {
		return proto.WriteOperationV2_MODE_OVERWRITE, nil
	} else if strings.EqualFold(mode, "OverwritePartition") {
		return proto.WriteOperationV2_MODE_OVERWRITE_PARTITIONS, nil
	} else if strings.EqualFold(mode, "Append") {
		return proto.WriteOperationV2_MODE_APPEND, nil
	} else if strings.EqualFold(mode, "Replace") {
		return proto.WriteOperationV2_MODE_REPLACE, nil
	} else if strings.EqualFold(mode, "CreateOrReplace") {
		return proto.WriteOperationV2_MODE_CREATE_OR_REPLACE, nil
	} else {
		return 0, sparkerrors.WithType(fmt.Errorf("unsupported saveTableV2 mode: %s", mode), sparkerrors.InvalidInputError)
	}
}

func (w *dataFrameWriterImpl) SaveAsStream(ctx context.Context) error {
	var format string
	if w.formatSource != "" {
		format = w.formatSource
	}
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_WriteStreamOperationStart{
					WriteStreamOperationStart: &proto.WriteStreamOperationStart{
						Input:   w.relation,
						Options: w.options,
						Format:  format,
						Trigger: &proto.WriteStreamOperationStart_Once{
							Once: true,
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

func (w *dataFrameWriterImpl) SaveAsStreamTable(ctx context.Context, table string) error {
	var format string
	if w.formatSource != "" {
		format = w.formatSource
	}
	plan := &proto.Plan{
		OpType: &proto.Plan_Command{
			Command: &proto.Command{
				CommandType: &proto.Command_WriteStreamOperationStart{
					WriteStreamOperationStart: &proto.WriteStreamOperationStart{
						Input:   w.relation,
						Options: w.options,
						Format:  format,
						Trigger: &proto.WriteStreamOperationStart_Once{
							Once: true,
						},
						SinkDestination: &proto.WriteStreamOperationStart_TableName{
							TableName: table,
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
