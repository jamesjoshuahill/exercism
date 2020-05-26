package encode

import (
	"strconv"
	"strings"
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
	var runLength []rune
	for _, r := range s {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			runLength = append(runLength, r)
		default:
			l := 1
			if len(runLength) > 0 {
				var err error
				l, err = strconv.Atoi(string(runLength))
				if err != nil {
					panic(err)
				}
				runLength = nil
			}
			for l > 0 {
				b.WriteRune(r)
				l--
			}
		}
	}
	return b.String()
}
