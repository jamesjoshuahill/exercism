package leap

func IsLeapYear(n int) bool {
	year := Year(n)
	return year.isLeapYear()
}

type Year int

func (year Year) isFourthYear() bool {
	return int(year)%4 == 0
}

func (year Year) isCentury() bool {
	return int(year)%100 == 0
}

func (year Year) isQuadricentennial() bool {
	return int(year)%400 == 0
}

func (year Year) isLeapYear() bool {
	return year.isFourthYear() && (!year.isCentury() || year.isQuadricentennial())
}
