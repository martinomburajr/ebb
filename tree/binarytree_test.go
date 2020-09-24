package tree

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestBinaryTree_FromMathematicalExpression(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		bt   BinaryTree
		args args
		want BinaryTree
	}{
		{"empty", BinaryTree{}, args{expression: "0"}, BTree0},
		{"0", BinaryTree{}, args{expression: "0"}, BTree0},
		{"x+1 -> (x+1)", BinaryTree{}, args{expression: "x+1"}, BTree1},
		{"x*x+1 -> ((x*x)+1)", BinaryTree{}, args{expression: "x*x+1"}, BTree2},
		{"x*x+1/3 -> ((x*x)+1)/3", BinaryTree{}, args{expression: "x*x+1/3"}, BTree3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFromExpression(tt.args.expression)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromExpression() = %v, want %v", got.ToMathematicalString(), tt.want.ToMathematicalString())
			}

			got.ToMathematicalString()

		})
	}
}

func TestBinaryTree_TerminalCount(t *testing.T) {
	tests := []struct {
		name string
		bt   BinaryTree
		want int
	}{
		{"0", BTree0, 1},
		{"x+1", BTree1, 2},
		{"x*x+1", BTree2, 3},
		{"x*x+1/3", BTree3, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.TerminalCount(); got != tt.want {
				t.Errorf("TerminalCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_TerminalIndices(t *testing.T) {
	tests := []struct {
		name string
		bt   BinaryTree
		want []int
	}{
		//{"0 -> (0)", BTree0, []int{0}},
		{"x+1 -> (x+1)", BTree1, []int{1, 3}},
		{"x*x+1 -> ((x*x)+1)", BTree2, []int{2, 4, 7}},
		{"x*x+1/3 -> ((x*x)+1)/3", BTree3, []int{3, 5, 8, 11}},
		{"x*x+1/3-6 -> (((x*x)+1)/3)-6)", BTree4, []int{4, 6, 9, 12, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.TerminalIndices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TerminalIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Terminals(t *testing.T) {
	tests := []struct {
		name string
		bt   BinaryTree
		want []BinaryTreeNode
	}{
		{"0", BTree0Exp, []BinaryTreeNode{BTree0Exp[0]}},
		{"x+1 -> (x+1)", BTree1, []BinaryTreeNode{BTree1[1], BTree1[3]}},
		{"x*x+1 -> ((x*x)+1)", BTree2, []BinaryTreeNode{BTree2[2], BTree2[4], BTree2[7]}},
		{"x*x+1/3 -> ((x*x)+1)/3", BTree3, []BinaryTreeNode{BTree3[3], BTree3[5], BTree3[8], BTree3[11]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.Terminals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Terminals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_NumNonTerminals(t *testing.T) {
	tests := []struct {
		name string
		bt   BinaryTree
		want int
	}{
		{"0", BTree0, 0},
		{"x+1", BTree1, 1},
		{"x*x+1", BTree2, 2},
		{"x*x+1/3", BTree3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.NonTerminalCount(); got != tt.want {
				t.Errorf("NonTerminalCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_NonTerminalIndices(t *testing.T) {
	tests := []struct {
		name string
		bt   BinaryTree
		want []int
	}{
		//{"0", BTree0, []int{}},
		{"x+1 -> (x+1)", BTree1, []int{2}},
		{"x*x+1 -> ((x*x)+1)", BTree2, []int{3, 6}},
		{"x*x+1/3 -> ((x*x)+1)/3", BTree3, []int{4, 7, 10}},
		{"x*x+1/3-6 -> (((x*x)+1)/3)-6", BTree4, []int{5, 8, 11, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.NonTerminalIndices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nonTerminalIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_NonTerminals(t *testing.T) {
	tests := []struct {
		name string
		bt   BinaryTree
		want []BinaryTreeNode
	}{
		{"0", BTree0Exp, []BinaryTreeNode{}},
		{"x+1 -> (x+1)", BTree1, []BinaryTreeNode{BTree1[2]}},
		{"x*x+1 -> ((x*x)+1)", BTree2, []BinaryTreeNode{BTree2[3], BTree2[6]}},
		{"x*x+1/3 -> ((x*x)+1)/3", BTree3, []BinaryTreeNode{BTree3[4], BTree3[7], BTree3[10]}},
		{"x*x+1/3-2 -> (((x*x)+1)/3-2)", BTree4, []BinaryTreeNode{BTree4[5], BTree4[8], BTree4[11], BTree4[14]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.NonTerminals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NonTerminals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_CalculateNewLength(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		bt   BinaryTree
		args args
		want int
	}{
		{"empty", BinaryTree{}, args{expression: "0"}, 1},
		{"x+1 -> (x+1)", BinaryTree{}, args{expression: "x+1"}, 5},
		{"x*x+1 -> ((x*x)+1)", BinaryTree{}, args{expression: "x*x+1"}, 9},
		{"x*x+1/3 -> ((x*x)+1)/3", BinaryTree{}, args{expression: "x*x+1/3"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateNewLength(tt.args.expression); got != tt.want {
				t.Errorf("CalculateNewLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type args struct {
		terminals    []rune
		nonTerminals []rune
		length       int
	}
	tests := []struct {
		name string
		args args
		want BinaryTree
	}{
		{"length: 0 (only terminal)", args{BTreeTerminals, BTreeNonTerminals, 0}, BTree0Exp},
		{"length: 1", args{BTreeTerminals, BTreeNonTerminals, 1}, BTree1},
		{"length: 2", args{BTreeTerminals, BTreeNonTerminals, 2}, BTree2},
		{"length: 3", args{BTreeTerminals, BTreeNonTerminals, 3}, BTree3},
		{"length: 10", args{BTreeTerminals, BTreeNonTerminals, 10}, BTree3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomTree(tt.args.terminals, tt.args.nonTerminals, tt.args.length)
			// Check Got Length. This checks against the number of non-terminals and uses an equation 4*NT + 1 to get
			// the total length of the expression including parentheses.
			var correctLength int = 0
			if tt.args.length < 1 {
				correctLength = 1
			} else {
				correctLength = (4 * tt.args.length) + 1
			}
			if len(got) != correctLength {
				t.Errorf("GenerateRandomTree() Length Mismatch = %v, want %v", len(got), correctLength)
			}

			// Check IDs are properly allocated
			for i := 0; i < len(got); i++ {
				if got[i].key != i {
					t.Errorf("GenerateRandomTree() Key Mismatch = %v, want %v", got[i].key, i)
				}
			}
			t.Logf("RandomTree: %s", got.ToMathematicalString())
		})
	}
}

func Test_calculateEnclosingParentheses(t *testing.T) {
	type args struct {
		bt               BinaryTree
		nonTerminalIndex int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{

		{"", args{BTree1, 2}, 0, 4},
		{"", args{BTree2, 3}, 1, 5},
		{"", args{BTree2, 6}, 0, 8},
		{"", args{BTree3, 4}, 2, 6},
		{"", args{BTree3, 7}, 1, 9},
		{"", args{BTree3, 11}, 0, 13},
		{"", args{BTree4, 5}, 3, 7},
		{"", args{BTree4, 8}, 2, 10},
		{"", args{BTree4, 11}, 1, 13},
		{"", args{BTree4, 14}, 0, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateEnclosingParentheses(tt.args.bt, tt.args.nonTerminalIndex)
			if got != tt.want {
				t.Errorf("calculateEnclosingParentheses() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateEnclosingParentheses() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_calculateNTPos(t *testing.T) {
	type args struct {
		bt               BinaryTree
		nonTerminalIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{BTree1, 2}, 0},
		{"", args{BTree2, 3}, 0},
		{"", args{BTree2, 6}, 1},
		{"", args{BTree3, 4}, 0},
		{"", args{BTree3, 7}, 1},
		{"", args{BTree3, 10}, 2},
		{"", args{BTree4, 5}, 0},
		{"", args{BTree4, 8}, 1},
		{"", args{BTree4, 11}, 2},
		{"", args{BTree4, 14}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateNTPos(len(tt.args.bt), tt.args.nonTerminalIndex); got != tt.want {
				t.Errorf("calculateNTPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateTPos(t *testing.T) {
	type args struct {
		bt            BinaryTree
		terminalIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{BTree0, 1}, 0},
		{"", args{BTree1, 1}, 0},
		{"", args{BTree1, 3}, 1},
		{"", args{BTree2, 2}, 0},
		{"", args{BTree2, 4}, 1},
		{"", args{BTree2, 7}, 2},
		{"", args{BTree3, 3}, 0},
		{"", args{BTree3, 5}, 1},
		{"", args{BTree3, 8}, 2},
		{"", args{BTree3, 11}, 3},
		{"", args{BTree4, 4}, 0},
		{"", args{BTree4, 6}, 1},
		{"", args{BTree4, 9}, 2},
		{"", args{BTree4, 12}, 3},
		{"", args{BTree4, 15}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTPos(tt.args.bt, tt.args.terminalIndex); got != tt.want {
				t.Errorf("calculateTPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test the random generation and mutation of trees for quality
func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		length := 100

		fmt.Printf("\n\t\t\t################ EXPRESSION %d ################\n", i)
		tree := GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, rand.Intn(length)).
			ApplyOperatorOnTerminal('/', 'x').
			MutateNonTerminal(BTreeTerminals)

		tree = tree.AddSubTreeRandomly(GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, rand.Intn(length))).
			ApplyOperatorOnTerminal('-', '5').
			MutateTerminal(BTreeTerminals)
		tree = tree.AddSubTreeRandomly(GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, rand.Intn(length))).
			MutateNonTerminal(BTreeNonTerminals)
		tree = tree.AddSubTreeRandomly(GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, rand.Intn(length))).
			MutateTerminal(BTreeTerminals)
		tree = tree.AddSubTreeRandomly(GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, rand.Intn(length))).
			ApplyOperatorOnTerminal('+', 'x').
			MutateTerminal(BTreeTerminals)
		tree = tree.AddSubTreeRandomly(GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, rand.Intn(length))).
		ApplyOperatorOnTerminal('*', '3')
		tree = tree.AddSubTreeRandomly(GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, rand.Intn(length)))

		fmt.Printf("Tree(%d) ==> %s\n", i, tree.ToMathematicalString())

		errors := evaluateTreeCorrectness(tree)
		if len(errors) > 0 {
			fmt.Printf("Error Summary:\n")
			for i := range errors {
				str := fmt.Sprintf("====> %s\n", errors[i].Error())
				fmt.Printf(Fata(str))
			}
		}
	}
}
