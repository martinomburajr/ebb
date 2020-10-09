package evolution

import "math/rand"

func (bt BinaryTree) MutateTerminal(terminalSet []rune) BinaryTree {
	// 1. Select random terminal
	index, terminal := bt.RandomTerminal()
	// 2. Select random terminalSet
	randomTerminalSetIndex := rand.Intn(len(terminalSet))
	// 3. Perform replace
	bt[index] = BinaryTreeNode{
		key:   terminal.key,
		value: terminalSet[randomTerminalSetIndex],
	}

	return bt
}

// MutateNonTerminal will attempt to mutate a non-terminal given the tree is large enough to contain a non-terminal
// in the first place
func (bt BinaryTree) MutateNonTerminal(nonTerminalSet []rune) BinaryTree {
	if len(bt) <= 4 {
		return bt
	}

	// 1. Select random nonTerminal
	index, nonTerminal := bt.RandomNonTerminal()
	// 2. Select random terminalSet
	randomTerminalSetIndex := rand.Intn(len(nonTerminalSet))
	// 3. Perform replace
	bt[index] = BinaryTreeNode{
		key:   nonTerminal.key,
		value: nonTerminalSet[randomTerminalSetIndex],
	}

	return bt
}

// TODO - FIX
func (bt BinaryTree) ReplaceBranch(newBT BinaryTree) BinaryTree {
	// 1. Select random nonTerminal
	index, _ := bt.RandomTerminal()

	finalTree := make([]BinaryTreeNode, len(newBT)+len(bt)-1)

	rem := len(bt) - index
	// bt
	for i := 0; i < index; i++ {
		finalTree[i] = bt[i]
	}

	//newBT
	outerI := 0
	for i := index; i < len(newBT); i++ {
		finalTree[i] = newBT[outerI]
		outerI++
	}

	//newBT
	for i := index + len(newBT); i < rem; i++ {
		finalTree[i] = bt[rem+i]
	}

	return finalTree
}

func (bt BinaryTree) Fell() BinaryTree {
	return BinaryTree{{key: 0, value: '0'}}
}

// ApplyOperatorOnTerminal - The operator is any mathematical operator i.e '*', '/', '-', '+' rune. This
// operator applies itself to the current tree and operates on the supplied terminal.
// E.g. If bt is ((x+1)-3) an operator of '/' and terminal of 'x' will
// result in the following tree. (((x+1)-3)/x)
func (bt BinaryTree) ApplyOperatorOnTerminal(operator rune, terminal rune) BinaryTree {
	lenBT := len(bt)
	newTreeLen := lenBT + 4
	newTree := make([]BinaryTreeNode, newTreeLen)

	copy(newTree[1:], bt)
	newTree[0] =  BinaryTreeNode{key:   0, value: '('}

	if lenBT == 2 {
		newTree[1].key = 1

		newTree[lenBT+1] = BinaryTreeNode{key: lenBT, value: operator}
		newTree[lenBT+2] = BinaryTreeNode{key: lenBT + 1, value: terminal}
		newTree[lenBT+3] = BinaryTreeNode{key: lenBT + 2, value: ')'}

		return newTree
	}

	// fix keys
	for i := 1; i < lenBT+1; i++ {
		newTree[i].key = i
	}

	newTree[lenBT+1] = BinaryTreeNode{key: lenBT+1, value: operator}
	newTree[lenBT+2] = BinaryTreeNode{key: lenBT + 2, value: terminal}
	newTree[lenBT+3] = BinaryTreeNode{key: lenBT + 3, value: ')'}

	return newTree
}

//func prependNode(x BinaryTree, y BinaryTreeNode) BinaryTree {
//	lastIndex := len(x) - 1
//	end := x[lastIndex]
//	x = append(x, BinaryTreeNode{})
//
//	stack := [2]BinaryTreeNode{}
//
//	stack[0] = x[0]
//	stack[1] = x[1]
//	for i := 1; i < len(x)-1; i++ {
//		x[i] = stack[0]
//		x[i].key=i
//
//		// update stack
//		stack[0] = stack[1]
//		stack[1] = x[i+1]
//	}
//	x[lastIndex] = end
//	x[lastIndex].key = lastIndex
//
//	//copy(x[1:], x)
//	x[0] = y
//	return x
//}

