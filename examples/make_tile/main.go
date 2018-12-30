// Port of http://members.shaw.ca/el.supremo/MagickWand/make_tile.htm to Go
package main

import "gopkg.in/gographics/imagick.v3/imagick"

// make-tile creates a tileable image from an input image.
// ( +clone -flop ) +append  ( +clone -flip ) -append -resize 50%
func make_tile(mw *imagick.MagickWand, outfile string) {
	mwc := mw.Clone()
	mwc.FlopImage()
	mw.AddImage(mwc)
	mwc.Destroy()
	mwc = mw.AppendImages(false)
	mwf := mwc.Clone()
	mwf.FlipImage()
	mwc.AddImage(mwf)
	mwf.Destroy()
	mwf = mwc.AppendImages(true)

	w := mwf.GetImageWidth()
	h := mwf.GetImageHeight()
	// 1 = Don't blur or sharpen image
	mwf.ResizeImage(w/2, h/2, imagick.FILTER_LANCZOS)
	mwf.WriteImage(outfile)
	mwf.Destroy()
	mwc.Destroy()
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.SetSize(100, 100)
	mw.ReadImage("plasma:red-yellow")
	make_tile(mw, "tile_plasma.png")
	mw.Destroy()

	mw = imagick.NewMagickWand()
	mw.SetSize(100, 100)
	mw.ReadImage("xc:")
	mw.AddNoiseImage(imagick.NOISE_RANDOM, 0)
	mw.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TILE)
	mw.BlurImage(0, 10)
	mw.NormalizeImage()
	make_tile(mw, "tile_random.png")
	mw.Destroy()
}
