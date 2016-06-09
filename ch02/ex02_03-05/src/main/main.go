package main

import (
	"popcount"
)

func main() {
	popcount.MeasurePopCountPerformance("Table", popcount.PopCountWithTable)
	popcount.MeasurePopCountPerformance("TableAndLoop", popcount.PopCountWithTableAndLoop)
	popcount.MeasurePopCountPerformance("LowestBitLoop", popcount.PopCountWithLowestBitLoop)
	popcount.MeasurePopCountPerformance("BitRemoval", popcount.PopCountWithBitRemoval)
}