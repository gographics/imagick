// Port from http://members.shaw.ca/el.supremo/MagickWand/tilt_shift.htm to Go
package main

import "github.com/gographics/imagick/imagick"

func main() {
	// Note that the colours are stored as separate *normalized* RGB components
	arglist := []float64{
		0, 0,
		// RGB black
		0, 0, 0,
		// The y coordinate is filled in later
		0, -1,
		// RGB white
		1, 1, 1,
	}
	// arguments for MagickFunctionImage
	funclist := []float64{4, -4, 1}
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	mw.ReadImage("beijing_md.jpg")
	// fill in the Y coordinate now that we can get the image dimensions
	arglist[6] = float64(mw.GetImageHeight() - 1)
	mw.SigmoidalContrastImage(true, 15, imagick.QUANTUM_RANGE*30/100)
	cw := mw.Clone()
	defer cw.Destroy()
	cw.SparseColorImage(imagick.CHANNELS_RGB, imagick.INTERPOLATE_BARYCENTRIC_COLOR, arglist)
	// Do the polynomial function
	cw.FunctionImage(imagick.FUNCTION_POLYNOMIAL, funclist)
	// -set option:compose:args 15
	if err := cw.SetImageArtifact("compose:args", "15"); err != nil {
		panic(err)
	}
	mw.CompositeImage(cw, imagick.COMPOSITE_OP_BLUR, 0, 0)
	mw.WriteImage("beijing_model.jpg")
}
