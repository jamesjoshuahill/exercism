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
	t := Tonic(tonic)
	startingPitch := t.Pitch()
	scale := t.Scale()
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	s := make([]string, len(interval))
	var j int
	for i, n := range scale {
		if startingPitch == n {
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

// Tonic represents the starting note of a scale.
type Tonic string

// Pitch returns the tonic as a pitch.
func (t Tonic) Pitch() string {
	var p string
	for i, r := range t {
		if i == 0 {
			r = unicode.ToUpper(r)
		}
		p += string(r)
	}
	return p
}

// Scale returns the chromatic scale for the tonic.
func (t Tonic) Scale() []string {
	s := scaleWithSharps
	switch t {
	case "F", "Bb", "Eb", "Db", "d", "g", "bb":
		s = scaleWithFlats
	}
	return s
}
