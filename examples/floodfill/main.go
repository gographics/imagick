// Port of http://members.shaw.ca/el.supremo/MagickWand/floodfill.htm to Go
package main

import (
	"math"
	"os"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	var err error

	mw := imagick.NewMagickWand()

	// fillcolor and bordercolor
	fc, bc := imagick.NewPixelWand(), imagick.NewPixelWand()

	fc.SetColor("none")
	bc.SetColor("white")

	err = mw.ReadImage("logo:")
	if err != nil {
		panic(err)
	}

	mw.SetImageChannelMask(imagick.CHANNELS_RGB | imagick.CHANNEL_ALPHA)

	// Floodfill uses a fuzz param value that is dependent on the
	// quantum color depth of ImageMagick. So we have to convert
	// from 1% to a color depth fuzz value.
	_, depth := imagick.GetQuantumDepth()
	fuzz := math.Pow(2, float64(depth)) * .01

	// The bordercolor (with fuzz of 1% applied) is replaced by the fill
	// colour starting at the given coordinate - in this case 0, 0.
	// Normally the last argument is false so that the colours are matched
	// but if it is true then it floodfills any pixel that does *not* match
	// the target color
	mw.FloodfillPaintImage(fc, fuzz, bc, 0, 0, false)

	mw.DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}
}
