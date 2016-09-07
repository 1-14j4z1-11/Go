package equal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEqualFloat(t *testing.T) {
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{1.0, 1.0, true},
		{2.0, 1, false},
		{1.0 + 1e-10, 1.0, true},
		{1.0 + 1e-9, 1.0, false},
		{2.0 + 1.0i + 1e-10i, 2.0 + 1.0i, true},
		{2.0 + 1.0i + 1e-9i, 2.0 + 1.0i, false},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}

func TestEqual(t *testing.T) {
	one, oneAgain, two := 1, 1, 2

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice CycleSlice
	cycleSlice = append(cycleSlice, cycleSlice)

	ch1, ch2 := make(chan int), make(chan int)
	var ch1ro <-chan int = ch1

	type mystring string

	var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{1, 1, true},
		{1, 2, false},
		{1, 1.0, false},
		{"foo", "foo", true},
		{"foo", "bar", false},
		{mystring("foo"), "foo", false},
		{[]string{"foo"}, []string{"foo"}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string{}, []string(nil), true},
		{cycleSlice, cycleSlice, true},
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3}},
			true,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3, 4}},
			false,
		},
		{
			map[string][]int{},
			map[string][]int(nil),
			true,
		},
		{&one, &one, true},
		{&one, &two, false},
		{&one, &oneAgain, true},
		{new(bytes.Buffer), new(bytes.Buffer), true},
		{cyclePtr1, cyclePtr1, true},
		{cyclePtr2, cyclePtr2, true},
		{cyclePtr1, cyclePtr2, true},
		{(func())(nil), (func())(nil), true},
		{(func())(nil), func() {}, false},
		{func() {}, func() {}, false},
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}, true},
		{[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
		{ch1, ch1, true},
		{ch1, ch2, false},
		{ch1ro, ch1, false},
		{&iface1, &iface1, true},
		{&iface1, &iface2, false},
		{&iface1Again, &iface1, true},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}

func Example_equal() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))
	fmt.Println(Equal([]string(nil), []string{}))
	fmt.Println(Equal(map[string]int(nil), map[string]int{}))
}

func Example_equalCycle() {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a))
	fmt.Println(Equal(b, b))
	fmt.Println(Equal(c, c))
	fmt.Println(Equal(a, b))
	fmt.Println(Equal(a, c))
}
