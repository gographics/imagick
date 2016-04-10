package types

// Returns true if the object is a verified C object
type Verifier interface {
	IsVerified() bool
}
