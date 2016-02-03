package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""

	for index, arg := range os.Args[0:] {
		s += sep + strconv.Itoa(index) + " : " + arg + "\n"
	}

	fmt.Println(s);
}
