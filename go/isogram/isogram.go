package isogram

import "unicode"

// IsIsogram returns true if s is an isogram.
func IsIsogram(s string) bool {
	tally := map[rune]struct{}{}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}

		c := unicode.ToLower(r)
		if _, ok := tally[c]; ok {
			return false
		}
		tally[c] = struct{}{}
	}

	return true
}
