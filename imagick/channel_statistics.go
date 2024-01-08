// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"

//"fmt"
//"unsafe"

type ChannelStatistics struct {
	// TODO: https://github.com/gographics/imagick/issues/301
	//   As of ImageMagick 7.1.1-24 they have introduced fields
	//   of type "long double". This causes a cgo compilation error.
	//   Disabling this reference for now until the struct is actually
	//   used somewhere in these bindings.
	//   When needed, A c-shim type will need to be declared to mirror
	//   ChannelStatistics with a double instead of long double fields,
	//   and a helper to convert between.
	//cs *C.ChannelStatistics
}
