package imageutil

import (
	"image"
	"image/color"
	"math"
	"os"
)

func ReadFile(path string) (image.Image, error) {
	file, fileError := os.Open(path)

	if fileError != nil {
		return nil, fileError
	}

	img, _, imgError := image.Decode(file)
	return img, imgError
}

func SuperSampling(src image.Image, dstWidth, dstHeight int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, dstWidth, dstHeight))
	scaleX := float64(dstWidth) / float64(src.Bounds().Size().X)
	scaleY := float64(dstHeight) / float64(src.Bounds().Size().Y)

	for x := 0; x < dstWidth; x++ {
		for y := 0; y < dstHeight; y++ {
			pixel := calcSubpixel(src, float64(x) / scaleX, float64(x + 1) / scaleX, float64(y) / scaleY, float64(y + 1) / scaleY)
			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func calcSubpixel(src image.Image, begX, endX, begY, endY float64) color.Color {
	var r, g, b, a float64
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y

	for x := int(begX); x <= int(endX + 1); x++ {
		for y := int(begY); y <= int(endY + 1); y++ {
			if x >= width || y >= height {
				continue
			}

			dx, dy := 1.0, 1.0

			if float64(x) < begX {
				dx -= begX - float64(x)
			}
			if float64(x) > endX - 1 {
				dx -= math.Min(float64(x) - endX + 1, 1)
			}
			if float64(y) < begY {
				dy -= begY - float64(y)
			}
			if float64(y) > endY - 1 {
				dy -= math.Min(float64(y) - endY + 1, 1)
			}

			dxy := dx * dy

			r += float64(getR(src.At(x, y))) * dxy
			g += float64(getG(src.At(x, y))) * dxy
			b += float64(getB(src.At(x, y))) * dxy
			a += float64(getA(src.At(x, y))) * dxy
		}
	}

	area := (endX - begX) * (endY - begY)
	r /= area
	g /= area
	b /= area
	a /= area

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func getR(pixel color.Color) uint8 {
	value, _, _, _ := pixel.RGBA()
	return uint8((value >> 8) & 0xFF)
}

func getG(pixel color.Color) uint8 {
	_, value, _, _ := pixel.RGBA()
	return uint8((value >> 8) & 0xFF)
}

func getB(pixel color.Color) uint8 {
	_, _, value, _ := pixel.RGBA()
	return uint8((value >> 8) & 0xFF)
}

func getA(pixel color.Color) uint8 {
	_, _, _, value := pixel.RGBA()
	return uint8((value >> 8) & 0xFF)
}


