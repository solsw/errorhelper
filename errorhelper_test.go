package errorhelper

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/solsw/builtinhelper"
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

func TestMust(t *testing.T) {
	got := func() (err error) {
		defer func() {
			builtinhelper.PanicToError(recover(), &err)
		}()
		Must(nil)
		return nil
	}()
	if got != nil {
		t.Errorf("Must = %v, want 'nil'", got)
	}
}

func TestMust_panic(t *testing.T) {
	const must_error = "Must error"
	got := func() (err error) {
		defer func() {
			builtinhelper.PanicToError(recover(), &err)
		}()
		Must(errors.New(must_error))
		return nil
	}()
	want := must_error
	if !reflect.DeepEqual(got.Error(), want) {
		t.Errorf("Must_panic = %v, want %v", got, want)
	}
}
