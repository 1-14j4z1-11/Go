package main

import (
	"fmt"
	"testing"
)

var t0 *testing.T
var target *ByteCounter
var testCase string

//////////////////////////////////////////////////////////////

func TestEmpy(t *testing.T) {
	bc := setup(t, "Empty")

	bc.Write([]byte(""))
	assertCount(0, 0)

	tearDown()
}

func TestSingleWord(t *testing.T) {
	bc := setup(t, "SingleWord")

	bc.Write([]byte("TEST"))
	assertCount(1, 1)

	tearDown()
}

func TestMultiWord(t *testing.T) {
	bc := setup(t, "MultiWord")

	bc.Write([]byte("aaa, bb, ccc, dd, eee"))
	assertCount(5, 1)

	tearDown()
}

func TestMultiLine(t *testing.T) {
	bc := setup(t, "MultiLine")

	bc.Write([]byte("aaa, bb\nccc, dd\neee\n"))
	assertCount(5, 3)

	tearDown()
}

func TestMultiWrite(t *testing.T) {
	bc := setup(t, "MultiLine")

	bc.Write([]byte("aaa, bb\n ccc, dd\n eee\n"))
	assertCount(5, 3)

	bc.Write([]byte("1"))
	assertCount(6, 4)

	bc.Write([]byte("2"))
	assertCount(7, 5)

	bc.Write([]byte("3 4 5\n6"))
	assertCount(11, 7)

	tearDown()
}

//////////////////////////////////////////////////////////////

func setup(t *testing.T, testCase0 string) *ByteCounter {
	t0 = t
	target = new(ByteCounter)
	testCase = testCase0

	return target
}

func tearDown() {
	t0 = nil
	target = nil
	testCase = ""
}

func assertCount(words int, lines int) {
	assertTrue(target.Words() == words, fmt.Sprintf("Unexpected words : %d != %d", target.Words(), words))
	assertTrue(target.Lines() == lines, fmt.Sprintf("Unexpected lines : %d != %d", target.Lines(), lines))
}

func assertTrue(result bool, msg string) {
	if t0 == nil {
		panic(0)
	}

	if !result {
		t0.Errorf("Test Failed [%s] : %s", testCase, msg)
	}
}
