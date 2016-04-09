package types

// ImageMagick objects common interface
type IMagick interface {
	Clearer
	Destroyer
	Verifier
	Increaser
	Decreaser
}
