// Port of http://members.shaw.ca/el.supremo/MagickWand/trans_paint.htm to Go
package main

import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.ReadImage("logo:")

	// A larger fuzz value allows more colours "near" white to be
	// modified. A fuzz of zero only allows an exact match with the
	// given colour
	// Set up the pixelwand containing the colour to be "targeted"
	// by transparency
	target := imagick.NewPixelWand()
	target.SetColor("white")
	// Change the transparency of all colours which match target (with
	// fuzz applied). In this case they are made completely transparent (0)
	// but you can set this to any value from 0 to 1.
	mw.TransparentPaintImage(target, 0, 10, false)
	mw.WriteImage("logo_white.png")
}
