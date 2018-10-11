package main

import (
	"math"

	"github.com/fogleman/gg"
	"flag"

	. "ggg/2DImage/base"
)

func init() {
	InitBG()
	InitSavePath()
	flag.Float64Var(&X1, "x1", 500, "五角心中心x")
	flag.Float64Var(&Y1, "y1", 500, "五角心中心y")
	flag.Float64Var(&Radius, "radius", 400, "原型半径")
	flag.Float64Var(&LineWidth, "lineWidth", 5, "线宽")

	flag.Parse()
}

type Point struct {
	X, Y float64
}

func Polygon(n int, x, y, r float64) []Point {
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
	}
	return result
}

func main() {
	n := 5
	points := Polygon(n, X1, Y1, Radius)
	dc := gg.NewContext(Width, Height)
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGB(ColorR, ColorG, ColorB)
	dc.FillPreserve()
	dc.Stroke()
	dc.SavePNG(PngSavePath)
}
