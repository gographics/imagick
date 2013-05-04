// Port of http://members.shaw.ca/el.supremo/MagickWand/affine.htm to Go
package main

import (
	"github.com/gobinds/imagick/imagick"
	"math"
)

func DegreesToRadians(deg float64) (rad float64) {
	rad = deg * math.Pi / 180.0
	return
}

func RadiansToDegrees(rad float64) (deg float64) {
	deg = rad * 180.0 / math.Pi
	return
}

// Each of these affine functions assumes that the input arrays
// have been correctly defined to be arrays of 6 (or more) doubles

// Initialize an affine array to arbitrary values
func set_affine(a []float64, sx, rx, ry, sy, tx, ty float64) []float64 {
	a[0] = sx
	a[1] = rx
	a[2] = ry
	a[3] = sy
	a[4] = tx
	a[5] = ty
	return a
}

// Set the affine array to translate by (x,y)
func set_translate_affine(t []float64, x, y float64) []float64 {
	t[0] = 1
	t[1] = 0
	t[2] = 0
	t[3] = 1
	t[4] = x
	t[5] = y
	return t
}

// Set the affine array to scale the image by sx,sy
func set_scale_affine(s []float64, sx, sy float64) []float64 {
	s[0] = sx
	s[1] = 0
	s[2] = 0
	s[3] = sy
	s[4] = 0
	s[5] = 0
	return s
}

// Set the affine array to rotate image by 'degrees' clockwise
func set_rotate_affine(r []float64, degrees float64) []float64 {
	r[0] = math.Cos(DegreesToRadians(math.Mod(degrees, 360.0)))
	r[1] = math.Sin(DegreesToRadians(math.Mod(degrees, 360.0)))
	r[2] = -math.Sin(DegreesToRadians(math.Mod(degrees, 360.0)))
	r[3] = math.Cos(DegreesToRadians(math.Mod(degrees, 360.0)))
	r[4] = 0
	r[5] = 0
	return r
}

// Multiply two affine arrays and return the result.
func affine_multiply(c, a, b []float64) []float64 {
	c[0] = a[0]*b[0] + a[1]*b[2]
	c[1] = a[0]*b[1] + a[1]*b[3]
	c[2] = a[2]*b[0] + a[3]*b[2]
	c[3] = a[2]*b[1] + a[3]*b[3]
	c[4] = a[4]*b[0] + a[5]*b[2] + b[4]
	c[5] = a[4]*b[1] + a[5]*b[3] + b[5]
	return c
}

// Remember that these operations are done with respect to the
// origin of the image which is the TOP LEFT CORNER.
func main() {
	example1()
	example2()
	example3()
	example4()
}

// Example 1.
// Rotate logo: by 90 degrees (about the origin), scale by 50 percent and
// then move the image 240 in the x direction
func example1() {
	imagick.Initialize()
	defer imagick.Terminate()

	s := make([]float64, 6)
	r := make([]float64, 6)
	t := make([]float64, 6)
	temp := make([]float64, 6)
	result := make([]float64, 6)

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	mw.ReadImage("logo:")

	// Set up the affine matrices
	// rotate 90 degrees clockwise
	set_rotate_affine(r, 90)
	// scale by .5 in x and y
	set_scale_affine(s, 0.5, 0.5)
	// translate to (240,0)
	set_translate_affine(t, 240, 0)
	// now multiply them - note the order in
	// which they are specified - in particular beware that
	// temp = r*s is NOT necessarily the same as temp = s*r

	//first do the rotation and scaling
	// temp = r*s
	affine_multiply(temp, r, s)
	// now the translation
	// result = temp*t
	affine_multiply(result, temp, t)

	// and then apply the result to the image
	mw.DistortImage(imagick.DISTORTION_AFFINE_PROJECTION, result, false)
	mw.WriteImage("logo_affine_1.jpg")
}

// Example 2
// Rotate logo: 30 degrees around the point (300,100)
// Since rotation is done around the origin, we must translate
// the point (300,100) up to the origin, do the rotation, and
// then translate back again
//
func example2() {
	imagick.Initialize()
	defer imagick.Terminate()

	t1 := make([]float64, 6)
	r := make([]float64, 6)
	t2 := make([]float64, 6)
	temp := make([]float64, 6)
	result := make([]float64, 6)

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	mw.ReadImage("logo:")

	// Initialize the required affines
	// translate (300,100) to origin
	set_translate_affine(t1, -300, -100)
	// rotate clockwise by 30 degrees
	set_rotate_affine(r, 30)
	// translate back again
	set_translate_affine(t2, 300, 100)

	// Now multiply the affine sequence
	// temp = t1*r
	affine_multiply(temp, t1, r)
	// result = temp*t2
	affine_multiply(result, temp, t2)

	mw.DistortImage(imagick.DISTORTION_AFFINE_PROJECTION, result, false)
	mw.WriteImage("logo_affine_2.jpg")
}

