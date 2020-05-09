package isogram

import "unicode"

// IsIsogram returns true if s is an isogram.
func IsIsogram(s string) bool {
	tally := map[rune]bool{}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}

		c := unicode.ToLower(r)
		if tally[c] {
			return false
		}
		tally[c] = true
	}

	return true
}
