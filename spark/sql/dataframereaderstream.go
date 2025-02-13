package sql

// DataFrameReaderStream supports reading data from storage and returning a data frame.
// TODO needs to implement other methods like Option(), Schema(), and also "strong typed"
// reading (e.g. Parquet(), Orc(), Csv(), etc.
type DataFrameReaderStream interface {
	// Format specifies data format (data source type) for the underlying data, e.g. parquet.
	Format(source string) DataFrameReaderStream
	// Load reads the underlying data and returns a data frame.
	Load(path string) (DataFrame, error)
	// Table Reads a table from the underlying data source.
	Table(name string) (DataFrame, error)
	Option(key, value string) DataFrameReaderStream
}

// dataFrameReaderImpl is an implementation of DataFrameReader interface.
type dataFrameReaderStreamImpl struct {
	sparkSession *sparkSessionImpl
	formatSource string
	options      map[string]string
}

// newDataFrameReaderStream creates a new DataFrameReader
func newDataFrameReaderStream(session *sparkSessionImpl) DataFrameReaderStream {
	return &dataFrameReaderStreamImpl{
		sparkSession: session,
	}
}

func (w *dataFrameReaderStreamImpl) Format(source string) DataFrameReaderStream {
	w.formatSource = source
	return w
}

func (w *dataFrameReaderStreamImpl) Option(key, value string) DataFrameReaderStream {
	if w.options == nil {
		w.options = make(map[string]string)
	}
	w.options[key] = value
	return w
}

func (w *dataFrameReaderStreamImpl) Load(path string) (DataFrame, error) {
	var format string
	if w.formatSource != "" {
		format = w.formatSource
	}
	if w.options == nil {
		return NewDataFrame(w.sparkSession, newReadStreamWithFormatAndPath(path, format)), nil
	}
	return NewDataFrame(w.sparkSession, newReadStreamWithFormatAndPathAndOptions(path, format, w.options)), nil
}

func (w *dataFrameReaderStreamImpl) Table(name string) (DataFrame, error) {
	if w.options == nil {
		return NewDataFrame(w.sparkSession, newReadStreamTableRelation(name)), nil
	}
	return NewDataFrame(w.sparkSession, newReadStreamTableRelationAndOptions(name, w.options)), nil
}
