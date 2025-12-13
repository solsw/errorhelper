package errorhelper

// Pointer to ZeroError is a do-nothing [error].
//
// [error]: https://go.dev/ref/spec#Errors
type ZeroError struct{}

// Error implements the [error] interface.
//
// [error]: https://pkg.go.dev/builtin#error
func (*ZeroError) Error() string {
	return "zero error"
}

// ErrZeroNil is a nil pointer to [ZeroError].
var ErrZeroNil *ZeroError = nil
