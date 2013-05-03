package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type StyleType int

const (
	STYLE_UNDEFINED StyleType = C.UndefinedStyle
	STYLE_NORMAL    StyleType = C.NormalStyle
	STYLE_ITALIC    StyleType = C.ItalicStyle
	STYLE_OBLIQUE   StyleType = C.ObliqueStyle
	STYLE_ANYSTYLE  StyleType = C.AnyStyle
)
