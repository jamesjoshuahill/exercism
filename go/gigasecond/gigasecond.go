package gigasecond

import "time"

const testVersion = 4

func AddGigasecond(t time.Time) time.Time {
	return t.AddDate(31, 0, 251).
		Add(time.Hour * 1).
		Add(time.Minute * 46).
		Add(time.Second * 40).
		UTC()
}
