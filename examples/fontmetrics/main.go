// Port of http://members.shaw.ca/el.supremo/MagickWand/fontmetrics.htm to Go
package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

// Set up the drawingwand "dw" for the given font name, font size, and colour.
// If the font or size changes sx get the new width of a space
// (the magickwand is required if it is necessary to query the font metrics)
func draw_setfont(mw *imagick.MagickWand, dw *imagick.DrawingWand, font string, size float64, colour string, sx *float64) {
	sflag := false

	if len(font) > 0 {
		dw.SetFont(font)
		sflag = true
	}

	if len(colour) > 0 {
		pw := imagick.NewPixelWand()
		pw.SetColor(colour)
		dw.SetFillColor(pw)
		pw.Destroy()
		sflag = true
	}

	if size > 0 {
		dw.SetFontSize(size)
	}

	// If either the font or the fontsize (or both) have changed
	// we need to get the size of a space again
	if sflag {
		fm := mw.QueryFontMetrics(dw, " ")
		*sx = fm.TextWidth
	}
}

// sx is the width of a space in the current font and fontsize.
// If the font or fontsize is changed a new value for the space
// width must be obtained before calling this again (by calling draw_setfont)

// The easiest (?) way to handle vertical text placement is
// not to use gravity at all because then you know that all the text is
// placed relative to ImageMagick's (0,0) coordinate which is the top left of
// the screen and the baseline will be y=0.
func draw_metrics(mw *imagick.MagickWand, dw *imagick.DrawingWand, dx *float64, dy, sx float64, text string) {
	mw.AnnotateImage(dw, *dx, dy, 0, text)
	mw.DrawImage(dw)

	// get the font metrics
	fm := mw.QueryFontMetrics(dw, text)
	if fm != nil {
		// Adjust the new x coordinate
		*dx += fm.TextWidth + sx
	}
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	// Current coordinates of text
	var dx, dy float64
	// Width of a space in current font/size
	var sx float64

	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()

	// Set the size of the image
	mw.SetSize(300, 100)
	//mw.SetImageAlphaChannel(imagick.CHANNEL_ALPHA)
	mw.ReadImage("xc:white")

	// DO NOT SET GRAVITY - it makes text placement more complicated
	// (unless it does exactly what you wanted anyway).

	// Start near the left edge
	dx = 10

	// If we know the largest font we're using, we can set the y coordinate
	// fairly accurately. In this case it is the 72 point Times font, so to
	// place the text such that the largest almost touches the top of the image
	// we just add the text height to the descender to give the coordinate for
	// our baseline.
	// In this case the largest is the 72 point Times-New-Roman so I'll use that
	// to compute the baseline
	dw.SetFontSize(72)
	dw.SetFont("Times-New-Roman")
	fm := mw.QueryFontMetrics(dw, "M")
	dy = fm.CharacterHeight + fm.Descender
	// Note that we must free up the fontmetric array once we're done with it

	// this
	draw_setfont(mw, dw, "Arial", 48, "#40FF80", &sx)
	draw_metrics(mw, dw, &dx, dy, sx, "this")

	// is
	// A NULL signals to draw_setfont that the font stays the same
	draw_setfont(mw, dw, "", 24, "#8040BF", &sx)
	draw_metrics(mw, dw, &dx, dy, sx, "is")

	// my
	draw_setfont(mw, dw, "Times-New-Roman", 18, "#BF00BF", &sx)
	draw_metrics(mw, dw, &dx, dy, sx, "my")

	// text
	draw_setfont(mw, dw, "", 72, "#0F0FBF", &sx)
	draw_metrics(mw, dw, &dx, dy, sx, "text")
	mw.DrawImage(dw)

	// Now write the magickwand image
	mw.WriteImage("metric1.gif")
}
