package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(commaSignedFloat(""))
	fmt.Println(commaSignedFloat("0"))
	fmt.Println(commaSignedFloat("12"))
	fmt.Println(commaSignedFloat("123"))
	fmt.Println(commaSignedFloat("123456"))
	fmt.Println(commaSignedFloat("1234567"))
	fmt.Println(commaSignedFloat("-1"))
	fmt.Println(commaSignedFloat("-0.1234"))
	fmt.Println(commaSignedFloat("+12.345"))
	fmt.Println(commaSignedFloat("123.4567"))
	fmt.Println(commaSignedFloat("123456.789"))
	fmt.Println(commaSignedFloat("+1234567.890"))
}

func commaSignedFloat(s string) string {
	const splitLength = 3
	pattern := regexp.MustCompile(`([+-]?)(\d+)(\.\d+)?`)

	if !pattern.MatchString(s) {
		return s
	}

	groups := pattern.FindSubmatch([]byte(s))
	return fmt.Sprintf("%s%s%s", groups[1], comma(string(groups[2])), groups[3])
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
