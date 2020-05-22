package reverse

import (
	"strings"
	"unicode/utf8"
)

// Reverse returns the runes in s in reverse order.
func Reverse(s string) string {
	var b strings.Builder
	b.Grow(utf8.RuneCountInString(s))
	for i := len(s); i > 0; {
		r, size := utf8.DecodeLastRuneInString(s[:i])
		b.WriteRune(r)
		i -= size
	}
	return b.String()
}
