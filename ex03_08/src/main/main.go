package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"mycomplex"
	"os"
)

func main() {
	fractal("complex128.png", mandelbrot_Complex128)
	fractal("complex64.png", mandelbrot_Complex64)
	fractal("bigfloat.png", mandelbrot_BigFloatComplex)
	fractalWithBigRat("bigrat.png", mandelbrot_BigRatComplex)
}

func fractal(path string, function func(r, i float64) color.Color) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y :=  float64(py) / height * (ymax - ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			img.Set(px, py, function(x, y))
		}
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	png.Encode(file, img)
}

func fractalWithBigRat(path string, function func(r, i *big.Rat) color.Color) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := big.NewRat(int64(py), height * (ymax - ymin) + ymin)

		for px := 0; px < width; px++ {
			x := big.NewRat(int64(px), width * (xmax - xmin) + xmin)
			img.Set(px, py, function(x, y))
		}
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	png.Encode(file, img)
}

func mandelbrot_Complex128(r, i float64) color.Color {
	const iterations = 200
	const contrast = 15

	z := complex(r, i)
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			c := 255 - contrast * n
			return color.RGBA{c, 0, 255 - c, 255}
		}
	}

	return color.Black
}

func mandelbrot_Complex64(r, i float64) color.Color {
	const iterations = 200
	const contrast = 15

	z := complex(float32(r), float32(i))
	var v complex64

	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		abs := math.Sqrt(float64(real(v) * real(v) + imag(v) * imag(v)))
		if abs > 2 {
			c := 255 - contrast * n
			return color.RGBA{c, 0, 255 - c, 255}
		}
	}

	return color.Black
}

func mandelbrot_BigFloatComplex(r, i float64) color.Color {
	const iterations = 50
	const contrast = 15

	z := mycomplex.NewBigFloatComplex(r, i)
	v := mycomplex.NewBigFloatComplex(0, 0)

	for n := uint8(0); n < iterations; n++ {
		v = v.Mul(v).Add(z)
		if v.Abs() > 2 {
			c := 255 - contrast * n
			return color.RGBA{c, 0, 255 - c, 255}
		}
	}

	return color.Black
}

func mandelbrot_BigRatComplex(r, i *big.Rat) color.Color {
	const iterations = 5
	const contrast = 15

	z := mycomplex.NewBigRatComplex(r, i)
	v := mycomplex.NewBigRatComplex(mycomplex.NewBigRat(), mycomplex.NewBigRat())

	for n := uint8(0); n < iterations; n++ {
		v = v.Mul(v).Add(z)
		if v.Abs() > 2 {
			c := 255 - contrast * n
			return color.RGBA{c, 0, 255 - c, 255}
		}
	}

	return color.Black
}

func ratToFloat64(r *big.Rat) float64 {
	x, _ := r.Float64()
	return x
}
