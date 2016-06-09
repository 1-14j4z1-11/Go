package main

import (
	"fmt"
)

const arraySize = 5

func main() {
	a := [arraySize]int{1, 2, 3, 4, 5}

	fmt.Printf("%v\n", a)
	reverseArray(&a)
	fmt.Printf("%v\n", a)
}

func reverseArray(array *[arraySize]int) {
	for i, j := 0, len(array) - 1; i < j; i, j = i + 1, j - 1 {
		array[i], array[j] = array[j], array[i]
	}
}