package json

import (
	"fmt"
	"testing"
)

type inner struct {
	X int
	Y int
}

type sample struct {
	B bool
	F float64
	I interface{}
	S []string
	M map[string]*inner
}

func TestJson(t *testing.T) {
	v := sample{
		B: false,
		F: 1.2,
		I: inner{2, 3},
		S: []string{"A", "B", "C"},
		M: map[string]*inner{"1": &inner{X: 1, Y: 2}, "2": &inner{X: 1, Y: 2}},
	}
	bytes, _ := Marshal(v)
	fmt.Println(string(bytes))
}

func TestJsonContainsNilValue(t *testing.T) {
	v := sample{
		B: false,
		F: 1.2,
		I: nil,
		S: nil,
		M: nil,
	}
	m := map[string]interface{}{
		"A": 1,
		"B": 2,
		"C": nil,
	}

	bytes, _ := Marshal(v)
	fmt.Println(string(bytes))

	bytes, _ = Marshal(m)
	fmt.Println(string(bytes))
}
