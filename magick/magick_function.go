package magick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type MagickFunction int

const (
	UndefinedFunction MagickFunction = iota
	PolynomialFunction
	SinusoidFunction
	ArcsinFunction
	ArctanFunction
)

var magickFunctionStrings = map[MagickFunction]string{
	UndefinedFunction:  "UndefinedFunction",
	PolynomialFunction: "PolynomialFunction",
	SinusoidFunction:   "SinusoidFunction",
	ArcsinFunction:     "ArcsinFunction",
	ArctanFunction:     "ArctanFunction",
}

func (cst *MagickFunction) String() string {
	if v, ok := magickFunctionStrings[MagickFunction(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownFunction[%d]", *cst)
}
