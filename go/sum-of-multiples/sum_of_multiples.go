package summultiples

// SumMultiples returns the sum of all the unique multiples of the specified
// divisors up to but not including limit.
func SumMultiples(limit int, divisors ...int) int {
	multiples := map[int]bool{}
	for _, d := range divisors {
		if d == 0 {
			continue
		}

		for i := 0; i < limit; i += d {
			if i < limit {
				multiples[i] = true
			}
		}
	}

	var sum int
	for m := range multiples {
		sum += m
	}
	return sum
}
