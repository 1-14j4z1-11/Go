package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}

func testCase(s sort.Interface) {
	if IsPalindrome(s) {
		fmt.Printf("[%v] is palindrome\n", s)
	} else {
		fmt.Printf("[%v] is not palindrome\n", s)
	}
}

//////////////////////////////////////////////////

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

func (runes sortableRunes) String() string {
	return string(runes)
}

//////////////////////////////////////////////////

func main() {
	testCase(sortableRunes("TEST"))
	testCase(sortableRunes("TesttseT"))
	testCase(sortableRunes(""))
}
