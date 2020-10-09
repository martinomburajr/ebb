package evolution

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type BinaryTreeNode struct {
	key   int
	value rune
}

//BinaryTree representation as a contiguously allocated array.
type BinaryTree []BinaryTreeNode

// Clone performs a hard clone and allocates a new BinaryTree with the same internal structure as the one being cloned
func (bt BinaryTree) Clone() BinaryTree {
	dst := make([]BinaryTreeNode, len(bt))
	copy(dst, bt)

	return dst
}

func (bt BinaryTree) TerminalCount() int {
	return ((len(bt) - 1) / 4) + 1
}

func (bt BinaryTree) FirstTerminal() int {
	return firstTerminalIndex(len(bt))
}

func (bt BinaryTree) FirstNonTerminal() int {
	return firstNonTerminalIndex(len(bt))
}

// Root returns the index of the root
func (bt BinaryTree) Root() int {
	if len(bt) <= 4 {
		return -1
	}

	return len(bt) - 3
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

// RandomTerminal returns a random terminals index and value
func (bt BinaryTree) RandomTerminal() (int, BinaryTreeNode) {
	randTerminalIndex := randomTerminalIndex(len(bt))

	return randTerminalIndex, bt[randTerminalIndex]
}

// RandomNonTerminal returns a random non-terminals index and value
func (bt BinaryTree) RandomNonTerminal() (int, BinaryTreeNode) {
	if len(bt) <= 4 {
		panic("RandomNonTerminal: tree is less than 3 characters. Cannot have non-terminal")
	}

	randNonTerminalIndex := randomNonTerminalIndex(len(bt))

	return randNonTerminalIndex, bt[randNonTerminalIndex]
}

// NewFromExpression takes in a mathematical expression string and converts it into a BinaryTree
func NewFromExpression(expression string) BinaryTree {
	if len(expression) < 1 {
		panic("NewFromExpression: expression cannot be empty")
	}

	if len(expression) == 1 {
		return BinaryTree{
			BinaryTreeNode{
				key:   0,
				value: rune(expression[0]),
			},
		}
	}

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

	var bt BinaryTree = make([]BinaryTreeNode, len(expression))

	for i := 0; i < len(expression); i++ {
		bt[i] = BinaryTreeNode{key: i, value: rune(expression[i])}
	}

	errors := bt.EvaluateTreeCorrectness()
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

// NewFromParenthesizedExpression creates a Binary tree from a mathematical expression that MUST contain parentheses.
// The format of the parentheses must be lead heavy i.e all the open parentheses '(' accumulate at the start
func NewFromParenthesizedExpressionUnsafe(expression string) BinaryTree {
	if len(expression) < 1 {
		panic("NewFromParenthesizedExpression: expression cannot be empty")
	}

	var bt BinaryTree = make([]BinaryTreeNode, len(expression))

	for i := 0; i < len(expression); i++ {
		bt[i] = BinaryTreeNode{key: i, value: rune(expression[i])}
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

func (bt BinaryTree) ToString() string {
	builder := strings.Builder{}

	builder.WriteRune('[')
	for i := 0; i < len(bt); i++ {
		builder.WriteRune('{')
		builder.WriteString(strconv.FormatInt(int64(bt[i].key), 10))
		builder.WriteRune(' ')
		builder.WriteRune(bt[i].value)
		builder.WriteRune('}')
		builder.WriteRune(' ')
	}
	builder.WriteRune(']')

	return builder.String()
}

// GenerateRandomTree creates a tree given a set of terminals and nonTerminals to use. It generates a
// tree of size nonTerminalCount. nonTerminalCount does not refer to the final size of the binaryTree but rather the number of non-terminals
// the final tree will contain. A length of 0 returns a tree expression with a single terminal, a length > 0 returns a tree
// expression with length non-terminals
func GenerateRandomTree(nonTerminalCount int, terminals, nonTerminals []rune) BinaryTree {
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

// NewRandomTreeFromIVarCount creates a new tree with polDegree number of independent variable (x) scattered across
// the set of terminals.
func NewRandomTreeFromIVarCount(ivarCount int, maxAdditionalNTCount int, terminals, nonTerminals []rune) BinaryTree {
	nonTerminalCount := ivarCount + rand.Intn(maxAdditionalNTCount)

	randomTree := GenerateRandomTree(nonTerminalCount, terminals, nonTerminals)

	terminalIndices := randomTree.TerminalIndices()

	rand.Shuffle(len(terminalIndices), func(i, j int) {
		terminalIndices[i], terminalIndices[j] = terminalIndices[j], terminalIndices[i]
	})

	currXCount := 0

	for currXCount < ivarCount {
		currXCount = 0

		for i := 0; i < ivarCount; i++ {
			if randomTree[terminalIndices[i]].value == 'x' {
				currXCount++
			} else {
				randomTree[terminalIndices[i]].value = 'x'
			}
		}
	}

	randomTree.Sanitize()

	return randomTree
}

// NewRandomTreeFromPolDegreeCount creates a new tree with polDegree number of independent variable (x) scattered across
// the set of terminals.
func NewRandomTreeFromPolDegreeCount(polDegree int, maxAdditionalNTCount int, terminals, nonTerminals []rune) (BinaryTree, error) {
	nonTerminalCount := polDegree + rand.Intn(maxAdditionalNTCount)

	randomTree := GenerateRandomTree(nonTerminalCount, terminals, nonTerminals)

	terminalIndices := randomTree.TerminalIndices()

	rand.Shuffle(len(terminalIndices), func(i, j int) {
		terminalIndices[i], terminalIndices[j] = terminalIndices[j], terminalIndices[i]
	})

	currXCount := 0

	for currXCount < polDegree {
		currXCount = 0

		for i := 0; i < len(terminalIndices); i++ {
			if randomTree[terminalIndices[i]].value == 'x' && randomTree[terminalIndices[i]-1].value == '*' {
				currXCount++
			} else {
				firstTerminalIndex := firstTerminalIndex(len(randomTree))
				if terminalIndices[i] != firstTerminalIndex {
					randomTree[terminalIndices[i]-1].value = '*'
					randomTree[terminalIndices[i]].value = 'x'
					currXCount++
				}
			}

			if polDegree == currXCount {
				break
			}
		}
	}

	randomTree.Sanitize()

	treeErrors := randomTree.EvaluateTreeCorrectness()
	if len(treeErrors) > 0 {
		return nil, fmt.Errorf(treeErrors.ToString(randomTree))
	}

	return randomTree, nil
}

// Sanitize swaps any explicit '/0' with '/1'
func (bt BinaryTree) Sanitize() BinaryTree {

	for i := 0; i < len(bt)-1; i++ {
		if bt[i].value == '/' && bt[i+1].value == '0' {
			bt[i+1].value = '1'
		}
	}

	return bt
}

//NakedExpressionCount is a mathematical expression without the helpful parentheses.
func (bt BinaryTree) NakedExpressionCount() int {
	return (len(bt) + 1) / 2
}
