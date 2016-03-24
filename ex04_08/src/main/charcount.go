package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type RuneType string

const (
	Letter	RuneType = "Letter"
	Mark	RuneType = "Mark"
	Number	RuneType = "Number"
	Punct	RuneType = "Punct"
	Space	RuneType = "Space"
	Other	RuneType = "Other"
)

func main() {
	counts := make(map[rune]int)
	types := make(map[RuneType]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		counts[r]++
		types[typeOfRune(r)]++
	}
	fmt.Println("rune\tcount")
	for r, n := range counts {
		fmt.Printf("%q\t%d\n", r, n)
	}
	fmt.Println("\ntype\tcount")
	for t, n := range types {
		fmt.Printf("%s\t%d\n", t, n)
	}
}

func typeOfRune(r rune) RuneType {
	if unicode.IsLetter(r) {
		return Letter
	} else if unicode.IsMark(r) {
		return Mark
	} else if unicode.IsNumber(r) {
		return Number
	} else if unicode.IsPunct(r) {
		return Punct
	} else if unicode.IsSpace(r) {
		return Space
	} else {
		return Other
	}
}
