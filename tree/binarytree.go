package tree

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type BinaryTreeNode struct {
	key   int
	value rune
}

//BinaryTree representation as a contiguously allocated array.
type BinaryTree []BinaryTreeNode

//NakedExpressionCount is a mathematical expression without the helpful parentheses.
func (bt BinaryTree) NakedExpressionCount() int {
	return (len(bt) + 1) / 2
}

func (bt BinaryTree) TerminalCount() int {
	return ((len(bt) - 1) / 4) + 1
}

func (bt BinaryTree) FirstTerminal() (index int) {
	return firstTerminalIndex(len(bt))
}

func (bt BinaryTree) FirstNonTerminal() (index int) {
	return firstNonTerminalIndex(len(bt))
}

// TerminalIndices returns a list of indices of terminals of the given BinaryTree
// note the binary tree already contains the '(' and ')' runes where appropriate
func (bt BinaryTree) TerminalIndices() []int {
	lenBT := len(bt)
	if lenBT < 1 {
		panic("TerminalIndices: BinaryTree cannot be empty")
	}
	if lenBT == 1 {
		return []int{0}
	}

	numTerminals := bt.TerminalCount()
	indices := make([]int, numTerminals)

	indices[0] = (lenBT - 1) / 4
	indices[1] = (lenBT-1)/4 + 2

	for i := 2; i < numTerminals; i++ {
		indices[i] = indices[i-1] + 3
	}

	return indices
}

func (bt BinaryTree) Terminals() []BinaryTreeNode {
	lenBT := len(bt)
	if lenBT < 1 {
		panic("Terminals: cannot be empty")
	}

	if lenBT == 1 {
		return []BinaryTreeNode{bt[0]}
	}

	numTerminals := bt.TerminalCount()
	nodes := make([]BinaryTreeNode, numTerminals)

	nodes[0] = bt[(lenBT-1)/4]
	nodes[1] = bt[(lenBT-1)/4+2]

	for i := 2; i < numTerminals; i++ {
		nodes[i] = bt[nodes[i-1].key+3]
	}

	return nodes
}

func (bt BinaryTree) NonTerminalCount() int {
	return (len(bt) - 1) / 4
}

func (bt BinaryTree) NonTerminalIndices() []int {
	if len(bt) < 1 {
		panic("nonTerminalIndices: BinaryTree cannot be empty")
	}

	if len(bt) == 1 {
		return []int{}
	}

	return nonTerminalIndices(len(bt))
}

func (bt BinaryTree) NonTerminals() []BinaryTreeNode {
	if len(bt) < 1 {
		panic("Terminals: cannot be empty")
	}

	if len(bt) == 1 {
		return []BinaryTreeNode{}
	}

	expressionCount := bt.NakedExpressionCount()
	numNonTerminals := int(float64(expressionCount) / 2)
	nodes := make([]BinaryTreeNode, numNonTerminals)

	nodes[0] = bt[int(math.Ceil(float64(expressionCount)/2))]

	for i := 1; i < numNonTerminals; i++ {
		nodes[i] = bt[nodes[i-1].key+3]
	}

	return nodes
}

// CalculateNewLength calculates the length of the BinaryTree generated from the expression. This includes
// added parentheses.
func CalculateNewLength(expression string) int {
	if len(expression) < 1 {
		panic("CalculateNewLength: expression cannot be empty")
	}
	if len(expression) == 1 {
		return 1
	}

	length := len(expression)
	return length + length - 1
}

func (bt BinaryTree) RandomTerminal() (int, BinaryTreeNode) {
	indices := bt.TerminalIndices()
	if len(indices) == 1 {
		return indices[0], bt[indices[0]]
	}

	randIndex := rand.Intn(len(indices))
	return indices[randIndex], bt[indices[randIndex]]
}

func (bt BinaryTree) RandomNonTerminal() (int, BinaryTreeNode) {
	if len(bt) <= 4 {
		panic("RandomNonTerminal: tree is less than 3 characters. Cannot have non-terminal")
	}

	indices := bt.NonTerminalIndices()
	if len(indices) == 1 {
		return indices[0], bt[indices[0]]
	}

	randIndex := rand.Intn(len(indices))
	return indices[randIndex], bt[indices[randIndex]]
}

