package bob

import (
	"regexp"
	"strings"
)

const testVersion = 2

func Hey(phrase string) string {
	greeting := Greeting{phrase: phrase}

	if greeting.IsShout() {
		return "Whoa, chill out!"
	}
	if greeting.IsQuestion() {
		return "Sure."
	}
	if greeting.IsSilence() {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

type Greeting struct {
	phrase string
}

func (g Greeting) IsShout() bool {
	return g.containsUppercase() && g.allUppercase()
}

func (g Greeting) IsQuestion() bool {
	endsWithQuestionMark := `\?\s*\z`
	return g.matches(endsWithQuestionMark)
}

func (g Greeting) IsSilence() bool {
	allNonWordCharacters := `\A\W*\z`
	return g.matches(allNonWordCharacters)
}

func (g Greeting) containsUppercase() bool {
	return g.matches(`[A-Z]`)
}

func (g Greeting) allUppercase() bool {
	return g.phrase == strings.ToUpper(g.phrase)
}

func (g Greeting) matches(pattern string) bool {
	matcher := regexp.MustCompile(pattern)
	return matcher.Match([]byte(g.phrase))
}
