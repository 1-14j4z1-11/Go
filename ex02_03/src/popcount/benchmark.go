package popcount

import (
	"fmt"
	"performance"
	"time"
)

type PopCountFunction func(x uint64) int

type testCase struct {
	input uint64
	expected int
}

func newTestCase(input uint64, expected int) *testCase {
	obj := new(testCase)
	obj.input = input
	obj.expected = expected

	return obj
}

var testCases []testCase

func init() {
	testCases = append(testCases, *newTestCase(0, 0))
	testCases = append(testCases, *newTestCase(128, 1))
	testCases = append(testCases, *newTestCase(1023, 10))
	testCases = append(testCases, *newTestCase(3039518059, 19))
	testCases = append(testCases, *newTestCase(4294967295, 32))
}

func MeasurePopCountPerformance(title string, function PopCountFunction) {
	failed := 0
	dur := time.Duration(0)

	for _, tc := range testCases {
		var result int
		dur += performance.MeasurePerformance(func() { result = function(tc.input) })

		if result != tc.expected {
			failed++
		}
	}

	fmt.Printf("<%s>\n\tFailed = %d / %d\n\tTime = %s\n", title, failed, len(testCases), dur.String())

}