// Example 3
// Reflect the image about a line passing through the origin.
// If the line makes an angle of D degrees with the horizontal
// then this can be done by rotating the image by -D degrees so
// that the line is now (in effect) the x axis, reflect the image
// across the x axis, and then rotate everything back again.
// In this example, rather than just picking an arbitrary angle,
// let's say that we want the "logo:" image to be reflected across
// it's own major diagonal. Although we know the logo: image is
// 640x480 let's also generalize the code slightly so that it
// will still work if the name of the input image is changed.
// If the image has a width "w" and height "h", then the angle between
// the x-axis and the major diagonal is atan(h/w) (NOTE that this
// result is in RADIANS!)
// For this example I will also retain the original dimensions of the
// image so that anything that is reflected outside the borders of the
// input image is lost
func example3() {
	imagick.Initialize()
	defer imagick.Terminate()

	r1 := make([]float64, 6)
	reflect := make([]float64, 6)
	r2 := make([]float64, 6)
	temp := make([]float64, 6)
	result := make([]float64, 6)
	var angle_degrees float64

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	mw.ReadImage("logo:")
	w := mw.GetImageWidth()
	h := mw.GetImageHeight()

	// Just convert the radians to degrees. This way I don't have
	// to write a function which sets up an affine rotation for an
	// argument specified in radians rather than degrees.
	// You can always change this.
	angle_degrees = RadiansToDegrees(math.Atan(float64(h) / float64(w)))
	// Initialize the required affines
	// Rotate diagonal to the x axis
	set_rotate_affine(r1, -angle_degrees)
	// Reflection affine (about x-axis)
	// In this case there isn't a specific function to set the
	// affine array (like there is for rotation and scaling)
	// so use the function which sets an arbitrary affine
	set_affine(reflect, 1, 0, 0, -1, 0, 0)
	// rotate image back again
	set_rotate_affine(r2, angle_degrees)

	// temp = r1*reflect
	affine_multiply(temp, r1, reflect)
	// result = temp*r2
	affine_multiply(result, temp, r2)

	mw.DistortImage(imagick.DISTORTION_AFFINE_PROJECTION, result, false)
	mw.WriteImage("logo_affine_3.jpg")
}

// Example 4
// Create a rotated gradient
// See: http://www.imagemagick.org/discourse-server/viewtopic.php?f=1&t=12707
// The affine in this one is essentially the same as the one in Example 2 but
// this example has a different use for the result
func example4() {
	imagick.Initialize()
	defer imagick.Terminate()

	t1 := make([]float64, 6)
	r := make([]float64, 6)
	t2 := make([]float64, 6)
	temp := make([]float64, 6)
	result := make([]float64, 6)

	// Dimensions of the final rectangle
	w := uint(600)
	h := uint(100)
	// angle of clockwise rotation
	theta := 15.0 // degrees
	// Convert theta to radians
	rad_theta := DegreesToRadians(theta)
	// Compute the dimensions of the rectangular gradient
	gw := (uint)(float64(w)*math.Cos(rad_theta) + float64(h)*math.Sin(rad_theta) + 0.5)
	gh := (uint)(float64(w)*math.Sin(rad_theta) + float64(h)*math.Cos(rad_theta) + 0.5)
	// Don't let the rotation make the gradient rectangle any smaller
	// than the required output
	if gw < w {
		gw = w
	}
	if gh < h {
		gh = h
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	mw.SetSize(gw, gh)
	mw.ReadImage("gradient:white-black")

	// Initialize the required affines
	// translate centre of gradient to origin
	set_translate_affine(t1, float64(-gw)/2.0, float64(-gh)/2.0)
	// rotate clockwise by theta degrees
	set_rotate_affine(r, theta)
	// translate back again
	set_translate_affine(t2, float64(gw)/2.0, float64(gh)/2.0)

	// Now multiply the affine sequences
	// temp = t1*r
	affine_multiply(temp, t1, r)
	// result = temp*t2
	affine_multiply(result, temp, t2)

	mw.DistortImage(imagick.DISTORTION_AFFINE_PROJECTION, result, false)
	// Get the size of the distorted image and crop out the middle

	nw := mw.GetImageWidth()
	nh := mw.GetImageHeight()
	mw.CropImage(w, h, int((nw-w)/2), int((nh-h)/2))
	mw.WriteImage("rotgrad_2.png")
}
