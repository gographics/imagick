package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

type InterlaceType int

const (
	INTERLACE_UNDEFINED InterlaceType = C.UndefinedInterlace
	INTERLACE_NO        InterlaceType = C.NoInterlace
	INTERLACE_LINE      InterlaceType = C.LineInterlace
	INTERLACE_PLANE     InterlaceType = C.PlaneInterlace
	INTERLACE_PARTITION InterlaceType = C.PartitionInterlace
	INTERLACE_GIF       InterlaceType = C.GIFInterlace
	INTERLACE_JPEG      InterlaceType = C.JPEGInterlace
	INTERLACE_PNG       InterlaceType = C.PNGInterlace
)
