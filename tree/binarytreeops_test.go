package tree

import (
	"reflect"
	"testing"
)

func TestBinaryTree_AddRandomSubTreeNonTerminal(t *testing.T) {
	type args struct {
		tree BinaryTree
	}
	tests := []struct {
		name string
		bt   BinaryTree
		args args
		want BinaryTree
	}{
		{"BTree1 + BTree1", BTree1, args{BTree1}, BTree1},
		{"BTree1 + BTree1", BTree2, args{BTree1}, BTree1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.AddSubTreeRandomly(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSubTreeRandomly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addRandomSubTreeNonTerminal(t *testing.T) {
	type args struct {
		old              BinaryTree
		new              BinaryTree
		nonTerminalIndex int
	}
	tests := []struct {
		name string
		args args
		want BinaryTree
	}{
		{"", args{BTree1, BTree1, 2}, BTree1},
		{"", args{BTree2, BTree1, 3}, BTree2_BTree1_NTP0},
		{"", args{BTree2, BTree1, 6}, BTree1},
		{"", args{BTree3, BTree1, 4}, BTree3__BTree1_NTP0},
		{"", args{BTree3, BTree1, 7}, BTree3__BTree1_NTP1},
		{"", args{BTree3, BTree1, 10}, BTree1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNTNT(tt.args.old, tt.args.new, tt.args.nonTerminalIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNTNT()\n\tgott:%v\n\twant:%v", got, tt.want)
			}
		})
	}
}
