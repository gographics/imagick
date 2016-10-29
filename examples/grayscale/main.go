// Port of http://members.shaw.ca/el.supremo/MagickWand/grayscale.htm to Go
package main

import (
	"fmt"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	pw := imagick.NewPixelWand()
	pw.SetColor("white")

	mw := imagick.NewMagickWand()

	// Create a 100x100 image with a default of white
	mw.NewImage(100, 100, pw)

	// Get a new pixel iterator
	iterator := mw.NewPixelIterator()

	for y := 0; y < int(mw.GetImageHeight()); y++ {
		// Get the next row of the image as an array of PixelWands
		pixels := iterator.GetNextIteratorRow()
		if len(pixels) == 0 {
			break
		}
		// Set the row of wands to a simple gray scale gradient
		for x, pixel := range pixels {
			if !pixel.IsVerified() {
				panic("unverified pixel")
			}
			gray := x * 255 / 100
			hex := fmt.Sprintf("#%02x%02x%02x", gray, gray, gray)
			if ret := pixel.SetColor(hex); !ret {
				panic("Could not set color in pixel")
			}
		}
		// Sync writes the pixels back to the mw
		if err := iterator.SyncIterator(); err != nil {
			panic(err)
		}
	}

	mw.WriteImage("bits_demo.gif")
}
