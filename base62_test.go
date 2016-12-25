package base62

import (
	"testing"
)

var testCases = []struct {
	n   uint64
	b62 string
}{
	{0, "0"},
	{10, "a"},
	{630, "aa"},
	{1097900471, "1ciG47"},
}

func TestEncode(t *testing.T) {
	for _, tc := range testCases {
		t.Logf("Testing %d\n", tc.n)
		if v := Encode(tc.n); v != tc.b62 {
			t.Fatalf("%d encoded to %s (should be %s)", tc.n, v, tc.b62)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range testCases {
		t.Logf("Testing %s\n", tc.b62)
		if n := Decode(tc.b62); n != tc.n {
			t.Fatalf("%s decoded to %d (should be %d)", tc.b62, n, tc.n)
		}
	}
}
