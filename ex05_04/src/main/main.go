package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

const(

)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks : %v\n", err)
		os.Exit(1)
	}
	for t, link := range visit(nil, doc) {
		fmt.Printf("%s : %s\n", t, link)
	}
}

func visit(links map[string]string, n *html.Node) map[string]string {
	if links == nil {
		links = make(map[string]string)
	}

	if n.Type == html.ElementNode {
		linkType := ""
		link := ""
		var err error

		switch(n.Data) {
			case "a":
				link, err = getAttr(n, "href")
				linkType = "hyperlink"
			case "img":
				link, err = getAttr(n, "src")
				linkType = "img"
			case "link":
				link, err = getAttr(n, "href")
				linkType = "style"
			case "script":
				link, err = getAttr(n, "type")
				linkType = "script"
			default:
				err = fmt.Errorf("NotFound")
		}

		if err == nil {
			links[linkType] = link
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func getAttr(n *html.Node, key string) (string, error) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, nil
		}
	}

	return "", fmt.Errorf("Not Found")
}
