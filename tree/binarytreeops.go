package tree

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
	newTree := make([]BinaryTreeNode, lenBT+4)
	newTree[0] = BinaryTreeNode{
		key:   0,
		value: '(',
	}

	for i := 1; i < lenBT+1; i++ {
		newTree[i] = bt[i-1]
		newTree[i].key = i
	}

	newTree[lenBT+1] = BinaryTreeNode{key: lenBT + 1, value: operator}
	newTree[lenBT+2] = BinaryTreeNode{key: lenBT + 2, value: terminal}
	newTree[lenBT+3] = BinaryTreeNode{key: lenBT + 3, value: ')'}

	return newTree
}

// AddSubTreeRandomly swaps an incoming tree in a random non-terminal position. The incoming tree can be a single terminal
func (bt BinaryTree) AddSubTreeRandomly(tree BinaryTree) BinaryTree {
	// Guards against invokking a nonTerminal that may not exist given the size of bt
	if len(bt) <= 4 {
		return tree
	}

	nonTerminalIndex, _ := bt.RandomNonTerminal()
	return swapNTNT(bt, tree, nonTerminalIndex)
}


func (bt BinaryTree) AddRandomSubTreeTerminal(tree BinaryTree) BinaryTree {
	bt.Terminals()

	return BinaryTree{}
}
