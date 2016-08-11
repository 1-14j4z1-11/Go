package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	testcases := []struct {
		str string
		sep string
		exp []string
	}{
		{ "", "", []string{ } },
		{ "A", "", []string{ "A" } },
		{ "A", ",", []string{ "A" } },
		{ "A,B,C,D", ":", []string{ "A,B,C,D" } },
		{ "A,B,C,D", ",", []string{ "A", "B", "C", "D" } },
		{ "aaa,bbb,ccc", ",", []string{ "aaa", "bbb", "ccc" } },
		{ "aaa,bbb:ccc,ddd:", ",", []string{ "aaa", "bbb:ccc", "ddd:" } },
	}

	for _, tc := range testcases {
		words := strings.Split(tc.str, tc.sep)

		if len(words) != len(tc.exp) {
			t.Errorf("Split(\"%s\", \"%s\") returns %d words, want %d", tc.str, tc.sep, len(words), len(tc.exp))
		}

		for i := 0; i < len(words); i++ {
			if words[i] != tc.exp[i] {
				t.Errorf("Split(\"%s\", \"%s\")[%d] = %s, want %s", tc.str, tc.sep, i, words[i], tc.exp[i])
			}
		}
	}

}
