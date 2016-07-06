package main

import (
	"fmt"
	"math"
	"io"
	"os"
	"sync"
	"time"
)

const (
	width, height	= 900, 480
	cells			= 300
	xyrange			= 30.0
	xyscale			= width / 2 / xyrange
	zscale			= height * 0.4
	angle			= math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type biFunc func(float64, float64) float64

func main() {
	f := func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return math.Sin(r) / r;
	};

	writerSeq, err := os.Create("seq.svg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err);
		return
	}

	writerPal, err := os.Create("pal.svg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err);
		return
	}

	timeSeq := performance(func() { writeSVG(writerSeq, f) })
	timePal := performance(func() { writeSVGPallarel(writerPal, f)})

	/*
	 * Seq  = 613.0351ms
	 * Pal  = 414.0236ms
	 */
	fmt.Printf("Seq  = %v\n", timeSeq)
	fmt.Printf("Pal  = %v\n", timePal)
}

func performance(action func()) time.Duration {
	start := time.Now()
	action()
	end := time.Now()
	return end.Sub(start)
}

func writeSVG(writer io.Writer, f biFunc) {
	writeSVGHeader(writer)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			writeSVGPolygon(writer, i, j, f)
		}
	}
}

func writeSVGPallarel(writer io.Writer, f biFunc) {
	writeSVGHeader(writer)

	writerChan := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for s := range writerChan {
			writer.Write([]byte(s))
		}
	}()

	wg.Add(cells)
	for i := 0; i < cells; i++ {
		i := i
		go func() {
			for j := 0; j < cells; j++ {
				writeSVGPolygonWithChannel(writerChan, i, j, f)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func writeSVGHeader(writer io.Writer) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width=0.7' " +
		"width='%d' height='%d'>\n", width, height)
}

func writeSVGPolygonWithChannel(writerChan chan <- string, i, j int, f biFunc) {
	ax, ay := corner(i + 1, j, f)
	bx, by := corner(i, j, f)
	cx, cy := corner(i, j + 1, f)
	dx, dy := corner(i + 1, j + 1, f)

	if	!isFinite(ax) || !isFinite(ay) ||
		!isFinite(bx) || !isFinite(by) ||
		!isFinite(cx) || !isFinite(cy) ||
		!isFinite(dx) || !isFinite(dy) {
		return
	}

	writerChan <- fmt.Sprintf("<polygon points = '%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
		ax, ay, bx, by, cx, cy, dx, dy)
}

func writeSVGPolygon(writer io.Writer, i, j int, f biFunc) {
	ax, ay := corner(i + 1, j, f)
	bx, by := corner(i, j, f)
	cx, cy := corner(i, j + 1, f)
	dx, dy := corner(i + 1, j + 1, f)

	if	!isFinite(ax) || !isFinite(ay) ||
		!isFinite(bx) || !isFinite(by) ||
		!isFinite(cx) || !isFinite(cy) ||
		!isFinite(dx) || !isFinite(dy) {
		return
	}

	fmt.Fprintf(writer, "<polygon points = '%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
		ax, ay, bx, by, cx, cy, dx, dy)
}

func corner(i, j int, f biFunc) (float64, float64) {

	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	z := f(x, y)

	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale

	return sx, sy
}

func isFinite(x float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0)
}
