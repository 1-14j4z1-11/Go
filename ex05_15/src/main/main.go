package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d\n", max())
	fmt.Printf("%d\n", min())
	fmt.Printf("%d\n", max(100))
	fmt.Printf("%d\n", min(100))
	fmt.Printf("%d\n", max(1, 2, 3, 4))
	fmt.Printf("%d\n", min(1, 2, 3, 4))

	fmt.Printf("%d\n", max1(100))
	fmt.Printf("%d\n", min1(100))
	fmt.Printf("%d\n", max1(1, 2, 3, 4))
	fmt.Printf("%d\n", min1(1, 2, 3, 4))
}

func max(vals ...int) int {
	max := math.MinInt32
	for _, val := range vals {
		if max < val {
			max = val
		}
	}

	return max
}

func min(vals ...int) int {
	min := math.MaxInt32
	for _, val := range vals {
		if min > val {
			min = val
		}
	}

	return min
}

func max1(val0 int, vals ...int) int {
	max := val0
	for _, val := range vals {
		if max < val {
			max = val
		}
	}

	return max
}

func min1(val0 int, vals ...int) int {
	min := val0
	for _, val := range vals {
		if min > val {
			min = val
		}
	}

	return min
}
