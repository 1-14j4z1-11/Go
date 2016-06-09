package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

func main() {
	str1 := "Test"
	str2 := "test"

	c1 := sha256.Sum256([]byte(str1))
	c2 := sha256.Sum256([]byte(str2))

	fmt.Printf("C1 = %x\n", c1)
	fmt.Printf("C2 = %x\n", c2)
	fmt.Printf("Different bits = %d / %d\n", compareBits(c1, c2), 32 * 8)
}

func compareBits(b1, b2 [32]byte) int {
	length := int(math.Min(float64(len(b1)), float64(len(b2))))
	sub := make([]byte, length)

	for i := 0; i < length; i++ {
		sub[i] = b1[i] ^ b2[i]

	}

	return countBits(sub)
}

func countBits(bits []byte) int {
	count := 0

	for _, x := range bits {
		for x != 0 {
			count++
			x &= (x - 1)
		}
	}

	return count
}
