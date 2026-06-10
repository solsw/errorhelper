# errorhelper
[![GitHub](https://img.shields.io/badge/git-github.com-green)](https://github.com/solsw/errorhelper)

Package `errorhelper` contains [error](https://go.dev/ref/spec#Errors)-related helpers.

## Installation

```bash
go get github.com/solsw/errorhelper
```

```go
import "github.com/solsw/errorhelper"
```

## API

### CallerError

```go
func CallerError(err error, params ...any) error
```

CallerError prepends the provided error's `Error()` with the caller function/method name
and the provided parameters. If there are no parameters, a colon is used as the separator.
Parameters are formatted with no automatic separators, so include any desired spacing or
punctuation in the parameter values (e.g. `": "` or `" "`). If `err` is nil, nil is returned.
The original error is wrapped, so it remains accessible via `errors.Is`/`errors.As`.

CallerError must be called directly from the function whose name should appear in the error;
wrapping it in a helper shifts the call stack and produces the wrong name.

```go
func doSomething() error {
	err := errors.New("something failed")
	return errorhelper.CallerError(err)
	// error message: "doSomething:something failed"
}

func doSomethingElse(id int) error {
	err := errors.New("something failed")
	return errorhelper.CallerError(err, " id=", id, ": ")
	// error message: "doSomethingElse id=42: something failed"
}
```

### PackageCallerError

```go
func PackageCallerError(err error, params ...any) error
```

PackageCallerError works like [CallerError](#callererror), but the caller function/method name
is prepended with the package name.

```go
func doSomething() error {
	err := errors.New("something failed")
	return errorhelper.PackageCallerError(err)
	// error message: "mypackage.doSomething:something failed"
}
```

### PanicToError

```go
func PanicToError(panicArg any, err *error)
```

PanicToError converts a [panic](https://go.dev/ref/spec#Handling_panics) to an error
in the following way:

- if the surrounding function panics with an `error`, this error is returned;
- if the surrounding function panics with a `string`, an error wrapping this string is returned;
- if the surrounding function panics with a `fmt.Stringer`, an error wrapping the `String()` result is returned;
- if the surrounding function panics with an `encoding.TextMarshaler` and `MarshalText()` succeeds, an error wrapping the marshaled text is returned;
- otherwise the panic is reraised.

Cases are matched in the order listed; a type implementing both `fmt.Stringer` and
`encoding.TextMarshaler` is handled by the `fmt.Stringer` case.

PanicToError must be called from a [defer](https://go.dev/ref/spec#Defer_statements) statement:

```go
func Example() (err error) {
	defer func() {
		errorhelper.PanicToError(recover(), &err)
	}()
	// code that may panic
	return nil
}
```

### UnwrapErrors

```go
func UnwrapErrors(err error) []error
```

UnwrapErrors returns the result of calling the `Unwrap` method on `err`, in case err's type
contains an `Unwrap` method that returns `[]error` (such as errors created by
[errors.Join](https://pkg.go.dev/errors#Join) or
[fmt.Errorf](https://pkg.go.dev/fmt#Errorf) with multiple `%w` verbs).
Otherwise, or if the `Unwrap` method returns a single error, UnwrapErrors returns nil.

```go
joined := errors.Join(err1, err2)
errs := errorhelper.UnwrapErrors(joined)
// errs is []error{err1, err2}
```

### Must0, Must, Must2

```go
func Must0(err error)
func Must[R any](r R, err error) R
func Must2[R1, R2 any](r1 R1, r2 R2, err error) (R1, R2)
```

The `Must` helpers [panic](https://pkg.go.dev/builtin#panic) with `err` if `err` is not nil.
Otherwise:

- `Must0` simply returns;
- `Must` returns the single value `r`;
- `Must2` returns the pair of values `r1`, `r2`.

They are useful for initialization or in tests, where an error is not expected and
should stop the program:

```go
re := errorhelper.Must(regexp.Compile(`\d+`))
```

### ErrZeroNil

```go
var ErrZeroNil = new(zeroError)
```

ErrZeroNil is a sentinel error that signals a do-nothing (no-op) condition.
It is a non-nil error value with the message `"zero error"`;
use [errors.Is](https://pkg.go.dev/errors#Is) to check for it.

```go
if errors.Is(err, errorhelper.ErrZeroNil) {
	// nothing to do
}
```
