package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("n must be greater than zero")
	}

	var count int
	for {
		if n == 1 {
			return count, nil
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
	}
}
