package main

import (
	"fmt"
	"image"
	"imageutil"
	"image/png"
	"os"
)

func main() {
	const srcPath = "Lenna.png"

	src, err := imageutil.ReadFile(srcPath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open image %s\n", srcPath)
		return
	}

	superSamplingProcess("050x050y.png", src, 0.5, 0.5)
	superSamplingProcess("150x150y.png", src, 1.5, 1.5)
	superSamplingProcess("200x067y.png", src, 2.0, 0.67)
}

func superSamplingProcess(path string, src image.Image, scaleX, scaleY float64) {

	width := int(float64(src.Bounds().Size().X) * scaleX)
	height := int(float64(src.Bounds().Size().Y) * scaleY)

	dst := imageutil.SuperSampling(src, width, height)

	file, err := os.Create(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open image %s\n", path)
	}

	png.Encode(file, dst)
}
