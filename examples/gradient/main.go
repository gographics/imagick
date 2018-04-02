package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

// Demonstrate creating a gradient image including artifact mods.
func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.SetSize(300,300)

	fc := imagick.NewPixelWand()
	fc.SetColor("none")
	mw.NewImage(300, 300, fc)
	mw.SetImageArtifact("gradient:angle", "35")

	mw.GradientImage(imagick.GRADIENT_TYPE_LINEAR, imagick.SPREAD_METHOD_PAD, "#ff0000", "#0000ff")

	mw.WriteImage("out.png")

}
