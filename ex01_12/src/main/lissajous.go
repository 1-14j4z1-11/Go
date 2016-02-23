package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

const (
	backColorIndex = 0
	foreColorIndex = 1
)

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xFF}, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

type LissajousParameter struct {
	cycle int
	size int
	nframes int
}

func newLissajousParameter(cycle int, size int, nframes int) *LissajousParameter {
	obj := new(LissajousParameter)
	obj.cycle, obj.size, obj.nframes = cycle, size, nframes
	return obj
}

func lissajous(out io.Writer, param *LissajousParameter) {
	const (
		res 	= 0.001
		delay	= 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount : param.nframes}
	phase := 0.0

	for i := 0; i < param.nframes; i++ {
		rect := image.Rect(0, 0, 2 * param.size + 1, 2 * param.size + 1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(param.cycle) * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(param.size + int(x * float64(param.size) + 0.5), param.size + int(y * float64(param.size) + 0.5), foreColorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
