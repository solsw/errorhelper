package errorhelper

import (
	"errors"
	"testing"
)

func TestErrZeroNil_notNil(t *testing.T) {
	var err error = ErrZeroNil
	if err == nil {
		t.Error("ErrZeroNil assigned to error interface must not be nil")
	}
}

func TestErrZeroNil_Is(t *testing.T) {
	var err error = ErrZeroNil
	if !errors.Is(err, ErrZeroNil) {
		t.Error("errors.Is(error(ErrZeroNil), ErrZeroNil) must be true")
	}
}

func TestErrZeroNil_Error(t *testing.T) {
	if got := ErrZeroNil.Error(); got != "zero error" {
		t.Errorf("ErrZeroNil.Error() = %q, want %q", got, "zero error")
	}
}
