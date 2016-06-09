package fractal

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"io"
)

type FractalParam struct {
	Xmin	float64
	Ymin	float64
	Xmax	float64
	Ymax	float64
	Width	int
	Height	int
	Root	int
}

// フラクタル用パラメータを生成する
func NewFractalParam(xmin float64, ymin float64, xmax float64, ymax float64, width int, height int, root int) *FractalParam {
	param := new(FractalParam)

	param.Xmin = xmin
	param.Ymin = ymin
	param.Xmax = xmax
	param.Ymax = ymax
	param.Width = width
	param.Height = height
	param.Root = root

	return param
}

// フラクタルをPNGファイルに書き込む
func WriteFractalPng(writer io.Writer, param *FractalParam) {

	img := image.NewRGBA(image.Rect(0, 0, param.Width, param.Height))

	for py := 0; py < param.Height; py++ {
		y := float64(py) / float64(param.Height) * (param.Ymax - param.Ymin) + param.Ymin

		for px := 0; px < param.Width; px++ {
			x := float64(px) / float64(param.Width) * (param.Xmax - param.Xmin) + param.Xmin
			z := complex(x, y)
			img.Set(px, py, newtonMethod(z, makeRootRecursion(param.Root)))
		}
	}

	png.Encode(writer, img)
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
