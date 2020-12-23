package utils

import "time"

/* Time interface */
type TimeInterface interface {
	FromUnix(int64) time.Time
	Now() time.Time
	NowUnix() int64
}

/* Implementation of Time Interface */
type timeImplementation struct{}

func (*timeImplementation) FromUnix(i int64) time.Time {
	return time.Unix(i, 0)
}

func (*timeImplementation) Now() time.Time {
	return time.Now().UTC()
}

func (t *timeImplementation) NowUnix() int64 {
	return t.Now().Unix()
}
