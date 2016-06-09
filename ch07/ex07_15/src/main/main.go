package main

import (
	"bufio"
	"eval"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Expression > ")
	reader := bufio.NewReader(os.Stdin)

	line, _, err := reader.ReadLine()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	expr, vars, err := eval.Parse(string(line))

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	var env eval.Env = make(map[eval.Var]float64)

	if len(vars) > 0 {
		fmt.Println("\n< Define Vars >")
	}

	for _, v := range vars {
		fmt.Printf("%v = ", v)
		line, _, err := reader.ReadLine()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}

		val, err := strconv.ParseFloat(string(line), 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}

		env[v] = val
	}

	varSet := make(map[eval.Var]bool)

	for _, v := range vars {
		varSet[v] = true
	}

	err = expr.Check(varSet)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	fmt.Printf("\nValue = %f", expr.Eval(env))
}
