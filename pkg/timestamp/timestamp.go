package timestamp

import "time"

func Now() int64 {
	return time.Now().UnixMicro()
}

func HourBeforeNow(h int) int64 {
	return time.Now().Add(time.Duration(-1*h) * time.Hour).UnixMicro()
}

func SecondBeforeNow(s int) int64 {
	return time.Now().Add(time.Duration(-1*s) * time.Second).UnixMicro()
}
