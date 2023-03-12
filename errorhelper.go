package errorhelper

// UnwrapErrors returns the result of calling the [Unwrap] method on 'err', if err's
// type contains an [Unwrap] method returning []error.
// Otherwise, UnwrapErrors returns nil.
//
// UnwrapErrors returns nil if the [Unwrap] method returns error.
//
// [Unwrap]: https://pkg.go.dev/errors#pkg-overview
func UnwrapErrors(err error) []error {
	u, ok := err.(interface {
		Unwrap() []error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}
