package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	breadthFirst(func(item string) []string {
		files, err := ioutil.ReadDir(item)
		if err != nil {
			return nil
		}

		var paths []string
		for _, file := range files {
			path := item + "/" + file.Name()
			paths = append(paths, path)
			fmt.Printf("%s\n", path)
		}

		return paths
	}, []string{ "." })
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
