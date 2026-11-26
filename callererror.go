package errorhelper

import (
	"fmt"
	"path"
	"strings"

	"github.com/solsw/runtimehelper"
)

func callerErrorPrim(err error, withPackage bool, params ...any) error {
	if err == nil {
		return nil
	}
	return wrapCallerError(err, runtimehelper.NthCallerName(3), withPackage, params...)
}

func wrapCallerError(err error, symbol string, withPackage bool, params ...any) error {
	if symbol == "" {
		return err
	}
	var name string
	if withPackage {
		name = packageCallerName(symbol)
	} else {
		name = runtimehelper.JustFunctionName(symbol)
	}
	format := "%s"
	a := []any{name}
	if len(params) > 0 {
		format += strings.Repeat("%v", len(params))
		a = append(a, params...)
	} else {
		format += ":"
	}
	format += "%w"
	a = append(a, err)
	return fmt.Errorf(format, a...)
}

func packageCallerName(symbol string) string {
	name := runtimehelper.JustPackageFunctionName(symbol)
	pkg, function, _ := strings.Cut(name, ".")
	version, ok := strings.CutPrefix(pkg, "v")
	if !ok || version == "" || version == "1" || version[0] == '0' || strings.Trim(version, "0123456789") != "" {
		return name
	}
	// For a terminal /v2 or higher, take the directory before the version.
	// Type arguments can contain slashes, so exclude them from the path lookup.
	head, _, _ := strings.Cut(symbol, "[")
	if dir, _ := path.Split(head); dir != "" {
		return path.Base(dir) + "." + function
	}
	return name
}

// CallerError prepends the provided error's Error() with caller function/method name
// and the provided parameters. If there are no parameters, colon is used.
// Parameters are formatted with no automatic separators, so include any desired
// spacing or punctuation in the parameter values (e.g. ": " or " ").
// If 'err' is nil, nil is returned.
// If the caller name is unavailable, CallerError returns 'err'.
//
// CallerError must be called directly from the function whose name should appear
// in the error; wrapping it in a helper shifts the call stack and produces the wrong name.
func CallerError(err error, params ...any) error {
	return callerErrorPrim(err, false, params...)
}

// PackageCallerError prepends the provided error's Error() with caller function/method name
// prepended with the package label and the provided parameters. The label comes from the last
// import path component, skipping a terminal major version such as /v2.
// If there are no parameters, colon is used.
// Parameters are formatted with no automatic separators, so include any desired
// spacing or punctuation in the parameter values (e.g. ": " or " ").
// If 'err' is nil, nil is returned.
// If the caller name is unavailable, PackageCallerError returns 'err'.
//
// PackageCallerError must be called directly from the function whose name should appear
// in the error; wrapping it in a helper shifts the call stack and produces the wrong name.
func PackageCallerError(err error, params ...any) error {
	return callerErrorPrim(err, true, params...)
}
