package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage : %s <url> [tagName...]\n", os.Args[0])
		return
	}

	doc := parseHTML(os.Args[1])
	if doc == nil {
		return
	}

	searched := ElementByTagName(doc, os.Args[2:]...)

	for _, n := range searched {
		fmt.Printf("%v\n", n)
	}
}

func ElementByTagName(n *html.Node, tagNames ...string) []*html.Node {
	var foundNodes []*html.Node

	startElement := func(n *html.Node) bool {
		if len(tagNames) == 0 || containsTagName(n, tagNames) {
			foundNodes = append(foundNodes, n)
		}

		return true
	}

	endElement := func(n *html.Node) bool {
		return true
	}

	forEachNode(n, startElement, endElement)
	return foundNodes
}

func parseHTML(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprint(os.Stderr, "getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %s: as HTML: %v", url, err)
	}
	return doc
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil && !pre(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && !post(n) {
		return
	}
}

func containsTagName(n *html.Node, tagNames []string) bool {
	for _, tagName := range tagNames {
		if n.Type == html.ElementNode && n.Data == tagName {
			return true
		}
	}

	return false
}
