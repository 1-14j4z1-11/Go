package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	type inner struct {
		X int
		Y int
	}
	type sample struct {
		B bool
		F float64
		I interface{}
		S []string
		M map[string]inner
	}

	v := sample{
		B: false,
		F: 1.2,
		I: inner{2, 3},
		S: []string{"A", "B", "C"},
		M: map[string]inner{"1": inner{X: 1, Y: 2}, "2": inner{X: 1, Y: 2}},
	}
	bytes, _ := Marshal(v)
	fmt.Println(string(bytes))

	var s sample
	if err := json.Unmarshal(bytes, &s); err != nil {
		t.Errorf("Unmarshal failed : %v", err)
	}
}
