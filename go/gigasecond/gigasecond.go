package gigasecond

import "time"

const (
	testVersion = 4
	billion     = 1e9
)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(billion * time.Second)
}
