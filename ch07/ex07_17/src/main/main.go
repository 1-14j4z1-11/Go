package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(separateElems(stack), " "), tok)
			}
		}
	}
}

func containsAll(elems []xml.StartElement, y []string) bool {
	x := separateElems(elems)

	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}

	return false
}

func separateElems(elems []xml.StartElement) []string {
	vals := make([]string, 0)

	for _, e := range elems {
		vals = append(vals, e.Name.Local)
		for _, a := range e.Attr {
			vals = append(vals, a.Name.Local)
			vals = append(vals, a.Value)
		}
	}

	return vals
}
