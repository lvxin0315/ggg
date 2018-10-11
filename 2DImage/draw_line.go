package main

import (
	"github.com/fogleman/gg"
	"flag"
	. "ggg/2DImage/base"
)

func init() {
	InitBG()
	InitSavePath()

	flag.Float64Var(&X1, "x1", 100, "直线起点x")
	flag.Float64Var(&Y1, "y1", 100, "直线起点y")
	flag.Float64Var(&X2, "x2", 900, "直线终点x")
	flag.Float64Var(&Y2, "y2", 900, "直线终点y")
	flag.Float64Var(&LineWidth, "lineWidth", 5, "线宽")

	flag.Parse()
}

func main() {
	dc := gg.NewContext(Width, Height)
	dc.SetLineWidth(LineWidth)
	dc.SetRGB(ColorR, ColorG, ColorB)
	dc.DrawLine(X1, Y1, X2, Y2)
	dc.Stroke()
	dc.SavePNG(PngSavePath)
}
