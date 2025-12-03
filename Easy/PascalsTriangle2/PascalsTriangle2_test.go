package PascalsTriangle2

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test1", args{rowIndex: 3}, []int{1, 3, 3, 1}},
		{"Test2", args{rowIndex: 0}, []int{1}},
		{"Test3", args{rowIndex: 1}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
