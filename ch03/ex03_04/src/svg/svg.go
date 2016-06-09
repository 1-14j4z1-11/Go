package svg

import (
	"image/color"
	"fmt"
	"math"
	"io"
)

const (
	angle			= math.Pi / 6
	colorScale		= 1
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type BiFunction func(float64, float64) float64

type SVGParam struct {
	Width		int
	Height		int
	Cells		int
	XYrange		float64
	MaxColor	color.Color
	MinColor	color.Color
	Function	BiFunction
}

func NewSVGParam(width int, height int, cells int, xyrange float64, maxColor color.Color, minColor color.Color, f BiFunction) *SVGParam {
	param := new(SVGParam)
	param.Width = width
	param.Height = height
	param.Cells = cells
	param.XYrange = xyrange
	param.MaxColor = maxColor
	param.MinColor = minColor
	param.Function = f

	return param
}

func WriteSVG(writer io.Writer, param *SVGParam) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width=0.7' " +
		"width='%d' height='%d'>\n", param.Width, param.Height)

	for i := 0; i < param.Cells; i++ {
		for j := 0; j< param.Cells; j++ {
			ax, ay := corner(i + 1, j, param)
			bx, by := corner(i, j, param)
			cx, cy := corner(i, j + 1, param)
			dx, dy := corner(i + 1, j + 1, param)

			if	!isFinite(ax) || !isFinite(ay) ||
				!isFinite(bx) || !isFinite(by) ||
				!isFinite(cx) || !isFinite(cy) ||
				!isFinite(dx) || !isFinite(dy) {
				continue
			}

			r, g, b := getColor(i, j, param)

			fmt.Fprintf(writer, "<polygon points = '%g,%g,%g,%g,%g,%g,%g,%g' stroke = 'rgb(%d, %d, %d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
}

func corner(i, j int, param *SVGParam) (float64, float64) {
	x := param.XYrange * (float64(i) / float64(param.Cells) - 0.5)
	y := param.XYrange * (float64(j) / float64(param.Cells) - 0.5)

	z := param.Function(x, y)

	xyscale	:= float64(param.Width) / 2.0 / param.XYrange
	zscale := float64(param.Height) * 0.4

	sx := float64(param.Width) / 2 + (x - y) * cos30 * xyscale
	sy := float64(param.Height) / 2 + (x + y) * sin30 * xyscale - z * zscale

	return sx, sy
}

func getColor(i, j int, param *SVGParam) (int, int, int) {
	x := param.XYrange * (float64(i) / float64(param.Cells) - 0.5)
	y := param.XYrange * (float64(j) / float64(param.Cells) - 0.5)
	z := param.Function(x, y)

	a := math.Min(1.0, math.Max(-1.0, z * colorScale)) * 0.5 + 0.5

	r1, g1, b1, _ := param.MaxColor.RGBA()
	r2, g2, b2, _ := param.MinColor.RGBA()

	r := float64(r1 & 0xFF) * a + float64(r2 & 0xFF) * (1.0 - a)
	g := float64(g1 & 0xFF) * a + float64(g2 & 0xFF) * (1.0 - a)
	b := float64(b1 & 0xFF) * a + float64(b2 & 0xFF) * (1.0 - a)

	return int(r), int(g), int(b)
}

func isFinite(x float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0)
}
