package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	hours   int
	minutes int
}

func Time(hours, minutes int) Clock {
	if hours < 0 {
		return Time(hours+24, minutes)
	}
	if hours >= 24 {
		return Time(hours-24, minutes)
	}
	if minutes < 0 {
		return Time(hours-1, minutes+60)
	}
	if minutes >= 60 {
		return Time(hours+1, minutes-60)
	}
	return Clock{hours, minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func (c Clock) Add(minutes int) Clock {
	return Time(c.hours, c.minutes+minutes)
}
