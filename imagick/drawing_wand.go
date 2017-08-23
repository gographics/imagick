// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"
import (
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

type DrawingWand struct {
	dw   *C.DrawingWand
	init sync.Once
}

func newDrawingWand(cdw *C.DrawingWand) *DrawingWand {
	dw := &DrawingWand{dw: cdw}
	runtime.SetFinalizer(dw, Destroy)
	dw.IncreaseCount()

	return dw
}

// Returns a drawing wand required for all other methods in the API.
func NewDrawingWand() *DrawingWand {
	return newDrawingWand(C.NewDrawingWand())
}

// Clears resources associated with the drawing wand.
func (dw *DrawingWand) Clear() {
	C.ClearDrawingWand(dw.dw)
	runtime.KeepAlive(dw)
}

// Makes an exact copy of the specified wand.
func (dw *DrawingWand) Clone() *DrawingWand {
	ret := newDrawingWand(C.CloneDrawingWand(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Frees all resources associated with the drawing wand. Once the drawing wand
// has been freed, it should not be used and further unless it re-allocated.
func (dw *DrawingWand) Destroy() {
	if dw.dw == nil {
		return
	}

	dw.init.Do(func() {
		dw.dw = C.DestroyDrawingWand(dw.dw)
		relinquishMemory(unsafe.Pointer(dw.dw))
		dw.dw = nil

		dw.DecreaseCount()
	})
}

// Increase DrawingWand ref counter and set according "can`t be terminated status"
func (dw *DrawingWand) IncreaseCount() {
	atomic.AddInt64(&drawingWandCounter, int64(1))
	unsetCanTerminate()
}

// Decrease DrawingWand ref counter and set according "can be terminated status"
func (dw *DrawingWand) DecreaseCount() {
	atomic.AddInt64(&drawingWandCounter, int64(-1))
	setCanTerminate()
}

// Adjusts the current affine transformation matrix with the specified affine
// transformation matrix. Note that the current affine transform is adjusted
// rather than replaced.
//
// affine: Affine matrix parameters
//
func (dw *DrawingWand) Affine(affine *AffineMatrix) {
	C.DrawAffine(dw.dw, affine.ptr())
	runtime.KeepAlive(dw)
}

// Draws text on the image.
// x: x ordinate to left of text
// y: y ordinate to text baseline
// text: text to draw
func (dw *DrawingWand) Annotation(x, y float64, text string) {
	cstext := (*C.uchar)((unsafe.Pointer)(C.CString(text)))
	defer C.free(unsafe.Pointer(cstext))
	C.DrawAnnotation(dw.dw, C.double(x), C.double(y), cstext)
	runtime.KeepAlive(dw)
}

// Draws an arc falling within a specified bounding rectangle on the image.
//
// sx:  starting x ordinate of bounding rectangle
//
// sy:  starting y ordinate of bounding rectangle
//
// ex:  ending x ordinate of bounding rectangle
//
// ey:  ending y ordinate of bounding rectangle
//
// sd:  starting degrees of rotation
//
// ed:  ending degrees of rotation
//
func (dw *DrawingWand) Arc(sx, sy, ex, ey, sd, ed float64) {
	C.DrawArc(dw.dw, C.double(sx), C.double(sy), C.double(ex), C.double(ey), C.double(sd), C.double(ed))
	runtime.KeepAlive(dw)
}

// Draws a bezier curve through a set of points on the image.
func (dw *DrawingWand) Bezier(coordinates []PointInfo) {
	ccoordinates := [1 << 16]C.PointInfo{}
	for k, v := range coordinates {
		ccoordinates[k] = C.PointInfo{C.double(v.X), C.double(v.Y)}
	}
	C.DrawBezier(dw.dw, C.size_t(len(coordinates)), (*C.PointInfo)(&ccoordinates[0]))
	runtime.KeepAlive(dw)
}

// Draws a circle on the image.
//
// ox: origin x ordinate
//
// oy: origin y ordinate
//
// px: perimeter x ordinate
//
// py: perimeter y ordinate
//
func (dw *DrawingWand) Circle(ox, oy, px, py float64) {
	C.DrawCircle(dw.dw, C.double(ox), C.double(oy), C.double(px), C.double(py))
	runtime.KeepAlive(dw)
}

// Composites an image onto the current image, using the specified composition
// operator, specified position, and at the specified size.
//
// compose: composition operator
//
// x: x ordinate of top left corner
//
// y: y ordinate of top left corner
//
// width: Width to resize image to prior to compositing. Specify zero to use
// existing width.
//
// height: Height to resize image to prior to compositing. Specify zero to use
// existing height.
//
// mw: Image to composite is obtained from this wand.
//
func (dw *DrawingWand) Composite(compose CompositeOperator, x, y, width, height float64, mw *MagickWand) error {
	ok := C.DrawComposite(dw.dw, C.CompositeOperator(compose), C.double(x), C.double(y), C.double(width), C.double(height), mw.mw)
	return dw.getLastErrorIfFailed(ok)
}

// Draws color on image using the current fill color, starting at specified
// position, and using specified paint method. The available paint methods are:
//
// x: x ordinate.
//
// y: y ordinate.
//
// pm: paint method. PointMethod: Recolors the target pixel. ReplaceMethod:
// Recolor any pixel that matches the target pixel. FloodfillMethod: Recolors
// target pixels and matching neighbors. ResetMethod: Recolor all pixels.
//
func (dw *DrawingWand) Color(x, y float64, pm PaintMethod) {
	C.DrawColor(dw.dw, C.double(x), C.double(y), C.PaintMethod(pm))
	runtime.KeepAlive(dw)
}

// Adds a comment to a vector output stream.
func (dw *DrawingWand) Comment(comment string) {
	cscomment := C.CString(comment)
	defer C.free(unsafe.Pointer(cscomment))
	C.DrawComment(dw.dw, cscomment)
	runtime.KeepAlive(dw)
}

// Draws an ellipse on the image.
//
// ox: origin x ordinate
//
// oy: origin y ordinate
//
// rx: radius in x
//
// ry: radius in y
//
// start: starting rotation in degrees
//
// end: ending rotation in degrees
//
func (dw *DrawingWand) Ellipse(ox, oy, rx, ry, start, end float64) {
	C.DrawEllipse(dw.dw, C.double(ox), C.double(oy), C.double(rx), C.double(ry), C.double(start), C.double(end))
	runtime.KeepAlive(dw)
}

// Returns the border color used for drawing bordered objects.
func (dw *DrawingWand) GetBorderColor() (pw *PixelWand) {
	pw = NewPixelWand()
	C.DrawGetBorderColor(dw.dw, pw.pw)
	runtime.KeepAlive(dw)
	return
}

// Obtains the current clipping path ID.
func (dw *DrawingWand) GetClipPath() string {
	cscp := C.DrawGetClipPath(dw.dw)
	runtime.KeepAlive(dw)
	defer relinquishMemory(unsafe.Pointer(cscp))
	return C.GoString(cscp)
}

// Returns the current polygon fill rule to be used by the clipping path.
func (dw *DrawingWand) GetClipRule() FillRule {
	ret := FillRule(C.DrawGetClipRule(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the interpretation of clip path units.
func (dw *DrawingWand) GetClipUnits() ClipPathUnits {
	ret := ClipPathUnits(C.DrawGetClipUnits(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the fill color used for drawing filled objects.
func (dw *DrawingWand) GetFillColor() (pw *PixelWand) {
	pw = NewPixelWand()
	C.DrawGetFillColor(dw.dw, pw.pw)
	runtime.KeepAlive(dw)
	return
}

// Returns the opacity used when drawing using the fill color or fill texture.
// Fully opaque is 1.0.
func (dw *DrawingWand) GetFillOpacity() float64 {
	ret := float64(C.DrawGetFillOpacity(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the fill rule used while drawing polygons.
func (dw *DrawingWand) GetFillRule() FillRule {
	ret := FillRule(C.DrawGetFillRule(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns a string specifying the font used when annotating with text.
func (dw *DrawingWand) GetFont() string {
	csfont := C.DrawGetFont(dw.dw)
	runtime.KeepAlive(dw)
	defer relinquishMemory(unsafe.Pointer(csfont))
	return C.GoString(csfont)
}

// Returns the font family to use when annotating with text.
func (dw *DrawingWand) GetFontFamily() string {
	csfamily := C.DrawGetFontFamily(dw.dw)
	runtime.KeepAlive(dw)
	defer relinquishMemory(unsafe.Pointer(csfamily))
	return C.GoString(csfamily)
}

// Gets the image X and Y resolution.
func (dw *DrawingWand) GetFontResolution() (x, y float64, err error) {
	ok := C.DrawGetFontResolution(dw.dw, (*C.double)(&x), (*C.double)(&y))
	err = dw.getLastErrorIfFailed(ok)
	return
}

// Returns the font stretch used when annotating with text.
func (dw *DrawingWand) GetFontStretch() StretchType {
	ret := StretchType(C.DrawGetFontStretch(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the font style used when annotating with text.
func (dw *DrawingWand) GetFontStyle() StyleType {
	ret := StyleType(C.DrawGetFontStyle(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the font weight used when annotating with text.
func (dw *DrawingWand) GetFontWeight() uint {
	ret := uint(C.DrawGetFontWeight(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the text placement gravity used when annotating with text.
func (dw *DrawingWand) GetGravity() GravityType {
	ret := GravityType(C.DrawGetGravity(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the opacity used when drawing with the fill or stroke color or
// texture. Fully opaque is 1.0.
func (dw *DrawingWand) GetOpacity() float64 {
	ret := float64(C.DrawGetOpacity(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the current stroke antialias setting. Stroked outlines are
// antialiased by default. When antialiasing is disabled stroked pixels are
// thresholded to determine if the stroke color or underlying canvas color
// should be used.
func (dw *DrawingWand) GetStrokeAntialias() bool {
	ret := 1 == C.DrawGetStrokeAntialias(dw.dw)
	runtime.KeepAlive(dw)
	return ret
}

// Returns the color used for stroking object outlines.
func (dw *DrawingWand) GetStrokeColor() (pw *PixelWand) {
	pw = NewPixelWand()
	C.DrawGetStrokeColor(dw.dw, pw.pw)
	runtime.KeepAlive(dw)
	return
}

// Returns an array representing the pattern of dashes and gaps used to stroke
// paths (see SetStrokeDashArray). The array must be freed once it is no longer
// required by the user.
func (dw *DrawingWand) GetStrokeDashArray() (nums []float64) {
	count := C.size_t(0)
	p := C.DrawGetStrokeDashArray(dw.dw, &count)
	runtime.KeepAlive(dw)
	nums = sizedDoubleArrayToFloat64Slice(p, count)
	return
}

// Returns the offset into the dash pattern to start the dash.
func (dw *DrawingWand) GetStrokeDashOffset() float64 {
	ret := float64(C.DrawGetStrokeDashOffset(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the shape to be used at the end of open subpaths when they are
// stroked. Values of LineCap are UndefinedCap, ButtCap, RoundCap, and
// SquareCap.
func (dw *DrawingWand) GetStrokeLineCap() LineCap {
	ret := LineCap(C.DrawGetStrokeLineCap(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the shape to be used at the corners of paths (or other vector
// shapes) when they are stroked. Values of LineJoin are UndefinedJoin,
// MiterJoin, RoundJoin, and BevelJoin.
func (dw *DrawingWand) GetStrokeLineJoin() LineJoin {
	ret := LineJoin(C.DrawGetStrokeLineJoin(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the miter limit. When two line segments meet at a sharp angle and
// miter joins have been specified for 'lineJoin', it is possible for the
// miter to extend far beyond the thickness of the line stroking the path.
// The miterLimit' imposes a limit on the ratio of the miter length to the
// 'lineWidth'.
func (dw *DrawingWand) GetStrokeMiterLimit() uint {
	ret := uint(C.DrawGetStrokeMiterLimit(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the opacity of stroked object outlines.
func (dw *DrawingWand) GetStrokeOpacity() float64 {
	ret := float64(C.DrawGetStrokeOpacity(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the width of the stroke used to draw object outlines.
func (dw *DrawingWand) GetStrokeWidth() float64 {
	ret := float64(C.DrawGetStrokeWidth(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the alignment applied when annotating with text.
func (dw *DrawingWand) GetTextAlignment() AlignType {
	ret := AlignType(C.DrawGetTextAlignment(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns the current text antialias setting, which determines whether text
// is antialiased. Text is antialiased by default.
func (dw *DrawingWand) GetTextAntialias() bool {
	ret := 1 == C.DrawGetTextAntialias(dw.dw)
	runtime.KeepAlive(dw)
	return ret
}

// Returns the decoration applied when annotating with text.
func (dw *DrawingWand) GetTextDecoration() DecorationType {
	ret := DecorationType(C.DrawGetTextDecoration(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns a string which specifies the code set used for text annotations.
func (dw *DrawingWand) GetTextEncoding() string {
	cstr := C.DrawGetTextEncoding(dw.dw)
	runtime.KeepAlive(dw)
	defer relinquishMemory(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Gets the spacing between characters in text.
func (dw *DrawingWand) GetTextKerning() float64 {
	ret := float64(C.DrawGetTextKerning(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Gets the spacing between lines in text.
func (dw *DrawingWand) GetTextInterlineSpacing() float64 {
	ret := float64(C.DrawGetTextInterwordSpacing(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Gets the spacing between words in text.
func (dw *DrawingWand) GetTextInterwordSpacing() float64 {
	ret := float64(C.DrawGetTextInterwordSpacing(dw.dw))
	runtime.KeepAlive(dw)
	return ret
}

// Returns a string which specifies the vector graphics generated by any
// graphics calls made since the wand was instantiated.
func (dw *DrawingWand) GetVectorGraphics() string {
	cstr := C.DrawGetVectorGraphics(dw.dw)
	runtime.KeepAlive(dw)
	defer relinquishMemory(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

// Returns the color of a background rectangle to place under text annotations.
func (dw *DrawingWand) GetTextUnderColor() (pw *PixelWand) {
	pw = NewPixelWand()
	C.DrawGetTextUnderColor(dw.dw, pw.pw)
	runtime.KeepAlive(dw)
	return
}

// Draws a line on the image using the current stroke color, stroke opacity,
// and stroke width.
//
//sx: starting x ordinate
//
//sy: starting y ordinate
//
//ex: ending x ordinate
//
//ey: ending y ordinate
//
func (dw *DrawingWand) Line(sx, sy, ex, ey float64) {
	C.DrawLine(dw.dw, C.double(sx), C.double(sy), C.double(ex), C.double(ey))
	runtime.KeepAlive(dw)
}

// Deprecated in ImageMagick-7.x
// Use Alpha()
func (dw *DrawingWand) Matte(x, y float64, pmethod PaintMethod) {
	C.DrawAlpha(dw.dw, C.double(x), C.double(y), C.PaintMethod(pmethod))
	runtime.KeepAlive(dw)
}

// Paints on the image's opacity channel in order to set effected pixels to
// transparent. to influence the opacity of pixels. The available paint
// methods are:
//
// 	ResetMethod: Select all pixels.
//
// 	PointMethod: Select the target pixel
//
// 	ReplaceMethod: Select any pixel that matches the target pixel.
//
// 	FloodfillMethod: Select the target pixel and matching neighbors.
//
// 	FillToBorderMethod: Select the target pixel and neighbors not matching
//                      border color.
//
// x, y: x, y ordinates
// pmethod: paint method
func (dw *DrawingWand) Alpha(x, y float64, pmethod PaintMethod) {
	C.DrawAlpha(dw.dw, C.double(x), C.double(y), C.PaintMethod(pmethod))
	runtime.KeepAlive(dw)
}

// Adds a path element to the current path which closes the current subpath by
// drawing a straight line from the current point to the current subpath's most
// recent starting point (usually, the most recent moveto point).
func (dw *DrawingWand) PathClose() {
	C.DrawPathClose(dw.dw)
	runtime.KeepAlive(dw)
}

// Draws a cubic Bezier curve from the current point to (x,y) using (x1,y1) as
// the control point at the beginning of the curve and (x2,y2) as the control
// point at the end of the curve using absolute coordinates. At the end of the
// command, the new current point becomes the final (x,y) coordinate pair used
// in the polybezier.
//
// x1, y1: x, y ordinates of control point for curve beginning
//
// x2, y2: x, y ordinates of control point for curve ending
//
// x, y: x, y ordinates of the end of the curve
//
func (dw *DrawingWand) PathCurveToAbsolute(x1, y1, x2, y2, x, y float64) {
	C.DrawPathCurveToAbsolute(dw.dw, C.double(x1), C.double(y1), C.double(x2), C.double(y2), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a cubic Bezier curve from the current point to (x,y) using (x1,y1) as
// the control point at the beginning of the curve and (x2,y2) as the control
// point at the end of the curve using relative coordinates. At the end of the
// command, the new current point becomes the final (x,y) coordinate pair used
// in the polybezier.
//
// x1, y1: x, y ordinates of control point for curve beginning
//
// x2, y2: x, y ordinates of control point for curve ending
//
// x, y: x, y ordinates of the end of the curve
//
func (dw *DrawingWand) PathCurveToRelative(x1, y1, x2, y2, x, y float64) {
	C.DrawPathCurveToRelative(dw.dw, C.double(x1), C.double(y1), C.double(x2), C.double(y2), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a quadratic Bezier curve from the current point to (x,y) using (x1,y1)
// as the control point using absolute coordinates. At the end of the command,
// the new current point becomes the final (x,y) coordinate pair used in the
// polybezier.
//
// x1, y1: ordinates of the control point
//
// x, y: ordinates of final point
//
func (dw *DrawingWand) PathCurveToQuadraticBezierAbsolute(x1, y1, x, y float64) {
	C.DrawPathCurveToQuadraticBezierAbsolute(dw.dw, C.double(x1), C.double(y1), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a quadratic Bezier curve from the current point to (x,y) using (x1,y1)
// as the control point using relative coordinates. At the end of the command,
// the new current point becomes the final (x,y) coordinate pair used in the
// polybezier.
// x1, y1: ordinates of the control point
// x, y: ordinates of final point
func (dw *DrawingWand) PathCurveToQuadraticBezierRelative(x1, y1, x, y float64) {
	C.DrawPathCurveToQuadraticBezierRelative(dw.dw, C.double(x1), C.double(y1), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a quadratic Bezier curve (using absolute coordinates) from the current
// point to (x,y). The control point is assumed to be the reflection of the
// control point on the previous command relative to the current point. (If
// there is no previous command or if the previous command was not a
// PathCurveToQuadraticBezierAbsolute, PathCurveToQuadraticBezierRelative,
// PathCurveToQuadraticBezierSmoothAbsolute or
// PathCurveToQuadraticBezierSmoothRelative, assume the control point is
// coincident with the current point.). At the end of the command, the new
// current point becomes the final (x,y) coordinate pair used in the polybezier.
//
//x, y: ordinates of final point
//
func (dw *DrawingWand) PathCurveToQuadraticBezierSmoothAbsolute(x, y float64) {
	C.DrawPathCurveToQuadraticBezierSmoothAbsolute(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a quadratic Bezier curve (using relative coordinates) from the current
// point to (x,y). The control point is assumed to be the reflection of the
// control point on the previous command relative to the current point. (If
// there is no previous command or if the previous command was not a
// PathCurveToQuadraticBezierAbsolute, PathCurveToQuadraticBezierRelative,
// PathCurveToQuadraticBezierSmoothAbsolute or
// PathCurveToQuadraticBezierSmoothRelative, assume the control point is
// coincident with the current point.). At the end of the command, the new
// current point becomes the final (x,y) coordinate pair used in the polybezier.
//
//x, y: ordinates of final point
//
func (dw *DrawingWand) PathCurveToQuadraticBezierSmoothRelative(x, y float64) {
	C.DrawPathCurveToQuadraticBezierSmoothRelative(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a cubic Bezier curve from the current point to (x,y) using absolute
// coordinates. The first control point is assumed to be the reflection of the
// second control point on the previous command relative to the current point.
// (If there is no previous command or if the previous command was not an
// PathCurveToAbsolute, PathCurveToRelative, PathCurveToSmoothAbsolute or
// PathCurveToSmoothRelative, assume the first control point is coincident
// with the current point.) (x2,y2) is the second control point (i.e., the
// control point at the end of the curve). At the end of the command, the new
// current point becomes the final (x,y) coordinate pair used in the polybezier.
//
// x2, y2: ordinates of second control point
//
// x, y: ordinates of termination point
//
func (dw *DrawingWand) PathCurveToSmoothAbsolute(x2, y2, x, y float64) {
	C.DrawPathCurveToSmoothAbsolute(dw.dw, C.double(x2), C.double(y2), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a cubic Bezier curve from the current point to (x,y) using relative
// coordinates. The first control point is assumed to be the reflection of the
// second control point on the previous command relative to the current point.
// (If there is no previous command or if the previous command was not an
// PathCurveToAbsolute, PathCurveToRelative, PathCurveToSmoothAbsolute or
// PathCurveToSmoothRelative, assume the first control point is coincident
// with the current point.) (x2,y2) is the second control point (i.e., the
// control point at the end of the curve). At the end of the command, the new
// current point becomes the final (x,y) coordinate pair used in the polybezier.
//
// x2, y2: ordinates of second control point
//
// x, y: ordinates of termination point
//
func (dw *DrawingWand) PathCurveToSmoothRelative(x2, y2, x, y float64) {
	C.DrawPathCurveToSmoothRelative(dw.dw, C.double(x2), C.double(y2), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws an elliptical arc from the current point to (x, y) using absolute
// coordinates. The size and orientation of the ellipse are defined by two
// radii (rx, ry) and an xAxisRotation, which indicates how the ellipse as a
// whole is rotated relative to the current coordinate system. The center (cx,
// cy) of the ellipse is calculated automagically to satisfy the constraints
// imposed by the other parameters. largeArcFlag and sweepFlag contribute to
// the automatic calculations and help determine how the arc is drawn. If
// largeArcFlag is true then draw the larger of the available arcs. If
// sweepFlag is true, then draw the arc matching a clock-wise rotation.
//
// rx, ry: x, y radius
//
// xAxisRotation: indicates how the ellipse as a whole is rotated relative to
// the current coordinate system
//
// largeArcFlag: If true then draw the larger of the available arcs
//
// sweepFlag: If true then draw the arc matching a clock-wise rotation
//
func (dw *DrawingWand) PathEllipticArcAbsolute(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64) {
	C.DrawPathEllipticArcAbsolute(dw.dw, C.double(rx), C.double(ry), C.double(xAxisRotation), b2i(largeArcFlag), b2i(sweepFlag), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws an elliptical arc from the current point to (x, y) using relative
// coordinates. The size and orientation of the ellipse are defined by two
// radii (rx, ry) and an xAxisRotation, which indicates how the ellipse as a
// whole is rotated relative to the current coordinate system. The center (cx,
// cy) of the ellipse is calculated automagically to satisfy the constraints
// imposed by the other parameters. largeArcFlag and sweepFlag contribute to
// the automatic calculations and help determine how the arc is drawn. If
// largeArcFlag is true then draw the larger of the available arcs. If
// sweepFlag is true, then draw the arc matching a clock-wise rotation.
//
// rx, ry: x, y radius
//
// xAxisRotation: indicates how the ellipse as a whole is rotated relative to
// the current coordinate system
//
// largeArcFlag: If true then draw the larger of the available arcs
//
// sweepFlag: If true then draw the arc matching a clock-wise rotation
//
func (dw *DrawingWand) PathEllipticArcRelative(rx, ry, xAxisRotation float64, largeArcFlag, sweepFlag bool, x, y float64) {
	C.DrawPathEllipticArcRelative(dw.dw, C.double(rx), C.double(ry), C.double(xAxisRotation), b2i(largeArcFlag), b2i(sweepFlag), C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Terminates the current path.
func (dw *DrawingWand) PathFinish() {
	C.DrawPathFinish(dw.dw)
	runtime.KeepAlive(dw)
}

// Draws a line path from the current point to the given coordinate using
// absolute coordinates. The coordinate then becomes the new current point.
//
// x, y: target x and y ordinates
func (dw *DrawingWand) PathLineToAbsolute(x, y float64) {
	C.DrawPathLineToAbsolute(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a line path from the current point to the given coordinate using
// relative coordinates. The coordinate then becomes the new current point.
//
// x, y: target x and y ordinates
//
func (dw *DrawingWand) PathLineToRelative(x, y float64) {
	C.DrawPathLineToRelative(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a horizontal line path from the current point to the target point
// using absolute coordinates. The target point then becomes the new current
// point.
//
// x: target x ordinate
func (dw *DrawingWand) PathLineToHorizontalAbsolute(x float64) {
	C.DrawPathLineToHorizontalAbsolute(dw.dw, C.double(x))
	runtime.KeepAlive(dw)
}

// Draws a horizontal line path from the current point to the target point
// using relative coordinates. The target point then becomes the new current
// point.
//
// x: target x ordinate
func (dw *DrawingWand) PathLineToHorizontalRelative(x float64) {
	C.DrawPathLineToHorizontalRelative(dw.dw, C.double(x))
	runtime.KeepAlive(dw)
}

// Draws a vertical line path from the current point to the target point using
// absolute coordinates. The target point then becomes the new current point.
//
// y: target y ordinate
func (dw *DrawingWand) PathLineToVerticalAbsolute(y float64) {
	C.DrawPathLineToVerticalAbsolute(dw.dw, C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a vertical line path from the current point to the target point using
// relative coordinates. The target point then becomes the new current point.
//
// y: target y ordinate
func (dw *DrawingWand) PathLineToVerticalRelative(y float64) {
	C.DrawPathLineToVerticalRelative(dw.dw, C.double(y))
	runtime.KeepAlive(dw)
}

// Starts a new sub-path at the given coordinate using absolute coordinates.
// The current point then becomes the specified coordinate.
//
// x, y: target x and y ordinates
func (dw *DrawingWand) PathMoveToAbsolute(x, y float64) {
	C.DrawPathMoveToAbsolute(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Starts a new sub-path at the given coordinate using relative coordinates.
// The current point then becomes the specified coordinate.
//
// x, y: target x and y ordinates
func (dw *DrawingWand) PathMoveToRelative(x, y float64) {
	C.DrawPathMoveToRelative(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Declares the start of a path drawing list which is terminated by a matching
// PathFinish() command. All other Path commands must be enclosed between a
// PathStart() and a PathFinish() command. This is because path drawing
// commands are subordinate commands and they do not function by themselves.
func (dw *DrawingWand) PathStart() {
	C.DrawPathStart(dw.dw)
	runtime.KeepAlive(dw)
}

// Draws a point using the current fill color.
//
// x, y: target x, y coordinates
func (dw *DrawingWand) Point(x, y float64) {
	C.DrawPoint(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Draws a polygon using the current stroke, stroke width, and fill color or
// texture, using the specified array of coordinates.
func (dw *DrawingWand) Polygon(coordinates []PointInfo) {
	ccoordinates := [1 << 16]C.PointInfo{}
	for k, v := range coordinates {
		ccoordinates[k] = C.PointInfo{C.double(v.X), C.double(v.Y)}
	}
	C.DrawPolygon(dw.dw, C.size_t(len(coordinates)), (*C.PointInfo)(&ccoordinates[0]))
	runtime.KeepAlive(dw)
}

// Draws a polyline using the current stroke, stroke width, and fill color or
// texture, using the specified array of coordinates.
func (dw *DrawingWand) Polyline(coordinates []PointInfo) {
	ccoordinates := [1 << 16]C.PointInfo{}
	for k, v := range coordinates {
		ccoordinates[k] = C.PointInfo{C.double(v.X), C.double(v.Y)}
	}
	C.DrawPolyline(dw.dw, C.size_t(len(coordinates)), (*C.PointInfo)(&ccoordinates[0]))
	runtime.KeepAlive(dw)
}

// Terminates a clip path definition.
func (dw *DrawingWand) PopClipPath() {
	C.DrawPopClipPath(dw.dw)
	runtime.KeepAlive(dw)
}

// Terminates a definition list.
func (dw *DrawingWand) PopDefs() {
	C.DrawPopDefs(dw.dw)
	runtime.KeepAlive(dw)
}

// Terminates a pattern definition.
func (dw *DrawingWand) PopPattern() error {
	ok := C.DrawPopPattern(dw.dw)
	return dw.getLastErrorIfFailed(ok)
}

// Starts a clip path definition which is comprized of any number of drawing
// commands and terminated by a DrawPopClipPath() command.
//
// clipMaskId: string identifier to associate with the clip path for later use.
func (dw *DrawingWand) PushClipPath(clipMaskId string) {
	cstr := C.CString(clipMaskId)
	defer C.free(unsafe.Pointer(cstr))
	C.DrawPushClipPath(dw.dw, cstr)
	runtime.KeepAlive(dw)
}

// Indicates that commands up to a terminating PopDefs() command create named
// elements (e.g. clip-paths, textures, etc.) which may safely be processed
// earlier for the sake of efficiency.
func (dw *DrawingWand) PushDefs() {
	C.DrawPushDefs(dw.dw)
	runtime.KeepAlive(dw)
}

// Indicates that subsequent commands up to a PopPattern() command comprise the
// definition of a named pattern. The pattern space is assigned top left corner
// coordinates, a width and height, and becomes its own drawing space. Anything
// which can be drawn may be used in a pattern definition. Named patterns may
// be used as stroke or brush definitions.
//
// patternId: pattern identification for later reference
//
// x, y: ordinates of top left corner
//
// width, height of pattern space
func (dw *DrawingWand) PushPattern(patternId string, x, y, width, height float64) error {
	cstr := C.CString(patternId)
	defer C.free(unsafe.Pointer(cstr))
	ok := C.DrawPushPattern(dw.dw, cstr, C.double(x), C.double(y), C.double(width), C.double(height))
	return dw.getLastErrorIfFailed(ok)
}

// Draws a rectangle given two coordinates and using the current stroke, stroke
// width, and fill settings.
//
// x1, y1: ordinates of first coordinate
//
// x2, y2: ordinates of second coordinate
func (dw *DrawingWand) Rectangle(x1, y1, x2, y2 float64) {
	C.DrawRectangle(dw.dw, C.double(x1), C.double(y1), C.double(x2), C.double(y2))
	runtime.KeepAlive(dw)
}

// Resets the vector graphics associated with the specified wand.
func (dw *DrawingWand) ResetVectorGraphics() {
	C.DrawResetVectorGraphics(dw.dw)
	runtime.KeepAlive(dw)
}

// Applies the specified rotation to the current coordinate space.
//
// degrees: degrees of rotation
func (dw *DrawingWand) Rotate(degrees float64) {
	C.DrawRotate(dw.dw, C.double(degrees))
	runtime.KeepAlive(dw)
}

// Draws a rounted rectangle given two coordinates, x & y corner radiuses and
// using the current stroke, stroke width, and fill settings.
//
// x1, y1: ordinates of first coordinate
//
// x2, y2: ordinates of second coordinate
//
// rx, ry: radius of corner in horizontal and vertical directions
func (dw *DrawingWand) RoundRectangle(x1, y1, x2, y2, rx, ry float64) {
	C.DrawRoundRectangle(dw.dw, C.double(x1), C.double(y1), C.double(x2), C.double(y2), C.double(rx), C.double(ry))
	runtime.KeepAlive(dw)
}

// Adjusts the scaling factor to apply in the horizontal and vertical
// directions to the current coordinate space.
//
// x: horizontal scale factor
//
// y: vertical scale factor
//
func (dw *DrawingWand) Scale(x, y float64) {
	C.DrawScale(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Sets the border color to be used for drawing bordered objects.
func (dw *DrawingWand) SetBorderColor(borderWand *PixelWand) {
	C.DrawSetBorderColor(dw.dw, borderWand.pw)
	runtime.KeepAlive(dw)
}

// Associates a named clipping path with the image. Only the areas drawn on by
// the clipping path will be modified as C.ssize_t(as) it remains in effect.
// clipMaskId: name of clipping path to associate with image
func (dw *DrawingWand) SetClipPath(clipMaskId string) error {
	cstr := C.CString(clipMaskId)
	defer C.free(unsafe.Pointer(cstr))
	ok := C.DrawSetClipPath(dw.dw, cstr)
	return dw.getLastErrorIfFailed(ok)
}

// Set the polygon fill rule to be used by the clipping path.
func (dw *DrawingWand) SetClipRule(fillRule FillRule) {
	C.DrawSetClipRule(dw.dw, C.FillRule(fillRule))
	runtime.KeepAlive(dw)
}

// Sets the interpretation of clip path units.
// clipUnits: units to use
func (dw *DrawingWand) SetClipUnits(clipUnits ClipPathUnits) {
	C.DrawSetClipUnits(dw.dw, C.ClipPathUnits(clipUnits))
	runtime.KeepAlive(dw)
}

// Sets the fill color to be used for drawing filled objects.
func (dw *DrawingWand) SetFillColor(fillWand *PixelWand) {
	C.DrawSetFillColor(dw.dw, fillWand.pw)
	runtime.KeepAlive(dw)
}

// Sets the opacity to use when drawing using the fill color or fill texture.
// Fully opaque is 1.0.
func (dw *DrawingWand) SetFillOpacity(opacity float64) {
	C.DrawSetFillOpacity(dw.dw, C.double(opacity))
	runtime.KeepAlive(dw)
}

// Sets the image resolution.
//
// xRes, yRes: the image x and y resolutions
func (dw *DrawingWand) SetFontResolution(xRes, yRes float64) error {
	ok := C.DrawSetFontResolution(dw.dw, C.double(xRes), C.double(yRes))
	return dw.getLastErrorIfFailed(ok)
}

// Sets the opacity to use when drawing using the fill or stroke color or
// texture. Fully opaque is 1.0.
func (dw *DrawingWand) SetOpacity(opacity float64) {
	C.DrawSetOpacity(dw.dw, C.double(opacity))
	runtime.KeepAlive(dw)
}

// Sets the URL to use as a fill pattern for filling objects. Only local URLs
// ("#identifier") are supported at this time. These local URLs are normally
// created by defining a named fill pattern with PushPattern/PopPattern.
//
// fillUrl: URL to use to obtain fill pattern.
func (dw *DrawingWand) SetFillPatternURL(fillUrl string) error {
	cstr := C.CString(fillUrl)
	defer C.free(unsafe.Pointer(cstr))
	ok := C.DrawSetFillPatternURL(dw.dw, cstr)
	return dw.getLastErrorIfFailed(ok)
}

// Sets the fill rule to use while drawing polygons.
func (dw *DrawingWand) SetFillRule(fillRule FillRule) {
	C.DrawSetFillRule(dw.dw, C.FillRule(fillRule))
	runtime.KeepAlive(dw)
}

// Sets the fully-sepecified font to use when annotating with text.
func (dw *DrawingWand) SetFont(fontName string) error {
	csFontName := C.CString(fontName)
	defer C.free(unsafe.Pointer(csFontName))
	ok := C.DrawSetFont(dw.dw, csFontName)
	return dw.getLastErrorIfFailed(ok)
}

// Sets the font family to use when annotating with text.
func (dw *DrawingWand) SetFontFamily(fontFamily string) error {
	csFontFamily := C.CString(fontFamily)
	defer C.free(unsafe.Pointer(csFontFamily))
	ok := C.DrawSetFontFamily(dw.dw, csFontFamily)
	return dw.getLastErrorIfFailed(ok)
}

// Sets the font pointsize to use when annotating with text.
//
// pointSize: text pointsize
func (dw *DrawingWand) SetFontSize(pointSize float64) {
	C.DrawSetFontSize(dw.dw, C.double(pointSize))
	runtime.KeepAlive(dw)
}

// Sets the font stretch to use when annotating with text. The AnyStretch
// enumeration acts as a wild-card "don't care" option.
func (dw *DrawingWand) SetFontStretch(fontStretch StretchType) {
	C.DrawSetFontStretch(dw.dw, C.StretchType(fontStretch))
	runtime.KeepAlive(dw)
}

// Sets the font style to use when annotating with text. The AnyStyle
// enumeration acts as a wild-card "don't care" option.
func (dw *DrawingWand) SetFontStyle(style StyleType) {
	C.DrawSetFontStyle(dw.dw, C.StyleType(style))
	runtime.KeepAlive(dw)
}

// Sets the font weight to use when annotating with text.
//
// fontWeight: font weight (valid range 100-900)
func (dw *DrawingWand) SetFontWeight(fontWeight uint) {
	C.DrawSetFontWeight(dw.dw, C.size_t(fontWeight))
	runtime.KeepAlive(dw)
}

// Sets the text placement gravity to use when annotating with text.
func (dw *DrawingWand) SetGravity(gravity GravityType) {
	C.DrawSetGravity(dw.dw, C.GravityType(gravity))
	runtime.KeepAlive(dw)
}

// Sets the color used for stroking object outlines.
func (dw *DrawingWand) SetStrokeColor(strokeWand *PixelWand) {
	C.DrawSetStrokeColor(dw.dw, strokeWand.pw)
	runtime.KeepAlive(dw)
}

// Sets the pattern used for stroking object outlines.
//
// strokeUrl: URL specifying pattern ID (e.g. "#pattern_id")
func (dw *DrawingWand) SetStrokePatternURL(strokeUrl string) error {
	csStrokeUrl := C.CString(strokeUrl)
	defer C.free(unsafe.Pointer(csStrokeUrl))
	ok := C.DrawSetStrokePatternURL(dw.dw, csStrokeUrl)
	return dw.getLastErrorIfFailed(ok)
}

// Controls whether stroked outlines are antialiased. Stroked outlines are
// antialiased by default. When antialiasing is disabled stroked pixels are
// thresholded to determine if the stroke color or underlying canvas color
// should be used.
//
// antialias: set to false to disable antialiasing
func (dw *DrawingWand) SetStrokeAntialias(antialias bool) {
	C.DrawSetStrokeAntialias(dw.dw, b2i(antialias))
	runtime.KeepAlive(dw)
}

// Specifies the pattern of dashes and gaps used to stroke paths. The stroke
// dash array represents an array of numbers that specify the lengths of
// alternating dashes and gaps in pixels. If an odd number of values is
// provided, then the list of values is repeated to yield an even number of
// values. To remove an existing dash array, pass an empty slice. A typical
// stroke dash array might contain the members 5 3 2.
func (dw *DrawingWand) SetStrokeDashArray(dash []float64) error {
	if len(dash) == 0 {
		ok := C.DrawSetStrokeDashArray(dw.dw, C.size_t(0), nil)
		return dw.getLastErrorIfFailed(ok)
	}
	cdash := [1 << 16]C.double{}
	for k, v := range dash {
		cdash[k] = C.double(v)
	}
	ok := C.DrawSetStrokeDashArray(dw.dw, C.size_t(len(dash)), (*C.double)(&cdash[0]))
	return dw.getLastErrorIfFailed(ok)
}

// Specifies the offset into the dash pattern to start the dash.
func (dw *DrawingWand) SetStrokeDashOffset(offset float64) {
	C.DrawSetStrokeDashOffset(dw.dw, C.double(offset))
	runtime.KeepAlive(dw)
}

// Specifies the shape to be used at the end of open subpaths when they are
// stroked.
func (dw *DrawingWand) SetStrokeLineCap(lineCap LineCap) {
	C.DrawSetStrokeLineCap(dw.dw, C.LineCap(lineCap))
	runtime.KeepAlive(dw)
}

// Specifies the shape to be used at the corners of paths (or other vector
// shapes) when they are stroked.
func (dw *DrawingWand) SetStrokeLineJoin(lineJoin LineJoin) {
	C.DrawSetStrokeLineJoin(dw.dw, C.LineJoin(lineJoin))
	runtime.KeepAlive(dw)
}

// Specifies the miter limit. When two line segments meet at a sharp angle and
// miter joins have been specified for 'lineJoin', it is possible for the miter
// to extend far beyond the thickness of the line stroking the path. The
// miterLimit' imposes a limit on the ratio of the miter length to the
// 'lineWidth'.
func (dw *DrawingWand) SetStrokeMiterLimit(miterLimit uint) {
	C.DrawSetStrokeMiterLimit(dw.dw, C.size_t(miterLimit))
	runtime.KeepAlive(dw)
}

// Specifies the opacity of stroked object outlines.
//
// opacity: stroke opacity. The value 1.0 is opaque.
func (dw *DrawingWand) SetStrokeOpacity(opacity float64) {
	C.DrawSetStrokeOpacity(dw.dw, C.double(opacity))
	runtime.KeepAlive(dw)
}

// Sets the width of the stroke used to draw object outlines.
func (dw *DrawingWand) SetStrokeWidth(width float64) {
	C.DrawSetStrokeWidth(dw.dw, C.double(width))
	runtime.KeepAlive(dw)
}

// Specifies a text alignment to be applied when annotating with text.
func (dw *DrawingWand) SetTextAlignment(alignment AlignType) {
	C.DrawSetTextAlignment(dw.dw, C.AlignType(alignment))
	runtime.KeepAlive(dw)
}

// Controls whether text is antialiased. Text is antialiased by default.
func (dw *DrawingWand) SetTextAntialias(antialias bool) {
	C.DrawSetTextAntialias(dw.dw, b2i(antialias))
	runtime.KeepAlive(dw)
}

// Specifies a decoration to be applied when annotating with text.
func (dw *DrawingWand) SetTextDecoration(decoration DecorationType) {
	C.DrawSetTextDecoration(dw.dw, C.DecorationType(decoration))
	runtime.KeepAlive(dw)
}

// Specifies the code set to use for text annotations. The only character
// encoding which may be specified at this time is "UTF-8" for representing
// Unicode as a sequence of bytes. Specify an empty string to set text
// encoding to the system's default. Successful text annotation using Unicode
// may require fonts designed to support Unicode.
func (dw *DrawingWand) SetTextEncoding(encoding string) {
	csencoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(csencoding))
	C.DrawSetTextEncoding(dw.dw, csencoding)
	runtime.KeepAlive(dw)
}

// Sets the spacing between characters in text.
func (dw *DrawingWand) SetTextKerning(kerning float64) {
	C.DrawSetTextKerning(dw.dw, C.double(kerning))
	runtime.KeepAlive(dw)
}

// Sets the spacing between line in text.
func (dw *DrawingWand) SetTextInterlineSpacing(spacing float64) {
	C.DrawSetTextInterlineSpacing(dw.dw, C.double(spacing))
	runtime.KeepAlive(dw)
}

// Sets the spacing between words in text.
func (dw *DrawingWand) SetTextInterwordSpacing(spacing float64) {
	C.DrawSetTextInterwordSpacing(dw.dw, C.double(spacing))
	runtime.KeepAlive(dw)
}

// Specifies the color of a background rectangle to place under text
// annotations.
func (dw *DrawingWand) SetTextUnderColor(underWand *PixelWand) {
	C.DrawSetTextUnderColor(dw.dw, underWand.pw)
	runtime.KeepAlive(dw)
}

// Sets the vector graphics associated with the specified wand. Use this method
// with GetVectorGraphics() as a method to persist the vector graphics state.
func (dw *DrawingWand) SetVectorGraphics(xml string) error {
	csxml := C.CString(xml)
	defer C.free(unsafe.Pointer(csxml))
	ok := C.DrawSetVectorGraphics(dw.dw, csxml)
	return dw.getLastErrorIfFailed(ok)
}

// Skews the current coordinate system in the horizontal direction.
//
// degrees: number of degrees to skew the coordinates
func (dw *DrawingWand) SkewX(degrees float64) {
	C.DrawSkewX(dw.dw, C.double(degrees))
	runtime.KeepAlive(dw)
}

// Skews the current coordinate system in the vertical direction.
//
// degrees: number of degrees to skew the coordinates
func (dw *DrawingWand) SkewY(degrees float64) {
	C.DrawSkewY(dw.dw, C.double(degrees))
	runtime.KeepAlive(dw)
}

// Applies a translation to the current coordinate system which moves the
// coordinate system origin to the specified coordinate.
//
// x, y: new x, y ordinate for coordinate system origin
func (dw *DrawingWand) Translate(x, y float64) {
	C.DrawTranslate(dw.dw, C.double(x), C.double(y))
	runtime.KeepAlive(dw)
}

// Sets the overall canvas size to be recorded with the drawing vector data.
// Usually this will be specified using the same size as the canvas image.
// When the vector data is saved to SVG or MVG formats, the viewbox is use to
// specify the size of the canvas image that a viewer will render the vector
// data on.
//
// x1: left x ordinate
//
// y1: top y ordinate
//
// x2: right x ordinate
//
// y2: bottom y ordinate
func (dw *DrawingWand) SetViewbox(x1, y1, x2, y2 float64) {
	C.DrawSetViewbox(dw.dw, C.double(x1), C.double(y1), C.double(x2), C.double(y2))
	runtime.KeepAlive(dw)
}

// Returns true if the wand is verified as a drawing wand.
func (dw *DrawingWand) IsVerified() bool {
	if dw.dw == nil {
		return false
	}
	ret := 1 == C.IsDrawingWand(dw.dw)
	runtime.KeepAlive(dw)
	return ret
}

// Returns the current drawing wand.
func (dw *DrawingWand) PeekDrawingWand() *DrawInfo {
	ret := &DrawInfo{C.PeekDrawingWand(dw.dw)}
	runtime.KeepAlive(dw)
	return ret
}

// Destroys the current drawing wand and returns to the previously pushed
// drawing wand. Multiple drawing wands may exist. It is an error to attempt
// to pop more drawing wands than have been pushed, and it is proper form to
// pop all drawing wands which have been pushed.
func (dw *DrawingWand) PopDrawingWand() error {
	ok := C.PopDrawingWand(dw.dw)
	return dw.getLastErrorIfFailed(ok)
}

// Clones the current drawing wand to create a new drawing wand. The original
// drawing wand(s) may be returned to by invoking PopDrawingWand(). The drawing
// wands are stored on a drawing wand stack. For every Pop there must have
// already been an equivalent Push.
func (dw *DrawingWand) PushDrawingWand() error {
	ok := C.PushDrawingWand(dw.dw)
	return dw.getLastErrorIfFailed(ok)
}
