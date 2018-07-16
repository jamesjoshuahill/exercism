// Package twofer generates two-fers
package twofer

import "fmt"

// ShareWith returns a two-fer with 'you',
// or name if not empty string.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
