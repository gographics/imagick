// Port of http://members.shaw.ca/el.supremo/MagickWand/floodfill.htm to Go
package main

import (
	"github.com/gographics/imagick/imagick"
	"os"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	var err error

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// fillcolor and bordercolor
	fc, bc := imagick.NewPixelWand(), imagick.NewPixelWand()
	defer fc.Destroy()
	defer bc.Destroy()

	fc.SetColor("none")
	bc.SetColor("white")

	err = mw.ReadImage("logo:")
	if err != nil {
		panic(err)
	}

	rgba := imagick.RedChannel | imagick.GreenChannel | imagick.BlueChannel | imagick.AlphaChannel

	// The bordercolor (with fuzz of 20 applied) is replaced by the fill
	// colour starting at the given coordinate - in this case 0, 0.
	// Normally the last argument is false so that the colours are matched
	// but if it is true then it floodfills any pixel that does *not* match
	// the target color
	mw.FloodfillPaintImage(rgba, fc, 20, bc, 0, 0, false)

	mw.DisplayImage(os.Getenv("DYSPLAY"))
	if err != nil {
		panic(err)
	}
}
