package evolution

import (
	"testing"
)

var (
	longestXExpr = "((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((0*0)+1)*8)/7)+9)-4)+8)*9)/5)+5)*1)-1)*8)-9)/4)/2)*2)-3)/5)+3)/1)/x)/8)*0)/x)+6)/7)/2)*4)/x)+8)/2)/4)+4)+9)*x)*5)-0)*6)/2)+6)-x)-2)/1)/2)-x)*x)+5)+5)-6)+3)*9)*3)-5)/5)/x)+x)*0)+9)+2)*7)/6)*9)*7)+8)-4)-2)/2)-1)+1)-4)/1)/x)*6)*2)-x)+9)-3)/5)-3)/8)-3)+5)-2)/5)*5)/5)+1)/3)*9)+6)+2)/8)/7)+9)+3)/9)-6)/7)*1)*4)-3)-x)/4)/8)*6)+1)+7)/8)/7)*x)-2)/8)-6)-3)+0)/4)/2)/3)/2)/x)*1)+7)/1)/2)+x)+4)*4)+x)/2)/6)*2)-1)*9)+0)-2)-5)-4)+x)/3)/3)-5)/2)*5)-5)-0)-9)*1)+2)+0)+0)-8)-3)*0)/8)*5)-4)-x)/1)*8)-8)/4)*7)-4)-5)+3)-8)+5)*x)/7)-0)*x)*2)/8)+8)-4)*8)-2)+1)/9)+5)+2)-4)-0)/3)+4)*1)-3)-2)/7)/x)-5)+x)*3)"
)

func TestCalculateRolling(t *testing.T) {
	type args struct {
		bt   BinaryTree
		xVal float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"1", args{NewFromParenthesizedExpression("1"), 3.0}, 1.0, false},
		{"x", args{NewFromParenthesizedExpression("x"), 3.0}, 3.0, false},
		{"x", args{NewFromParenthesizedExpression("x"), -1}, -1.0, false},
		{"(1+1)", args{NewFromParenthesizedExpression("(1+1)"), 3.0}, 2.0, false},
		{"(1+x)", args{NewFromParenthesizedExpression("(1+x)"), 3.0}, 4.0, false},
		{"(x+1)", args{NewFromParenthesizedExpression("(x+1)"), 3.0}, 4.0, false},
		{"(1-x)", args{NewFromParenthesizedExpression("(1-x)"), 4.0}, -3.0, false},
		{"((x-x)+x)", args{NewFromParenthesizedExpression("((x-x)+x)"), 4.0}, 4.0, false},
		{"((x-x)+x)", args{NewFromParenthesizedExpression("((x-x)+x)"), -4.0}, -4, false},
		{"((x-x)+2)", args{NewFromParenthesizedExpression("((x-x)+2)"), -4.0}, 2, false},
		{"(((((((((((((0*0)/5)/6)/9)-4)-4)/2)*6)-7)/6)/2)+9)*3)",
			args{NewFromParenthesizedExpression("(((((((((((((0*0)/5)/6)/9)-4)-4)/2)*6)-7)/6)/2)+9)*3)"), -4.0},
			19.25, false},
		{"((((((((((((((((2/2)/1)/8)-6)+4)/0)-9)-6)-1)*6)*4)/1)/9)-8)-5)*3)",
			args{NewFromParenthesizedExpression("((((((((((((((((2/2)/1)/8)-6)+4)/0)-9)-6)-1)*6)*4)/1)/9)-8)-5)*3)"), -4.0},
			0, true},
		{"((((((((7-7)*2)+6)-8)*8)+4)+x)*3)",
			args{NewFromParenthesizedExpression("((((((((7-7)*2)+6)-8)*8)+4)+x)*3)"), -4.0},
			-48, false},
		{"((((((((((((5/5)/7)/3)*9)-0)+4)/5)/4)*6)-8)+x)*2)",
			args{NewFromParenthesizedExpression("((((((((((((5/5)/7)/3)*9)-0)+4)/5)/4)*6)-8)+x)*2)"), -4.0},
			-21.34285714285714, false},
		{"(((((((((7-7)*x)*6)*2)/7)/5)-9)+x)*3)",
			args{NewFromParenthesizedExpression("(((((((((7-7)*x)*6)*2)/7)/5)-9)+x)*3)"), -4.0},
			-39, false},
		{longestXExpr,
			args{NewFromParenthesizedExpression(longestXExpr), -4.0},
			-26.705782312925173, false},
		{"((0/x)-x)", args{NewFromParenthesizedExpression("((0/x)-x)"), 4.0}, -4.0, false},
		{"(((0-0)/x)-x)", args{NewFromParenthesizedExpression("(((0-0)/x)-x)"), 4.0}, -4.0, false},
		{"(((8/x)+7)/x)", args{NewFromParenthesizedExpression("(((8/x)+7)/x)"), 0.0}, 0.0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.bt, tt.args.xVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
