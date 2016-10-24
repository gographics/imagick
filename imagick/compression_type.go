// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type CompressionType int

const (
	COMPRESSION_UNDEFINED     CompressionType = C.UndefinedCompression
	COMPRESSION_NO            CompressionType = C.NoCompression
	COMPRESSION_BZIP          CompressionType = C.BZipCompression
	COMPRESSION_DXT1          CompressionType = C.DXT1Compression
	COMPRESSION_DXT3          CompressionType = C.DXT3Compression
	COMPRESSION_DXT5          CompressionType = C.DXT5Compression
	COMPRESSION_FAX           CompressionType = C.FaxCompression
	COMPRESSION_GROUP4        CompressionType = C.Group4Compression
	COMPRESSION_JPEG          CompressionType = C.JPEGCompression
	COMPRESSION_JPEG2000      CompressionType = C.JPEG2000Compression
	COMPRESSION_LOSSLESS_JPEG CompressionType = C.LosslessJPEGCompression
	COMPRESSION_LZW           CompressionType = C.LZWCompression
	COMPRESSION_RLE           CompressionType = C.RLECompression
	COMPRESSION_ZIP           CompressionType = C.ZipCompression
	COMPRESSION_ZIPS          CompressionType = C.ZipSCompression
	COMPRESSION_PIZ           CompressionType = C.PizCompression
	COMPRESSION_PXR24         CompressionType = C.Pxr24Compression
	COMPRESSION_B44           CompressionType = C.B44Compression
	COMPRESSION_B44A          CompressionType = C.B44ACompression
	COMPRESSION_LZMA          CompressionType = C.LZMACompression
	COMPRESSION_JBIG1         CompressionType = C.JBIG1Compression
	COMPRESSION_JBIG2         CompressionType = C.JBIG2Compression
)
