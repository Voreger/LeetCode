package Sqrt

import "testing"

func TestMySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{4}, 2},
		{"Test2", args{8}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySqrt(tt.args.x); got != tt.want {
				t.Errorf("MySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
