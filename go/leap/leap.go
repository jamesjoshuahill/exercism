package leap

type Year int

func (year Year) isMultipleOf(divisor int) bool {
	return int(year)%divisor == 0
}

func IsLeapYear(n int) bool {
	year := Year(n)
	fourthYear := year.isMultipleOf(4)
	century := year.isMultipleOf(100)
	quadricentennial := year.isMultipleOf(400)

	return fourthYear && (quadricentennial || !century)
}
