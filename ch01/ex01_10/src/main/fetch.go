package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<- ch)
	}
	fmt.Printf("fetched %d urls, %.2fs elasped\n", len(os.Args[1:]), time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	url = getURL(url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	out, err_open := os.Create(fmt.Sprintf("./body_%s.txt", time.Now().Format("20060102-030405")))

	if err_open != nil {
		ch <- fmt.Sprint(err_open)
	}

	nbytes, err := io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func getURL(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}

/*
<実行結果>

>main http://amazon.com
2.56s  413849 http://amazon.com
fetched 1 urls, 2.56s elasped

>main http://amazon.com
1.68s  220308 http://amazon.com
fetched 1 urls, 1.68s elasped

>main http://amazon.com
1.75s  220209 http://amazon.com
fetched 1 urls, 1.75s elasped

>main http://amazon.com
2.44s  416340 http://amazon.com
fetched 1 urls, 2.44s elasped
*/