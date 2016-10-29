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
	mwb := imagick.NewMagickWand()

	// Set the hsl and hsb images to the same size as the input image
	mwl.SetSize(width, height)
	mwb.SetSize(width, height)
	// Even though we won't be reading these images they must be initialized
	// to something
	mwb.ReadImage("xc:none")
	mwl.ReadImage("xc:none")

	// Create iterators for each image
	imw := mw.NewPixelIterator()
	imwl := mwl.NewPixelIterator()
	imwb := mwb.NewPixelIterator()

	for y := 0; y < int(height); y++ {
		// Get the next row from each image
		pmw := imw.GetNextIteratorRow()
		pmwl := imwl.GetNextIteratorRow()
		pmwb := imwb.GetNextIteratorRow()
		for x := 0; x < int(width); x++ {
			// Get the RGB quanta from the source image
			qr := pmw[x].GetRedQuantum()
			qg := pmw[x].GetGreenQuantum()
			qb := pmw[x].GetBlueQuantum()

			// Convert the source quanta to HSB
			bh, bs, bb := imagick.ConvertRGBToHSB(qr, qg, qb)
			bb *= 0.5
			qrb, qgb, qbb := imagick.ConvertHSBToRGB(bh, bs, bb)
			// Set the pixel in the HSB output image
			pmwb[x].SetRedQuantum(qrb)
			pmwb[x].SetGreenQuantum(qgb)
			pmwb[x].SetBlueQuantum(qbb)

			// Convert the source quanta to HSL
			lh, ls, ll := imagick.ConvertRGBToHSL(qr, qg, qb)
			ll *= 0.5
			qrl, qgl, qbl := imagick.ConvertHSLToRGB(lh, ls, ll)
			// Set the pixel in the HSL output image
			pmwl[x].SetRedQuantum(qrl)
			pmwl[x].SetGreenQuantum(qgl)
			pmwl[x].SetBlueQuantum(qbl)
		}
		// Sync writes the pixels back to the magick wands
		imwl.SyncIterator()
		imwl.SyncIterator()
	}
	// write the results
	mwb.WriteImage("logo_hsb.jpg")
	mwl.WriteImage("logo_hsl.jpg")

	// Clean up the iterators and magick wands
	imw.Destroy()
	imwl.Destroy()
	imwb.Destroy()
	mw.Destroy()
	mwl.Destroy()
	mwb.Destroy()
}
