package charcount

import (
	"unicode"
	"unicode/utf8"
)

func CharCount(text string) map[rune]int {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	for _, r := range text {
		n := utf8.RuneLen(r)
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return counts
}
