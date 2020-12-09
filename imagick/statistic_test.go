package imagick

import "testing"

func TestImageEntropy(t *testing.T) {
	Initialize()
	defer func(t *testing.T) {
		checkGC(t)
		Terminate()
	}(t)

	mw := NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(`logo:`); err != nil {
		t.Fatal(err.Error())
	}

	if entropy, err := mw.GetImageEntropy(); err != nil {
		t.Fatal(err.Error())
	} else {
		if entropy <= 0 || entropy > 0.3 {
			t.Errorf("Unexpected entropy value: %v", entropy)
		}
	}
}
