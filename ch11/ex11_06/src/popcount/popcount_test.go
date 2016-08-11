package popcount

import (
	"testing"
)

func BenchmarkPopCountWithTable(b *testing.B) {
	setup(b)
	measurePopCount(b, WithTable)
}

func BenchmarkPopCountWithTableAndLoop(b *testing.B) {
	setup(b)
	measurePopCount(b, WithTableAndLoop)
}

func BenchmarkPopCountWithLowestBitLoop(b *testing.B) {
	setup(b)
	measurePopCount(b, WithLowestBitLoop)
}

func BenchmarkPopCountWithBitRemoval(b *testing.B) {
	setup(b)
	measurePopCount(b, WithBitRemoval)
}

func setup(b *testing.B) {
	b.StopTimer()
}

func measurePopCount(b *testing.B, f func(uint64) int) {
	x := makeNum()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f(x)
	}
	b.StopTimer()
}

func makeNum() uint64 {
	return 0xAAAAAAAAAAAAAAAA
}
