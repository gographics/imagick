// Port of http://members.shaw.ca/el.supremo/MagickWand/cyclops.htm to Go
package main

import (
	"math"
	"os"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	var err error
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	bg := imagick.NewPixelWand()
	fg := imagick.NewPixelWand()

	err = mw.ReadImage("cyclops_sm.gif")
	if err != nil {
		panic(err)
	}

	bg.SetColor("white")
	mw.SetBackgroundColor(bg)
	mw.BorderImage(bg, 1, 1, imagick.COMPOSITE_OP_COPY)

	fg.SetColor("none")

	// Floodfill uses a fuzz param value that is dependent on the
	// quantum color depth of ImageMagick. So we have to convert
	// from 10% to a color depth fuzz value.
	_, depth := imagick.GetQuantumDepth()
	fuzz := math.Pow(2, float64(depth)) * .10

	// Floodfill the "background" colour with the "foreground" colour
	// starting at coordinate 0,0 using a fuzz of 20
	mw.FloodfillPaintImage(fg, fuzz, bg, 0, 0, false)
	mw.ShaveImage(1, 1)

	mw.DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}
}
