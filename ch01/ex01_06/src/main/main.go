package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func makePalette() []color.Color {
	var (
		backColor = color.RGBA{0x00, 0x00, 0x00, 0xFF}
		foreColor1 = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
		foreColor2 = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	)

	palette := []color.Color{backColor}
	foreCount := 64

	for i := 0; i < foreCount; i++ {
		alpha := float64(i) / (float64(foreCount) - 1)
		color := color.RGBA{
			makeColor(foreColor1.R, foreColor2.R, alpha),
			makeColor(foreColor1.G, foreColor2.G, alpha),
			makeColor(foreColor1.B, foreColor2.B, alpha),
			makeColor(foreColor1.A, foreColor2.A, alpha)}
		palette = append(palette, color)
	}

	return palette
}

func makeColor(value1 uint8, value2 uint8, alpha float64) uint8 {
	return uint8(float64(value1) * alpha + float64(value2) * (1.0 - alpha))
}

func lissajous(out io.Writer) {
	const (
		cycle	= 5
		res 	= 0.001
		size	= 100
		nframes	= 64
		delay	= 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount : nframes}
	phase := 0.0
	palette := makePalette()

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycle * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), uint8(i % (len(palette) - 1) + 1))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
