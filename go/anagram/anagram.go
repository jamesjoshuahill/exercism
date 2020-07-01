package anagram

import (
	"strings"
	"unicode"
)

// Detect returns a list of candidates that are anagrams of subject.
func Detect(subject string, candidates []string) []string {
	count := countRunes(subject)

	var anagrams []string
	for _, c := range candidates {
		if len(subject) != len(c) || strings.EqualFold(subject, c) {
			continue
		}
		if count == countRunes(c) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func countRunes(s string) [26]int {
	var count [26]int
	for _, r := range s {
		count[unicode.ToLower(r)-'a']++
	}
	return count
}
