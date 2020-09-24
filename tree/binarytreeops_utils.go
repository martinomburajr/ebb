package tree


func swapNTNT(old, new BinaryTree, nonTerminalIndex int) BinaryTree {
	lParen, rParen := calculateEnclosingParentheses(old, nonTerminalIndex)
	lenOld := len(old)
	lenNew := len(new)

	finalTreeLength := (lenOld - 1) - (rParen - lParen) + lenNew
	newTree := make([]BinaryTreeNode, finalTreeLength)

	start := lParen
	mid := lenNew
	end := finalTreeLength - (mid + start)

	// First few elements from old
	for i := 0; i < start; i++ {
		newTree[i] = old[i]
	}

	for i := 0; i < mid; i++ {
		newTree[i+start] = new[i]
		newTree[i+start].key = i + start
	}

	// Last section from old
	for i := 0; i < end; i++ {
		newTree[i+mid+start] = old[i+rParen+1]
		newTree[i+mid+start].key = i + mid + start
	}

	return newTree
}

// swapTT swaps a terminal in a the old tree for a new terminal
func swapTT(old BinaryTree, terminalIndex int, terminal rune) BinaryTree {
	old[terminalIndex].value = terminal

	return old
}