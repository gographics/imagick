// Port of http://members.shaw.ca/el.supremo/MagickWand/cyclops.htm to Go
package main

import (
	"github.com/gographics/imagick/imagick"
	"os"
)

func main() {
	var err error
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	bg := imagick.NewPixelWand()
	defer bg.Destroy()
	fg := imagick.NewPixelWand()
	defer fg.Destroy()

	err = mw.ReadImage("http://www.imagemagick.org/Usage/images/cyclops_sm.gif")
	if err != nil {
		panic(err)
	}

	bg.SetColor("white")
	mw.BorderImage(bg, 1, 1)
	mw.SetImageAlphaChannel(imagick.SetAlphaChannel)

	fg.SetColor("none")
	channel := imagick.RGBChannels | imagick.AlphaChannel

	// Floodfill the "background" colour with the "foreground" colour
	// starting at coordinate 0,0 using a fuzz of 20
	mw.FloodfillPaintImage(channel, fg, 20, bg, 0, 0, false)
	mw.ShaveImage(1, 1)

	mw.DisplayImage(os.Getenv("DYSPLAY"))
	if err != nil {
		panic(err)
	}
}
