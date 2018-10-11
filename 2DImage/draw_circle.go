package main

import (
	"github.com/fogleman/gg"
	"flag"
	. "ggg/2DImage/base"
)

func init() {
	InitBG()
	InitSavePath()
	flag.Float64Var(&Radius, "radius", 400, "原型半径")
	flag.Parse()
}

func main() {
	dc := gg.NewContext(Width, Height)
	dc.SetRGB(ColorR, ColorG, ColorB)
	dc.DrawCircle(float64(Width/2), float64(Width/2), Radius)
	dc.Fill()
	dc.SavePNG(PngSavePath)
}
