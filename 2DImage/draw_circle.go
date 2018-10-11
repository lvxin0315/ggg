package main

import (
	"github.com/fogleman/gg"
	"flag"
)

var width int
var height int
var radius float64

var colorR float64
var colorG float64
var colorB float64

func init()  {
	flag.IntVar(&width,"width",1000,"图片宽度")
	flag.IntVar(&height,"height",1000,"图片高度")
	flag.Float64Var(&radius,"radius",400,"原型半径")
	flag.Float64Var(&colorR, "colorR", 255, "背景RGB")
	flag.Float64Var(&colorG, "colorG", 255, "背景RGB")
	flag.Float64Var(&colorB, "colorB", 255, "背景RGB")
	flag.Parse()
}

func main() {
	dc := gg.NewContext(width,height)
	dc.DrawCircle(	float64(width / 2), float64(width / 2), radius)
	dc.SetRGB(colorR, colorG, colorB)

	dc.Fill()
	dc.SavePNG("out.png")
}