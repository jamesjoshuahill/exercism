package triangle

import "math"

const testVersion = 3

// Code this function.
func KindFromSides(a, b, c float64) Kind {
	if invalidSide(a) || invalidSide(b) || invalidSide(c) {
		return NaT
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func invalidSide(side float64) bool {
	if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
		return true
	}
	return false
}

type Kind string

const (
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)
