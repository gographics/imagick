// Port of http://members.shaw.ca/el.supremo/MagickWand/clipmask.htm to Go
package main

import (
	"github.com/gobinds/imagick/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	dest := imagick.NewMagickWand()
	defer dest.Destroy()

	src := imagick.NewMagickWand()
	defer src.Destroy()

	mask := imagick.NewMagickWand()
	defer mask.Destroy()

	dest.SetSize(100, 100)
	src.SetSize(100, 100)

	if err := dest.ReadImage("tile:tile_water.jpg"); err != nil {
		panic(err)
	}

	if err := mask.ReadImage("mask_bite.png"); err != nil {
		panic(err)
	}

	// When you create a mask, you use white for those parts that you want
	// to show through and black for those which must not show through.
	// But internally it's the opposite so the mask must be negated
	mask.NegateImage(false)
	dest.SetImageClipMask(mask)

	if err := src.ReadImage("tile:tile_disks.jpg"); err != nil {
		panic(err)
	}

	// This does the src (overlay) over the dest (background)
	dest.CompositeImage(src, imagick.COMPOSITE_OP_OVER, 0, 0)
	dest.WriteImage("clip_out.jpg")
}
