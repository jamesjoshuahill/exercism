package piglatin

import "strings"

const suffix = "ay"

func Sentence(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		words[i] = word(w)
	}
	return strings.Join(words, " ")
}

func word(s string) string {
	switch s[0] {
	case 'a', 'e', 'i', 'o', 'u':
		return s + suffix
	}

	if len(s) >= 3 {
		switch s[:3] {
		case "sch", "thr", "squ":
			return s[3:] + s[:3] + suffix
		}
	}

	switch s[:2] {
	case "yt", "xr":
		return s + suffix
	case "ch", "rh", "sq", "th", "qu":
		return s[2:] + s[:2] + suffix
	}

	return s[1:] + s[:1] + suffix
}
