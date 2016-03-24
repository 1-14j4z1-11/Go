package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma(""))
	fmt.Println(comma("0"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}

func comma(s string) string {
	const splitLength = 3

	var buffer bytes.Buffer
	offset := len(s) % splitLength;
	buffer.WriteString(s[0:offset])

	for i := offset; i < len(s); i++ {
		if i != 0 && i % splitLength == offset {
			buffer.WriteRune(',')
		}
		buffer.WriteByte(s[i]);
	}

	return buffer.String()
}
