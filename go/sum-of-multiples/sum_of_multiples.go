package summultiples

// SumMultiples returns the sum of all the unique multiples of the specified
// divisors up to but not including limit.
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	multiples := map[int]struct{}{}
	for _, d := range divisors {
		if d == 0 {
			continue
		}

		for i := 0; i < limit; i += d {
			if _, ok := multiples[i]; !ok {
				multiples[i] = struct{}{}
				sum += i
			}
		}
	}
	return sum
}
