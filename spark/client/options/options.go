package options

type SparkClientOptions struct {
	ReattachExecution bool
	UserAgent         string
	UserId            string
}

var DefaultSparkClientOptions = SparkClientOptions{
	ReattachExecution: false,
}

func NewSparkClientOptions(reattach bool) SparkClientOptions {
	return SparkClientOptions{
		ReattachExecution: reattach,
	}
}
