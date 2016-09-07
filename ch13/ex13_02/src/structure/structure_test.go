package structure

import (
	"testing"
)

func TestCircular(t *testing.T) {
	type sample struct {
		Inner interface{}
	}

	var obj sample
	obj.Inner = obj

	m := make(map[string]interface{})
	m["key"] = m

	testcase(t, 1.0, false)
	testcase(t, obj, true)
	testcase(t, m, true)
}

func testcase(t *testing.T, obj interface{}, result bool) {
	if IsCircular(obj) != result {
		t.Errorf("IsCircular() is %v, want %v", !result, result)
	}
}
