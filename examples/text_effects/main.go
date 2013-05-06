// Port of http://members.shaw.ca/el.supremo/MagickWand/text_effects.htm to Go
package main

import "github.com/gographics/imagick/imagick"

// Given a pattern name (which MUST have a leading #) and a pattern file,
// set up a pattern URL for later reference in the specified drawing wand
// Currently only used in Text Effect 2
func setTilePattern(dw *imagick.DrawingWand, pattern_name, pattern_file string) {
	tw := imagick.NewMagickWand()
	defer tw.Destroy()

	tw.ReadImage(pattern_file)
	// Read the tile's width and height
	w := tw.GetImageWidth()
	h := tw.GetImageHeight()

	dw.PushPattern(pattern_name[1:], 0, 0, float64(w), float64(h))
	dw.Composite(imagick.COMPOSITE_OP_SRC_OVER, 0, 0, 0, 0, tw)
	dw.PopPattern()
	dw.SetFillPatternURL(pattern_name)
}

var dargs = []float64{120.}
var d_args = []float64{-0.02, 0.0, 0.0, 1.02, 0.0, 0.0, -0.5, 1.9}

func main() {
	textEffect1()
	textEffect2()
	textEffect3()
	textEffect4()
	textEffect5And6()
	textEffect7()
	textEffect8()
}

// Text effect 1 - shadow effect using MagickShadowImage
// NOTE - if an image has a transparent background, adding a border of any colour other
// than "none" will remove all the transparency and replace it with the border's colour
func textEffect1() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	pw.SetColor("none")
	// Create a new transparent image
	mw.NewImage(350, 100, pw)
	// Set up a 72 point white font
	pw.SetColor("white")
	dw.SetFillColor(pw)
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Add a black outline to the text
	pw.SetColor("black")
	dw.SetStrokeColor(pw)
	// Turn antialias on - not sure this makes a difference
	dw.SetTextAntialias(true)
	// Now draw the text
	dw.Annotation(25, 65, "Magick")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	// Trim the image down to include only the text
	mw.TrimImage(0)
	// equivalent to the command line +repage
	mw.ResetImagePage("")
	// Make a copy of the text image
	cw := mw.Clone()
	// Set the background colour to blue for the shadow
	pw.SetColor("blue")
	mw.SetImageBackgroundColor(pw)
	// Opacity is a real number indicating (apparently) percentage
	mw.ShadowImage(70, 4, 5, 5)
	// Composite the text on top of the shadow
	mw.CompositeImage(cw, imagick.COMPOSITE_OP_OVER, 5, 5)
	cw.Destroy()
	cw = imagick.NewMagickWand()
	defer cw.Destroy()
	// Create a new image the same size as the text image and put a solid colour
	// as its background
	pw.SetColor("rgb(125,215,255)")
	cw.NewImage(mw.GetImageWidth(), mw.GetImageHeight(), pw)
	// Now composite the shadowed text over the plain background
	cw.CompositeImage(mw, imagick.COMPOSITE_OP_OVER, 0, 0)
	// and write the result
	cw.WriteImage("text_shadow.png")
}

// Text effect 2 - tiled text using the builtin checkerboard pattern
// Anthony's Tiled Font effect
func textEffect2() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	setTilePattern(dw, "#check", "pattern:checkerboard")
	pw.SetColor("lightblue")
	// Create a new transparent image
	mw.NewImage(320, 100, pw)
	// Set up a 72 point font
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Now draw the text
	dw.Annotation(28, 68, "Magick")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	// Trim the image
	mw.TrimImage(0)
	// Add a transparent border
	pw.SetColor("lightblue")
	mw.BorderImage(pw, 5, 5)
	// and write it
	mw.WriteImage("text_pattern.png")
}

// Text effect 3 -  arc font (similar to http://www.imagemagick.org/Usage/fonts/#arc)
func textEffect3() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	// Create a 320x100 lightblue canvas
	pw.SetColor("lightblue")
	mw.NewImage(320, 100, pw)
	// Set up a 72 point font
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Now draw the text
	dw.Annotation(25, 65, "Magick")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	mw.DistortImage(imagick.DISTORTION_ARC, dargs, false)
	// Trim the image
	mw.TrimImage(0)
	// Add the border
	pw.SetColor("lightblue")
	mw.BorderImage(pw, 10, 10)
	// and write it
	mw.WriteImage("text_arc.png")
}

