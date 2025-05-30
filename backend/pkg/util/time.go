package util

import (
	"time"

	"github.com/lib/pq"
)

func MarshalTime(t pq.NullTime) int64 {
	if !t.Valid {
		return 0
	}
	return t.Time.Unix()
}

func UnmarshalTime(t int64) pq.NullTime {
	return pq.NullTime{
		Valid: t != 0,
		Time:  time.Unix(t, 0),
	}
}
