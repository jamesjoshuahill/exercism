package leap

func IsLeapYear(n int) bool {
	year := Year(n)
	return year.isLeapYear()
}

type Year int

func (year Year) isMultipleOf(divisor int) bool {
	return int(year)%divisor == 0
}

func (year Year) isFourthYear() bool {
	return year.isMultipleOf(4)
}

func (year Year) isCentury() bool {
	return year.isMultipleOf(100)
}

func (year Year) isQuadricentennial() bool {
	return year.isMultipleOf(400)
}

func (year Year) isLeapYear() bool {
	return year.isQuadricentennial() || (year.isFourthYear() && !year.isCentury())
}
