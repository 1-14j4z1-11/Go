package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

const(

)

func main() {
	if(len(os.Args) != 2) {
		fmt.Printf("Usage : %s <url>", os.Args[0])
		return
	}

	words, images, err := CountWordsAndImages(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	fmt.Printf("words = %d, images = %d", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return visit(n)
}

func visit(n *html.Node) (int, int) {
	w, i := 0, 0

	if n.Type == html.TextNode {
		w += countWords(n.Data)

	} else if n.Type == html.ElementNode && n.Data == "img" {
		i++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w0, i0 := visit(c)
		w += w0
		i += i0
	}
	return w, i
}

func countWords(line string) int {
	pat := regexp.MustCompile(`[\s\r\n]*[^\s\r\n]+[\s\r\n]*`)
	words := pat.FindAllString(line, -1)
	return len(words)
}
