package imagick

import "gopkg.in/gographics/imagick.v3/imagick/types"

// Destroy instance of Destroyer
// If GOGC=off you should call obj.Destroy() manually
func Destroy(d types.Destroyer) {
	d.Destroy()
}
