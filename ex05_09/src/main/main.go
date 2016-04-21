package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const(

)

func main() {
	testCase("", "", func(s string) string { return s })
	testCase("", "$", func(s string) string { return s })
	testCase("bar", "$foo", func(s string) string { return "bar" })
	testCase("FOO", "$foo", strings.ToUpper)
	testCase("A B c d", "$a $b c d", strings.ToUpper)
	testCase("ABC DEF Gh I ", "$abc $Def Gh I $", strings.ToUpper)
}

func testCase(expected string, input string, f func(string) string) {
	if expected != expand(input, f) {
		fmt.Fprintf(os.Stderr, "Failed : expected = '%s', actual = '%s'\n", expected, expand(input, f))
	} else {
		fmt.Printf("OK : expected = '%s', actual = '%s'\n", expected, expand(input, f))
	}
}

func expand(s string, f func(string) string) string {
	replaceFunc := func(s string) string {
		s = f(s)
		return strings.TrimPrefix(s, "$")
	}

	pat := regexp.MustCompile("\\$([^\\s]*)")
	return pat.ReplaceAllStringFunc(s, replaceFunc)
}
