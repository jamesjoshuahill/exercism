package anagram

import (
	"sort"
	"unicode"
	"unicode/utf8"
)

// Detect returns a list of candidates that are anagrams of subject.
func Detect(subject string, candidates []string) []string {
	lowercaseSubject := lowercaseRunes(subject)
	sortedSubject := lowercaseRunes(subject)
	sortRunes(sortedSubject)

	var anagrams []string
	for _, c := range candidates {
		if len(subject) != len(c) {
			continue
		}
		r := lowercaseRunes(c)
		if isEqual(lowercaseSubject, r) {
			continue
		}
		sortRunes(r)
		if isEqual(sortedSubject, r) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func lowercaseRunes(s string) []rune {
	runes := make([]rune, utf8.RuneCountInString(s))
	for i, r := range s {
		runes[i] = unicode.ToLower(r)
	}
	return runes
}

func sortRunes(runes []rune) {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
}

func isEqual(a, b []rune) bool {
	for i, r := range b {
		if a[i] != r {
			return false
		}
	}
	return true
}
