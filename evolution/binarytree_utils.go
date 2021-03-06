package evolution

import (
	"fmt"
	"math/rand"
	"strings"
)

// nonTerminalIndices returns a list of non-terminal indices for a given verifiable tree of length treeLen
func nonTerminalIndices(treeLen int) []int {
	if treeLen <= 4 {
		panic("nonTerminalIndices: tree is too small")
	}

	L := treeLen - 1

	positions := make([]int, L/4)
	for i := 0; i < L/4; i++ {
		positions[i] = ((L / 4) + 1) + (3 * i)
	}

	return positions
}

//treeLenFromNonTerminals calculates how long a padded tree is from the number of non-terminals
func treeLenFromNonTerminals(ntCount int) int {
	return (4 * ntCount) + 1
}

// randomNonTerminalIndex returns a random index containing a non-terminal
func randomNonTerminalIndex(treeLen int) int {
	if treeLen <= 5 {
		return 2
	}

	totalNTs := (treeLen - 1) / 4

	randIndex := rand.Intn(totalNTs)

	return (totalNTs + 1) + (3 * randIndex)
}

// randomTerminalIndex returns a random index containing a terminal
func randomTerminalIndex(treeLen int) int {
	if treeLen <= 1 {
		return 0
	}

	totalTs := ((treeLen - 1) / 4) + 1

	if totalTs <= 1 {
		panic(totalTs)
	}

	randIndex := rand.Intn(totalTs - 1)

	i := (totalTs + 2) + (3 * randIndex) - 1
	return i
}

func isIndexRoot(index, treeLen int) bool {
	if index == (treeLen - 3) {
		return true
	}

	return false
}

// calculateEnclosingParenthesesIndices calculates the index of the surrounding parentheses of the given nonTerminal.
func calculateEnclosingParenthesesIndices(bt BinaryTree, nonTerminalIndex int) (int, int) {
	// Calculates if the supplied nonTerminalIndex is the first non-terminal in a padded expression
	btLen := len(bt)
	L := btLen - 1

	i := calculateNTPos(btLen, nonTerminalIndex)

	if (4*nonTerminalIndex)-4 == btLen-1 {
		return nonTerminalIndex - 2, nonTerminalIndex + 2
	}

	NT0 := (L / 4) + 1
	NTi := ((L / 4) + 1) + 3*i

	left := (NT0 - 2) - (NTi-NT0)/3
	right := nonTerminalIndex + 2
	return left, right
}

func calculateNumberOfParentheses(parenthesisedLength int) int {
	return (parenthesisedLength - 1) / 2
}

// calculateNTPos calculates the index value of the set of nonTerminals for the given nonTerminalIndex
// e.g. In a padded expression ((x*x)+1) the non terminal set is [*,+]. Given the index of a
// NT in the padded expression e.g. 3 for '*' and 6 for '+' will  return the associated index in the non-terminal set.
// Therefore a nonTerminalIndex of 3 give 0 and a nonTerminalIndex of 6 gives 1.
func calculateNTPos(treeLen int, nonTerminalIndex int) int {
	L := treeLen - 1
	return (nonTerminalIndex - ((L / 4) + 1)) / 3
}

// calculateTPos calculates the index value of the set of terminals for the given terminalIndex. It is similar to calculateNTPos
// but for terminals. The return value states whether the terminalIndex in the BinaryTree is the first terminal, second terminal etc.
// in the terminal set
func calculateTPos(bt BinaryTree, terminalIndex int) int {
	L := len(bt) - 1

	if len(bt) <= 3 {
		return 0
	}

	if terminalIndex == (L / 4) {
		return 0
	}

	num := (L / 4) + 2

	return ((terminalIndex - num) / 3) + 1
}

type BinaryTreeErrors []error

func (b BinaryTreeErrors) ToString(tree BinaryTree) string {
	if len(b) > 0 {
		sb := strings.Builder{}

		for i := 0; i < len(b); i++ {
			sb.WriteRune('\t')
			sb.WriteString(b[i].Error())
			sb.WriteRune('\n')
		}

		err := fmt.Sprintf("Tree Correctness Error:\nTree: %s\n\t%s", tree.ToMathematicalString(), sb.String())

		return err
	}

	return ""
}

