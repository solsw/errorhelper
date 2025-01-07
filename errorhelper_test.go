package errorhelper

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestUnwrapErrors(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		{name: "00",
			args: args{err: nil},
			want: nil,
		},
		{name: "01",
			args: args{err: errors.New("error")},
			want: nil,
		},
		{name: "02",
			args: args{err: fmt.Errorf("%w", errors.New("error"))},
			want: nil,
		},
		{name: "1",
			args: args{err: fmt.Errorf("%w%w", errors.New("error1"), errors.New("error2"))},
			want: []error{errors.New("error1"), errors.New("error2")},
		},
		{name: "2",
			args: args{err: errors.Join(errors.New("error1"), errors.New("error2"))},
			want: []error{errors.New("error1"), errors.New("error2")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnwrapErrors(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMust0_nil(t *testing.T) {
	got := func() (err error) {
		defer func() {
			PanicToError(recover(), &err)
		}()
		Must0(nil)
		return nil
	}()
	if got != nil {
		t.Errorf("Must0 = %v, want 'nil'", got)
	}
}

func TestMust0_panic(t *testing.T) {
	const must_error = "Must error"
	got := func() (err error) {
		defer func() {
			PanicToError(recover(), &err)
		}()
		Must0(errors.New(must_error))
		return nil
	}()
	want := must_error
	if !reflect.DeepEqual(got.Error(), want) {
		t.Errorf("Must0_panic = %v, want %v", got, want)
	}
}

func TestMust_nil(t *testing.T) {
	if got := Must[any](nil, nil); got != nil {
		t.Errorf("Must[any]() = %v, want 'nil'", got)
	}
}

func TestMust_int(t *testing.T) {
	want := 23
	if got := Must(23, nil); !reflect.DeepEqual(got, want) {
		t.Errorf("Must[int]() = %v, want %v", got, want)
	}
}

func TestMust_panic(t *testing.T) {
	const must_error = "Must error"
	got := func() (err error) {
		defer func() {
			PanicToError(recover(), &err)
		}()
		Must(23, errors.New(must_error))
		return nil
	}()
	want := must_error
	if !reflect.DeepEqual(got.Error(), want) {
		t.Errorf("Must_panic = %v, want %v", got, want)
	}
}

func TestMust_panic2(t *testing.T) {
	const must_error = "Must error"
	got := func() (err error) {
		defer func() {
			PanicToError(recover(), &err)
		}()
		Must[any](nil, errors.New(must_error))
		return nil
	}()
	want := must_error
	if !reflect.DeepEqual(got.Error(), want) {
		t.Errorf("Must_panic2 = %v, want %v", got, want)
	}
}

func TestMust2_nil(t *testing.T) {
	got := func() (err error) {
		defer func() {
			PanicToError(recover(), &err)
		}()
		Must2(1, "one", nil)
		return nil
	}()
	if got != nil {
		t.Errorf("Must2 = %v, want 'nil'", got)
	}
}

func TestMust2_int_string(t *testing.T) {
	want_int, want_string := 1, "one"
	if got_int, got_string := Must2(1, "one", nil); got_int != want_int || got_string != want_string {
		t.Errorf("Must2[int, string] = (%v, %v) want (%v, %v)", got_int, got_string, want_int, want_string)
	}
}

func TestMust2_panic(t *testing.T) {
	const must_error = "Must error"
	got := func() (err error) {
		defer func() {
			PanicToError(recover(), &err)
		}()
		Must2(1, "one", errors.New(must_error))
		return nil
	}()
	want := must_error
	if !reflect.DeepEqual(got.Error(), want) {
		t.Errorf("Must2_panic = %v, want %v", got, want)
	}
}

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
		{name: "error1",
			args:        args{err: error1},
			wantErr:     true,
			expectedErr: error1,
		},
		{name: "error1 with params",
			args:        args{err: error1, params: []any{"param1", "param2"}},
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

type errMethod struct {
}

func (errMethod) Error1() error {
	return CallerError(error1)
}

func (*errMethod) Error2() error {
	return CallerError(error1)
}

func ExampleCallerError() {
	errf := CallerError(error1)
	fmt.Println(errf.Error())
	ermth := errMethod{}
	errm1 := ermth.Error1()
	fmt.Println(errm1.Error())
	errm2 := ermth.Error2()
	fmt.Println(errm2.Error())
	// Output:
	// ExampleCallerError: error1
	// errMethod.Error1: error1
	// (*errMethod).Error2: error1
}
