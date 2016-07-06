package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage %s <url...>", os.Args[0])
		return
	}

	if res, err := mirroredQuery(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		fmt.Println(res)
	}
}

func mirroredQuery(urls []string) (res string, rErr error) {
	results := make(chan string, len(urls))
	timeout := make(chan struct{})
	cancel := make(chan struct{})

	for _, url := range urls {
		go func() {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}
			req.Cancel = cancel

			if res, err := http.DefaultClient.Do(req); err == nil {
				defer res.Body.Close()

				buf := new(bytes.Buffer)
				io.Copy(buf, res.Body)
				results <- buf.String()
			}
		}()
	}

	go func() {
		time.Sleep(time.Duration(10 * time.Second))
		close(timeout)
	}()

	select {
		case <-timeout:
			close(cancel)
			return "", fmt.Errorf("Timeout")
		case r := <- results:
			close(cancel)
			return r, nil
	}
}
