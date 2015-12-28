package imagick

import "github.com/gographics/imagick/imagick/types"

// Destroy instance of Destroyer
// If GOGC=off you should call obj.Destroy() manually
func Destroy(d types.Destroyer) {
	d.Destroy()
}
