package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d", square(3))
}

func square(x int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = x * x
		}
	}()

	panic(0)
}
