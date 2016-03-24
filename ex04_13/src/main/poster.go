package main

import (
	"fmt"
	"io"
	"net/http"
	"omdb"
	"os"
)

const(
	imagePath = "poster.jpg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <title>", os.Args[0])
		return
	}

	title := os.Args[1]

	info, err := omdb.FetchMovieInfo(title)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	resp, err := http.Get(info.PosterURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	file, err := os.Create(imagePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	io.Copy(file, resp.Body)
	resp.Body.Close()
	file.Close()
}
