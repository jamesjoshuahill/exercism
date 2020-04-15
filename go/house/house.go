package house

import "strings"

var phrases = []string{
	"This is",
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
	"the house that Jack built.",
}

func Verse(i int) string {
	verse := []string{phrases[0]}
	verse = append(verse, phrases[len(phrases)-(i):]...)
	return strings.Join(verse, " ")
}

func Song() string {
	var song []string
	for i := 1; i <= 12; i++   {
		song = append(song, Verse(i))
	}
	return strings.Join(song, "\n\n")
}
