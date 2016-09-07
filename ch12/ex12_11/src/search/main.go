package main

import (
	"fmt"
	"params"
)

type query struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func testcase(labels []string, max int, exact bool) {
	q := query{Labels:labels, MaxResults:max, Exact:exact}
	fmt.Printf("%s\n", params.Pack("http://example.com/", q))
}

func main() {
	testcase([]string{"golang"}, 0, false)
	testcase([]string{"golang", "programming"}, 11, true)
}
