package main

import (
	"fmt"
	"tempconv"
)

func main() {
	c := tempconv.Celsius(20)
	f := tempconv.Fahrenheit(160)
	k := tempconv.Kelven(450)

	fmt.Printf("C : %v -> F : %v, K : %v\n", c, c.ToF(), c.ToK())
	fmt.Printf("F : %v -> C : %v, K : %v\n", f, f.ToC(), f.ToK())
	fmt.Printf("K : %v -> C : %v, F : %v\n", k, k.ToC(), k.ToF())
}