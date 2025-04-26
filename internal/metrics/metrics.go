package metrics

type Metrics interface {
}

func New() Metrics {
	return &metrics{}
}

type metrics struct {
}
