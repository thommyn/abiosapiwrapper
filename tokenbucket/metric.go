package tokenbucket

type Metric interface {
	Diff(from int64, to int64) int64
}

type TimeMetric struct {}

func NewTimeMetric() Metric {
	return &TimeMetric{}
}

func (t TimeMetric) Diff(from int64, to int64) int64 {
	return to-from
}
