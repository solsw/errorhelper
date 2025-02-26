package errorhelper

import (
	"errors"
	"fmt"
	"testing"
)

var error1 = errors.New("error1")

func TestCallerError(t *testing.T) {
	type args struct {
		err    error
		params []any
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
	}{
		{name: "nil",
			args:    args{err: nil},
			wantErr: false,
		},
		{name: "error",
			args:        args{err: error1},
			wantErr:     true,
			expectedErr: error1,
		},
		{name: "error with params",
			args:        args{err: error1, params: []any{" param1", " param2", " "}},
			wantErr:     true,
			expectedErr: error1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CallerError(tt.args.err, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallerError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				t.Logf("%v", err)
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("CallerError() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
		})
	}
}

func TestPackageCallerError(t *testing.T) {
	type args struct {
		err    error
		params []any
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
	}{
		{name: "nil",
			args:    args{err: nil},
			wantErr: false,
		},
		{name: "error",
			args:        args{err: error1},
			wantErr:     true,
			expectedErr: error1,
		},
		{name: "error with params",
			args:        args{err: error1, params: []any{" param1", " param2", " "}},
			wantErr:     true,
			expectedErr: error1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PackageCallerError(tt.args.err, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("PackageCallerError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				t.Logf("%v", err)
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("PackageCallerError() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
		})
	}
}

type errMethod struct {
}

func (errMethod) Error11() error {
	return CallerError(error1, ": ")
}

func (*errMethod) Error12() error {
	return CallerError(error1, "--")
}

func (errMethod) Error21() error {
	return PackageCallerError(error1, ": ")
}

func (*errMethod) Error22() error {
	return PackageCallerError(error1, "--")
}

func ExampleCallerError() {
	errf1 := CallerError(error1)
	fmt.Println(errf1.Error())
	errf2 := CallerError(error1, " ")
	fmt.Println(errf2.Error())
	ermth := errMethod{}
	errm1 := ermth.Error11()
	fmt.Println(errm1.Error())
	errm2 := ermth.Error12()
	fmt.Println(errm2.Error())
	// Output:
	// ExampleCallerError:error1
	// ExampleCallerError error1
	// errMethod.Error11: error1
	// (*errMethod).Error12--error1
}

func ExamplePackageCallerError() {
	errf1 := PackageCallerError(error1)
	fmt.Println(errf1.Error())
	errf2 := PackageCallerError(error1, " ")
	fmt.Println(errf2.Error())
	ermth := errMethod{}
	errm1 := ermth.Error21()
	fmt.Println(errm1.Error())
	errm2 := ermth.Error22()
	fmt.Println(errm2.Error())
	// Output:
	// errorhelper.ExamplePackageCallerError:error1
	// errorhelper.ExamplePackageCallerError error1
	// errorhelper.errMethod.Error21: error1
	// errorhelper.(*errMethod).Error22--error1
}
