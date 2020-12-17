package utils

import "time"

type Timer interface {
	Now() time.Time
}

func NewRealTimer() Timer {
	return new(RealTimer)
}
type RealTimer struct {

}

func (t RealTimer) Now() time.Time {
	return time.Now()
}