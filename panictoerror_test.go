package errorhelper

import (
	"errors"
	"testing"
)

func func_no_panic() (err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	return nil
}

func TestPanicToError_no_panic(t *testing.T) {
	err := func_no_panic()
	if err != nil {
		t.Errorf("PanicToError_no_panic: err != nil")
	}
}

var err_func_error = errors.New("func_error")

func func_error() (err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	panic(err_func_error)
}

func TestPanicToError_error(t *testing.T) {
	err := func_error()
	if err != err_func_error {
		t.Errorf("PanicToError_error: err != err_func_error")
	}
}

func func_string() (err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	panic("func_string")
}

func TestPanicToError_string(t *testing.T) {
	err := func_string()
	if err.Error() != "func_string" {
		t.Errorf(`PanicToError_string: err != "func_string"`)
	}
}

type stringer string

func (s stringer) String() string {
	return string(s)
}

func func_Stringer() (err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	panic(stringer("func_Stringer"))
}

func TestPanicToError_Stringer(t *testing.T) {
	err := func_Stringer()
	if err.Error() != "func_Stringer" {
		t.Errorf(`PanicToError_Stringer: err != "func_Stringer"`)
	}
}

type textMarshaler string

func (t textMarshaler) MarshalText() ([]byte, error) {
	return []byte(t), nil
}

func func_TextMarshaler() (err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	panic(textMarshaler("func_TextMarshaler"))
}

func TestPanicToError_TextMarshaler(t *testing.T) {
	err := func_TextMarshaler()
	if err.Error() != "func_TextMarshaler" {
		t.Errorf(`PanicToError_TextMarshaler: err != "func_TextMarshaler"`)
	}
}
