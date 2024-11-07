package slicex

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T any, R any] struct {
		slice []T
		f     func(int, T) R
	}
	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want []R
	}
	tests := []testCase[int, string]{
		{
			name: "convert from []int to []string",
			args: args[int, string]{
				slice: []int{1, 2, 3},
				f: func(_ int, i int) string {
					return strconv.Itoa(i)
				},
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.slice, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
