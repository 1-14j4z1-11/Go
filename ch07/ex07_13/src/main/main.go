package main

import (
	"eval"
	"fmt"
	"os"
)

func main() {
	testCase("x + y + 10")
	testCase("sin(x + 0.5 * PI) + pow(2, 3.1)")
	testCase("sqrt(x * x + y * y)")
}

func testCase(text string) {
	fmt.Printf("BaseText : %s\n", text)

	expr, err := eval.Parse(text)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err)
		return
	}

	fmt.Printf("Parse    : %v\n", expr)
	expr, err = eval.Parse(expr.String())

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err)
		return
	}

	fmt.Printf("Reparse  : %v\n\n", expr)
}
