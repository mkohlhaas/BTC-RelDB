package protocol

import (
	"time"
)

type Timestamp int64

func (t Timestamp) String() string {
	t1 := time.Unix(int64(t), 0)
	return t1.Format(time.UnixDate)
}

func Now() Timestamp {
	return Timestamp(time.Now().Unix())
}
