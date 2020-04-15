package foodchain

import (
	"strings"
)

const (
	firstLine    = "I know an old lady who swallowed a "
	lastLine     = "I don't know why she swallowed the fly. Perhaps she'll die."
	spider       = "spider"
	spiderPhrase = "wriggled and jiggled and tickled inside her"
)

var parts = map[int]struct {
	animal, comment string
}{
	1: {animal: "fly"},
	2: {animal: spider, comment: "It " + spiderPhrase + "."},
	3: {animal: "bird", comment: "How absurd to swallow a bird!"},
	4: {animal: "cat", comment: "Imagine that, to swallow a cat!"},
	5: {animal: "dog", comment: "What a hog, to swallow a dog!"},
	6: {animal: "goat", comment: "Just opened her throat and swallowed a goat!"},
	7: {animal: "cow", comment: "I don't know how she swallowed a cow!"},
	8: {animal: "horse", comment: "She's dead, of course!"},
}

func Song() string {
	return Verses(1, 8)
}

func Verses(first, last int) string {
	b := strings.Builder{}
	for i := first; i <= last; i++ {
		verse(&b, i)
		if i != last {
			b.WriteString("\n\n")
		}
	}
	return b.String()
}

func Verse(i int) string {
	b := strings.Builder{}
	verse(&b, i)
	return b.String()
}

func verse(b *strings.Builder, i int) {
	firstPart(b, i)
	if i < 8 {
		repeatedParts(b, i)
	}
}

func firstPart(b *strings.Builder, i int) {
	b.WriteString(firstLine)
	b.WriteString(parts[i].animal)
	b.WriteString(".")
	if i > 1 {
		b.WriteString("\n")
		b.WriteString(parts[i].comment)
	}
}

func repeatedParts(b *strings.Builder, i int) {
	if i == 1 {
		b.WriteString("\n")
		b.WriteString(lastLine)
		return
	}

	swallowedLine(b, parts[i].animal, parts[i-1].animal)
	repeatedParts(b, i-1)
}

func swallowedLine(b *strings.Builder, animal, next string) {
	b.WriteString("\nShe swallowed the ")
	b.WriteString(animal)
	b.WriteString(" to catch the ")
	b.WriteString(next)
	if next == spider {
		b.WriteString(" that ")
		b.WriteString(spiderPhrase)
	}
	b.WriteString(".")
}
