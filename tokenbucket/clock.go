package tokenbucket

import "time"

type Timer interface {
	Now() int64
}

type Clock struct {}

func NewClock() Timer {
	return &Clock{}
}

func (t Clock) Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
