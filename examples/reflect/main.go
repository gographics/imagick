// Port of http://members.shaw.ca/el.supremo/MagickWand/reflect.htm to Go
package main

import "gopkg.in/gographics/imagick.v1/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.ReadImage("logo:")

	// We know that logo: is 640x480 but in the general case
	// we need to get the dimensions of the image
	w := mw.GetImageWidth()
	h := mw.GetImageHeight()

	// +matte is the same as -alpha off
	// This does it the "new" way but if your IM doesn't have this
	// then MagickSetImageMatte(mw,MagickFalse); can be used
	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	// clone the input image
	mwr := mw.Clone()

	// Resize it
	mwr.ResizeImage(w, h/2, imagick.FILTER_LANCZOS, 1)
	// Flip the image over to form the reflection
	mwr.FlipImage()

	// Create the gradient image which will be used as the alpha
	// channel in the reflection image
	mwg := imagick.NewMagickWand()
	mwg.SetSize(w, h/2)
	mwg.ReadImage("gradient:white-black")

	// Copy the gradient in to the alpha channel of the reflection image
	mwr.CompositeImage(mwg, imagick.COMPOSITE_OP_COPY_OPACITY, 0, 0)

	// Add the reflection image to the wand which holds the original image
	mw.AddImage(mwr)

	// Append the reflection to the bottom (MagickTrue) of the original image
	mwout := mw.AppendImages(true)

	// and save the result
	mwout.WriteImage("logo_reflect.png")
}
