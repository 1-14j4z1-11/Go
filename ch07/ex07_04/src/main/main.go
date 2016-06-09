package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type stringReader struct {
	str   string
	count int
}

func (sr *stringReader) Read(p []byte) (int, error) {
	if sr.count >= len(sr.str) {
		return 0, io.EOF
	}

	var reading []byte

	if len(p) < len(sr.str)-sr.count {
		reading = []byte(sr.str)[sr.count : sr.count+len(p)]
	} else {
		reading = []byte(sr.str)[sr.count:]
	}

	sr.count += len(reading)
	copy(p, reading)

	return len(reading), nil
}

func NewReader(s string) io.Reader {
	reader := new(stringReader)
	reader.str = s
	return reader
}

func main() {
	reader := NewReader(htmlString())
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks : %v\n", err)
		os.Exit(1)
	}
	for eType, count := range visit(nil, doc) {
		fmt.Printf("%s : %d\n", eType, count)
	}
}

func htmlString() string {
	return `<html>
<body>

<h1>1</h1>
<h2>2</h2>
<h3>3</h3>

<p>p</p>

</body>
</html>`
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
