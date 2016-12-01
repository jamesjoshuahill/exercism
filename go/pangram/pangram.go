package pangram

import "strings"

const (
	testVersion = 1
	alphabet    = "abcdefghijklmnopqrstuvwxyz"
)

func IsPangram(input string) bool {
	lowercase := strings.ToLower(input)

	var count int
	for _, character := range alphabet {
		if strings.Contains(lowercase, string(character)) {
			count++
		}
	}
	return count >= 26
}
