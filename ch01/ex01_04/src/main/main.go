package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	lines := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Printf("Usage : %s <input_files> ...", os.Args[0])
		return;
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error : %v", err)
			continue
		}
		scanLines(arg, f, lines)
		f.Close()
	}
}

func scanLines(path string, f *os.File, lines map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		lines[input.Text()] = append(lines[input.Text()], path)
	}
}
