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
