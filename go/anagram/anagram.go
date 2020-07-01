package anagram

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Detect returns a list of candidates that are anagrams of subject.
func Detect(subject string, candidates []string) []string {
	sorted := sortRunes(subject)

	var anagrams []string
	for _, c := range candidates {
		if len(subject) != len(c) || strings.EqualFold(subject, c) {
			continue
		}
		r := sortRunes(c)
		if isEqual(sorted, r) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func sortRunes(s string) []rune {
	runes := make([]rune, utf8.RuneCountInString(s))
	for i, r := range s {
		runes[i] = unicode.ToLower(r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}

func isEqual(a, b []rune) bool {
	for i, r := range b {
		if a[i] != r {
			return false
		}
	}
	return true
}