// AddSubTreeRandomly swaps an incoming tree in a random non-terminal position. The incoming tree can be a single terminal
func (bt BinaryTree) AddSubTreeRandomly(tree BinaryTree) BinaryTree {
	// Guards against invoking a nonTerminal that may not exist given the size of bt
	if len(bt) <= 4 {
		return tree
	}

	nonTerminalIndex, _ := bt.RandomNonTerminal()
	return swapNTNT(bt, tree, nonTerminalIndex)
}

func (bt BinaryTree) AppendRandomOperation(terminals, nonTerminals []rune) BinaryTree {
	randTerminalIndex := rand.Intn(len(terminals))
	randNonTerminalIndex := rand.Intn(len(nonTerminals))

	lenBT := len(bt)
	newTree := make([]BinaryTreeNode, lenBT+4)

	newTree[0] = BinaryTreeNode{
		key:   0,
		value: '(',
	}

	for i := 1; i < lenBT+1; i++ {
		newTree[i] = bt[i-1]
		newTree[i].key = i
	}

	newTree[lenBT+1] = BinaryTreeNode{key: lenBT + 1, value: nonTerminals[randNonTerminalIndex]}
	newTree[lenBT+2] = BinaryTreeNode{key: lenBT + 2, value: terminals[randTerminalIndex]}
	newTree[lenBT+3] = BinaryTreeNode{key: lenBT + 3, value: ')'}

	return newTree
}

func (bt BinaryTree) DeleteNonTerminal() BinaryTree {
	// IGNORE trees without non-terminals
	treeLen := len(bt)
	if treeLen <= 4 {
		return bt
	}
	if treeLen == 5 {
		return BinaryTree{{key: 0, value: '0'}}
	}

	nonTerminalIndex := randomNonTerminalIndex(treeLen)

	isRoot := isIndexRoot(nonTerminalIndex, treeLen)

	if isRoot {
		return BinaryTree{{key: 0, value: '0'}}
	}

	lParen, rParen := calculateEnclosingParenthesesIndices(bt, nonTerminalIndex)

	newTreeLen := treeLen - (rParen - lParen)

	tree := make([]BinaryTreeNode, newTreeLen)

	for i := 0; i < lParen; i++ {
		tree[i] = bt[i]
		tree[i].key = i
	}

	tree[lParen] = BinaryTreeNode{key: lParen, value: '0'}

	rem := treeLen - rParen - 1
	for i := 0; i < rem; i++ {
		treeI := i + lParen + 1
		btI := i + rParen + 1

		tree[treeI] = BinaryTreeNode{key: treeI, value: bt[btI].value}
	}

	return tree
}

// DeleteTerminal attempts to ensure that whichever terminal is selected, based on the preceding operation, the terminal is set to a value
// that embodies what would happen if the terminal was not there. e.g. if the terminal is preceded by a '+' or '-' sign it will set the terminal to 0. if it is
// preceded by a '*' or '/', the terminal is set to a value of zero. These have the effect of nullifying the influence of the operation.
func (bt BinaryTree) DeleteTerminal() BinaryTree {
	randTerminalIndex := randomTerminalIndex(len(bt))

	if len(bt) > 1 {
		switch bt[randTerminalIndex-1].value {
		case '/':
			bt[randTerminalIndex].value = '1'

			return bt
		case '*':
			bt[randTerminalIndex].value = '1'

			return bt
		case '+':
			bt[randTerminalIndex].value = '0'

			return bt
		case '-':
			bt[randTerminalIndex].value = '0'

			return bt
		}

		return bt
	}

	bt[randTerminalIndex].value = '0'

	return bt
}

// contains checks to see if a binary tree contains a certain rune. If not it returns -1 otherwise
// it returns the first index of that item.
func (bt BinaryTree) contains(needle rune) int {
	for i := 0; i < len(bt); i++ {
		if bt[i].value == needle {
			return i
		}
	}

	return -1
}

// containsAll checks to see if a binary tree contains a certain rune. It returns a slice of all indices
// where the rune was found
func (bt BinaryTree) containsAll(needle rune) []int {
	ans := make([]int, 0)

	for i := 0; i < len(bt); i++ {
		if bt[i].value == needle {
			ans = append(ans, i)
		}
	}

	return ans
}
