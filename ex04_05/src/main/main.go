package main

import (
	"fmt"
)

func main() {
	action([]string{"", "", "", "A", "A"})
	action([]string{})
	action([]string{"A", "B", "B", "C", "B", "B", "B"})
	action([]string{"ABC", "ABC", "ABC", "ABC", "ABC"})
	action([]string{"AA", "BB", "BB", "CC", "BB", "BB", "CC", "CC", "BB", "BB"})
}

func action(s []string) {
	fmt.Printf("%v\n", removeDuplicates(s))
}

func removeDuplicates(s []string) []string {
	if len(s) == 0 {
		return s
	}

	prev := s[0]
	offset := 0

	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			offset++;
		} else {
			s[i - offset] = s[i]
		}

		prev = s[i]
	}

	return s[:len(s) - offset]
}