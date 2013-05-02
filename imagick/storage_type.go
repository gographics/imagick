package imagick

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
	UndefinedPixel StorageType = C.UndefinedPixel
	CharPixel      StorageType = C.CharPixel
	DoublePixel    StorageType = C.DoublePixel
	FloatPixel     StorageType = C.FloatPixel
	IntegerPixel   StorageType = C.IntegerPixel
	LongPixel      StorageType = C.LongPixel
	QuantumPixel   StorageType = C.QuantumPixel
	ShortPixel     StorageType = C.ShortPixel
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
