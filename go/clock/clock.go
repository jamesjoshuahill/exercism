package clock

import (
  "fmt"
)

const TestVersion = 2

type Clock struct {
  hours int
  minutes int
}

func (clock Clock) String() string {
  return fmt.Sprintf("%02v:%02v", clock.hours, clock.minutes)
}

func (clock Clock) Add(minutes int) Clock {
  return clock
}

func Time(hours, minutes int) Clock {
  hours += minutes / 60
  minutes = minutes % 60
  hours = hours % 24
  if minutes < 0 {
    minutes += 60
    hours -= 1
  }
  if hours < 0 {
    hours += 24
  }
  return Clock{hours, minutes}
}
