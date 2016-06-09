package main

import (
	"fmt"
	"unicode"
)

func main() {
	action("")
	action(" ")
	action("       ")
	action("A \t \t \tA")
	action(" AA BB CC DD ")
	action(" \nAA\r\vBB\t\fCC\v\vDD    ")
	action(" \v\t\n\f\rAA \v\t\n\f\r")
}

func action(s string) {
	fmt.Printf("[%v]\n", removeDuplicateSpace(s))
}

func removeDuplicateSpace(s string) string {
	bytes := []byte(s)
	bytes = removeDuplicateSpaceInBytes(bytes)
	return string(bytes)
}

func removeDuplicateSpaceInBytes(s []byte) []byte {
	const space = byte(' ')

	if len(s) == 0 {
		return s
	}

	prevIsSpace := isSpace(s[0])
	offset := 0

	for i := 1; i < len(s); i++ {
		if prevIsSpace && isSpace(s[i]) {
			offset++;
		} else if isSpace(s[i]) {
			s[i - offset] = space
		} else {
			s[i - offset] = s[i]
		}

		prevIsSpace = isSpace(s[i])
	}

	return s[:len(s) - offset]
}

func isSpace(b byte) bool {
	return unicode.IsSpace(rune(b))
}