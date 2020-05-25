// Package proverb generates proverbs.
package proverb

import "fmt"

// Proverb generates lines of the proverb from a list of rhymes.
func Proverb(rhyme []string) []string {
	p := make([]string, len(rhyme))
	for i := range rhyme {
		if i == len(rhyme)-1 {
			p[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			p[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
	}
	return p
}
