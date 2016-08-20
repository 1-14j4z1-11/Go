package charcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input	string
		want	map[rune]int
	} {
		{ "", map[rune]int{} },
		{ " ", map[rune]int{ ' ':1 } },
		{ "\n\n\n\r\r\t\t\t\b\b", map[rune]int{ '\n':3, '\r':2, '\t':3, '\b':2 } },
		{ "ABCDABCABA", map[rune]int{ 'A':4, 'B':3, 'C':2, 'D':1 } },
		{ "あいうえお\nあいう\nあかさ", map[rune]int{ 'あ':3, 'い':2, 'う':2, 'え':1, 'お':1, 'か':1, 'さ':1, '\n':2 } },
	}

	for _, test := range tests {
		got := CharCount(test.input)
		want := test.want

		for r, n := range got {
			if want[r] != n {
				t.Errorf("count %c = %d of CharCount(%s), want %d", r, n, got, want[r])
			}
		}
		for r, n := range want {
			if got[r] != n {
				t.Errorf("count %c = %d of CharCount(%s), want %d", r, n, got, want[r])
			}
		}
	}
}
