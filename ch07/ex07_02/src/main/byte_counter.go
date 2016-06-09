package main

import (
	"bufio"
	"bytes"
)

type ByteCounter struct {
	words int
	lines int
}

func (b *ByteCounter) Words() int {
	return b.words
}

func (b *ByteCounter) Lines() int {
	return b.lines
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	if b == nil {
		return 0, nil
	}

	b.words += countItem(p, bufio.ScanWords)
	b.lines += countItem(p, bufio.ScanLines)
	return len(p), nil
}

func countItem(p []byte, splitFunc func(data []byte, eof bool) (int, []byte, error)) int {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunc)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}
