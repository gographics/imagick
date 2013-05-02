package imagick

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
	UndefinedFunction  MagickFunction = C.UndefinedFunction
	PolynomialFunction MagickFunction = C.PolynomialFunction
	SinusoidFunction   MagickFunction = C.SinusoidFunction
	ArcsinFunction     MagickFunction = C.ArcsinFunction
	ArctanFunction     MagickFunction = C.ArctanFunction
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
