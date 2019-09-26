package list

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

func buildTrie(children []*Node) *Trie {
	trie := NewTrie(3)
	if children != nil {
		trie.root.children = children
	}
	return trie
}

func TestTrie_Insert(t *testing.T) {
	tests := []struct {
		name   string
		word   string
		before []*Node
		want   []*Node
	}{
		{"add one letter", "a", nil, []*Node{{Letter: "a", Count: 1, index: 0, Word: "a", children: []*Node{}}}},
		{"add second letter", "b", []*Node{{Letter: "a", Count: 1, Word: "a", index: 0, children: []*Node{}}}, []*Node{
			{Letter: "a", Count: 1, Word: "a", index: 0, children: []*Node{}},
			{Letter: "b", Count: 1, Word: "b", index: 0, children: []*Node{}},
		}},

		{"add two letter", "ab", nil, []*Node{{
			Letter: "a", Count: 0, Word: "", index: 0, children: []*Node{
				{Letter: "b", Count: 1, Word: "ab", index: 0, children: []*Node{}},
			}},
		},
		},

		{"add three letter", "abc", nil, []*Node{{
			Letter: "a", Count: 0, Word: "", index: 0, children: []*Node{
				{Letter: "b", Count: 0, Word: "", index: 0, children: []*Node{
					{Letter: "c", Count: 1, Word: "abc", index: 0, children: []*Node{}},
				}},
			}},
		},
		},

		{"adding substring, should increase counter", "ab", []*Node{{
			Letter: "a", Count: 0, Word: "", index: 0, children: []*Node{
				{Letter: "b", Count: 0, Word: "", index: 0, children: []*Node{
					{Letter: "c", Count: 1, Word: "abc", index: 0, children: []*Node{}},
				}},
			}}},
			[]*Node{{
				Letter: "a", Count: 0, Word: "", index: 0, children: []*Node{
					{Letter: "b", Count: 1, Word: "ab", index: 0, children: []*Node{
						{Letter: "c", Count: 1, Word: "abc", index: 0, children: []*Node{}},
					}},
				}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := buildTrie(tt.before)
			want := buildTrie(tt.want).root
			trie.Insert(tt.word)
			if got := trie.root; !reflect.DeepEqual(got, want) {
				t.Errorf(pretty.Sprint("Insert() - ", got, ", want - ", want))
			}
		})
	}
}

func TestTrie_addToList(t *testing.T) {
	type fields struct {
		list []*Node
	}

	type want struct {
		index int
		list  []*Node
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want
	}{
		{"add first", fields{list: nil}, args{node: &Node{Word: "a", Count: 1, index: -1}}, want{
			index: 0, list: []*Node{{Word: "a", Count: 1, index: 0}},
		}},
		{"add second", fields{list: []*Node{{Word: "a", Count: 1, index: 0}}},
			args{node: &Node{Word: "b", Count: 1, index: -1}}, want{1, []*Node{
				{Word: "a", Count: 1, index: 0},
				{Word: "b", Count: 1, index: 1}},
			}},
		{"add third", fields{list: []*Node{
			{Word: "a", Count: 1, index: 0},
			{Word: "b", Count: 1, index: 1},
		}}, args{node: &Node{Word: "c", Count: 1, index: -1}}, want{2, []*Node{
			{Word: "a", Count: 1, index: 0},
			{Word: "b", Count: 1, index: 1},
			{Word: "c", Count: 1, index: 2},
		},
		}},
		{"add c shod order", fields{list: []*Node{
			{Word: "a", Count: 1},
			{Word: "b", Count: 1},
			{Word: "c", Count: 1},
		}}, args{node: &Node{Word: "c", Count: 2, index: 2}}, want{0, []*Node{
			{Word: "c", Count: 2, index: 0},
			{Word: "a", Count: 1, index: 1},
			{Word: "b", Count: 1, index: 2},
		},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie(3)
			if tt.fields.list != nil {
				trie.list = append(trie.list, tt.fields.list...)
			}
			node := tt.args.node
			if node.index >= 0 {
				node = tt.fields.list[node.index]
			}
			if got := trie.addToList(node); got != tt.want.index {
				t.Error(pretty.Sprintf("addToList() = %v, want %v, list: %v, should be:%v", got, tt.want.index, trie.list, tt.want.list))
			}
			if got := trie.list; !reflect.DeepEqual(got, tt.want.list) {
				t.Error(pretty.Sprint("List not Equal: addToList() = ", got, "want -", tt.want.list))
			}
		})
	}
}
