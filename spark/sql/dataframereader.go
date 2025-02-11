package sql

// DataFrameReader supports reading data from storage and returning a data frame.
// TODO needs to implement other methods like Option(), Schema(), and also "strong typed"
// reading (e.g. Parquet(), Orc(), Csv(), etc.
type DataFrameReader interface {
	// Format specifies data format (data source type) for the underlying data, e.g. parquet.
	Format(source string) DataFrameReader
	// Load reads the underlying data and returns a data frame.
	Load(path string) (DataFrame, error)
	// Reads a table from the underlying data source.
	Table(name string) (DataFrame, error)
	Option(key, value string) DataFrameReader
}

// dataFrameReaderImpl is an implementation of DataFrameReader interface.
type dataFrameReaderImpl struct {
	sparkSession *sparkSessionImpl
	formatSource string
	options      map[string]string
}

// NewDataframeReader creates a new DataFrameReader
func NewDataframeReader(session *sparkSessionImpl) DataFrameReader {
	return &dataFrameReaderImpl{
		sparkSession: session,
	}
}

func (w *dataFrameReaderImpl) Table(name string) (DataFrame, error) {
	return NewDataFrame(w.sparkSession, newReadTableRelation(name)), nil
}

func (w *dataFrameReaderImpl) Format(source string) DataFrameReader {
	w.formatSource = source
	return w
}

func (w *dataFrameReaderImpl) Load(path string) (DataFrame, error) {
	var format string
	if w.formatSource != "" {
		format = w.formatSource
	}
	if w.options == nil {
		return NewDataFrame(w.sparkSession, newReadWithFormatAndPath(path, format)), nil
	}
	return NewDataFrame(w.sparkSession, newReadWithFormatAndPathAndOptions(path, format, w.options)), nil
}

func (w *dataFrameReaderImpl) Option(key, value string) DataFrameReader {
	if w.options == nil {
		w.options = make(map[string]string)
	}
	w.options[key] = value
	return w
}
