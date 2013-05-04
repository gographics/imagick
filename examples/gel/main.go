// Port of http://members.shaw.ca/el.supremo/MagickWand/gel.htm to Go
package main

import (
	"github.com/gobinds/imagick/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	/* First step is to create the gel shape:
	convert -size 100x60 xc:none \
	          -fill red -draw 'circle    25,30  10,30' \
	                    -draw 'circle    75,30  90,30' \
	                    -draw 'rectangle 25,15  75,45' \
	          gel_shape.png
	*/

	/* Create a wand */
	mw := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()
	dw := imagick.NewDrawingWand()

	mw.SetSize(100, 60)
	mw.ReadImage("xc:none")

	pw.SetColor("red")
	dw.SetFillColor(pw)
	dw.Circle(25, 30, 10, 30)
	dw.Circle(75, 30, 90, 30)
	dw.Rectangle(25, 15, 75, 45)

	// Now we draw the Drawing wand on to the Magick Wand
	mw.DrawImage(dw)
	mw.WriteImage("gel_shape.png")

	dw.Destroy()
	pw.Destroy()
	mw.Destroy()

	/* Next step is to create the gel highlight:
	convert gel_shape.png \
	            \( +clone -fx A +matte  -blur 0x12  -shade 110x0 -normalize \
	               -sigmoidal-contrast 16,60% -evaluate multiply .5 \
	               -roll +5+10 +clone -compose Screen -composite \) \
	            -compose In  -composite  gel_highlight.png
	*/

	mw = imagick.NewMagickWand()
	mw.ReadImage("gel_shape.png")

	mwc := mw.Clone()
	mwf, err := mwc.FxImage("A")
	if err != nil {
		panic(err)
	}

	mwf.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	mwf.BlurImage(0, 12)
	mwf.ShadeImage(true, 110, 0)
	mwf.NormalizeImage()

	// The last argument is specified as a percentage on the command line
	// but is specified to the function as a percentage of the QuantumRange
	mwf.SigmoidalContrastImage(true, 16, 0.6*imagick.QUANTUM_RANGE)
	mwf.EvaluateImage(imagick.EVAL_OP_MULTIPLY, 0.5)
	mwf.RollImage(5, 10)

	mwc.Destroy()
	// The +clone operation copies the original but only so that
	// it can be used in the following composite operation, so we don't
	// actually need to do a clone, just reference the original image.
	mwf.CompositeImage(mw, imagick.COMPOSITE_OP_SCREEN, 0, 0)

	mw.CompositeImage(mwf, imagick.COMPOSITE_OP_IN, 0, 0)
	mw.WriteImage("gel_highlight.png")

	mw.Destroy()
	mwc.Destroy()
	mwf.Destroy()

	// Now create the gel border
	/*
		convert gel_highlight.png \
		            \( +clone -fx A  +matte -blur 0x2  -shade 0x90 -normalize \
		               -blur 0x2  -negate -evaluate multiply .4 -negate -roll -.5-1 \
		               +clone  -compose Multiply -composite \) \
		            -compose In  -composite  gel_border.png

	*/

	mw = imagick.NewMagickWand()
	mw.ReadImage("gel_highlight.png")
	mwc = mw.Clone()
	mwf, err = mwc.FxImage("A")
	if err != nil {
		panic(err)
	}
	mwf.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	mwf.BlurImage(0, 2)
	mwf.ShadeImage(true, 0, 90)
	mwf.NormalizeImage()
	mwf.BlurImage(0, 2)
	mwf.NegateImage(false)
	mwf.EvaluateImage(imagick.EVAL_OP_MULTIPLY, 0.4)
	mwf.NegateImage(false)
	mwf.RollImage(-1, -1)
	mwf.CompositeImage(mw, imagick.COMPOSITE_OP_MULTIPLY, 0, 0)
	mw.CompositeImage(mwf, imagick.COMPOSITE_OP_IN, 0, 0)
	mw.WriteImage("gel_border.png")

	mw.Destroy()
	mwc.Destroy()
	mwf.Destroy()

	// and finally the text and shadow effect
	/*
		convert gel_border.png \
		            -font Candice  -pointsize 24  -fill white  -stroke black \
		            -gravity Center  -annotate 0 "Gel"  -trim -repage 0x0+4+4 \
		            \( +clone -background navy -shadow 80x4+4+4 \) +swap \
		            -background none  -flatten    gel_button.png
	*/

	mw = imagick.NewMagickWand()
	dw = imagick.NewDrawingWand()
	pw = imagick.NewPixelWand()

	mw.ReadImage("gel_border.png")
	dw.SetFont("Lucida-Handwriting-Italic")
	dw.SetFontSize(24)
	pw.SetColor("white")
	dw.SetFillColor(pw)
	pw.SetColor("black")
	dw.SetStrokeColor(pw)
	dw.SetGravity(imagick.GRAVITY_CENTER)
	// It is important to notice here that MagickAnnotateImage renders the text on
	// to the MagickWand, NOT the DrawingWand. It only uses the DrawingWand for font
	// and colour information etc.
	mw.AnnotateImage(dw, 0, 0, 0, "Gel")
	mw.TrimImage(0)
	mw.ResetImagePage("0x0+4+4")
	mwc = mw.Clone()
	pw.SetColor("navy")
	mwc.SetImageBackgroundColor(pw)
	mwc.ShadowImage(80, 4, 4, 4)
	mwf = imagick.NewMagickWand()
	mwf.AddImage(mwc)
	mwf.AddImage(mw)

	mw.Destroy()
	mwc.Destroy()

	pw.SetColor("none")
	mwf.SetImageBackgroundColor(pw)
	mw = mwf.MergeImageLayers(imagick.IMAGE_LAYER_FLATTEN)
	mw.WriteImage("gel_button.png")

	mw.Destroy()
	mwc.Destroy()
	mwf.Destroy()
	dw.Destroy()
	pw.Destroy()
}
