package main

import (
	"fmt"
	"testing"
)

var t0 *testing.T
var testCase string

//////////////////////////////////////////////////////////////

func TestWithNilInterface(t *testing.T) {
	setup(t, "NilInterface")

	wrapper, counter := CountingWriter(nil)
	wrapper.Write([]byte("1234"))
	assertInt(0, *counter)

	tearDown()
}

func TestWithNilWriter(t *testing.T) {
	setup(t, "NilWriter")

	var writer *ByteCounter
	wrapper, counter := CountingWriter(writer)
	wrapper.Write([]byte("1234"))
	assertInt(0, *counter)

	tearDown()
}

func TestWithWriter(t *testing.T) {
	setup(t, "NilWriter")

	writer := new(ByteCounter)
	wrapper, counter := CountingWriter(writer)
	wrapper.Write([]byte("1234"))
	assertInt(4, *counter)
	wrapper.Write([]byte("12345"))
	assertInt(9, *counter)

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

func assertInt(expected int64, actual int64) {
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
