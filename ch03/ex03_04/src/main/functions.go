package main

import (
	"math"
)

func func0(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r;
}

func func1(x, y float64) float64 {
	r := math.Hypot(x, y)
	return r * math.Sin(r) / 64;
}

func func2(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 8
}

func func3(x, y float64) float64 {
	return (math.Pow(x, 4) - math.Pow(y, 4)) / 65536
}

func func4(x, y float64) float64 {
	return (y * math.Sin(x) + x * math.Sin(y)) / 256
}
