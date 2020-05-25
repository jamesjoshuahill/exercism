package accumulate

// Accumulate applies function f to each element of c and returns a new slice
// containing the results.
func Accumulate(c []string, f func(string) string) []string {
	a := make([]string, len(c))
	for i, s := range c {
		a[i] = f(s)
	}
	return a
}
