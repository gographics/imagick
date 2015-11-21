// Port of http://members.shaw.ca/el.supremo/MagickWand/round_mask.htm to Go
package main

import "github.com/gographics/imagick/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	lw := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()
	dw := imagick.NewDrawingWand()

	// Create the initial 640x480 transparent canvas
	pw.SetColor("none")
	mw.NewImage(640, 480, pw)

	pw.SetColor("white")
	dw.SetFillColor(pw)
	dw.RoundRectangle(15, 15, 624, 464, 15, 15)
	mw.DrawImage(dw)

	lw.ReadImage("logo:")
	// Note that MagickSetImageCompose is usually only used for the MagickMontageImage
	// function and isn't used or needed by MagickCompositeImage
	mw.CompositeImage(lw, imagick.COMPOSITE_OP_SRC_IN, 0, 0)

	/* Write the new image */
	mw.WriteImage("mask_result.png")
}
