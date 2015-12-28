// Port of http://members.shaw.ca/el.supremo/MagickWand/3dlogo.htm to Go
package main

import (
	"os"

	"github.com/gographics/imagick/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()
	dw := imagick.NewDrawingWand()

	if err := mw.SetSize(170, 100); err != nil {
		panic(err)
	}
	if err := mw.ReadImage("xc:black"); err != nil {
		panic(err)
	}

	pw.SetColor("white")
	dw.SetFillColor(pw)
	dw.Circle(50, 50, 13, 50)
	dw.Circle(120, 50, 157, 50)
	dw.Rectangle(50, 13, 120, 87)

	pw.SetColor("black")
	dw.SetFillColor(pw)
	dw.Circle(50, 50, 25, 50)
	dw.Circle(120, 50, 145, 50)
	dw.Rectangle(50, 25, 120, 75)

	pw.SetColor("white")
	dw.SetFillColor(pw)
	dw.Circle(60, 50, 40, 50)
	dw.Circle(110, 50, 130, 50)
	dw.Rectangle(60, 30, 110, 70)

	// Now we draw the Drawing wand on to the Magick Wand
	if err := mw.DrawImage(dw); err != nil {
		panic(err)
	}
	if err := mw.GaussianBlurImage(1, 1); err != nil {
		panic(err)
	}
	// Turn the matte of == +matte
	if err := mw.SetImageMatte(false); err != nil {
		panic(err)
	}
	if err := mw.WriteImage("logo_mask.png"); err != nil {
		panic(err)
	}

	mw.Destroy()
	dw.Destroy()
	pw.Destroy()

	mw = imagick.NewMagickWand()
	pw = imagick.NewPixelWand()
	dw = imagick.NewDrawingWand()

	mwc := imagick.NewMagickWand()

	mw.ReadImage("logo_mask.png")

	pw.SetColor("red")
	dw.SetFillColor(pw)

	dw.Color(0, 0, imagick.PAINT_METHOD_RESET)
	mw.DrawImage(dw)

	mwc.ReadImage("logo_mask.png")
	mwc.SetImageMatte(false)
	mw.CompositeImage(mwc, imagick.COMPOSITE_OP_COPY_OPACITY, 0, 0)

	// Annotate gets all the font information from the drawingwand
	// but draws the text on the magickwand
	// Get the first available "*Sans*" font
	fonts := mw.QueryFonts("*Sans*")
	dw.SetFont(fonts[0])
	dw.SetFontSize(36)
	pw.SetColor("white")
	dw.SetFillColor(pw)
	pw.SetColor("black")
	dw.SetStrokeColor(pw)
	dw.SetGravity(imagick.GRAVITY_CENTER)
	mw.AnnotateImage(dw, 0, 0, 0, "Ant")
	mw.WriteImage("logo_ant.png")

	mwc.Destroy()
	mw.Destroy()

	mw = imagick.NewMagickWand()

	if err := mw.ReadImage("logo_ant.png"); err != nil {
		panic(err)
	}

	mwf, err := mw.FxImage("A")
	if err != nil {
		panic(err)
	}

	//mw.SetImageMatte(false)

	// +matte is the same as -alpha off
	mwf.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	mwf.BlurImage(0, 6)
	mwf.ShadeImage(true, 110, 30)
	mwf.NormalizeImage()

	// ant.png  -compose Overlay -composite
	mwc = imagick.NewMagickWand()
	mwc.ReadImage("logo_ant.png")
	mwf.CompositeImage(mwc, imagick.COMPOSITE_OP_OVERLAY, 0, 0)
	mwc.Destroy()

	// ant.png  -matte  -compose Dst_In  -composite
	mwc = imagick.NewMagickWand()
	mwc.ReadImage("logo_ant.png")

	// -matte is the same as -alpha on
	// I don't understand why the -matte in the command line
	// does NOT operate on the image just read in (logo_ant.png in mwc)
	// but on the image before it in the list
	// It would appear that the -matte affects each wand currently in the
	// command list because applying it to both wands gives the same result
	mwf.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_SET)
	mwf.CompositeImage(mwc, imagick.COMPOSITE_OP_DST_IN, 0, 0)

	mwf.WriteImage("logo_ant_3D.png")

	mw.Destroy()
	mwc.Destroy()
	mwf.Destroy()

	/* Now for the shadow
	   convert ant_3D.png \( +clone -background navy -shadow 80x4+6+6 \) +swap \
	             -background none  -layers merge +repage ant_3D_shadowed.png
	*/

	mw = imagick.NewMagickWand()
	mw.ReadImage("logo_ant_3D.png")

	mwc = mw.Clone()

	pw.SetColor("navy")
	mwc.SetImageBackgroundColor(pw)

	mwc.ShadowImage(80, 4, 6, 6)

	// at this point
	// mw = ant_3D.png
	// mwc = +clone -background navy -shadow 80x4+6+6
	// To do the +swap I create a new blank MagickWand and then
	// put mwc and mw into it. ImageMagick probably doesn't do it
	// this way but it works here and that's good enough for me!
	mwf = imagick.NewMagickWand()
	mwf.AddImage(mwc)
	mwf.AddImage(mw)
	mwc.Destroy()

	pw.SetColor("none")
	mwf.SetImageBackgroundColor(pw)
	mwc = mwf.MergeImageLayers(imagick.IMAGE_LAYER_MERGE)
	mwc.WriteImage("logo_shadow_3D.png")

	mw.Destroy()
	mwc.Destroy()
	mwf.Destroy()

	/*
	   and now for the fancy background
	   convert ant_3D_shadowed.png \
	             \( +clone +repage +matte   -fx 'rand()' -shade 120x30 \
	                -fill grey70 -colorize 60 \
	                -fill lavender -tint 100 \) -insert 0 \
	             -flatten  ant_3D_bg.jpg
	*/
	mw = imagick.NewMagickWand()
	mw.ReadImage("logo_shadow_3D.png")

	mwc = mw.Clone()

	// +repage
	mwc.ResetImagePage("")

	// +matte is the same as -alpha off
	mwc.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	mwf, err = mwc.FxImage("rand()")
	if err != nil {
		panic(err)
	}

	mwf.ShadeImage(true, 120, 30)
	pw.SetColor("grey70")

	// It seems that this must be a separate pixelwand for Colorize to work!
	pwo := imagick.NewPixelWand()

	// AHA .. this is how to do a 60% colorize
	pwo.SetColor("rgb(60%,60%,60%)")
	mwf.ColorizeImage(pw, pwo)

	pw.SetColor("lavender")

	// and this is a 100% tint
	pwo.SetColor("rgb(100%,100%,100%)")
	mwf.TintImage(pw, pwo)

	mwc.Destroy()

	mwc = imagick.NewMagickWand()
	mwc.AddImage(mwf)
	mwc.AddImage(mw)

	mwf = mwc.MergeImageLayers(imagick.IMAGE_LAYER_FLATTEN)

	if err := mwf.DisplayImage(os.Getenv("DISPLAY")); err != nil {
		panic(err)
	}
}
