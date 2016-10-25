package hamming

import (
	"errors"
	"strings"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must have equal length to calculate hamming distance")
	}

	var distance int
	aReader := strings.NewReader(a)
	bReader := strings.NewReader(b)

	for index := 0; index < len(a); index++ {
		nextA, _ := aReader.ReadByte()
		nextB, _ := bReader.ReadByte()
		if nextA != nextB {
			distance++
		}
	}

	return distance, nil
}
