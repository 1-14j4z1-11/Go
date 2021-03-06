package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for index, url := range os.Args[1:] {
		 resp , err := http.Get(getURL(url))

		fmt.Printf("URL %d: %s\n", index, getURL(url))

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %d: %v\n", index, err)
			os.Exit(1)
		}

		fmt.Printf("StatusCode: %d\n", resp.StatusCode)
		fmt.Printf("ResponseBody:\n", resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)

		if(err != nil) {
			fmt.Fprintf(os.Stderr, "fetch %d: reading %s: %v\n", index, url, err)
		}
	}
}

func getURL(baseURL string) string {
	if strings.HasPrefix(baseURL, "http://") {
		return baseURL
	} else {
		return "http://" + baseURL
	}
}
