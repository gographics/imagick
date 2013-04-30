package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type EndianType int

const (
	UndefinedEndian EndianType = iota
	LSBEndian
	MSBEndian
)

var endianTypeStrings = map[EndianType]string{
	UndefinedEndian: "UndefinedEndian",
	LSBEndian:       "LSBEndian",
	MSBEndian:       "MSBEndian",
}

func (cst *EndianType) String() string {
	if v, ok := endianTypeStrings[EndianType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownEndianType[%d]", *cst)
}
