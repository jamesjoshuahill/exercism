package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last

	daysPerWeek = 7
)

func Day(s WeekSchedule, wd time.Weekday, m time.Month, y int) int {
	var start int
	switch s {
	case First, Second, Third, Fourth:
		start = int(s)*daysPerWeek + 1
	case Teenth:
		start = 13
	case Last:
		last := date(y, m+1, 0) // day _before_ the first day of next month
		start = last.Day() - daysPerWeek + 1
	}

	return findWeekday(start, y, m, wd)
}

func findWeekday(start, y int, m time.Month, wd time.Weekday) int {
	for d := start; d < start+daysPerWeek; d++ {
		date := date(y, m, d)
		if date.Weekday() == wd {
			return d
		}
	}
	return -1
}

func date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
