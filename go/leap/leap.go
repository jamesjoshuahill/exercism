package leap

import (
	"math"
)

func IsLeapYear(year int) bool {
	fourthYear := IsMultipleOf(4, year)
	century := IsMultipleOf(100, year)
	quadricentennial := IsMultipleOf(400, year)

	return fourthYear && (quadricentennial || !century)
}

func IsMultipleOf(y, x int) bool {
	return math.Remainder(float64(x), float64(y)) == 0
}
