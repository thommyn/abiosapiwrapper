package tokenbucket

type Metric interface {
	Diff(from int64, to int64) int64
}

type timeMetric struct {}

func NewTimeMetric() Metric {
	return &timeMetric{}
}

func (t timeMetric) Diff(from int64, to int64) int64 {
	return to-from
}
