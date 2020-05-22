package etl

import (
	"strings"
)

// Transform transforms a list of uppercase letters per score into a score per
// letter.
func Transform(input map[int][]string) map[string]int {
	o := map[string]int{}
	for s, letters := range input {
		for _, l := range letters {
			o[strings.ToLower(l)] = s
		}
	}
	return o
}
