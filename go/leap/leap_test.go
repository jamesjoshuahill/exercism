package leap

import (
	"testing"
)

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{1996, true, "leap year"},
	{1997, false, "non-leap year"},
	{1998, false, "non-leap even year"},
	{1900, false, "century"},
	{2400, true, "fourth century"},
	{2000, true, "Y2K"},
}

func TestLeapYears(t *testing.T) {
	for _, test := range testCases {
		observed := IsLeapYear(test.year)
		if observed != test.expected {
			t.Fatalf("%v is %s", test.year, test.description)
		}
	}
}

func BenchmarkLeapYears(b *testing.B) {
	b.StopTimer()
	for _, test := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			IsLeapYear(test.year)
		}

		b.StopTimer()
	}
}
