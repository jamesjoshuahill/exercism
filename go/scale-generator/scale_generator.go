package scale

import (
	"unicode"
)

var (
	scaleWithSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	scaleWithFlats  = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
)

// Scale returns a scale that starts at tonic and follows the interval pattern.
// Supported intervals are 'm': half step, 'M': whole step, 'A': augmented first.
// If no interval pattern is given the default is 12 half steps.
func Scale(tonic, interval string) []string {
	startingNote := convertToNote(tonic)
	scale := chromaticScale(tonic)
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	s := make([]string, len(interval))
	var j int
	for i, n := range scale {
		if startingNote == n {
			j = i
		}
	}
	for i, step := range interval {
		s[i] = scale[j%12]

		switch step {
		case 'm':
			j++
		case 'M':
			j += 2
		case 'A':
			j += 3
		}
	}
	return s
}

func convertToNote(tonic string) string {
	var note string
	for i, r := range tonic {
		switch i {
		case 0:
			note += string(unicode.ToUpper(r))
		case 1:
			note += string(r)
		}
	}
	return note
}

func chromaticScale(tonic string) []string {
	scale := scaleWithSharps
	switch tonic {
	case "F", "Bb", "Eb", "Db", "d", "g", "bb":
		scale = scaleWithFlats
	}
	return scale
}
