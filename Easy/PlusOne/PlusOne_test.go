package PlusOne

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"Test2", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"Test3", args{[]int{9}}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
