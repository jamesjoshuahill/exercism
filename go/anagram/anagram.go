package anagram

import (
	"strings"
	"unicode"
)

// Detect returns a list of candidates that are anagrams of subject.
func Detect(subject string, candidates []string) []string {
	p := product(subject)

	var anagrams []string
	for _, c := range candidates {
		if len(subject) != len(c) || strings.EqualFold(subject, c) {
			continue
		}
		if p == product(c) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func product(s string) int {
	p := 1
	for _, r := range s {
		p *= prime(unicode.ToLower(r))
	}
	return p
}

func prime(r rune) int {
	switch r {
	case 'a':
		return 2
	case 'b':
		return 3
	case 'c':
		return 5
	case 'd':
		return 7
	case 'e':
		return 11
	case 'f':
		return 13
	case 'g':
		return 17
	case 'h':
		return 19
	case 'i':
		return 23
	case 'j':
		return 29
	case 'k':
		return 31
	case 'l':
		return 37
	case 'm':
		return 41
	case 'n':
		return 43
	case 'o':
		return 47
	case 'p':
		return 53
	case 'q':
		return 59
	case 'r':
		return 61
	case 's':
		return 67
	case 't':
		return 71
	case 'u':
		return 73
	case 'v':
		return 79
	case 'w':
		return 83
	case 'x':
		return 89
	case 'y':
		return 97
	case 'z':
		return 101
	}
	return 0
}
