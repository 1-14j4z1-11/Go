package main

import (
	"fmt"
)

const (
	notFound = -1
)

var prereqs = map[string][]string {
	"algorithms" :				{ "data structures" },
	"calculus" :				{ "linear algebra" },
	"complilers" : {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures" :			{ "discreate math" },
	"databases" :				{ "data structures" },
	"discreate math" :			{ "intro to programming" },
	"formal languages" :		{ "discreate math" },
	"networks" :				{ "operating system" },
	"operating system" :		{ "data structures", "computer organization" },
	"programming languagea" :	{ "data structures", "computer organization" },
}

func main() {
	order := topoSort(prereqs)

	if !validateTopoOrder(order, prereqs) {
		fmt.Printf("<< Invalid order >>\n")
	} else {
		fmt.Printf("<< Valid order >>\n")
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i + 1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(sliceToMap(m[item]))
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key, _ := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order
}

func sliceToMap(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}

	return m
}

func validateTopoOrder(order []string, conditions map[string][]string) bool {
	var tryBackwardPath func(val string, valIndex int, seen []string) bool

	tryBackwardPath = func(val string, backwardRange int, seen []string) bool {
		indexInOrder := findIndex(order, val)

		if indexInOrder > backwardRange || findIndex(seen, val) != notFound {
			return false
		}
		seen = append(seen, val)

		for _, back := range conditions[val] {
			if !tryBackwardPath(back, indexInOrder, seen) {
				return false
			}
		}

		return true
	}

	for i, val := range order {
		if !tryBackwardPath(val, i, nil) {
			return false
		}
	}

	return true
}

func findIndex(s []string, target string) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}

	return notFound
}
