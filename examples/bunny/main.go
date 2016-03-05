// Port of lmembers.shaw.ca/el.supremo/MagickWand/bunny.htm to Go
package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	aw := imagick.NewMagickWand()
	defer aw.Destroy()

	// Read the first input image
	if err := mw.ReadImage("bunny_grass.gif"); err != nil {
		panic(err)
	}

	// We need a separate wand to do this bit in parentheses
	if err := aw.ReadImage("bunny_anim.gif"); err != nil {
		panic(err)
	}
	aw.ResetImagePage("0x0+5+15!")

	// Now we have to add the images in the aw wand on to the end
	// of the mw wand.
	mw.AddImage(aw)
	// We can now destroy the aw wand so that it can be used
	// for the next operation
	aw.Destroy()
	// -coalesce
	aw = mw.CoalesceImages()

	// do "-delete 0" by copying the images from the "aw" wand to
	// the "mw" wand but omit the first one
	// free up the mw wand and recreate it for this step
	mw.Destroy()
	mw = imagick.NewMagickWand()

	for i := 1; i < int(aw.GetNumberImages()); i++ {
		aw.SetIteratorIndex(i)
		tw := aw.GetImage()
		mw.AddImage(tw)
		tw.Destroy()
	}
	mw.ResetIterator()
	aw.Destroy()

	// -deconstruct
	// Anthony says that MagickDeconstructImages is equivalent
	// to MagickCompareImagesLayers so we'll use that
	aw = mw.CompareImageLayers(imagick.IMAGE_LAYER_COMPARE_ANY)
	// -loop 0
	aw.SetOption("loop", "0")

	// write the images into one file
	if err := aw.WriteImages("bunny_bgnd.gif", true); err != nil {
		panic(err)
	}
}
