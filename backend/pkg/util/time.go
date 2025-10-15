package util

import (
	"time"

	"github.com/lib/pq"
)

func UnmarshalTime(t int64) pq.NullTime {
	return pq.NullTime{
		Valid: t != 0,
		Time:  time.Unix(t, 0),
	}
}
