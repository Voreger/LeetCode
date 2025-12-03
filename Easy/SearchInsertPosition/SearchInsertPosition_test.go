package SearchInsertPosition

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"Test2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"Test3", args{[]int{1, 3, 5, 6}, 7}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
