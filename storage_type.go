package magick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type StorageType int

const (
	UndefinedPixel StorageType = iota
	CharPixel
	DoublePixel
	FloatPixel
	IntegerPixel
	LongPixel
	QuantumPixel
	ShortPixel
)

var storageTypeStrings = map[StorageType]string{
	UndefinedPixel: "UndefinedPixel",
	CharPixel:      "CharPixel",
	DoublePixel:    "DoublePixel",
	FloatPixel:     "FloatPixel",
	IntegerPixel:   "IntegerPixel",
	LongPixel:      "LongPixel",
	QuantumPixel:   "QuantumPixel",
	ShortPixel:     "ShortPixel",
}

func (cst *StorageType) String() string {
	if v, ok := storageTypeStrings[StorageType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownStorage[%d]", *cst)
}