// NewFromExpression takes in a mathematical expression string and converts it into a BinaryTree
func NewFromExpression(expression string) BinaryTree {
	if len(expression) < 1 {
		panic("NewFromExpression: expression cannot be empty")
	}

	if len(expression) == 1 {
		return BinaryTree{
			//BinaryTreeNode{
			//	key:   0,
			//	value: '(',
			//},
			BinaryTreeNode{
				key:   0,
				value: rune(expression[0]),
			},
			//BinaryTreeNode{
			//	key:   2,
			//	value: ')',
			//},
		}
	}

	// move len(expression) to single variable

	totalChars := len(expression) + len(expression) - 1
	bt := make([]BinaryTreeNode, totalChars)

	outerI := len(expression) / 2

	// setup outer-left pair of parentheses
	for i := 0; i < len(expression)/2; i++ {
		bt[i] = BinaryTreeNode{
			key:   i,
			value: '(',
		}
	}
	// populate next item
	bt[outerI] = BinaryTreeNode{
		key:   outerI,
		value: rune(expression[0]),
	}

	// populate remaining items from the start of expression
	counter := totalChars - outerI - 1
	innerCounter := 0
	for i := 0; i < counter; i++ {
		if i%3 < 2 {
			innerCounter++
			bt[outerI+1+i] = BinaryTreeNode{
				key:   outerI + 1 + i,
				value: rune(expression[innerCounter]),
			}
		} else {
			bt[outerI+1+i] = BinaryTreeNode{
				key:   outerI + 1 + i,
				value: ')',
			}
		}
	}

	return bt
}

// NewFromParenthesizedExpression creates a Binary tree from a mathematical expression that MUST contain parentheses.
// The format of the parentheses must be lead heavy i.e all the open parentheses '(' accumulate at the start
func NewFromParenthesizedExpression(expression string) BinaryTree {
	if len(expression) < 1 {
		panic("NewFromParenthesizedExpression: expression cannot be empty")
	}

	bt := make([]BinaryTreeNode, len(expression))

	for i := 0; i < len(expression); i++ {
		bt[i] = BinaryTreeNode{key: i, value: rune(expression[i])}
	}

	errors := evaluateTreeCorrectness(bt)
	if len(errors) > 0 {
		sb := strings.Builder{}

		for i := 0; i < len(errors); i++ {
			sb.WriteString(errors[i].Error())
			sb.WriteRune('\n')
		}

		err := fmt.Sprintf("Tree Correctness Error:\n\t%s", sb.String())
		panic(err)
	}

	return bt
}

func (bt BinaryTree) GetRandomSubTree() (parentIndex int32, node BinaryTreeNode) {
	return 0, BinaryTreeNode{}
}

func (bt BinaryTree) ToMathematicalString() string {
	builder := strings.Builder{}

	for i := 0; i < len(bt); i++ {
		builder.WriteRune(bt[i].value)
	}

	return builder.String()
}

// GenerateRandomTree creates a tree given a set of terminals and nonTerminals to use. It generates a
// tree of size nonTerminalCount. nonTerminalCount does not refer to the final size of the binaryTree but rather the number of non-terminals
// the final tree will contain. A length of 0 returns a tree expression with a single terminal, a length > 0 returns a tree
// expression with length non-terminals
func GenerateRandomTree(terminals, nonTerminals []rune, nonTerminalCount int) BinaryTree {
	if len(terminals) < 1 {
		panic("GenerateRandomTree: terminals cannot be empty")
	}

	if len(nonTerminals) < 1 && nonTerminalCount > 1 {
		panic("GenerateRandomTree: at least one nonTerminal required")
	}

	if nonTerminalCount == 0 {
		randIndex := rand.Intn(len(terminals))
		return BinaryTree{{
			key:   0,
			value: terminals[randIndex],
		}}
	}

	nonTerminalIndices := make([]int, nonTerminalCount)
	terminalIndices := make([]int, nonTerminalCount+1)
	terminalL := len(terminals)
	nonTerminalL := len(nonTerminals)

	for i := 0; i < nonTerminalCount; i++ {
		terminalIndices[i] = rand.Intn(terminalL)
		nonTerminalIndices[i] = rand.Intn(nonTerminalL)
	}

	// add to the tail of terminal
	terminalIndices[len(terminalIndices)-1] = rand.Intn(terminalL)

	// BEGIN POPULATING
	expressionLen := 4*nonTerminalCount + 1
	bt := make([]BinaryTreeNode, expressionLen)

	outerI := nonTerminalCount

	// setup outer-left pair of parentheses
	for i := 0; i < nonTerminalCount; i++ {
		bt[i] = BinaryTreeNode{
			key:   i,
			value: '(',
		}
	}
	// populate next item
	bt[outerI] = BinaryTreeNode{
		key:   outerI,
		value: rune(terminals[terminalIndices[0]]),
	}

	// populate remaining items from the start of expression
	counter := expressionLen - outerI - 1
	innerCounter := 0
	counterT := 0
	counterNT := 0

	for i := 0; i < counter; i++ {
		if i%3 < 2 {
			innerCounter++
			if innerCounter%2 == 0 {
				// terminal
				bt[outerI+1+i] = BinaryTreeNode{
					key:   outerI + 1 + i,
					value: rune(terminals[terminalIndices[counterT]]),
				}
				counterT++
			} else {
				// nonTerminal
				bt[outerI+1+i] = BinaryTreeNode{
					key:   outerI + 1 + i,
					value: rune(nonTerminals[nonTerminalIndices[counterNT]]),
				}
				counterNT++
			}

		} else {
			bt[outerI+1+i] = BinaryTreeNode{
				key:   outerI + 1 + i,
				value: ')',
			}
		}
	}

	return bt
}
