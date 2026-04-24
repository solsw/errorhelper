package errorhelper

type zeroError struct{}

func (*zeroError) Error() string {
	return "zero error"
}

// ErrZeroNil is a sentinel [error] that signals a do-nothing (no-op) condition.
// It is a non-nil error value; use [errors.Is] to check for it.
//
// [error]: https://go.dev/ref/spec#Errors
var ErrZeroNil = new(zeroError)
