// Port of http://members.shaw.ca/el.supremo/MagickWand/draw_shapes.htm to Go
package main

import (
	"github.com/gographics/imagick/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	cw := imagick.NewPixelWand()

	diameter := uint(640)
	radius := float64(diameter / 2)

	cw.SetColor("white")
	mw.NewImage(diameter, diameter, cw)

	dw.SetStrokeOpacity(1)

	// circle and rectangle
	dw.PushDrawingWand()

	// Hmmmm. Very weird. rgb(0,0,1) draws a black line around the edge
	// of the circle as it should. But rgb(0,0,0) or black don't.
	// AND if I remove the PixelSetColor then it draws a white boundary
	// around the rectangle (and presumably around the circle too)
	cw.SetColor("rgb(0,0,1)")

	dw.SetStrokeColor(cw)
	dw.SetStrokeWidth(4)
	dw.SetStrokeAntialias(true)
	cw.SetColor("red")

	//dw.SetStrokeOpacity(1)
	dw.SetFillColor(cw)

	dw.Circle(radius, radius, radius, radius*2)
	dw.Rectangle(50, 13, 120, 87)
	dw.PopDrawingWand()

	// rounded rectangle
	dw.PushDrawingWand()

	points := []imagick.PointInfo{
		{378.1, 81.72}, {381.1, 79.56}, {384.3, 78.12}, {387.6, 77.33},
		{391.1, 77.11}, {394.6, 77.62}, {397.8, 78.77}, {400.9, 80.57},
		{403.6, 83.02}, {523.9, 216.8}, {526.2, 219.7}, {527.6, 223},
		{528.4, 226.4}, {528.6, 229.8}, {528, 233.3}, {526.9, 236.5},
		{525.1, 239.5}, {522.6, 242.2}, {495.9, 266.3}, {493, 268.5},
		{489.7, 269.9}, {486.4, 270.8}, {482.9, 270.9}, {479.5, 270.4},
		{476.2, 269.3}, {473.2, 267.5}, {470.4, 265}, {350, 131.2},
		{347.8, 128.3}, {346.4, 125.1}, {345.6, 121.7}, {345.4, 118.2},
		{346, 114.8}, {347.1, 111.5}, {348.9, 108.5}, {351.4, 105.8},
		{378.1, 81.72},
	}

	dw.SetStrokeAntialias(true)
	dw.SetStrokeWidth(2.016)
	dw.SetStrokeLineCap(imagick.LINE_CAP_ROUND)
	dw.SetStrokeLineJoin(imagick.LINE_JOIN_ROUND)

	dw.SetStrokeDashArray(make([]float64, 0))
	cw.SetColor("rgb(0,0,128)")
	// If strokecolor is removed completely then the circle is not there
	dw.SetStrokeColor(cw)
	// But now I've added strokeopacity - 1=circle there 0=circle not there
	// If opacity is 1 the black edge around the rectangle is visible
	dw.SetStrokeOpacity(1)
	// No effect
	// dw.SetFillRule(imagick.FILL_EVEN_ODD)
	// this doesn't affect the circle
	cw.SetColor("#c2c280")
	dw.SetFillColor(cw)
	// 1=circle there 0=circle there but rectangle fill disappears
	// dw.SetFillOpacity(0)
	dw.Polygon(points)
	dw.SetStrokeOpacity(1)

	dw.PopDrawingWand()

	// yellow polygon
	dw.PushDrawingWand()
	points = []imagick.PointInfo{
		{540, 288}, {561.6, 216}, {547.2, 43.2}, {280.8, 36},
		{302.4, 194.4}, {331.2, 64.8}, {504, 64.8}, {475.2, 115.2},
		{525.6, 93.6}, {496.8, 158.4}, {532.8, 136.8}, {518.4, 180},
		{540, 172.8}, {540, 223.2}, {540, 288},
	}

	dw.SetStrokeAntialias(true)
	dw.SetStrokeWidth(5.976)
	dw.SetStrokeLineCap(imagick.LINE_CAP_ROUND)
	dw.SetStrokeLineJoin(imagick.LINE_JOIN_ROUND)
	dw.SetStrokeDashArray([]float64{})
	cw.SetColor("#4000c2")
	dw.SetStrokeColor(cw)
	dw.SetFillRule(imagick.FILL_EVEN_ODD)
	cw.SetColor("#ffff00")
	dw.SetFillColor(cw)
	dw.Polygon(points)
	dw.PopDrawingWand()

	// rotated and translated ellipse
	// The DrawEllipse function only draws the ellipse with
	// the major and minor axes orthogonally aligned. This also
	// applies to some of the other functions such as DrawRectangle.
	// If you want an ellipse that has the major axis rotated, you
	// have to rotate the coordinate system before the ellipse is
	// drawn. And you'll also want the ellipse somewhere on the
	// image rather than at the top left (where the 0,0 origin is
	// located) so before drawing the ellipse we move the origin to
	// wherever we want the centre of the ellipse to be and then
	// rotate the coordinate system by the angle of rotation we wish
	// to apply to the ellipse and *then* we draw the ellipse.
	// NOTE that doing all this within PushDrawingWand()/PopDrawingWand()
	// means that the coordinate system will be restored after
	// the PopDrawingWand
	dw.PushDrawingWand()

	cw.SetColor("rgb(0,0,1)")

	dw.SetStrokeColor(cw)
	dw.SetStrokeWidth(2)
	dw.SetStrokeAntialias(true)
	cw.SetColor("orange")
	//dw.DrawSetStrokeOpacity(1)
	dw.SetFillColor(cw)
	// Be careful of the order in which you meddle with the
	// coordinate system! Rotating and then translating is
	// not the same as translating then rotating
	dw.Translate(radius/2, 3*radius/2)
	dw.Rotate(-30)
	dw.Ellipse(0, 0, radius/8, 3*radius/8, 0, 360)
	dw.PopDrawingWand()

	// A line from the centre of the circle
	// to the top left edge of the image
	dw.Line(0, 0, radius, radius)

	mw.DrawImage(dw)
	mw.WriteImage("chart_test.jpg")
}
