package tree

import (
	"math/rand"
	"testing"
	"time"
)

var ans float64
var err error

func BenchmarkBinaryTree_CalculateRolling_NT16(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ReportAllocs()

	bt := NewFromParenthesizedExpression("((((((((((((((((2/2)/1)/8)-6)+4)/2)-9)-6)-1)*6)*4)/1)/9)-8)-5)*3)")
	for i := 0; i < b.N; i++ {
		ans, err = CalculateRolling(bt, -4)
	}
}

func BenchmarkBinaryTree_CalculateRolling_NT194(b *testing.B) {
	b.ReportAllocs()
	bt := NewFromParenthesizedExpression(longestXExpr)

	for i := 0; i < b.N; i++ {
		ans, err = CalculateRolling(bt, -4)
	}
}