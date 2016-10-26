package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 1

func abbreviate(phrase string) string {
	words := regexp.MustCompile(`[A-Z]+[a-z]*|[a-z]+`)
	matches := words.FindAllString(phrase, -1)
	var initials string
	for _, match := range matches {
		initials += string(match[0])
	}
	return strings.ToUpper(initials)
}
