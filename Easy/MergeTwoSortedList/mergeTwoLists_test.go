package MergeTwoSortedList

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Test1", args{list1: nil, list2: nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
