// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

import "testing"

func TestEnvironmentLifecycle(t *testing.T) {
	defer func() {
		if reco := recover(); reco != nil {
			t.Fatalf("MagickWand environment was not correctly inited or terminated: %q", reco)
		}
	}()

	Initialize()
	Initialize()

	Terminate()
	Terminate()
}
