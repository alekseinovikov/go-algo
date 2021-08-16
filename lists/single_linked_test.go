package lists

import (
	"reflect"
	"testing"
)

func TestNewSingleLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want List
	}{
		{"Creating empty list", &singleLinkedListNode{
			head:  true,
			value: nil,
			next:  nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingleLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingleLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedListNode_Add(t *testing.T) {
	tests := []struct {
		name  string
		elems []interface{}
		want  List
	}{
		{
			name:  "Addding one element",
			elems: []interface{}{"1"},
			want: &singleLinkedListNode{
				head:  true,
				value: nil,
				next: &singleLinkedListNode{
					head:  false,
					value: "1",
					next:  nil,
				},
			},
		},
		{
			name:  "Addding two elements",
			elems: []interface{}{"1", "2"},
			want: &singleLinkedListNode{
				head:  true,
				value: nil,
				next: &singleLinkedListNode{
					head:  false,
					value: "1",
					next: &singleLinkedListNode{
						head:  false,
						value: "2",
						next:  nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewSingleLinkedList()
			for _, elem := range tt.elems {
				list.Add(elem)
			}

			if !reflect.DeepEqual(list, tt.want) {
				t.Errorf("List.Add() = %v, want %v", list, tt.want)
			}
		})
	}
}

func TestSingleLinkedListNode_GetAt(t *testing.T) {
	list := NewSingleLinkedList()
	list.Add("1")
	list.Add("2")
	list.Add("3")

	tests := []struct {
		name  string
		list  List
		index int
		found bool
		want  interface{}
	}{
		{
			name:  "Has a few elements returns by index",
			list:  list,
			index: 0,
			found: true,
			want:  "1",
		},
		{
			name:  "Has a few elements returns by index",
			list:  list,
			index: 1,
			found: true,
			want:  "2",
		},
		{
			name:  "Has a few elements returns by index",
			list:  list,
			index: 2,
			found: true,
			want:  "3",
		},
		{
			name:  "Has a few elements returns by index",
			list:  list,
			index: 99,
			found: false,
			want:  nil,
		},
		{
			name:  "Has a few elements returns by index",
			list:  list,
			index: 3,
			found: false,
			want:  nil,
		},
		{
			name:  "Has a few elements returns by index",
			list:  list,
			index: -1,
			found: false,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, ok := tt.list.GetAt(tt.index)
			if !reflect.DeepEqual(found, tt.want) || !reflect.DeepEqual(ok, tt.found) {
				t.Errorf("List.Add() = %v, want %v", list, tt.want)
			}
		})
	}
}

func TestSingleLinkedListNode_AddAt(t *testing.T) {
	tests := []struct {
		name  string
		list  *singleLinkedListNode
		elem  interface{}
		index int
		want  *singleLinkedListNode
		err   error
	}{
		{
			name: "Empty list adding one",
			list: &singleLinkedListNode{
				head:  true,
				next:  nil,
				value: nil,
			},
			elems: []interface{}{"1"},
			want: &singleLinkedListNode{
				head:  true,
				value: nil,
				next: &singleLinkedListNode{
					value: "1",
				},
			},
		}, {},
	}
}