// Text effect 4 - bevelled font http://www.imagemagick.org/Usage/fonts/#bevel
func textEffect4() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	// Create a 320x100 canvas
	pw.SetColor("gray")
	mw.NewImage(320, 100, pw)
	// Set up a 72 point font
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Set up a 72 point white font
	pw.SetColor("white")
	dw.SetFillColor(pw)
	// Now draw the text
	dw.Annotation(25, 65, "Magick")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	// the "gray" parameter must be true to get the effect shown on Anthony's page
	mw.ShadeImage(true, 140, 60)
	pw.SetColor("yellow")
	dw.SetFillColor(pw)
	cpw := imagick.NewPixelWand()
	defer cpw.Destroy()
	cpw.SetColor("gold")
	mw.ColorizeImage(pw, cpw)
	// and write it
	mw.WriteImage("text_bevel.png")
}

// Text effect 5 and 6 - Plain text and then Barrel distortion
func textEffect5And6() {
	imagick.Initialize()
	defer imagick.Terminate()
	// This one uses d_args
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	// Create a 320x100 transparent canvas
	pw.SetColor("none")
	mw.NewImage(320, 100, pw)
	// Set up a 72 point font
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Now draw the text
	dw.Annotation(25, 65, "Magick")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	mw.WriteImage("text_plain.png")
	// Trim the image
	mw.TrimImage(0)
	// Add the border
	pw.SetColor("none")
	mw.BorderImage(pw, 10, 10)
	//mw.SetImageMatte(true)
	//mw.SetImageVirtualPixelMethod(TransparentVirtualPixelMethod)
	// 	d_args[0] = 0.1;d_args[1] = -0.25;d_args[2] = -0.25; [3] += .1
	// The first value should be positive. If it is negative the image is *really* distorted
	d_args[0] = 0.0
	d_args[1] = 0.0
	d_args[2] = 0.5
	// d_args[3] should normally be chosen such the sum of all 4 values is 1
	// so that the result is the same size as the original
	// You can override the sum with a different value
	// If the sum is greater than 1 the resulting image will be smaller than the original
	d_args[3] = 1 - (d_args[0] + d_args[1] + d_args[2])
	// Make the result image smaller so that it isn't as likely
	// to overflow the edges
	// d_args[3] += 0.1
	// 0.0,0.0,0.5,0.5,0.0,0.0,-0.5,1.9
	d_args[3] = 0.5
	d_args[4] = 0.0
	d_args[5] = 0.0
	d_args[6] = -0.5
	d_args[7] = 1.9
	// DON'T FORGET to set the correct number of arguments here
	mw.DistortImage(imagick.DISTORTION_BARREL, d_args, true)
	//mw.ResetImagePage("")
	// Trim the image again
	mw.TrimImage(0)
	// Add the border
	pw.SetColor("none")
	mw.BorderImage(pw, 10, 10)
	// and write it
	mw.WriteImage("text_barrel.png")
}

// Text effect 7 - Polar distortion
func textEffect7() {
	imagick.Initialize()
	defer imagick.Terminate()
	// This one uses d_args[0]
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	// Create a 320x200 transparent canvas
	pw.SetColor("none")
	mw.NewImage(320, 200, pw)
	// Set up a 72 point font
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Now draw the text
	dw.Annotation(25, 65, "Magick")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	d_args[0] = 0.0
	// DON'T FORGET to set the correct number of arguments here
	mw.DistortImage(imagick.DISTORTION_POLAR, d_args, true)
	//mw.ResetImagePage("")
	// Trim the image again
	mw.TrimImage(0)
	// Add the border
	pw.SetColor("none")
	mw.BorderImage(pw, 10, 10)
	// and write it
	mw.WriteImage("text_polar.png")
}

// Text effect 8 - Shepard's distortion
func textEffect8() {
	imagick.Initialize()
	defer imagick.Terminate()
	// This one uses d_args[0]
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	// Create a 320x200 transparent canvas
	pw.SetColor("none")
	mw.NewImage(640, 480, pw)
	// Set up a 72 point font
	dw.SetFont("Verdana-Bold-Italic")
	dw.SetFontSize(72)
	// Now draw the text
	dw.Annotation(50, 240, "Magick Rocks")
	// Draw the image on to the mw
	mw.DrawImage(dw)
	d_args[0] = 150.0
	d_args[1] = 190.0
	d_args[2] = 100.0
	d_args[3] = 290.0
	d_args[4] = 500.0
	d_args[5] = 200.0
	d_args[6] = 430.0
	d_args[7] = 130.0
	// DON'T FORGET to set the correct number of arguments here
	mw.DistortImage(imagick.DISTORTION_SHEPARDS, d_args, true)
	// Trim the image
	mw.TrimImage(0)
	// Add the border
	pw.SetColor("none")
	mw.BorderImage(pw, 10, 10)
	// and write it
	mw.WriteImage("text_shepards.png")
}
