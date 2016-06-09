package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s <url>\n", os.Args[0])
		return
	}

	f, n, err := fetch(os.Args[1])
	fmt.Printf("file=%s\nsize=%d\nerror=%v\n", f, n, err)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err != nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
