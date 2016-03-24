package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", rotate([]int{}, 1))
	fmt.Printf("%v\n", rotate([]int{1, 2, 3}, 1))
	fmt.Printf("%v\n", rotate([]int{1, 2, 3, 4, 5}, 2))
	fmt.Printf("%v\n", rotate([]int{}, 1))
	fmt.Printf("%v\n", rotate([]int{}, 1))
	fmt.Printf("%v\n", rotate([]int{}, 1))
}

func rotate(s []int, offset int) []int {
	length := len(s)
	for i := 0; i < len(s) - offset; i++ {
		s[i], s[(i + offset) % length] = s[(i + offset) % length], s[i]
	}

	return s
}