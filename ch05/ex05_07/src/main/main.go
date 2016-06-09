package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <url>\n", os.Args[0])
		return
	}

	doc := parseHTML(os.Args[1])
	if doc == nil {
		return
	}

	depth := 0
	isShortTag := false
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			isShortTag = (n.FirstChild == nil)

			if isShortTag {
				fmt.Printf("%*s<%s%s/>", depth * 2, "", n.Data, attrString(n))
			} else {
				fmt.Printf("%*s<%s%s>\n", depth * 2, "", n.Data, attrString(n))
			}
			depth++
		} else if n.Type == html.TextNode && shouldPrint(n) {
			if isShortTag {
				fmt.Printf("%s\n", n.Data)
			} else {
				fmt.Printf("%*s%s\n", depth * 2, "", n.Data)
			}
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			if !isShortTag {
				fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
			} else {
				fmt.Println()
			}
		}

		isShortTag = false
	}

	forEachNode(doc, startElement, endElement)
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

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func attrString(n *html.Node) string {
	s := ""

	for _, a := range n.Attr {
		s += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
	}

	return s
}

func shouldPrint(n *html.Node) bool {
	pat := regexp.MustCompile(`^[ \r\n]*$`)
	return !pat.MatchString(n.Data)
}
