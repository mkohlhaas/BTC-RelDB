package protocol

import "time"

type timestamp int64

func (ts timestamp) String() string {
	t := time.Unix(int64(ts), 0)
	return t.Format("02-Jan-06 15:04 MST")
}

func now() timestamp {
	return timestamp(time.Now().Unix())
}
