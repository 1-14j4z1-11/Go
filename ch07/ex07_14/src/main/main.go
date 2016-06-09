package main

import (
	"eval"
	"fmt"
	"os"
)

var env eval.Env = map[eval.Var]float64{
	"x": 1.0,
	"y": 0.5,
}

func main() {
	testCase("x + y + 10")
	testCase("sin(x + 0.5 * PI) + pow(2, 3.1)")
	testCase("sqrt(x * x + y * y)")
	testCase("pow(sin(x), 2) + pow(cos(x), 2)")
	testCase("max(x, y) - min(x, y)")
}

func testCase(text string) {
	fmt.Printf("BaseText : %s\n", text)

	expr, err := eval.Parse(text)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err)
		return
	}

	fmt.Printf("Parse    : %v\n", expr)
	fmt.Printf("Value    = %f\n", expr.Eval(env))
	expr, err = eval.Parse(expr.String())

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err)
		return
	}

	fmt.Printf("Reparse  : %v\n", expr)
	fmt.Printf("Value    = %f\n\n", expr.Eval(env))
}
