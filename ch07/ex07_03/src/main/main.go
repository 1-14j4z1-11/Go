package main

import (
	"fmt"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}

	str := "["

	if t.left != nil {
		str += t.left.String()
	}
	str += fmt.Sprintf("%d ", t.value)
	if t.right != nil {
		str += t.right.String()
	}

	return str + "]"
}

var valueSet map[int]bool = map[int]bool{
	1:  true,
	2:  true,
	3:  true,
	4:  true,
	5:  true,
	6:  true,
	7:  true,
	8:  true,
	9:  true,
	10: true,
}

func main() {
	t := createTree()
	fmt.Printf("%v", t)
}

func createTree() *tree {
	var t *tree
	for k, _ := range valueSet {
		t = add(t, k)
	}
	return t
}
