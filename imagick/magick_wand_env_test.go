// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import "testing"

func TestEnvironmentLifecycle(t *testing.T) {
	Initialize()
	if !initialized {
		t.Fatal("MagickWand environment was not initialized")
	}
	Initialize()
	if !initialized {
		t.Fatal("MagickWand environment was not idenpotently initialized")
	}
	Terminate()
	if initialized {
		t.Fatal("MagickWand environment was not terminated")
	}
	Terminate()
	if initialized {
		t.Fatal("MagickWand environment was not idenpotently terminated")
	}
}
