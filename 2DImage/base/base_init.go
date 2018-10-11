package base

import "flag"

func InitBG()  {
	flag.IntVar(&Width,"width",1000,"图片宽度")
	flag.IntVar(&Height,"height",1000,"图片高度")
	flag.Float64Var(&ColorR, "colorR", 255, "内容RGB")
	flag.Float64Var(&ColorG, "colorG", 255, "内容RGB")
	flag.Float64Var(&ColorB, "colorB", 255, "内容RGB")
}

func InitSavePath()  {
	flag.StringVar(&PngSavePath, "pngSavePath", "out.png", "文件保存path")
}

