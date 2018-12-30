// Port of http://members.shaw.ca/el.supremo/MagickWand/modulate.htm to Go
package main

import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()

	mw.ReadImage("logo:")
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	mwl := imagick.NewMagickWand()

	// Set the hsl image to the same size as the input image
	mwl.SetSize(width, height)
	// Even though we won't be reading the image it must be initialized
	// to something
	mwl.ReadImage("xc:none")

	// Create iterators for each image
	imw := mw.NewPixelIterator()
	imwl := mwl.NewPixelIterator()

	for y := 0; y < int(height); y++ {
		// Get the next row from each image
		pmw := imw.GetNextIteratorRow()
		pmwl := imwl.GetNextIteratorRow()
		for x := 0; x < int(width); x++ {
			// Get the RGB quanta from the source image
			qr := float64(pmw[x].GetRedQuantum())
			qg := float64(pmw[x].GetGreenQuantum())
			qb := float64(pmw[x].GetBlueQuantum())

			// Convert the source quanta to HSL
			lh, ls, ll := imagick.ConvertRGBToHSL(qr, qg, qb)
			ll *= 0.5
			qrl, qgl, qbl := imagick.ConvertHSLToRGB(lh, ls, ll)
			// Set the pixel in the HSL output image
			pmwl[x].SetRedQuantum(imagick.Quantum(qrl))
			pmwl[x].SetGreenQuantum(imagick.Quantum(qgl))
			pmwl[x].SetBlueQuantum(imagick.Quantum(qbl))
		}
		// Sync writes the pixels back to the magick wands
		imwl.SyncIterator()
	}
	// write the results
	mwl.WriteImage("logo_hsl.jpg")

	// Clean up the iterators and magick wands
	imw.Destroy()
	imwl.Destroy()
	mw.Destroy()
	mwl.Destroy()
}
