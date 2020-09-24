package tree

import (
	"math/rand"
	"testing"
	"time"
)

var (
	randomTerminal = BinaryTreeNode{}
	binTree        = BinaryTree{}
	mathExpression = ""
)

//func BenchmarkBinaryTree_FromRandomTerminal(b *testing.B) {
//	b.ReportAllocs()
//	expressionSet := GenerateRandomSymbolicExpressionSet(1)
//	tree1 := BinaryTree{}
//
//	tree1.FromSymbolicExpressionSet(expressionSet)
//	for i := 0; i < b.N; i++ {
//		randomTerminal, _ = tree1.RandomTerminal()
//	}
//}

func BenchmarkBinaryTree_NewFromExpression(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		binTree = NewFromExpression("x*x+1/3-4")
	}
}

func BenchmarkBinaryTree_GenerateRandomTree_0(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ReportAllocs()

	terminals := BTreeTerminals
	nonTerminals := BTreeNonTerminals

	for i := 0; i < b.N; i++ {
		binTree = GenerateRandomTree(terminals, nonTerminals, 0)
	}
}

func BenchmarkBinaryTree_GenerateRandomTree_10(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ReportAllocs()

	terminals := BTreeTerminals
	nonTerminals := BTreeNonTerminals

	for i := 0; i < b.N; i++ {
		binTree = GenerateRandomTree(terminals, nonTerminals, 10)
	}
}

func BenchmarkBinaryTree_GenerateRandomTree_1000(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ReportAllocs()

	terminals := BTreeTerminals
	nonTerminals := BTreeNonTerminals

	for i := 0; i < b.N; i++ {
		binTree = GenerateRandomTree(terminals, nonTerminals, 1000)
	}
}

func BenchmarkBinaryTree_GenerateRandomTree_1000000(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ReportAllocs()

	terminals := BTreeTerminals
	nonTerminals := BTreeNonTerminals

	for i := 0; i < b.N; i++ {
		binTree = GenerateRandomTree(terminals, nonTerminals, 1000000)
	}
}

var (
	openingParenIndex = 0
	closingParenIndex = 0
)

func BenchmarkCalculateEnclosingParentheses(b *testing.B) {
	nonTerminalCount := 1000
	length := treeLenFromNonTerminals(nonTerminalCount)
	middleNonTerminalIndex := nonTerminalIndices(length)
	b.Logf("tree size: %d\n", length)
	tree := GenerateRandomTree(BTreeTerminals, BTreeNonTerminals, nonTerminalCount)

	for i := 0; i < b.N; i++ {
		openingParenIndex, closingParenIndex = calculateEnclosingParentheses(tree, middleNonTerminalIndex[len(middleNonTerminalIndex)/2])
	}
}
