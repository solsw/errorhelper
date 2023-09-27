package errorhelper

// UnwrapErrors returns the result of calling the [Unwrap] method on 'err', if err's
// type contains an [Unwrap] method returning []error.
// Otherwise or if the [Unwrap] method returns error, UnwrapErrors returns nil.
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

// Must0 [panics] with 'err', if 'err' is not nil.
//
// [panics]: https://pkg.go.dev/builtin#panic
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// Must returns 'r' if 'err' is nil. Otherwise, it [panics] with 'err'.
//
// [panics]: https://pkg.go.dev/builtin#panic
func Must[R any](r R, err error) R {
	if err != nil {
		panic(err)
	}
	return r
}
