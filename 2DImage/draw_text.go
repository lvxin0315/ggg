package main

import (
	"github.com/fogleman/gg"
	"flag"
	. "ggg/2DImage/base"
)

func init() {
	InitBG()
	InitSavePath()

	flag.StringVar(&TextContent, "text", "文本内容\n第二行", "文本内容")
	flag.IntVar(&TextFontSize, "textFontSize", 100, "字体大小")
	flag.Parse()
}

func main() {
	dc := gg.NewContext(Width, Height)
	dc.SetRGB(ColorR, ColorG, ColorB)
	if err := dc.LoadFontFace("微软雅黑.ttf", float64(TextFontSize)); err != nil {
		panic(err)
	}

	dc.DrawStringWrapped(TextContent,
		float64(Width/2),
		float64(TextFontSize/4),
		0,
		0,
		float64(TextWidth),
		1.5,
		gg.AlignCenter)

	dc.SavePNG("out.png")
}
