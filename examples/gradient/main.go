package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

// Demonstrate creating a gradient image including artifact mods.
func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.SetSize(300, 300)

	fc := imagick.NewPixelWand()
	fc.SetColor("none")
	mw.NewImage(300, 300, fc)
	mw.SetImageArtifact("gradient:angle", "35")

	red := imagick.NewPixelWand()
	if ok := red.SetColor("#ff0000"); !ok {
		panic(mw.GetLastError())
	}

	green := imagick.NewPixelWand()
	if ok := green.SetColor("green"); !ok {
		panic(mw.GetLastError())
	}

	blue := imagick.NewPixelWand()
	if ok := blue.SetColor("blue"); !ok {
		panic(mw.GetLastError())
	}

	stops := []imagick.StopInfo{
		imagick.NewStopInfo(red.GetMagickColor(), 0),
		imagick.NewStopInfo(green.GetMagickColor(), 0.2),
		imagick.NewStopInfo(green.GetMagickColor(), 0.7),
		imagick.NewStopInfo(blue.GetMagickColor(), 1),
	}

	err := mw.GradientImage(imagick.GRADIENT_TYPE_LINEAR, imagick.SPREAD_METHOD_PAD, stops)
	if err != nil {
		panic(err)
	}

	mw.WriteImage("out.png")
	//mw.DisplayImage(os.Getenv("DISPLAY"))
}
