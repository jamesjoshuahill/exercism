package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode compresses a string using run-length encoding.
func RunLengthEncode(s string) string {
	b := &strings.Builder{}
	var prev rune
	var count int
	for _, r := range s {
		if r == prev {
			count++
		} else {
			writeRune(b, count, prev)
			prev = r
			count = 1
		}
	}
	writeRune(b, count, prev)
	return b.String()
}

func writeRune(b *strings.Builder, count int, r rune) {
	if count > 1 {
		b.WriteString(strconv.Itoa(count))
	}
	if count > 0 {
		b.WriteRune(r)
	}
}

// RunLengthDecode decompresses a run-length encoded string.
func RunLengthDecode(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		j := i
		for unicode.IsNumber(rune(s[j])) {
			j++
		}

		l := 1
		if j > i {
			var err error
			l, err = strconv.Atoi(s[i:j])
			if err != nil {
				panic(err)
			}
		}
		for l > 0 {
			b.WriteRune(rune(s[j]))
			l--
		}

		i = j + 1
	}
	return b.String()
}
