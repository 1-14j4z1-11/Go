package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks : %v\n", err)
		os.Exit(1)
	}
	for eType, count := range visit(nil, doc) {
		fmt.Printf("%s : %d\n", eType, count)
	}
}

func visit(counts map[string]int, n *html.Node) map[string]int {
	if counts == nil {
		counts = make(map[string]int)
	}

	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		counts = visit(counts, c)
	}
	return counts
}
