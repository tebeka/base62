package base62

import (
	"testing"
)

type testCase struct {
	in  uint64
	out string
}

var testCases = []testCase{
	{0, "0"},
	{1, "b"},
	{22, "w"},
	{4444, "bjG"},
	{7777777, "6Dwb"},
}

func TestEncode(t *testing.T) {
	for _, tc := range testCases {
		t.Logf("Testing %d\n", tc.in)
		if v := Encode(tc.in); v != tc.out {
			t.Fatalf("%d decoded to %s (should be %s)", tc.in, v, tc.out)
		}
	}
}
