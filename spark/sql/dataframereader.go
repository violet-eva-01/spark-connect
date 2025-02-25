package sql

// DataFrameReader supports reading data from storage and returning a data frame.
// TODO needs to implement other methods like Option(), Schema(), and also "strong typed"
// reading (e.g. Parquet(), Orc(), Csv(), etc.
type DataFrameReader interface {
	// Format specifies data format (data source type) for the underlying data, e.g. parquet.
	Format(source string) DataFrameReader
	// Load reads the underlying data and returns a data frame.
	Load() (DataFrame, error)
	LoadStream() (DataFrame, error)
	// Table Reads a table from the underlying data source.
	Table(name string) (DataFrame, error)
	TableStream(name string) (DataFrame, error)
	Option(key, value string) DataFrameReader
	Path(path ...string) DataFrameReader
}

// dataFrameReaderImpl is an implementation of DataFrameReader interface.
type dataFrameReaderImpl struct {
	sparkSession *sparkSessionImpl
	formatSource string
	path         []string
	options      map[string]string
}

// NewDataframeReader creates a new DataFrameReader
func NewDataframeReader(session *sparkSessionImpl) DataFrameReader {
	return &dataFrameReaderImpl{
		sparkSession: session,
	}
}

func (w *dataFrameReaderImpl) Table(name string) (DataFrame, error) {
	if w.options == nil {
		return NewDataFrame(w.sparkSession, newReadTableRelation(name)), nil
	}
	return NewDataFrame(w.sparkSession, newReadTableRelationAndOptions(name, w.options)), nil
}

func (w *dataFrameReaderImpl) Format(source string) DataFrameReader {
	w.formatSource = source
	return w
}

func (w *dataFrameReaderImpl) Path(path ...string) DataFrameReader {
	if len(w.path) > 0 {
		w.path = append(w.path, path...)
	} else {
		w.path = path
	}
	return w
}

func (w *dataFrameReaderImpl) Load() (DataFrame, error) {
	var format string
	if w.formatSource != "" {
		format = w.formatSource
	}
	if len(w.path) > 0 {
		if w.options == nil {
			return NewDataFrame(w.sparkSession, newReadWithFormatAndPath(w.path, format)), nil
		} else {
			return NewDataFrame(w.sparkSession, newReadWithFormatAndPathAndOptions(w.path, format, w.options)), nil
		}
	} else {
		if w.options == nil {
			return NewDataFrame(w.sparkSession, newReadWithFormat(format)), nil
		} else {
			return NewDataFrame(w.sparkSession, newReadWithFormatAndOptions(format, w.options)), nil
		}
	}
}

func (w *dataFrameReaderImpl) Option(key, value string) DataFrameReader {
	if w.options == nil {
		w.options = make(map[string]string)
	}
	w.options[key] = value
	return w
}

func (w *dataFrameReaderImpl) LoadStream() (DataFrame, error) {
	var format string
	if w.formatSource != "" {
		format = w.formatSource
	}
	if len(w.path) > 0 {
		if w.options == nil {
			return NewDataFrame(w.sparkSession, newReadStreamWithFormatAndPath(w.path, format)), nil
		} else {
			return NewDataFrame(w.sparkSession, newReadStreamWithFormatAndPathAndOptions(w.path, format, w.options)), nil
		}
	} else {
		if w.options == nil {
			return NewDataFrame(w.sparkSession, newReadStreamWithFormat(format)), nil
		} else {
			return NewDataFrame(w.sparkSession, newReadStreamWithFormatAndOptions(format, w.options)), nil
		}
	}
}

func (w *dataFrameReaderImpl) TableStream(name string) (DataFrame, error) {
	if w.options == nil {
		return NewDataFrame(w.sparkSession, newReadStreamTableRelation(name)), nil
	}
	return NewDataFrame(w.sparkSession, newReadStreamTableRelationAndOptions(name, w.options)), nil
}
