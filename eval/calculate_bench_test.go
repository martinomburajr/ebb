package eval

import (
	"testing"
)

var ans1 float64
var ans = 0.0
var expression = "1.23*2.1"
var expressionWMinus = "1.23*-2.1"
var expreessionLong = "(((1.23)*(1.34)+(2.3442))"
var expreessionLongWMinus = "(((1.23)*(-1.34)+(2.3442)))"
var expreessionLongg = "10.9*9.8*8.7*7.6*6.5*5.4*4.3*3.2*2.1*1"
var expressionVar = "x*10.9*9.8*8.7*7.6*6.5*5.4*4.3*3.2*2.1*1"
var expressionManyVar = "x*x*x*x*x*x*x*x*x*x*x*x*x"
var expressionManyVarXY = "x*y*x*y*x*y*x*y*x*y"
var exprLong = "(((((((((((((((((((((((((((5*5)/1)/3)-1)*1)*2)/2)*4)*3)*4)/3)*x)-9)*7)*x)-9)*8)/5)*4)+x)+x)/x)+3)*7)/5)+x)*3)"

func BenchmarkCalculate(b *testing.B) {
	b.ReportAllocs()
	var x1 float64
	for i := 0; i < b.N; i++ {
		x1, _ = Calculate(expreessionLongg)
		//if err != nil {
		//	b.Error(err)
		//}
	}
	b.Log(x1)
}

func BenchmarkCalculateWithVarX(b *testing.B) {
	b.ReportAllocs()
	var x1 float64

	for i := 0; i < b.N; i++ {
		ans, err := CalculateWithVar(exprLong, map[string]float64{"x": 12})
		if err != nil {
			b.Error(err)
		}
		x1 = ans
	}
	b.Log(x1)
}

//
//func BenchmarkCalculateWithManyVarX(b *testing.B) {
//	b.ReportAllocs()
//	var x1 float64
//	for i := 0; i < b.N; i++ {
//		ans, err := CalculateWithVar(exprLong, map[string]float64{"x": 12})
//		if err != nil {
//			b.Error(err)
//		}
//		x1 = ans
//	}
//	b.Log(x1)
//}
//
//func BenchmarkGValWithManyVarX(b *testing.B) {
//	b.ReportAllocs()
//	var x1 float64
//	for i := 0; i < b.N; i++ {
//		ans, err := gval.Evaluate(exprLong, map[string]float64{"x": 12})
//		if err != nil {
//			b.Error(err)
//		}
//		x1, err = utils.ConvertToFloat64(ans)
//		if err != nil {
//			b.Error(err)
//		}
//	}
//	b.Log(x1)
//}

//
//func BenchmarkCalculateWithManyVarXY(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		ans, err := CalculateWithVar(expressionManyVar, map[string]float64{"x": 10, "y": 5})
//		if err != nil {
//			b.Error(err)
//		}
//		ans1 = ans
//	}
//}
//
//func BenchmarkGValWithManyVarXY(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		ans, err := gval.Evaluate(expressionManyVar, map[string]float64{"x": 10, "y": 5})
//		if err != nil {
//			b.Error(err)
//		}
//		ans1, err = utils.ConvertToFloat64(ans)
//		if err != nil {
//			b.Error(err)
//		}
//	}
//}
//
//var negativeNumber = "0--1+2--2--3"
//var negativeNumberAns = ""
//
//func BenchmarkNegativeNumberParser(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		g := NegativeNumberParser(negativeNumber)
//		negativeNumberAns = g
//	}
//}
//
//var replacerAns = ""
//
//func BenchmarkMartinsReplace(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		replacerAns = MartinsReplace("x * x * x", " ", "")
//	}
//}
//
//func BenchmarkReplace(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		replacerAns = strings.ReplaceAll("x * x * x", " ", "")
//	}
//}
//
//var replacerLong = "x                       *                                  x                              *                                            x"
//
//func BenchmarkMartinsReplaceLong(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		replacerAns = MartinsReplace(replacerLong, " ", "")
//	}
//}
//
//func BenchmarkReplaceLong(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		replacerAns = strings.ReplaceAll(replacerLong, " ", "")
//	}
//}
