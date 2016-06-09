package main

import (
	"fmt"
)

func main() {
	action("")
	action("A")
	action("ABC")
	action("ABCDEFG")
	action("あ")
	action("あいうえお")
	action("ABCDEあいうえお")
	action("あいうえおABCDEFG")
	action("AあBいCうDえEお")
	action("あいうえおABCDEFGかきくけこFさ")
}

func action(s string) {
	fmt.Printf("[%v]\n", reverseString(s))
}

func reverseString(s string) string {
	return string(reverseBytesAsUTF8([]byte(s)))
}

func reverseBytesAsUTF8(bytes []byte) []byte {
	if len(bytes) == 0 {
		return bytes
	}

	// Byte列全体を反転
	reverseBytes(bytes, 0, len(bytes) - 1)

	// 各文字ごとに再度反転
	utf8End := len(bytes) - 1

	for i := len(bytes) - 1; i >= 0; i-- {
		if i == 0 || isUTF8FirstByte(bytes[i - 1]) {
			reverseBytes(bytes, i, utf8End)
			utf8End = i - 1
		}
	}

	return bytes
}

func reverseBytes(bytes []byte, beg, end int) {
	for i, j := beg, end; i < j; i, j = i + 1, j - 1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func isUTF8FirstByte(b byte) bool {
	return (0xC0 & b) != 0x80
}
