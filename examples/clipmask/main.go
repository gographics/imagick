// Port of http://members.shaw.ca/el.supremo/MagickWand/clipmask.htm to Go
package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	dest := imagick.NewMagickWand()
	src := imagick.NewMagickWand()
	mask := imagick.NewMagickWand()

	dest.SetSize(100, 100)
	src.SetSize(100, 100)

	if err := dest.ReadImage("tile:tile_water.jpg"); err != nil {
		panic(err)
	}

	if err := mask.ReadImage("mask_bite.png"); err != nil {
		panic(err)
	}

	// IM7 got rid of clip masks in the API, so we have to approach this by
	// way of using alpha channels.

	// We want to cut out a hole with the mask within the middle of the image.
	// This will allow the src image to be filled within the masked hole.
	mask.NegateImage(false)
	dest.CompositeImage(mask, imagick.COMPOSITE_OP_COPY_ALPHA, false, 0, 0)

	if err := src.ReadImage("tile:tile_disks.jpg"); err != nil {
		panic(err)
	}

	// This does the dst over the src
	dest.CompositeImage(src, imagick.COMPOSITE_OP_DST_OVER, false, 0, 0)
	dest.WriteImage("clip_out.jpg")
}
