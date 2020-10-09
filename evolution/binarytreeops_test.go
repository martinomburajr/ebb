package evolution

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var maxSize = 12

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

func TestBinaryTree_DeleteTerminal(t *testing.T) {
	numTests := 20
	maxSize := 12

	testExpressions := expressionGenerator(numTests, maxSize)

	for i := 0; i < numTests; i++ {
		for j := 0; j < len(testExpressions); j++ {
			testExpression := testExpressions[j]

			t.Run(testExpression.ToMathematicalString(), func(t *testing.T) {
				// Function/Method to test
				got := testExpression.DeleteTerminal()

				errors := got.EvaluateTreeCorrectness()
				if len(errors) > 0 {
					t.Errorf("DeleteTerminal - FormatErrors\n%s", errors.ToString(got))
				}

				// Test condition
				if got.contains('0') == -1 {
					t.Errorf("DeleteTerminal() = %v", got)
				}
			})
		}
	}

}

func TestBinaryTree_DeleteNonTerminal(t *testing.T) {
	numTests := 200
	maxSize = 20

	testExpressions := expressionGenerator(numTests, maxSize, "x", "(x+1)", "((x+4)/2)", "(((x-2)/5)+2)")

	for j := 0; j < numTests; j++ {
		testExpression := testExpressions[j]

		t.Run(testExpression.ToMathematicalString(), func(t *testing.T) {
			// Function/Method to test
			got := testExpression.DeleteNonTerminal()

			errors := got.EvaluateTreeCorrectness()
			if len(errors) > 0 {
				t.Fatalf("DeleteNonTerminal - FormatErrors\n%s", errors.ToString(got))
			}

			// Test condition
			if len(testExpression) > 4 {
				if testExpression.isShorterEq(got) {
					ba := beforeAndAfter(testExpression, got)

					t.Log(ba)
					t.Errorf("DeleteNonTerminal() = len(got): %d | len(ans): %d", len(got), len(testExpression))
				}
			}
		})
	}

}

// expressionGenerator creates a set of expressions. It allows one to add some parenthesized expressions that will be added
// to the head of list
func expressionGenerator(count, maxSize int, parenthesizedExprs ...string) []BinaryTree {
	rand.Seed(time.Now().UnixNano())

	parLen := len(parenthesizedExprs)
	finalLength := count + parLen
	trees := make([]BinaryTree, finalLength)

	for i := 0; i < parLen; i++ {
		trees[i] = NewFromParenthesizedExpression(parenthesizedExprs[i])
	}

	for i := 0; i < finalLength-parLen; i++ {
		size := rand.Intn(maxSize)
		trees[i+parLen] = GenerateRandomTree(size, BTreeTerminals, BTreeNonTerminals)
	}

	return trees
}

func beforeAndAfter(before, after BinaryTree) string {
	return fmt.Sprintf("\n\tbefore: %s\n\tafter: %s\n", before.ToMathematicalString(), after.ToMathematicalString())
}

func TestBinaryTree_ApplyOperatorOnTerminal(t *testing.T) {
	type args struct {
		operator rune
		terminal rune
	}
	tests := []struct {
		name string
		bt   BinaryTree
		args args
		want BinaryTree
	}{
		{"", NewFromParenthesizedExpression("0"), args{'*', 'x'}, NewFromParenthesizedExpression("(0*x)")},
		{"", NewFromParenthesizedExpression("(x+1)"), args{'*', 'x'}, NewFromParenthesizedExpression("((x+1)*x)")},
		{"", NewFromParenthesizedExpression("((x+1)*1)"), args{'*', 'x'}, NewFromParenthesizedExpression("(((x+1)*1)*x)")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.ApplyOperatorOnTerminal(tt.args.operator, tt.args.terminal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyOperatorOnTerminal() = %s, wamt %s\n%s, want %s", got.ToMathematicalString(), tt.want.ToMathematicalString(),   got.ToString(), tt.want.ToString())
			}
		})
	}
}

//func Test_prependNode(t *testing.T) {
//	type args struct {
//		x BinaryTree
//		y BinaryTreeNode
//	}
//	tests := []struct {
//		name string
//		args args
//		want BinaryTree
//	}{
//		//{"",args{NewFromParenthesizedExpressionUnsafe("x"), BinaryTreeNode{0, '('}}, NewFromParenthesizedExpressionUnsafe("(x")},
//		{"",args{NewFromParenthesizedExpressionUnsafe("(x+1)"), BinaryTreeNode{0, '('}}, NewFromParenthesizedExpressionUnsafe("((x+1)")},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := prependNode(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ApplyOperatorOnTerminal() = %s, wamt %s\n%s, want %s", got.ToMathematicalString(), tt.want.ToMathematicalString(),
//					got.ToString(), tt.want.ToString())
//			}
//		})
//	}
//}