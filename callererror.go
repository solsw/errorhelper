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
	s := path.Base(runtimehelper.NthCallerName(3))
	if s == "" {
		return err
	}
	s, _, _ = strings.Cut(s, "[")
	if !withPackage {
		_, s, _ = strings.Cut(s, ".")
	}
	format := "%s"
	a := []any{s}
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

// CallerError prepends the provided error error's Error() with caller function/method name
// and the provided parameters. If there are no parameters, colon is used.
// If 'err' is nil, nil is returned.
func CallerError(err error, params ...any) error {
	return callerErrorPrim(err, false, params...)
}

// PackageCallerError prepends the provided error's Error() with caller function/method name
// prepended with package name and the provided parameters. If there are no parameters, colon is used.
// If 'err' is nil, nil is returned.
func PackageCallerError(err error, params ...any) error {
	return callerErrorPrim(err, true, params...)
}
