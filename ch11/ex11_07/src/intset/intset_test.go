package intset

import (
	"math/rand"
	"testing"
)

/*
 * [uint32版]
 * BenchmarkAdd-8                  30000000                43.0 ns/op
 * BenchmarkUnion-8                 1000000              2622 ns/op
 * BenchmarkIntersect-8             1000000              2257 ns/op
 * BenchmarkDifference-8            1000000              1958 ns/op
 * BenchmarkSymmetricDifference-8   1000000              2126 ns/op
 *
 * [uint64版]
 * BenchmarkAdd-8                  30000000                43.0 ns/op
 * BenchmarkUnion-8                 1000000              1497 ns/op
 * BenchmarkIntersect-8             1000000              1200 ns/op
 * BenchmarkDifference-8            1000000              1050 ns/op
 * BenchmarkSymmetricDifference-8   1000000              1161 ns/op
 *
 * [map版]
 * BenchmarkAdd_Map-8                      20000000                89.8 ns/op
 * BenchmarkUnion_Map-8                      500000              2170 ns/op
 * BenchmarkIntersect_Map-8                  500000              2370 ns/op
 * BenchmarkDifference_Map-8                 500000              2346 ns/op
 * BenchmarkSymmetricDifference_Map-8       1000000              2556 ns/op
 */

func BenchmarkAdd(b *testing.B) {
	set := new(IntSet)
	for i := 0; i < b.N; i++ {
		x := makeNum()
		set.Add(x)
	}
}

func BenchmarkUnion(b *testing.B) {
	b.StopTimer()
	set := makeSet()
	measureOperator(b, set.UnionWith)
}

func BenchmarkIntersect(b *testing.B) {
	b.StopTimer()
	set := makeSet()
	measureOperator(b, set.IntersectWith)
}

func BenchmarkDifference(b *testing.B) {
	b.StopTimer()
	set := makeSet()
	measureOperator(b, set.DifferenceWith)
}

func BenchmarkSymmetricDifference(b *testing.B) {
	b.StopTimer()
	set := makeSet()
	measureOperator(b, set.SymmetricDifference)
}

///////////////////////////////////////////////////////////////

func BenchmarkAdd_Map(b *testing.B) {
	set := make(map[int]bool)
	for i := 0; i < b.N; i++ {
		x := makeNum()
		set[x] = true
	}
}

func BenchmarkUnion_Map(b *testing.B) {
	b.StopTimer()
	set := makeSet_Map()
	measureOperator_Map(b, func(other map[int]bool) {
		for k, _ := range other {
			set[k] = true
		}
	})
}

func BenchmarkIntersect_Map(b *testing.B) {
	b.StopTimer()
	set := makeSet_Map()
	measureOperator_Map(b, func(other map[int]bool) {
		for k, _ := range other {
			set[k] = set[k] && true
		}
	})
}

func BenchmarkDifference_Map(b *testing.B) {
	b.StopTimer()
	set := makeSet_Map()
	measureOperator_Map(b, func(other map[int]bool) {
		for k, _ := range other {
			set[k] = false
		}
	})
}

func BenchmarkSymmetricDifference_Map(b *testing.B) {
	b.StopTimer()
	set := makeSet_Map()
	measureOperator_Map(b, func(other map[int]bool) {
		for k, _ := range other {
			set[k] = !set[k]
		}
	})
}

///////////////////////////////////////////////////////////////

func makeSet() *IntSet {
	set := new(IntSet)
	n := rand.Intn(32)
	for i := 0; i < n; i++ {
		set.Add(makeNum())
	}
	return set
}

func measureOperator(b *testing.B, op func(s *IntSet)) {
	for i := 0; i < b.N; i++ {
		x := makeSet()
		b.StartTimer()
		op(x)
		b.StopTimer()
	}
}

func makeSet_Map() map[int]bool {
	set := make(map[int]bool)
	n := rand.Intn(32)
	for i := 0; i < n; i++ {
		set[makeNum()] = true
	}
	return set
}

func measureOperator_Map(b *testing.B, op func(map[int]bool)) {
	for i := 0; i < b.N; i++ {
		x := makeSet_Map()
		b.StartTimer()
		op(x)
		b.StopTimer()
	}
}

func makeNum() int {
	return rand.Intn(0xFFFF)
}
