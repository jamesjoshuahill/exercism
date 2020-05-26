package encode

import "testing"

func TestRunLengthEncode(t *testing.T) {
	for _, test := range encodeTests {
		t.Run(test.description, func(t *testing.T) {
			if actual := RunLengthEncode(test.input); actual != test.expected {
				t.Fatalf("FAIL - RunLengthEncode(%s) = %q, expected %q.",
					test.input, actual, test.expected)
			}
		})
	}
}

func TestRunLengthDecode(t *testing.T) {
	for _, test := range decodeTests {
		t.Run(test.description, func(t *testing.T) {
			if actual := RunLengthDecode(test.input); actual != test.expected {
				t.Fatalf("FAIL - RunLengthDecode(%s) = %q, expected %q.",
					test.input, actual, test.expected)
			}
		})
	}
}

func TestRunLengthEncodeDecode(t *testing.T) {
	for _, test := range encodeDecodeTests {
		t.Run(test.description, func(t *testing.T) {
			if actual := RunLengthDecode(RunLengthEncode(test.input)); actual != test.expected {
				t.Fatalf("FAIL - RunLengthDecode(RunLengthEncode(%s)) = %q, expected %q.",
					test.input, actual, test.expected)
			}
		})
	}
}
