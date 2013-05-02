package imagick
/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"
import (
	"unsafe"
)

type DrawingWand struct {
	dw *DrawingWand
}

// Returns a drawing wand required for all other methods in the API.
func NewDrawingWand() *DrawingWand {
	return &DrawingWand{C.NewDrawingWand()}
}

// Clears resources associated with the drawing wand.
func (dw *DrawingWand) Clear() {
	C.ClearDrawingWand(dw.dw)
}
// Makes an exact copy of the specified wand.
func (dw *DrawingWand) Clone() *DrawingWand {
	return &DrawingWand{C.CloneDrawingWand(dw.dw)}
}

// Frees all resources associated with the drawing wand. Once the drawing wand 
// has been freed, it should not be used and further unless it re-allocated.
func (dw *DrawingWand) Destroy() {
	dw.dw = DestroyDrawingWand(dw.dw)
}

// Adjusts the current affine transformation matrix with the specified affine 
// transformation matrix. Note that the current affine transform is adjusted 
// rather than replaced.
// 
// affine: Affine matrix parameters
// 
func (dw *DrawingWand) DrawAffine(affine *AffineMatrix) {
	C.DrawAffine(dw.dw, affine.am)
}

// Draws text on the image.
// x: x ordinate to left of text
// y: y ordinate to text baseline
// text: text to draw
func (dw *DrawingWand) DrawAnnotation(x, y float64, text string) {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	C.DrawAnnotation(dw.dw, C.double(x), C.double(y), cstext)
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
func (dw *DrawingWand) DrawArc(sx, sy, ex, ey, sd, ed float64) {
	C.DrawArc(dw.dw, C.double(sx),C.double(sy), C.double(ex),C.double(ey),C.double(sd),C.double(ed))
}

// Draws a bezier curve through a set of points on the image.
func (dw *DrawingWand) DrawBezier(coordinates []*PointInfo) {
	// TODO will this work? probably not
	C.DrawBezier(dw.dw, C.size_t(len(coordinates)), (*C.PointInfo)(&coordinates[0].pi)
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
func (dw *DrawingWand) DrawCircle(ox, oy, px, py float64) {
	C.DrawCircle(dw.dw, C.double(ox),C.double(oy),C.double(px), C.double(py))
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
func (dw *DrawingWand) DrawComposite(compose CompositeOperator, x, y, width, height float64, mw *MagickWand) error {
	C.DrawComposite(dw.dw, C.CompositeOperator(compose),C.double(x),C.double(y),C.double(width),C.double(height),mw.mw)
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
func (dw *DrawingWand) DrawColor(x, y float64, pm PaintMethod) {
	C.DrawColor(dw.dw, C.double(x),C.double(y), C.PaintMethod(pm))
}


//DrawComment
//DrawComment() adds a comment to a vector output stream.
func (dw *DrawingWand)
//C.DrawComment(dw.dw, const char *comment)
//comment
//comment text
//DrawEllipse
//DrawEllipse() draws an ellipse on the image.
func (dw *DrawingWand)
//C.DrawEllipse(dw.dw, C.double(ox),C.double(oy),
//C.double(rx),C.double(ry),C.double(start),C.double(end))
//ox
//origin x ordinate
//oy
//origin y ordinate
//rx
//radius in x
//ry
//radius in y
//start
//starting rotation in degrees
//end
//ending rotation in degrees
//DrawGetBorderColor
//DrawGetBorderColor() returns the border color used for drawing bordered objects.
func (dw *DrawingWand)
//C.DrawGetBorderColor(dw.dw,
//PixelWand *border_color)
//border_color
//Return the border color.
//DrawGetClipPath
//DrawGetClipPath() obtains the current clipping path ID. The value returned must be deallocated by the user when it is no longer needed.
func (dw *DrawingWand)
//char *DrawGetClipPath(dw.dw)
//DrawGetClipRule
//DrawGetClipRule() returns the current polygon fill rule to be used by the clipping path.
func (dw *DrawingWand)
//FillRule DrawGetClipRule(dw.dw)
//DrawGetClipUnits
//DrawGetClipUnits() returns the interpretation of clip path units.
func (dw *DrawingWand)
//ClipPathUnits DrawGetClipUnits(dw.dw)
//DrawGetException
//DrawGetException() returns the severity, reason, and description of any error that occurs when using other methods in this API.
func (dw *DrawingWand)
//char *DrawGetException(const DrawWand *wand,
//ExceptionType *severity)
//severity
//the severity of the error is returned here.
//DrawGetExceptionType
//DrawGetExceptionType() the exception type associated with the wand. If no exception has occurred, UndefinedExceptionType is returned.
func (dw *DrawingWand)
//ExceptionType DrawGetExceptionType(const DrawWand *wand)
//DrawGetFillColor
//DrawGetFillColor() returns the fill color used for drawing filled objects.
func (dw *DrawingWand)
//C.DrawGetFillColor(dw.dw,
//PixelWand *fill_color)
//fill_color
//Return the fill color.
//DrawGetFillOpacity
//DrawGetFillOpacity() returns the opacity used when drawing using the fill color or fill texture. Fully opaque is 1.0.
func (dw *DrawingWand)
//double DrawGetFillOpacity(dw.dw)
//DrawGetFillRule
//DrawGetFillRule() returns the fill rule used while drawing polygons.
func (dw *DrawingWand)
//FillRule DrawGetFillRule(dw.dw)
//DrawGetFont
//DrawGetFont() returns a null-terminaged string specifying the font used when annotating with text. The value returned must be freed by the user when no longer needed.
func (dw *DrawingWand)
//char *DrawGetFont(dw.dw)
//DrawGetFontFamily
//DrawGetFontFamily() returns the font family to use when annotating with text. The value returned must be freed by the user when it is no longer needed.
func (dw *DrawingWand)
//char *DrawGetFontFamily(dw.dw)
//DrawGetFontResolution
//DrawGetFontResolution() gets the image X and Y resolution.
func (dw *DrawingWand)
//DrawBooleanType DrawGetFontResolution(dw.dw,
//double *x,double *y)
//x
//the x-resolution.
//y
//the y-resolution.
//DrawGetFontSize
//DrawGetFontSize() returns the font pointsize used when annotating with text.
func (dw *DrawingWand)
//double DrawGetFontSize(dw.dw)
//DrawGetFontStretch
//DrawGetFontStretch() returns the font stretch used when annotating with text.
func (dw *DrawingWand)
//StretchType DrawGetFontStretch(dw.dw)
//DrawGetFontStyle
//DrawGetFontStyle() returns the font style used when annotating with text.
func (dw *DrawingWand)
//StyleType DrawGetFontStyle(dw.dw)
//DrawGetFontWeight
//DrawGetFontWeight() returns the font weight used when annotating with text.
func (dw *DrawingWand)
//size_t DrawGetFontWeight(dw.dw)
//DrawGetGravity
//DrawGetGravity() returns the text placement gravity used when annotating with text.
func (dw *DrawingWand)
//GravityType DrawGetGravity(dw.dw)
//DrawGetOpacity
//DrawGetOpacity() returns the opacity used when drawing with the fill or stroke color or texture. Fully opaque is 1.0.
func (dw *DrawingWand)
//double DrawGetOpacity(dw.dw)
//DrawGetStrokeAntialias
//DrawGetStrokeAntialias() returns the current stroke antialias setting. Stroked outlines are antialiased by default. When antialiasing is disabled stroked pixels are thresholded to determine if the stroke color or underlying canvas color should be used.
func (dw *DrawingWand)
//MagickBooleanType DrawGetStrokeAntialias(dw.dw)
//DrawGetStrokeColor
//DrawGetStrokeColor() returns the color used for stroking object outlines.
func (dw *DrawingWand)
//C.DrawGetStrokeColor(dw.dw,
//PixelWand *stroke_color)
//stroke_color
//Return the stroke color.
//DrawGetStrokeDashArray
//DrawGetStrokeDashArray() returns an array representing the pattern of dashes and gaps used to stroke paths (see DrawSetStrokeDashArray). The array must be freed once it is no longer required by the user.
func (dw *DrawingWand)
//double *DrawGetStrokeDashArray(dw.dw,
//size_t *number_elements)
//number_elements
//address to place number of elements in dash array
//DrawGetStrokeDashOffset
//DrawGetStrokeDashOffset() returns the offset into the dash pattern to start the dash.
func (dw *DrawingWand)
//double DrawGetStrokeDashOffset(dw.dw)
//DrawGetStrokeLineCap
//DrawGetStrokeLineCap() returns the shape to be used at the end of open subpaths when they are stroked. Values of LineCap are UndefinedCap, ButtCap, RoundCap, and SquareCap.
func (dw *DrawingWand)
//LineCap DrawGetStrokeLineCap(dw.dw)
//DrawGetStrokeLineJoin
//DrawGetStrokeLineJoin() returns the shape to be used at the corners of paths (or other vector shapes) when they are stroked. Values of LineJoin are UndefinedJoin, MiterJoin, RoundJoin, and BevelJoin.
func (dw *DrawingWand)
//LineJoin DrawGetStrokeLineJoin(dw.dw)
//DrawGetStrokeMiterLimit
//DrawGetStrokeMiterLimit() returns the miter limit. When two line segments meet at a sharp angle and miter joins have been specified for 'lineJoin', it is possible for the miter to extend far beyond the thickness of the line stroking the path. The miterLimit' imposes a limit on the ratio of the miter length to the 'lineWidth'.
func (dw *DrawingWand)
//size_t DrawGetStrokeMiterLimit(dw.dw)
//DrawGetStrokeOpacity
//DrawGetStrokeOpacity() returns the opacity of stroked object outlines.
func (dw *DrawingWand)
//double DrawGetStrokeOpacity(dw.dw)
//DrawGetStrokeWidth
//DrawGetStrokeWidth() returns the width of the stroke used to draw object outlines.
func (dw *DrawingWand)
//double DrawGetStrokeWidth(dw.dw)
//DrawGetTextAlignment
//DrawGetTextAlignment() returns the alignment applied when annotating with text.
func (dw *DrawingWand)
//AlignType DrawGetTextAlignment(dw.dw)
//DrawGetTextAntialias
//DrawGetTextAntialias() returns the current text antialias setting, which determines whether text is antialiased. Text is antialiased by default.
func (dw *DrawingWand)
//MagickBooleanType DrawGetTextAntialias(dw.dw)
//DrawGetTextDecoration
//DrawGetTextDecoration() returns the decoration applied when annotating with text.
func (dw *DrawingWand)
//DecorationType DrawGetTextDecoration(dw.dw)
//DrawGetTextEncoding
//DrawGetTextEncoding() returns a null-terminated string which specifies the code set used for text annotations. The string must be freed by the user once it is no longer required.
func (dw *DrawingWand)
//char *DrawGetTextEncoding(dw.dw)
//DrawGetTextKerning
//DrawGetTextKerning() gets the spacing between characters in text.
func (dw *DrawingWand)
//double DrawGetTextKerning(dw.dw)
//DrawGetTextInterlineSpacing
//DrawGetTextInterlineSpacing() gets the spacing between lines in text.
func (dw *DrawingWand)
//double DrawGetTextInterwordSpacing(dw.dw)
//DrawGetTextInterwordSpacing
//DrawGetTextInterwordSpacing() gets the spacing between words in text.
func (dw *DrawingWand)
//double DrawGetTextInterwordSpacing(dw.dw)
//DrawGetVectorGraphics
//DrawGetVectorGraphics() returns a null-terminated string which specifies the vector graphics generated by any graphics calls made since the wand was instantiated. The string must be freed by the user once it is no longer required.
func (dw *DrawingWand)
//char *DrawGetVectorGraphics(dw.dw)
//DrawGetTextUnderColor
//DrawGetTextUnderColor() returns the color of a background rectangle to place under text annotations.
func (dw *DrawingWand)
//C.DrawGetTextUnderColor(dw.dw,
//PixelWand *under_color)
//under_color
//Return the under color.
//DrawLine
//DrawLine() draws a line on the image using the current stroke color, stroke opacity, and stroke width.
func (dw *DrawingWand)
//C.DrawLine(dw.dw, C.double(sx),C.double(sy),
//C.double(ex),C.double(ey))
//sx
//starting x ordinate
//sy
//starting y ordinate
//ex
//ending x ordinate
//ey
//ending y ordinate
//DrawMatte
//DrawMatte() paints on the image's opacity channel in order to set effected pixels to transparent. to influence the opacity of pixels. The available paint methods are:
func (dw *DrawingWand)
//PointMethod: Select the target pixel
//ReplaceMethod: Select any pixel that matches the target pixel.
//FloodfillMethod: Select the target pixel and matching neighbors.
//FillToBorderMethod: Select the target pixel and neighbors not matching
//border color.
//ResetMethod: Select all pixels.
//C.DrawMatte(dw.dw, C.double(x),C.double(y),
//const PaintMethod paint_method)
//x
//x ordinate
//y
//y ordinate
//paint_method
//paint method.
//DrawPathClose
//DrawPathClose() adds a path element to the current path which closes the current subpath by drawing a straight line from the current point to the current subpath's most recent starting point (usually, the most recent moveto point).
func (dw *DrawingWand)
//C.DrawPathClose(dw.dw)
//DrawPathCurveToAbsolute
//DrawPathCurveToAbsolute() draws a cubic Bezier curve from the current point to (x,y) using (x1,y1) as the control point at the beginning of the curve and (x2,y2) as the control point at the end of the curve using absolute coordinates. At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToAbsolute(dw.dw, C.double(x1),
//C.double(y1),C.double(x2),C.double(y2),C.double(x),
//C.double(y))
//x1
//x ordinate of control point for curve beginning
//y1
//y ordinate of control point for curve beginning
//x2
//x ordinate of control point for curve ending
//y2
//y ordinate of control point for curve ending
//x
//x ordinate of the end of the curve
//y
//y ordinate of the end of the curve
//DrawPathCurveToRelative
//DrawPathCurveToRelative() draws a cubic Bezier curve from the current point to (x,y) using (x1,y1) as the control point at the beginning of the curve and (x2,y2) as the control point at the end of the curve using relative coordinates. At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToRelative(dw.dw, C.double(x1),
//C.double(y1),C.double(x2),C.double(y2),C.double(x),
//C.double(y))
//x1
//x ordinate of control point for curve beginning
//y1
//y ordinate of control point for curve beginning
//x2
//x ordinate of control point for curve ending
//y2
//y ordinate of control point for curve ending
//x
//x ordinate of the end of the curve
//y
//y ordinate of the end of the curve
//DrawPathCurveToQuadraticBezierAbsolute
//DrawPathCurveToQuadraticBezierAbsolute() draws a quadratic Bezier curve from the current point to (x,y) using (x1,y1) as the control point using absolute coordinates. At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToQuadraticBezierAbsolute(dw.dw, 
//C.double(x1),C.double(y1),C.double(x),C.double(y))
//x1
//x ordinate of the control point
//y1
//y ordinate of the control point
//x
//x ordinate of final point
//y
//y ordinate of final point
//DrawPathCurveToQuadraticBezierRelative
//DrawPathCurveToQuadraticBezierRelative() draws a quadratic Bezier curve from the current point to (x,y) using (x1,y1) as the control point using relative coordinates. At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToQuadraticBezierRelative(dw.dw, 
//C.double(x1),C.double(y1),C.double(x),C.double(y))
//x1
//x ordinate of the control point
//y1
//y ordinate of the control point
//x
//x ordinate of final point
//y
//y ordinate of final point
//DrawPathCurveToQuadraticBezierSmoothAbsolute
//DrawPathCurveToQuadraticBezierSmoothAbsolute() draws a quadratic Bezier curve (using absolute coordinates) from the current point to (x,y). The control point is assumed to be the reflection of the control point on the previous command relative to the current point. (If there is no previous command or if the previous command was not a DrawPathCurveToQuadraticBezierAbsolute, DrawPathCurveToQuadraticBezierRelative, DrawPathCurveToQuadraticBezierSmoothAbsolute or DrawPathCurveToQuadraticBezierSmoothRelative, assume the control point is coincident with the current point.). At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToQuadraticBezierSmoothAbsolute(
//DrawingWand *wand,C.double(x),C.double(y))
//x
//x ordinate of final point
//y
//y ordinate of final point
//DrawPathCurveToQuadraticBezierSmoothRelative
//DrawPathCurveToQuadraticBezierSmoothRelative() draws a quadratic Bezier curve (using relative coordinates) from the current point to (x,y). The control point is assumed to be the reflection of the control point on the previous command relative to the current point. (If there is no previous command or if the previous command was not a DrawPathCurveToQuadraticBezierAbsolute, DrawPathCurveToQuadraticBezierRelative, DrawPathCurveToQuadraticBezierSmoothAbsolute or DrawPathCurveToQuadraticBezierSmoothRelative, assume the control point is coincident with the current point.). At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToQuadraticBezierSmoothRelative(dw.dw, 
//C.double(x),C.double(y))
//x
//x ordinate of final point
//y
//y ordinate of final point
//DrawPathCurveToSmoothAbsolute
//DrawPathCurveToSmoothAbsolute() draws a cubic Bezier curve from the current point to (x,y) using absolute coordinates. The first control point is assumed to be the reflection of the second control point on the previous command relative to the current point. (If there is no previous command or if the previous command was not an DrawPathCurveToAbsolute, DrawPathCurveToRelative, DrawPathCurveToSmoothAbsolute or DrawPathCurveToSmoothRelative, assume the first control point is coincident with the current point.) (x2,y2) is the second control point (i.e., the control point at the end of the curve). At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToSmoothAbsolute(dw.dw, 
//C.double(x2),C.double(y2),C.double(x),C.double(y))
//x2
//x ordinate of second control point
//y2
//y ordinate of second control point
//x
//x ordinate of termination point
//y
//y ordinate of termination point
//DrawPathCurveToSmoothRelative
//DrawPathCurveToSmoothRelative() draws a cubic Bezier curve from the current point to (x,y) using relative coordinates. The first control point is assumed to be the reflection of the second control point on the previous command relative to the current point. (If there is no previous command or if the previous command was not an DrawPathCurveToAbsolute, DrawPathCurveToRelative, DrawPathCurveToSmoothAbsolute or DrawPathCurveToSmoothRelative, assume the first control point is coincident with the current point.) (x2,y2) is the second control point (i.e., the control point at the end of the curve). At the end of the command, the new current point becomes the final (x,y) coordinate pair used in the polybezier.
func (dw *DrawingWand)
//C.DrawPathCurveToSmoothRelative(dw.dw, 
//C.double(x2),C.double(y2),C.double(x),C.double(y))
//x2
//x ordinate of second control point
//y2
//y ordinate of second control point
//x
//x ordinate of termination point
//y
//y ordinate of termination point
//DrawPathEllipticArcAbsolute
//DrawPathEllipticArcAbsolute() draws an elliptical arc from the current point to (x, y) using absolute coordinates. The size and orientation of the ellipse are defined by two radii (rx, ry) and an xAxisRotation, which indicates how the ellipse as a whole is rotated relative to the current coordinate system. The center (cx, cy) of the ellipse is calculated automagically to satisfy the constraints imposed by the other parameters. largeArcFlag and sweepFlag contribute to the automatic calculations and help determine how the arc is drawn. If largeArcFlag is true then draw the larger of the available arcs. If sweepFlag is true, then draw the arc matching a clock-wise rotation.
func (dw *DrawingWand)
//C.DrawPathEllipticArcAbsolute(dw.dw, 
//C.double(rx),C.double(ry),C.double(x_axis_rotation),
//const MagickBooleanType large_arc_flag,
//const MagickBooleanType sweep_flag,C.double(x),C.double(y))
//rx
//x radius
//ry
//y radius
//x_axis_rotation
//indicates how the ellipse as a whole is rotated relative to the current coordinate system
//large_arc_flag
//If non-zero (true) then draw the larger of the available arcs
//sweep_flag
//If non-zero (true) then draw the arc matching a clock-wise rotation
//DrawPathEllipticArcRelative
//DrawPathEllipticArcRelative() draws an elliptical arc from the current point to (x, y) using relative coordinates. The size and orientation of the ellipse are defined by two radii (rx, ry) and an xAxisRotation, which indicates how the ellipse as a whole is rotated relative to the current coordinate system. The center (cx, cy) of the ellipse is calculated automagically to satisfy the constraints imposed by the other parameters. largeArcFlag and sweepFlag contribute to the automatic calculations and help determine how the arc is drawn. If largeArcFlag is true then draw the larger of the available arcs. If sweepFlag is true, then draw the arc matching a clock-wise rotation.
func (dw *DrawingWand)
//C.DrawPathEllipticArcRelative(dw.dw, 
//C.double(rx),C.double(ry),C.double(x_axis_rotation),
//const MagickBooleanType large_arc_flag,
//const MagickBooleanType sweep_flag,C.double(x),C.double(y))
//rx
//x radius
//ry
//y radius
//x_axis_rotation
//indicates how the ellipse as a whole is rotated relative to the current coordinate system
//large_arc_flag
//If non-zero (true) then draw the larger of the available arcs
//sweep_flag
//If non-zero (true) then draw the arc matching a clock-wise rotation
//DrawPathFinish
//DrawPathFinish() terminates the current path.
func (dw *DrawingWand)
//C.DrawPathFinish(dw.dw)
//DrawPathLineToAbsolute
//DrawPathLineToAbsolute() draws a line path from the current point to the given coordinate using absolute coordinates. The coordinate then becomes the new current point.
func (dw *DrawingWand)
//C.DrawPathLineToAbsolute(dw.dw, C.double(x),
//C.double(y))
//x
//target x ordinate
//y
//target y ordinate
//DrawPathLineToRelative
//DrawPathLineToRelative() draws a line path from the current point to the given coordinate using relative coordinates. The coordinate then becomes the new current point.
func (dw *DrawingWand)
//C.DrawPathLineToRelative(dw.dw, C.double(x),
//C.double(y))
//x
//target x ordinate
//y
//target y ordinate
//DrawPathLineToHorizontalAbsolute
//DrawPathLineToHorizontalAbsolute() draws a horizontal line path from the current point to the target point using absolute coordinates. The target point then becomes the new current point.
func (dw *DrawingWand)
//C.DrawPathLineToHorizontalAbsolute(dw.dw, 
//C.double(x))
//x
//target x ordinate
//DrawPathLineToHorizontalRelative
//DrawPathLineToHorizontalRelative() draws a horizontal line path from the current point to the target point using relative coordinates. The target point then becomes the new current point.
func (dw *DrawingWand)
//C.DrawPathLineToHorizontalRelative(dw.dw, 
//C.double(x))
//x
//target x ordinate
//DrawPathLineToVerticalAbsolute
//DrawPathLineToVerticalAbsolute() draws a vertical line path from the current point to the target point using absolute coordinates. The target point then becomes the new current point.
func (dw *DrawingWand)
//C.DrawPathLineToVerticalAbsolute(dw.dw, 
//C.double(y))
//y
//target y ordinate
//DrawPathLineToVerticalRelative
//DrawPathLineToVerticalRelative() draws a vertical line path from the current point to the target point using relative coordinates. The target point then becomes the new current point.
func (dw *DrawingWand)
//C.DrawPathLineToVerticalRelative(dw.dw, 
//C.double(y))
//y
//target y ordinate
//DrawPathMoveToAbsolute
//DrawPathMoveToAbsolute() starts a new sub-path at the given coordinate using absolute coordinates. The current point then becomes the specified coordinate.
func (dw *DrawingWand)
//C.DrawPathMoveToAbsolute(dw.dw, C.double(x),
//C.double(y))
//x
//target x ordinate
//y
//target y ordinate
//DrawPathMoveToRelative
//DrawPathMoveToRelative() starts a new sub-path at the given coordinate using relative coordinates. The current point then becomes the specified coordinate.
func (dw *DrawingWand)
//C.DrawPathMoveToRelative(dw.dw, C.double(x),
//C.double(y))
//x
//target x ordinate
//y
//target y ordinate
//DrawPathStart
//DrawPathStart() declares the start of a path drawing list which is terminated by a matching DrawPathFinish() command. All other DrawPath commands must be enclosed between a DrawPathStart() and a DrawPathFinish() command. This is because path drawing commands are subordinate commands and they do not function by themselves.
func (dw *DrawingWand)
//C.DrawPathStart(dw.dw)
//DrawPoint
//DrawPoint() draws a point using the current fill color.
func (dw *DrawingWand)
//C.DrawPoint(dw.dw, C.double(x),C.double(y))
//x
//target x coordinate
//y
//target y coordinate
//DrawPolygon
//DrawPolygon() draws a polygon using the current stroke, stroke width, and fill color or texture, using the specified array of coordinates.
func (dw *DrawingWand)
//C.DrawPolygon(dw.dw, 
//C.size_t(number_coordinates),const PointInfo *coordinates)
//number_coordinates
//number of coordinates
//coordinates
//coordinate array
//DrawPolyline
//DrawPolyline() draws a polyline using the current stroke, stroke width, and fill color or texture, using the specified array of coordinates.
func (dw *DrawingWand)
//C.DrawPolyline(dw.dw, 
//C.size_t(number_coordinates),const PointInfo *coordinates)
//number_coordinates
//number of coordinates
//coordinates
//coordinate array
//DrawPopClipPath
//DrawPopClipPath() terminates a clip path definition.
func (dw *DrawingWand)
//C.DrawPopClipPath(dw.dw)
//DrawPopDefs
//DrawPopDefs() terminates a definition list.
func (dw *DrawingWand)
//C.DrawPopDefs(dw.dw)
//DrawPopPattern
//DrawPopPattern() terminates a pattern definition.
func (dw *DrawingWand)
//MagickBooleanType DrawPopPattern(dw.dw)
//DrawPushClipPath
//DrawPushClipPath() starts a clip path definition which is comprized of any number of drawing commands and terminated by a DrawPopClipPath() command.
func (dw *DrawingWand)
//C.DrawPushClipPath(dw.dw, const char *clip_mask_id)
//clip_mask_id
//string identifier to associate with the clip path for later use.
//DrawPushDefs
//DrawPushDefs() indicates that commands up to a terminating DrawPopDefs() command create named elements (e.g. clip-paths, textures, etc.) which may safely be processed earlier for the sake of efficiency.
func (dw *DrawingWand)
//C.DrawPushDefs(dw.dw)
//DrawPushPattern
//DrawPushPattern() indicates that subsequent commands up to a DrawPopPattern() command comprise the definition of a named pattern. The pattern space is assigned top left corner coordinates, a width and height, and becomes its own drawing space. Anything which can be drawn may be used in a pattern definition. Named patterns may be used as stroke or brush definitions.
func (dw *DrawingWand)
//MagickBooleanType DrawPushPattern(dw.dw, 
//const char *pattern_id,C.double(x),C.double(y),
//C.double(width),C.double(height))
//pattern_id
//pattern identification for later reference
//x
//x ordinate of top left corner
//y
//y ordinate of top left corner
//width
//width of pattern space
//height
//height of pattern space
//DrawRectangle
//DrawRectangle() draws a rectangle given two coordinates and using the current stroke, stroke width, and fill settings.
func (dw *DrawingWand)
//C.DrawRectangle(dw.dw, C.double(x1),
//C.double(y1),C.double(x2),C.double(y2))
//x1
//x ordinate of first coordinate
//y1
//y ordinate of first coordinate
//x2
//x ordinate of second coordinate
//y2
//y ordinate of second coordinate
//DrawResetVectorGraphics
//DrawResetVectorGraphics() resets the vector graphics associated with the specified wand.
func (dw *DrawingWand)
//C.DrawResetVectorGraphics(dw.dw)
//DrawRotate
//DrawRotate() applies the specified rotation to the current coordinate space.
func (dw *DrawingWand)
//C.DrawRotate(dw.dw, C.double(degrees))
//degrees
//degrees of rotation
//DrawRoundRectangle
//DrawRoundRectangle() draws a rounted rectangle given two coordinates, x & y corner radiuses and using the current stroke, stroke width, and fill settings.
func (dw *DrawingWand)
//C.DrawRoundRectangle(dw.dw, double x1,double y1,
//double x2,double y2,double rx,double ry)
//x1
//x ordinate of first coordinate
//y1
//y ordinate of first coordinate
//x2
//x ordinate of second coordinate
//y2
//y ordinate of second coordinate
//rx
//radius of corner in horizontal direction
//ry
//radius of corner in vertical direction
//DrawScale
//DrawScale() adjusts the scaling factor to apply in the horizontal and vertical directions to the current coordinate space.
func (dw *DrawingWand)
//C.DrawScale(dw.dw, C.double(x),C.double(y))
//x
//horizontal scale factor
//y
//vertical scale factor
//DrawSetBorderColor
//DrawSetBorderColor() sets the border color to be used for drawing bordered objects.
func (dw *DrawingWand)
//C.DrawSetBorderColor(dw.dw, const PixelWand *border_wand)
//border_wand
//border wand.
//DrawSetClipPath
//DrawSetClipPath() associates a named clipping path with the image. Only the areas drawn on by the clipping path will be modified as C.ssize_t(as) it remains in effect.
func (dw *DrawingWand)
//MagickBooleanType DrawSetClipPath(dw.dw, 
//const char *clip_mask)
//clip_mask
//name of clipping path to associate with image
//DrawSetClipRule
//DrawSetClipRule() set the polygon fill rule to be used by the clipping path.
func (dw *DrawingWand)
//C.DrawSetClipRule(dw.dw, const FillRule fill_rule)
//fill_rule
//fill rule (EvenOddRule or NonZeroRule)
//DrawSetClipUnits
//DrawSetClipUnits() sets the interpretation of clip path units.
func (dw *DrawingWand)
//C.DrawSetClipUnits(dw.dw, 
//const ClipPathUnits clip_units)
//clip_units
//units to use (UserSpace, UserSpaceOnUse, or ObjectBoundingBox)
//DrawSetFillColor
//DrawSetFillColor() sets the fill color to be used for drawing filled objects.
func (dw *DrawingWand)
//C.DrawSetFillColor(dw.dw, const PixelWand *fill_wand)
//fill_wand
//fill wand.
//DrawSetFillOpacity
//DrawSetFillOpacity() sets the opacity to use when drawing using the fill color or fill texture. Fully opaque is 1.0.
func (dw *DrawingWand)
//C.DrawSetFillOpacity(dw.dw, C.double(fill_opacity))
//fill_opacity
//fill opacity
//DrawSetFontResolution
//DrawSetFontResolution() sets the image resolution.
func (dw *DrawingWand)
//MagickBooleanType DrawSetFontResolution(dw.dw, 
//C.double(x_resolution),C.double(y_resolution))
//x_resolution
//the image x resolution.
//y_resolution
//the image y resolution.
//DrawSetOpacity
//DrawSetOpacity() sets the opacity to use when drawing using the fill or stroke color or texture. Fully opaque is 1.0.
func (dw *DrawingWand)
//C.DrawSetOpacity(dw.dw, C.double(opacity))
//opacity
//fill opacity
//DrawSetFillPatternURL
//DrawSetFillPatternURL() sets the URL to use as a fill pattern for filling objects. Only local URLs ("#identifier") are supported at this time. These local URLs are normally created by defining a named fill pattern with DrawPushPattern/DrawPopPattern.
func (dw *DrawingWand)
//MagickBooleanType DrawSetFillPatternURL(dw.dw, 
//const char *fill_url)
//fill_url
//URL to use to obtain fill pattern.
//DrawSetFillRule
//DrawSetFillRule() sets the fill rule to use while drawing polygons.
func (dw *DrawingWand)
//C.DrawSetFillRule(dw.dw, const FillRule fill_rule)
//fill_rule
//fill rule (EvenOddRule or NonZeroRule)
//DrawSetFont
//DrawSetFont() sets the fully-sepecified font to use when annotating with text.
func (dw *DrawingWand)
//MagickBooleanType DrawSetFont(dw.dw, const char *font_name)
//font_name
//font name
//DrawSetFontFamily
//DrawSetFontFamily() sets the font family to use when annotating with text.
func (dw *DrawingWand)
//MagickBooleanType DrawSetFontFamily(dw.dw, 
//const char *font_family)
//font_family
//font family
//DrawSetFontSize
//DrawSetFontSize() sets the font pointsize to use when annotating with text.
func (dw *DrawingWand)
//C.DrawSetFontSize(dw.dw, C.double(pointsize))
//pointsize
//text pointsize
//DrawSetFontStretch
//DrawSetFontStretch() sets the font stretch to use when annotating with text. The AnyStretch enumeration acts as a wild-card "don't care" option.
func (dw *DrawingWand)
//C.DrawSetFontStretch(dw.dw, 
//const StretchType font_stretch)
//font_stretch
//font stretch (NormalStretch, UltraCondensedStretch, CondensedStretch, SemiCondensedStretch, SemiExpandedStretch, ExpandedStretch, ExtraExpandedStretch, UltraExpandedStretch, AnyStretch)
//DrawSetFontStyle
//DrawSetFontStyle() sets the font style to use when annotating with text. The AnyStyle enumeration acts as a wild-card "don't care" option.
func (dw *DrawingWand)
//C.DrawSetFontStyle(dw.dw, const StyleType style)
//style
//font style (NormalStyle, ItalicStyle, ObliqueStyle, AnyStyle)
//DrawSetFontWeight
//DrawSetFontWeight() sets the font weight to use when annotating with text.
func (dw *DrawingWand)
//C.DrawSetFontWeight(dw.dw, 
//C.size_t(font_weight))
//font_weight
//font weight (valid range 100-900)
//DrawSetGravity
//DrawSetGravity() sets the text placement gravity to use when annotating with text.
func (dw *DrawingWand)
//C.DrawSetGravity(dw.dw, const GravityType gravity)
//gravity
//positioning gravity (NorthWestGravity, NorthGravity, NorthEastGravity, WestGravity, CenterGravity, EastGravity, SouthWestGravity, SouthGravity, SouthEastGravity)
//DrawSetStrokeColor
//DrawSetStrokeColor() sets the color used for stroking object outlines.
func (dw *DrawingWand)
//C.DrawSetStrokeColor(dw.dw, 
//const PixelWand *stroke_wand)
//stroke_wand
//stroke wand.
//DrawSetStrokePatternURL
//DrawSetStrokePatternURL() sets the pattern used for stroking object outlines.
func (dw *DrawingWand)
//MagickBooleanType DrawSetStrokePatternURL(dw.dw, 
//const char *stroke_url)
//stroke_url
//URL specifying pattern ID (e.g. "#pattern_id")
//DrawSetStrokeAntialias
//DrawSetStrokeAntialias() controls whether stroked outlines are antialiased. Stroked outlines are antialiased by default. When antialiasing is disabled stroked pixels are thresholded to determine if the stroke color or underlying canvas color should be used.
func (dw *DrawingWand)
//C.DrawSetStrokeAntialias(dw.dw, 
//const MagickBooleanType stroke_antialias)
//stroke_antialias
//set to false (zero) to disable antialiasing
//DrawSetStrokeDashArray
//DrawSetStrokeDashArray() specifies the pattern of dashes and gaps used to stroke paths. The stroke dash array represents an array of numbers that specify the lengths of alternating dashes and gaps in pixels. If an odd number of values is provided, then the list of values is repeated to yield an even number of values. To remove an existing dash array, pass a zero number_elements argument and null dash_array. A typical stroke dash array might contain the members 5 3 2.
func (dw *DrawingWand)
//MagickBooleanType DrawSetStrokeDashArray(dw.dw, 
//C.size_t(number_elements),const double *dash_array)
//number_elements
//number of elements in dash array
//dash_array
//dash array values
//DrawSetStrokeDashOffset
//DrawSetStrokeDashOffset() specifies the offset into the dash pattern to start the dash.
func (dw *DrawingWand)
//C.DrawSetStrokeDashOffset(dw.dw, 
//C.double(dash_offset))
//dash_offset
//dash offset
//DrawSetStrokeLineCap
//DrawSetStrokeLineCap() specifies the shape to be used at the end of open subpaths when they are stroked. Values of LineCap are UndefinedCap, ButtCap, RoundCap, and SquareCap.
func (dw *DrawingWand)
//C.DrawSetStrokeLineCap(dw.dw, 
//const LineCap linecap)
//linecap
//linecap style
//DrawSetStrokeLineJoin
//DrawSetStrokeLineJoin() specifies the shape to be used at the corners of paths (or other vector shapes) when they are stroked. Values of LineJoin are UndefinedJoin, MiterJoin, RoundJoin, and BevelJoin.
func (dw *DrawingWand)
//C.DrawSetStrokeLineJoin(dw.dw, 
//const LineJoin linejoin)
//linejoin
//line join style
//DrawSetStrokeMiterLimit
//DrawSetStrokeMiterLimit() specifies the miter limit. When two line segments meet at a sharp angle and miter joins have been specified for 'lineJoin', it is possible for the miter to extend far beyond the thickness of the line stroking the path. The miterLimit' imposes a limit on the ratio of the miter length to the 'lineWidth'.
func (dw *DrawingWand)
//C.DrawSetStrokeMiterLimit(dw.dw, 
//C.size_t(miterlimit))
//miterlimit
//miter limit
//DrawSetStrokeOpacity
//DrawSetStrokeOpacity() specifies the opacity of stroked object outlines.
func (dw *DrawingWand)
//C.DrawSetStrokeOpacity(dw.dw, 
//C.double(stroke_opacity))
//stroke_opacity
//stroke opacity. The value 1.0 is opaque.
//DrawSetStrokeWidth
//DrawSetStrokeWidth() sets the width of the stroke used to draw object outlines.
func (dw *DrawingWand)
//C.DrawSetStrokeWidth(dw.dw, 
//C.double(stroke_width))
//stroke_width
//stroke width
//DrawSetTextAlignment
//DrawSetTextAlignment() specifies a text alignment to be applied when annotating with text.
func (dw *DrawingWand)
//C.DrawSetTextAlignment(dw.dw, const AlignType alignment)
//alignment
//text alignment. One of UndefinedAlign, LeftAlign, CenterAlign, or RightAlign.
//DrawSetTextAntialias
//DrawSetTextAntialias() controls whether text is antialiased. Text is antialiased by default.
func (dw *DrawingWand)
//C.DrawSetTextAntialias(dw.dw, 
//const MagickBooleanType text_antialias)
//text_antialias
//antialias boolean. Set to false (0) to disable antialiasing.
//DrawSetTextDecoration
//DrawSetTextDecoration() specifies a decoration to be applied when annotating with text.
func (dw *DrawingWand)
//C.DrawSetTextDecoration(dw.dw, 
//const DecorationType decoration)
//decoration
//text decoration. One of NoDecoration, UnderlineDecoration, OverlineDecoration, or LineThroughDecoration
//DrawSetTextEncoding
//DrawSetTextEncoding() specifies the code set to use for text annotations. The only character encoding which may be specified at this time is "UTF-8" for representing Unicode as a sequence of bytes. Specify an empty string to set text encoding to the system's default. Successful text annotation using Unicode may require fonts designed to support Unicode.
func (dw *DrawingWand)
//C.DrawSetTextEncoding(dw.dw, const char *encoding)
//encoding
//character string specifying text encoding
//DrawSetTextKerning
//DrawSetTextKerning() sets the spacing between characters in text.
func (dw *DrawingWand)
//C.DrawSetTextKerning(dw.dw, C.double(kerning))
//kerning
//text kerning
//DrawSetTextInterlineSpacing
//DrawSetTextInterlineSpacing() sets the spacing between line in text.
func (dw *DrawingWand)
//C.DrawSetTextInterlineSpacing(dw.dw, 
//C.double(interline_spacing))
//interline_spacing
//text line spacing
//DrawSetTextInterwordSpacing
//DrawSetTextInterwordSpacing() sets the spacing between words in text.
func (dw *DrawingWand)
//C.DrawSetTextInterwordSpacing(dw.dw, 
//C.double(interword_spacing))
//interword_spacing
//text word spacing
//DrawSetTextUnderColor
//DrawSetTextUnderColor() specifies the color of a background rectangle to place under text annotations.
func (dw *DrawingWand)
//C.DrawSetTextUnderColor(dw.dw, 
//const PixelWand *under_wand)
//under_wand
//text under wand.
//DrawSetVectorGraphics
//DrawSetVectorGraphics() sets the vector graphics associated with the specified wand. Use this method with DrawGetVectorGraphics() as a method to persist the vector graphics state.
func (dw *DrawingWand)
//MagickBooleanType DrawSetVectorGraphics(dw.dw, 
//const char *xml)
//xml
//DrawSkewX
//DrawSkewX() skews the current coordinate system in the horizontal direction.
func (dw *DrawingWand)
//C.DrawSkewX(dw.dw, C.double(degrees))
//degrees
//number of degrees to skew the coordinates
//DrawSkewY
//DrawSkewY() skews the current coordinate system in the vertical direction.
func (dw *DrawingWand)
//C.DrawSkewY(dw.dw, C.double(degrees))
//degrees
//number of degrees to skew the coordinates
//DrawTranslate
//DrawTranslate() applies a translation to the current coordinate system which moves the coordinate system origin to the specified coordinate.
func (dw *DrawingWand)
//C.DrawTranslate(dw.dw, C.double(x),
//C.double(y))
//x
//new x ordinate for coordinate system origin
//y
//new y ordinate for coordinate system origin
//DrawSetViewbox
//DrawSetViewbox() sets the overall canvas size to be recorded with the drawing vector data. Usually this will be specified using the same size as the canvas image. When the vector data is saved to SVG or MVG formats, the viewbox is use to specify the size of the canvas image that a viewer will render the vector data on.
func (dw *DrawingWand)
//C.DrawSetViewbox(dw.dw, C.ssize_t(x1),
//C.ssize_t(y1),C.ssize_t(x2),C.ssize_t(y2))
//x1
//left x ordinate
//y1
//top y ordinate
//x2
//right x ordinate
//y2
//bottom y ordinate
//IsDrawingWand
//IsDrawingWand() returns true if the wand is verified as a drawing wand.
func (dw *DrawingWand)
//MagickBooleanType IsDrawingWand(dw.dw)
//PeekDrawingWand
//PeekDrawingWand() returns the current drawing wand.
func (dw *DrawingWand)
//DrawInfo *PeekDrawingWand(dw.dw)
//PopDrawingWand
//PopDrawingWand() destroys the current drawing wand and returns to the previously pushed drawing wand. Multiple drawing wands may exist. It is an error to attempt to pop more drawing wands than have been pushed, and it is proper form to pop all drawing wands which have been pushed.
func (dw *DrawingWand)
//MagickBooleanType PopDrawingWand(dw.dw)
//PushDrawingWand
//PushDrawingWand() clones the current drawing wand to create a new drawing wand. The original drawing wand(s) may be returned to by invoking PopDrawingWand(). The drawing wands are stored on a drawing wand stack. For every Pop there must have already been an equivalent Push.
func (dw *DrawingWand)
//MagickBooleanType PushDrawingWand(dw.dw)
