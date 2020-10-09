package evolution

import (
	"math"
)

// Eval is a simple helper function that takes in an independent variable,
// uses the programs treeNode to compute the resultant value
//func (bt BinaryTree) EvalMulti(independentVariables IndependentVariableMap, expressionString string) (float64,
//	error) {
//	return EvaluateMathematicalExpression(expressionString, independentVariables)
//}

const (
	MathErrorGeneral = 0
	MathErrorInvalid = 1
)

// EvaluateMathematicalExpression evaluates a valid expression using the given independentVar
func EvaluateMathematicalExpression(tree BinaryTree, xVal float64) float64 {
	return Calculate(tree, xVal)
}

func Calculate(bt BinaryTree, xVal float64) float64 {
	index := bt.FirstTerminal()

	if len(bt) <= 4 {
		if len(bt) == 1 {
			if bt[0].value == 'x' {
				return xVal
			} else {
				return float64(bt[0].value - '0')
			}
		} else {
			panic("Calculate: tree is of obscure length > 1 but less than 4")
		}
	}

	var ans float64
	var ok bool

	ans, ok = eval(bt[index].value, bt[index+2].value, bt[index+1].value, xVal)
	if !ok {
		return math.NaN()
	}
	if math.IsNaN(ans) {
		return math.NaN()
	}

	counter := 1
	for i := index; i < len(bt)-index; i++ {
		nextTerminalIndex := (index + 2) + (3 * counter)
		nextOp := nextTerminalIndex - 1

		if nextTerminalIndex >= len(bt)-1 {
			return ans
		}

		var n float64
		var ok bool

		n, ok = evall(ans, bt[nextTerminalIndex].value, bt[nextOp].value, xVal)

		if !ok {
			return math.NaN()
		}

		ans = n
		counter++
	}

	return ans
}

func eval(l, r, op rune, xVal float64) (float64, bool) {
	newL := 0.0
	newR := 0.0

	if l == 'x' {
		newL = xVal
	} else {
		newL = float64(l - '0')
	}

	if r == 'x' {
		newR = xVal
	} else {
		newR = float64(r - '0')
	}

	switch op {
	case '+':
		return newL + newR, true
	case '-':
		return newL - newR, true
	case '*':
		return newL * newR, true
	case '/':
		if newR == 0 {
			return math.NaN(), false
		}
		return newL / newR, true
	}

	panic("eval: no valid operation")
}

func evall(l float64, r, op rune, xVal float64) (float64, bool) {
	newR := 0.0
	if r == 'x' {
		newR = xVal
	} else {
		newR = float64(r - '0')
	}

	switch op {
	case '+':
		return l + newR, true
	case '-':
		return l - newR, true
	case '*':
		return l * newR, true
	case '/':
		if float64(r-'0') == 0 {
			return math.NaN(), false
		}
		return l / newR, true
	}

	panic("eval: no valid operation")
}
