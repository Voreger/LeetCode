package Valid_Parentheses

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{s: "()"}, true},
		{"Test2", args{s: "()[]{}"}, true},
		{"Test3", args{s: "(]"}, false},
		{"Test4", args{s: "([])"}, true},
		{"Test5", args{s: "([)]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
