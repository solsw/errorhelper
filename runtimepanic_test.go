package errorhelper

import (
	"testing"
)

func hash_of_unhashable_type() int {
	mai := make(map[any]int)
	mai[[]int{}] = 0
	return mai[2]
}

func index_out_of_range() int {
	ii := []int{1}
	return ii[1]
}

func nil_map() int {
	f := func(m map[int]int) { m[1] = 2 }
	f(nil)
	return 3
}

func interface_conversion() int {
	f := func(a any) int { return a.(int) }
	return f("")
}

func nil_pointer_dereference() int {
	var f func()
	f()
	return 0
}

func divide_by_zero() int {
	z := 0
	return 1 / z
}

func nil_pointer_dereference2() int {
	var x *int = nil
	return *x
}

func TestRuntime_panic(t *testing.T) {
	tests := []struct {
		name      string
		panicFunc func() int
		want      string
	}{
		{name: "hash_of_unhashable_type",
			panicFunc: hash_of_unhashable_type,
			want:      "runtime error: hash of unhashable type []int",
		},
		{name: "index_out_of_range",
			panicFunc: index_out_of_range,
			want:      "runtime error: index out of range [1] with length 1",
		},
		{name: "nil_map",
			panicFunc: nil_map,
			want:      "assignment to entry in nil map",
		},
		{name: "interface_conversion",
			panicFunc: interface_conversion,
			want:      "interface conversion: interface {} is string, not int",
		},
		{name: "nil_pointer_dereference",
			panicFunc: nil_pointer_dereference,
			want:      "runtime error: invalid memory address or nil pointer dereference",
		},
		{name: "divide_by_zero",
			panicFunc: divide_by_zero,
			want:      "runtime error: integer divide by zero",
		},
		{name: "nil_pointer_dereference2",
			panicFunc: nil_pointer_dereference2,
			want:      "runtime error: invalid memory address or nil pointer dereference",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := func() (err error) {
				defer func() {
					PanicToError(recover(), &err)
				}()
				tt.panicFunc()
				return nil
			}()
			got := gotErr.Error()
			if got != tt.want {
				t.Errorf("Runtime_panic = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
