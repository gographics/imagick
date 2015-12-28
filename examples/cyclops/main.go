// Port of http://members.shaw.ca/el.supremo/MagickWand/cyclops.htm to Go
package main

import (
	"os"
	
	"github.com/gographics/imagick/imagick"
)

func main() {
	var err error
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	bg := imagick.NewPixelWand()
	fg := imagick.NewPixelWand()

	err = mw.ReadImage("cyclops_sm.gif")
	if err != nil {
		panic(err)
	}

	bg.SetColor("white")
	mw.BorderImage(bg, 1, 1)
	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_SET)

	fg.SetColor("none")
	channel := imagick.CHANNELS_RGB | imagick.CHANNEL_ALPHA

	// Floodfill the "background" colour with the "foreground" colour
	// starting at coordinate 0,0 using a fuzz of 20
	mw.FloodfillPaintImage(channel, fg, 20, bg, 0, 0, false)
	mw.ShaveImage(1, 1)

	mw.DisplayImage(os.Getenv("DISPLAY"))
	if err != nil {
		panic(err)
	}
}
