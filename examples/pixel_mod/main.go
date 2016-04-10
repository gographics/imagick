// Port of http://members.shaw.ca/el.supremo/MagickWand/pixel_mod.htm to Go
package main

import "gopkg.in/gographics/imagick.v2/imagick"

func main() {
	useDraw()
	usePixelIterator()
}

func useDraw() {
	imagick.Initialize()
	defer imagick.Terminate()

	/* Create a wand */
	mw := imagick.NewMagickWand()

	/* Read the input image */
	mw.ReadImage("logo:")
	fill := imagick.NewPixelWand()
	dw := imagick.NewDrawingWand()

	// Set the fill to "red" or you can do the same thing with this:
	fill.SetColor("red")
	dw.SetFillColor(fill)
	// Uses the current Fill as the colour of the point at 200,100
	dw.Point(200, 100)
	mw.DrawImage(dw)
	/* write it */
	mw.WriteImage("logo_pixel_drawingwand.gif")
}

func usePixelIterator() {
	imagick.Initialize()
	defer imagick.Terminate()

	/* Create a wand */
	mw := imagick.NewMagickWand()

	/* Read the input image */
	mw.ReadImage("logo:")
	// Get a one-pixel region at coordinate 200,100
	iterator := mw.NewPixelRegionIterator(200, 100, 1, 1)
	pixels := iterator.GetNextIteratorRow()
	// Modify the pixel
	pixels[0].SetColor("red")
	// then sync it back into the wand
	iterator.SyncIterator()
	mw.WriteImage("logo_pixel_iterator.gif")
}
