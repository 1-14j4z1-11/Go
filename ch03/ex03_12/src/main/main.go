package main

import (
	"fmt"
	"sort"
)

type sortableRunes []rune

func (runes sortableRunes) Len() int {
	return len(runes)
}

func (runes sortableRunes) Less(i, j int) bool {
	return runes[i] < runes[j]
}

func (runes sortableRunes) Swap(i, j int) {
	runes[i], runes[j] = runes[j], runes[i]
}

////////////////////////////////////////////////////

func main() {
	anagramTest("ABCD", "CDAB")
	anagramTest("ABCD", "ABCE")
	anagramTest("", "")
	anagramTest("A", "A")
	anagramTest("あいうえお", "おえいうあ")
}

func anagramTest(s1, s2 string) {
	if(areAnagram(s1, s2)) {
		fmt.Printf("'%s' '%s' are anagram\n", s1, s2)
	} else {
		fmt.Printf("'%s' '%s' are not anagram\n", s1, s2)
	}
}

func areAnagram(s1, s2 string) bool {
	r1 := sortableRunes(s1)
	r2 := sortableRunes(s2)

	if r1.Len() != r2.Len() {
		return false
	}

	sort.Sort(r1)
	sort.Sort(r2)

	for i := 0; i < r1.Len(); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}
