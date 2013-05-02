package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type CompressionType int

const (
	UndefinedCompression    CompressionType = C.UndefinedCompression
	NoCompression           CompressionType = C.NoCompression
	BZipCompression         CompressionType = C.BZipCompression
	DXT1Compression         CompressionType = C.DXT1Compression
	DXT3Compression         CompressionType = C.DXT3Compression
	DXT5Compression         CompressionType = C.DXT5Compression
	FaxCompression          CompressionType = C.FaxCompression
	Group4Compression       CompressionType = C.Group4Compression
	JPEGCompression         CompressionType = C.JPEGCompression
	JPEG2000Compression     CompressionType = C.JPEG2000Compression
	LosslessJPEGCompression CompressionType = C.LosslessJPEGCompression
	LZWCompression          CompressionType = C.LZWCompression
	RLECompression          CompressionType = C.RLECompression
	ZipCompression          CompressionType = C.ZipCompression
	ZipSCompression         CompressionType = C.ZipSCompression
	PizCompression          CompressionType = C.PizCompression
	Pxr24Compression        CompressionType = C.Pxr24Compression
	B44Compression          CompressionType = C.B44Compression
	B44ACompression         CompressionType = C.B44ACompression
	LZMACompression         CompressionType = C.LZMACompression
	JBIG1Compression        CompressionType = C.JBIG1Compression
	JBIG2Compression        CompressionType = C.JBIG2Compression
)

var compressionTypeStrings = map[CompressionType]string{
	UndefinedCompression:    "UndefinedCompression",
	NoCompression:           "NoCompression",
	BZipCompression:         "BZipCompression",
	DXT1Compression:         "DXT1Compression",
	DXT3Compression:         "DXT3Compression",
	DXT5Compression:         "DXT5Compression",
	FaxCompression:          "FaxCompression",
	Group4Compression:       "Group4Compression",
	JPEGCompression:         "JPEGCompression",
	JPEG2000Compression:     "JPEG2000Compression",
	LosslessJPEGCompression: "LosslessJPEGCompression",
	LZWCompression:          "LZWCompression",
	RLECompression:          "RLECompression",
	ZipCompression:          "ZipCompression",
	ZipSCompression:         "ZipSCompression",
	PizCompression:          "PizCompression",
	Pxr24Compression:        "Pxr24Compression",
	B44Compression:          "B44Compression",
	B44ACompression:         "B44ACompression",
	LZMACompression:         "LZMACompression",
	JBIG1Compression:        "JBIG1Compression",
	JBIG2Compression:        "JBIG2Compression",
}

func (ct *CompressionType) String() string {
	if v, ok := compressionTypeStrings[CompressionType(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
