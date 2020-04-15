package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("n must be greater than zero")
	}
	if n == 1 {
		return 0, nil
	}

	var next int
	if n%2 == 0 {
		next = n / 2
	} else {
		next = 3*n + 1
	}

	steps, err := CollatzConjecture(next)
	if err != nil {
		return -1, err
	}

	return 1 + steps, nil
}
