package main

import (
	"fmt"
	"math"
	"io"
	"os"
)

const (
	width, height	= 1200, 800
	cells			= 100
	xyrange			= 30.0
	xyscale			= width / 2 / xyrange
	zscale			= height * 0.4
	angle			= math.Pi / 6
	colorScale		= 255
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type biFunc func(float64, float64) float64

func main() {
	writeSVGFile("surface0.xml", func0);
	writeSVGFile("surface1.xml", func1);
	writeSVGFile("surface2.xml", func2);
	writeSVGFile("surface3.xml", func3);
	writeSVGFile("surface4.xml", func4);
}

func writeSVGFile(path string, f biFunc) {
	file, err := os.Create(path)

	if(err != nil) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	writeSVG(file, f)
}

func writeSVG(writer io.Writer, f biFunc) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width=0.7' " +
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j< cells; j++ {
			ax, ay := corner(i + 1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j + 1, f)
			dx, dy := corner(i + 1, j + 1, f)

			if	!isFinite(ax) || !isFinite(ay) ||
				!isFinite(bx) || !isFinite(by) ||
				!isFinite(cx) || !isFinite(cy) ||
				!isFinite(dx) || !isFinite(dy) {
				continue
			}

			r, g, b := color(i, j, f)

			fmt.Fprintf(writer, "<polygon points = '%g,%g,%g,%g,%g,%g,%g,%g' stroke = 'rgb(%d, %d, %d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
}

func corner(i, j int, f biFunc) (float64, float64) {

	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	z := f(x, y)

	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

	return sx, sy
}

func color(i, j int, f biFunc) (int, int, int) {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	z := f(x, y)

	r := math.Min(255.0, math.Max(0.0, z * colorScale + 128.0))
	g := 0.0
	b := 255.0 - r

	return int(r), int(g), int(b)
}


func isFinite(x float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0)
}
