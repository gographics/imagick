package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type InterlaceType int

const (
	UndefinedInterlace InterlaceType = C.UndefinedInterlace
	NoInterlace        InterlaceType = C.NoInterlace
	LineInterlace      InterlaceType = C.LineInterlace
	PlaneInterlace     InterlaceType = C.PlaneInterlace
	PartitionInterlace InterlaceType = C.PartitionInterlace
	GIFInterlace       InterlaceType = C.GIFInterlace
	JPEGInterlace      InterlaceType = C.JPEGInterlace
	PNGInterlace       InterlaceType = C.PNGInterlace
)

var interlaceTypeStrings = map[InterlaceType]string{
	UndefinedInterlace: "UndefinedInterlace",
	NoInterlace:        "NoInterlace",
	LineInterlace:      "LineInterlace",
	PlaneInterlace:     "PlaneInterlace",
	PartitionInterlace: "PartitionInterlace",
	GIFInterlace:       "GIFInterlace",
	JPEGInterlace:      "JPEGInterlace",
	PNGInterlace:       "PNGInterlace",
}

func (cst *InterlaceType) String() string {
	if v, ok := interlaceTypeStrings[InterlaceType(*cst)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownInterlaceType[%d]", *cst)
}
