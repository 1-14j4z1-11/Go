package main

import (
	"fmt"
)

func main() {
	testCase("", "")
	testCase("", "", "", "", "")
	testCase("ABC", "", "A", "B", "C")
	testCase("A,BC,D", ",", "A", "BC", "D")
}

func testCase(expected string, sep string, words ...string) {
	result := Join(sep, words...)

	if(expected != result) {
		panic(fmt.Sprintf("Unexpected result = %s, expected = %s", result, expected))
	} else {
		fmt.Printf("result = [%s]\n", result)
	}
}

func Join(sep string, words ...string) string {
	s := ""

	for i, w := range words {
		if i != 0 {
			s += sep
		}

		s += w
	}

	return s
}