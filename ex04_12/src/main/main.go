package main

import (
	"fmt"
	"xkcd"
	"os"
	"strings"
)

const(
	begNum = 1000
	endNum = 1100
	indexPath = "index.json"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <keyword>", os.Args[0])
		return
	}

	infos := getComicIndex(indexPath)
	keyword := os.Args[1]

	filtered := xkcd.Filtering(infos, func(info *xkcd.ComicInfo) bool {
		return strings.Contains(info.Transcript, keyword)
	})

	printComicInfos(keyword, filtered)
}

func getComicIndex(path string) []*xkcd.ComicInfo {
	if Exists(path) {
		infos, _ := xkcd.LoadComicInfo(path)
		return infos
	} else {
		infos := xkcd.FetchComicInfo(begNum, endNum, true)
		xkcd.SaveComicInfo(path, infos)
		return infos
	}
}

func printComicInfos(keyword string,items []*xkcd.ComicInfo) {
	fmt.Printf("%d items (keyword = %s)\n\n", len(items), keyword)
	for _, item := range items {
		fmt.Printf("Number : %d\n", item.Number)
		fmt.Printf("URL : %s\n", item.URL)
		fmt.Printf("Transcript : \n%s\n", item.Transcript)
		fmt.Println()
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
