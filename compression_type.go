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
	UndefinedCompression CompressionType = iota
	NoCompression
	BZipCompression
	DXT1Compression
	DXT3Compression
	DXT5Compression
	FaxCompression
	Group4Compression
	JPEGCompression
	JPEG2000Compression
	LosslessJPEGCompression
	LZWCompression
	RLECompression
	ZipCompression
	ZipSCompression
	PizCompression
	Pxr24Compression
	B44Compression
	B44ACompression
	LZMACompression
	JBIG1Compression
	JBIG2Compression
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
