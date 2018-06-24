package tokenbucket

import "time"

type Timer interface {
	Now() int64
}

type clock struct {}

func NewClock() Timer {
	return &clock{}
}

func (t clock) Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
