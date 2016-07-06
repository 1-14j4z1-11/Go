package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"gopl.io/ch5/links"
)

type linkInfo struct {
	Url		string
	Depth	int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage : %s [-depth <depth>] <url>", os.Args[0])
		return
	}

	rootUrl, depth := parseArgs(os.Args)

	saveList := make(chan string)
	worklist := make(chan []linkInfo)
	unseenLinks := make(chan linkInfo)
	cancel := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(worklist)
		close(cancel)
	}()

	go func() { worklist <- linkList([]string{ rootUrl }, 0) }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.Url)
				go func(link linkInfo) {
					saveList <- link.Url
					worklist <- linkList(foundLinks, link.Depth + 1)
				}(link)
			}
		}()
	}

	for i := 0; i < 20; i++ {
		go saveContents(saveList, cancel)
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.Url] && (link.Depth <= depth) && isSameDomain(rootUrl, link.Url) {
				seen[link.Url] = true
				unseenLinks <- link
			}
		}
	}
}

func linkList(urlList []string, depth int) []linkInfo {
	var infoList []linkInfo
	for _, url := range urlList {
		infoList = append(infoList, linkInfo{ Url:url, Depth:depth })
	}
	return infoList
}

func parseArgs(args []string) (string, int) {
	var url string
	depth := math.MaxInt32

	regex := regexp.MustCompile("^-depth=(\\d+)")

	for _, arg := range args {
		if regex.MatchString(arg) {
			depth, _ = strconv.Atoi(regex.FindStringSubmatch(arg)[1])
		} else {
			url = arg
		}
	}

	return url, depth
}

var urlPatettern = regexp.MustCompile("(?:http://|https://)([^/]+)(?:(?:/.+)*)")

func isSameDomain(root, target string) bool {
	rootDomain := urlPatettern.FindStringSubmatch(root)[1]
	targetDomain := urlPatettern.FindStringSubmatch(target)[1]
	rootDomainPattern, err := regexp.Compile(rootDomain)
	if err != nil {
		return false
	}
	return rootDomainPattern.MatchString(targetDomain)
}

func saveContents(urlChan <-chan string, cancel <-chan struct{}) {
	for url := range urlChan {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}
		req.Cancel = cancel

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer res.Body.Close()

		if file, err := createLocalFile(url); err == nil {
			defer file.Close()
			io.Copy(file, res.Body)
		}
	}
}

func createLocalFile(urlStr string) (*os.File, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	path := url.Path
	if path == "" || path == "/" {
		path = "index.html"
	}
	path = "." + path

	if err = os.MkdirAll(path, 0x777); err != nil {
		return nil, err
	}
	os.Remove(path)

	return os.Create(path)
}
