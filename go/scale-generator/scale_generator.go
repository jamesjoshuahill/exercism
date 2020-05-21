package scale

import (
	"strings"
)

var (
	scaleWithSharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	scaleWithFlats  = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
)

// Scale returns a scale that starts at tonic and follows the interval pattern.
// Supported intervals are 'm': half step, 'M': whole step, 'A': augmented first.
// If no interval pattern is given the default is 12 half steps.
func Scale(tonic, interval string) []string {
	t := Tonic(tonic)
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	s := make([]string, len(interval))
	c := t.ChromaticScale()
	var j int
	for i, step := range interval {
		s[i] = c[j]

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

// Tonic represents the starting note of a scale.
type Tonic string

// Pitch returns the tonic as a pitch.
func (t Tonic) Pitch() string {
	return strings.ToUpper(string(t[:1])) + string(t[1:])
}

// ChromaticScale returns the chromatic scale starting at tonic.
func (t Tonic) ChromaticScale() []string {
	s := scaleWithSharps
	switch t {
	case "F", "Bb", "Eb", "Db", "d", "g", "bb":
		s = scaleWithFlats
	}

	var start int
	for i, p := range s {
		if t.Pitch() == p {
			start = i
		}
	}

	return append(s[start:], s[:start]...)
}
