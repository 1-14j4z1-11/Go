package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage : %s <url> <attr_key>\n", os.Args[0])
		return
	}

	doc := parseHTML(os.Args[1])
	if doc == nil {
		return
	}

	searched := ElementByID(doc, os.Args[2])

	if searched != nil {
		fmt.Printf("%v\n", searched)
	} else {
		fmt.Println("Not found")
	}
}

func ElementByID(n *html.Node, id string) *html.Node {
	isFound := false
	var foundNode *html.Node

	startElement := func(n *html.Node) bool {
		if(isFound) {
			return false
		}

		for _, a := range n.Attr {
			if a.Key == id {
				isFound = true
				foundNode = n
				return false
			}
		}

		return true
	}

	endElement := func(n *html.Node) bool {
		return true
	}

	forEachNode(n, startElement, endElement)
	return foundNode
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
