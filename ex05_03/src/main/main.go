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
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isVisibleNode(c) {
			texts = visit(texts, c)
		}
	}
	return texts
}

func isVisibleNode(n *html.Node) bool {
	if(n.Type != html.ElementNode) {
		return true
	}

	return n.Data != "script" && n.Data != "style"
}
