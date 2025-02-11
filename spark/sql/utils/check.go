package utils

func WarnOnError(f func() error, h func(e error)) {
	if err := f(); err != nil {
		h(err)
	}
}
