package clock

import "fmt"

// Clock represents a time without dates.
type Clock struct {
	hour, minute int
}

// New returns a new Clock at hour:minute.
func New(hour, minute int) Clock {
	hour += minute / 60
	minute %= 60
	if minute < 0 {
		minute += 60
		hour -= 1
	}

	hour %= 24
	if hour < 0 {
		hour += 24
	}

	return Clock{hour: hour, minute: minute}
}

// String returns the time in the format "hh:mm".
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add returns a clock at time + minutes.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract returns a clock at time - minutes.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
