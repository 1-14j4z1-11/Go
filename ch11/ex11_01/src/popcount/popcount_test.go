package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input	uint64
		want	int
	} {
		{ 0, 0 },
		{ 0xFFFFFFFFFFFFFFFF, 64 },
		{ 1, 1 },
		{ 1 << 63, 1 },
		{ 1023, 10 },
		{ 0xAAAAAAAA, 16 },
	}

	for _, test := range tests {
		if got := PopCount(test.input); got != test.want {
			t.Errorf("PopCount(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}
