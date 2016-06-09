package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for index, url := range os.Args[1:] {
		 resp , err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %d: %v\n", index, err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)

		if(err != nil) {
			fmt.Fprintf(os.Stderr, "fetch %d: reading %s: %v\n", index, url, err)
		}
	}
}
