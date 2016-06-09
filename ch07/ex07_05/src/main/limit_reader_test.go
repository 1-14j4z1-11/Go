package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

var t0 *testing.T
var testCase string

//////////////////////////////////////////////////////////////

func TestWithNilInterface(t *testing.T) {
	setup(t, "NilInterface")

	wrapper := LimitReader(nil, 10)
	_, err := wrapper.Read(make([]byte, 5))

	assertTrue(err == io.EOF, "")

	tearDown()
}

func TestWithReader1(t *testing.T) {
	setup(t, "Reader1")

	reader := bytes.NewBuffer([]byte("12345678"))
	wrapper := LimitReader(reader, 5)
	buf := make([]byte, 10)
	n, err := wrapper.Read(buf)

	assertBytes([]byte("12345"), buf)
	assertInt(5, n)
	assertTrue(err == io.EOF, "")

	tearDown()
}

func TestWithReader2(t *testing.T) {
	setup(t, "Reader2")

	reader := bytes.NewBuffer([]byte("123456789"))
	wrapper := LimitReader(reader, 8)
	buf := make([]byte, 5)

	n, err := wrapper.Read(buf)
	assertBytes([]byte("12345"), buf)
	assertInt(5, n)
	assertTrue(err == nil, "Read1")

	n, err = wrapper.Read(buf)
	assertBytes([]byte("678"), buf)
	assertInt(3, n)
	assertTrue(err == io.EOF, "Read2")

	tearDown()
}

func TestWithReader3(t *testing.T) {
	setup(t, "Reader3")

	reader := bytes.NewBuffer([]byte("123456789"))
	wrapper := LimitReader(reader, 10)
	buf := make([]byte, 5)

	n, err := wrapper.Read(buf)
	assertBytes([]byte("12345"), buf)
	assertInt(5, n)
	assertTrue(err == nil, "Read1")

	n, err = wrapper.Read(buf)
	assertBytes([]byte("6789"), buf)
	assertInt(4, n)
	assertTrue(err == nil, "Read2")

	n, err = wrapper.Read(buf)
	assertTrue(err == io.EOF, "Read3")

	tearDown()
}

//////////////////////////////////////////////////////////////

func setup(t *testing.T, testCase0 string) {
	t0 = t
	testCase = testCase0
}

func tearDown() {
	t0 = nil
	testCase = ""
}

func assertBytes(expected []byte, actual []byte) {
	for i := 0; (i < len(expected)) && (i < len(actual)); i++ {
		if expected[i] != actual[i] {
			assertTrue(false, fmt.Sprintf("Mismatch []byte at %d", i))
		}
	}
}

func assertInt(expected int, actual int) {
	assertTrue(expected == actual, fmt.Sprintf("Unexpected value : %d != %d", expected, actual))
}

func assertTrue(result bool, msg string) {
	if t0 == nil {
		panic(0)
	}

	if !result {
		t0.Errorf("Test Failed [%s] : %s", testCase, msg)
	}
}
