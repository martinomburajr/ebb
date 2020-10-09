package evolution

import "testing"

func Test_randomTerminalIndex(t *testing.T) {
	type args struct {
		treeLen int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{1}, []int{0}},
		{"", args{5}, []int{1, 3}},
		{"", args{9}, []int{2, 4, 7}},
		{"", args{13}, []int{3, 5, 8}},
		{"", args{17}, []int{4, 6, 9, 12}},
		{"", args{21}, []int{5, 7, 10, 13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomTerminalIndex(tt.args.treeLen)

			found := 0
			for i := 0; i < len(tt.want); i++ {
				if got == tt.want[i] {
					found++
				}
			}

			if found == 0 {
				t.Errorf("randomTerminalIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomNonTerminalIndex(t *testing.T) {
	type args struct {
		treeLen int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{"", args{1}, []int{0}},
		{"", args{5}, []int{2}},
		{"", args{9}, []int{3, 6}},
		{"", args{17}, []int{5, 8, 11, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomNonTerminalIndex(tt.args.treeLen)

			found := 0
			for i := 0; i < len(tt.want); i++ {
				if got == tt.want[i] {
					found++
				}
			}

			if found == 0 {
				t.Errorf("randomTerminalIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
