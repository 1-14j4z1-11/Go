package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		x := i ^ (1 << uint(i))
		fmt.Printf("x = %d, popcount = %d\n", x, PopCount(uint64(x)))
	}
}
