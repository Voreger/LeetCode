package SmallestIntegerDivisibleByK

import "testing"

func TestSmallestRepunitDivByK(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{1}, 1},
		{"Test2", args{2}, -1},
		{"Test3", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestRepunitDivByK(tt.args.k); got != tt.want {
				t.Errorf("SmallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
