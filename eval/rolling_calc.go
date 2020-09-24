package eval

//
//func CalculateRolling(bt tree.BinaryTree, xVal float64) (float64, error) {
//	index := bt.FirstTerminal()
//	ans, ok := eval(bt[index].value, bt[index+2].value, bt[index+1].value)
//	if !ok {
//		return 0, fmt.Errorf("invalid operation - attempted divByZero")
//	}
//
//	counter := 0
//	for i := index; i < len(bt)-index; i++ {
//		nextTerminal := index + 3*counter
//		nextOp := nextTerminal - 1
//		n, ok := evalc(ans, bt[nextTerminal].value, bt[nextOp].value)
//
//		if !ok {
//			return 0, fmt.Errorf("invalid operation - attempted divByZero")
//		}
//
//		ans += n
//		counter++
//	}
//
//	return ans, nil
//}
//
//func eval(l, r, op rune) (float64, bool) {
//	switch op {
//	case '+':
//		return float64(l-'0') + float64(r-'0'), true
//	case '-':
//		return float64(l-'0') - float64(r-'0'), true
//	case '*':
//		return float64(l-'0') * float64(r-'0'), true
//	case '/':
//		if float64(r-'0') == 0 {
//			return math.NaN(), false
//		}
//		return float64(l-'0') / float64(r-'0'), true
//	}
//
//	panic("eval: no valid operation")
//}
//
//func evalc(l float64, r, op rune) (float64, bool) {
//	switch op {
//	case '+':
//		return l + float64(r-'0'), true
//	case '-':
//		return l - float64(r-'0'), true
//	case '*':
//		return l * float64(r-'0'), true
//	case '/':
//		if float64(r-'0') == 0 {
//			return math.NaN(), false
//		}
//		return l / float64(r-'0'), true
//	}
//
//	panic("eval: no valid operation")
//}
