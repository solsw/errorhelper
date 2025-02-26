package errorhelper

// UnwrapErrors returns the result of calling [Unwrap] method on 'err',
// in case if err's type contains an [Unwrap] method that returns []error.
// Otherwise or if the [Unwrap] method returns single error, UnwrapErrors returns nil.
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

// Must2 returns ('r1', 'r2') if 'err' is nil. Otherwise, it [panics] with 'err'.
//
// [panics]: https://pkg.go.dev/builtin#panic
func Must2[R1, R2 any](r1 R1, r2 R2, err error) (R1, R2) {
	if err != nil {
		panic(err)
	}
	return r1, r2
}
