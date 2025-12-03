package ClimbingStairs

import "testing"

func TestClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{2}, 2},
		{"Test2", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
