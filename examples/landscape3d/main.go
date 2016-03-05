// Port of http://members.shaw.ca/el.supremo/MagickWand/landscape_3d.htm to Go
package main

import "gopkg.in/gographics/imagick.v1/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	image := imagick.NewMagickWand()
	defer image.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	image.ReadImage("fract6.jpg")

	// scale it down
	w := image.GetImageWidth()
	h := image.GetImageHeight()

	pw.SetColor("transparent")
	if err := image.ShearImage(pw, 45, 0); err != nil {
		panic(err)
	}

	w = image.GetImageWidth()
	h = image.GetImageHeight()

	// scale it to make it look like it is laying down
	if err := image.ScaleImage(w, h/2); err != nil {
		panic(err)
	}

	// Get image stats
	w = image.GetImageWidth()
	h = image.GetImageHeight()

	// Make a blank canvas to draw on
	canvas := imagick.NewMagickWand()
	defer canvas.Destroy()
	// Use a colour from the input image
	pw, err := image.GetImagePixelColor(0, 0)
	if err != nil {
		panic(err)
	}
	canvas.NewImage(w, h*2, pw)

	offset := int(h)
	// The original drawing method was to go along each row from top to bottom so that
	// a line in the "front" (which is one lower down the picture) will be drawn over
	// one behind it.
	// The problem with this method is that every line is drawn even if it will be covered
	// up by a line "in front" of it later on.
	// The method used here goes up each column from left to right and only draws a line if
	// it is longer than everything drawn so far in this column and will therefore be visible.
	// With the new drawing method this takes 13 secs - the previous method took 59 secs
	// loop through all points in image
	for x := 0; x < int(w); x++ {
		// The PHP version created, used and destroyed the drawingwand inside
		// the inner loop but it is about 25% faster to do only the DrawLine
		// inside the loop
		line := imagick.NewDrawingWand()
		line_height := 0
		for y := int(h) - 1; y >= 0; y-- {
			// get (r,g,b) and grey value
			pw, err := image.GetImagePixelColor(int(x), int(y))
			if err != nil {
				continue
			}
			// 255* adjusts the rgb values to Q8 even if the IM being used is Q16
			r := (int)(255 * pw.GetRed())
			g := (int)(255 * pw.GetGreen())
			b := (int)(255 * pw.GetBlue())

			// Calculate grayscale - a divisor of 10-25 seems to work well.
			//			grey = (r+g+b)/25
			grey := (r + g + b) / 15
			//			grey = (r+g+b)/10
			// Only draw a line if it will show "above" what's already been done
			if line_height == 0 || line_height < grey {
				line.SetFillColor(pw)
				line.SetStrokeColor(pw)
				// Draw the part of the line that is visible
				start_y := y + offset - line_height
				end_y := y - grey + offset
				line.Line(float64(x), float64(start_y), float64(x), float64(end_y))
				line_height = grey
			}
			line_height--
		}
		// Draw the lines on the image
		canvas.DrawImage(line)
		line.Destroy()
	}
	canvas.ScaleImage(w-h, h*2)
	// write canvas
	canvas.WriteImage("3d_fractal.jpg")
}
