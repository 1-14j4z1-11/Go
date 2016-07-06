package main

import (
	"fmt"
	"log"
	"math"
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
		fmt.Printf("Usage : %s [-depth <depth>] <url...>", os.Args[0])
		return
	}

	urls, depth := parseArgs(os.Args)

	worklist := make(chan []linkInfo)
	unseenLinks := make(chan linkInfo)

	go func() { worklist <- linkList(urls, 0) }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.Url)
				go func(link linkInfo) { worklist <- linkList(foundLinks, link.Depth + 1) }(link)
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.Url] && (link.Depth <= depth) {
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

func parseArgs(args []string) ([]string, int) {
	var urls []string
	depth := math.MaxInt32

	regex := regexp.MustCompile("^-depth=(\\d+)")

	for _, arg := range args {
		if regex.MatchString(arg) {
			depth, _ = strconv.Atoi(regex.FindStringSubmatch(arg)[1])
		} else {
			urls = append(urls, arg)
		}
	}

	return urls, depth
}
