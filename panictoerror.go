package errorhelper

import (
	"encoding"
	"errors"
	"fmt"
)

// PanicToError converts [panic] to [error] in the following way:
//   - if surrounding function panics with an [error], this error is returned;
//   - if surrounding function panics with a [string], the error wrapping this string is returned;
//   - if surrounding function panics with a [fmt.Stringer], the error wrapping [fmt.Stringer.String] is returned;
//   - if surrounding function panics with an [encoding.TextMarshaler], the error wrapping [encoding.TextMarshaler.MarshalText] is returned;
//   - otherwise the panic is reraised.
//
// PanicToError must be called from a [defer] statement:
//
//	func Example() (err error) {
//	  defer func() {
//	    PanicToError(recover(), &err)
//	  }()
//	  return nil
//	}
//
// (See tests.)
//
// [defer]: https://go.dev/ref/spec#Defer_statements
// [error]: https://go.dev/ref/spec#Errors
// [panic]: https://go.dev/ref/spec#Handling_panics
// [string]: https://go.dev/ref/spec#String_types
func PanicToError(panicArg any, err *error) {
	switch v := panicArg.(type) {
	case nil:
		return
	case error:
		*err = v
		return
	case string:
		*err = errors.New(v)
		return
	case fmt.Stringer:
		*err = errors.New(v.String())
		return
	case encoding.TextMarshaler:
		if bb, e := v.MarshalText(); e == nil {
			*err = errors.New(string(bb))
			return
		}
	}
	panic(panicArg)
}