//EvaluateTreeCorrectness performs checks on a given BinaryTree. Should be used solely for testing and rarely during runtime
func (bt BinaryTree) EvaluateTreeCorrectness() BinaryTreeErrors {
	openingBraces := 0
	closingBraces := 0
	//numberOfOpenBraces := 0
	expectedOpeningBraces := numberOfStartingBraces(len(bt))
	errorQueue := make([]error, 0)

	if len(bt) < 1 {
		err := fmt.Errorf("tree size - got: 0 | want: >= 1")

		errorQueue = append(errorQueue, err)
	}

	for i := 0; i < len(bt); i++ {
		if bt[i].key != i {
			err := fmt.Errorf("key mismatch - got: %d | want: %d", bt[i].key, i)

			errorQueue = append(errorQueue, err)
		}

		if bt[i].value == '(' {
			openingBraces++
		}
		if bt[i].value == ')' {
			closingBraces++
		}

		if i < expectedOpeningBraces {
			if bt[i].value != '(' {
				err := fmt.Errorf(
					"found illegal character where opening parenthesis should be\n"+
						"found rune: %d | index: %d", bt[i].value, i)

				errorQueue = append(errorQueue, err)
			}
		}
	}

	// Check Braces
	if openingBraces != closingBraces {
		err := fmt.Errorf("brace mismatch -> opening: %d | closing: %d", openingBraces, closingBraces)

		errorQueue = append(errorQueue, err)
	}

	// Check Terminals and NonTerminals are in the correct place

	nonTerminalIndices := bt.NonTerminalIndices()
	terminalIndices := bt.TerminalIndices()

	for i := 0; i < len(nonTerminalIndices); i++ {
		nonTerminal := bt[nonTerminalIndices[i]].value
		if !isValidRune(nonTerminal, BTreeNonTerminals) {
			err := fmt.Errorf("nonTerminal not found in index: %d - found: %s", nonTerminalIndices[i], string(nonTerminal))

			errorQueue = append(errorQueue, err)
		}
	}

	for i := 0; i < len(terminalIndices); i++ {
		terminal := bt[terminalIndices[i]].value
		if !isValidRune(terminal, BTreeTerminals) {
			err := fmt.Errorf("terminal not found in index: %d - found: %s", terminalIndices[i], string(terminal))

			errorQueue = append(errorQueue, err)
		}
	}

	return errorQueue
}

// firstTerminalIndex returns the expected position of the first terminal
func firstTerminalIndex(treeLen int) int {
	return (treeLen - 1) / 4
}

// firstNonTerminalIndex returns the expected position of the first non-terminal
func firstNonTerminalIndex(treeLen int) int {
	return ((treeLen - 1) / 4) + 1
}

// numberOfStartingBraces returns the expected number of Starting Braces
func numberOfStartingBraces(treeLen int) int {
	return (treeLen - 1) / 4
}

// isValidRune checks to see if a needle rune exists in the specified haystack
func isValidRune(needle rune, haystack []rune) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

// IsEqual checks if value and key are identical for each index. -2 indicates that trees differ in length, -1 indicates
// trees are identical, any other number >= indicates the difference at that index.
func (bt BinaryTree) IsEqual(test BinaryTree) int {
	if len(bt) != len(test) {
		return -2
	}

	for i := 0; i < len(bt); i++ {
		if bt[i].value != test[i].value {
			return i
		}
		if bt[i].key != test[i].key {
			return i
		}
	}

	return -1
}

func (bt BinaryTree) isShorter(test BinaryTree) bool {
	return len(bt) < len(test)
}

func (bt BinaryTree) isLarger(test BinaryTree) bool {
	return len(bt) > len(test)
}

func (bt BinaryTree) isLargerEq(test BinaryTree) bool {
	return len(bt) >= len(test)
}

func (bt BinaryTree) isShorterEq(test BinaryTree) bool {
	return len(bt) <= len(test)
}

func (bt BinaryTree) isLengthEq(test BinaryTree) bool {
	return len(bt) == len(test)
}
