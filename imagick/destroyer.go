package imagick

// Destroyer represents destroyable types which require manual release of resources
type Destroyer interface {
	Destroy()
}

// Destroy instance of Destroyer
func Destroy(d Destroyer) {
	d.Destroy()
}
