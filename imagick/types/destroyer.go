package types

// Destroyer represents destroyable types which require manual release of resources
type Destroyer interface {
	Destroy()
}
