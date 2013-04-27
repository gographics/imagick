package magick

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
	0:  "UndefinedCompression",
	1:  "NoCompression",
	2:  "BZipCompression",
	3:  "DXT1Compression",
	4:  "DXT3Compression",
	5:  "DXT5Compression",
	6:  "FaxCompression",
	7:  "Group4Compression",
	8:  "JPEGCompression",
	9:  "JPEG2000Compression",
	10: "LosslessJPEGCompression",
	11: "LZWCompression",
	12: "RLECompression",
	13: "ZipCompression",
	14: "ZipSCompression",
	15: "PizCompression",
	16: "Pxr24Compression",
	17: "B44Compression",
	18: "B44ACompression",
	19: "LZMACompression",
	20: "JBIG1Compression",
	21: "JBIG2Compression",
}

func (ct *CompressionType) String() string {
	if v, ok := compressionTypeStrings[CompressionType(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
