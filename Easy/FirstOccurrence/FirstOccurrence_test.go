package FirstOccurrence

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{haystack: "sadbutsad", needle: "sad"}, 0},
		{"Test2", args{haystack: "leetcode", needle: "leeto"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
