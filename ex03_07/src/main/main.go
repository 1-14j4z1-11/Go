package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -4, -4, 4, 4
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newtonMethod(z, makeRootRecursion(12)))
		}
	}

	file, err := os.Create("newton.png")

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	png.Encode(file, img)
}

// n乗根を求める漸化式関数を作成する
func makeRootRecursion(n int) func(complex128) complex128 {
	root := complex(float64(n), 0)

	return func(z complex128) complex128 {
		return z - (cmplx.Pow(z, root) + 1) / (root * cmplx.Pow(z, root - 1))
	}
}

// ニュートン法を行う関数
func newtonMethod(initial complex128, recursion func(complex128) complex128) color.Color {
	const iterations = 200
	const theta = 1e-6
	const contrast = 10

	v := initial

	for n := uint8(0); n < iterations; n++ {
		v0 := v
		v = recursion(v)

		if cmplx.Abs(v - v0) < theta {
			c := n * contrast
			return color.RGBA{c, 0, 255 - c, 255}
		}
	}

	return color.Black
}

