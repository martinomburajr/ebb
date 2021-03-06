package evolution

//
//var i1 = &Individual{AverageFitness: 1, ID: 0}
//var i2 = &Individual{AverageFitness: 2, ID: 1}
//var iMax = &Individual{AverageFitness: math.MaxInt64, ID: 2}
//var imin1 = &Individual{AverageFitness: -1, ID: 3}
//
//func TestCalcTopIndividual(t *testing.T) {
//	type args struct {
//		individuals       []*Individual
//		fitnessComparator int
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *Individual
//		wantErr bool
//	}{
//		{"nil", args{nil, 0}, nil, true},
//		{"empty", args{[]*Individual{}, 0}, nil, true},
//		{"1", args{[]*Individual{i1}, 0}, i1, false},
//		{"2", args{[]*Individual{i1, i2}, 0}, i1, false},
//		{"3", args{[]*Individual{i1, i2, iMax}, 0}, i1, false},
//		{"4", args{[]*Individual{i1, i2, iMax, imin1}, 0}, imin1, false},
//		{"4", args{[]*Individual{imin1, i2, iMax, i1}, 0}, imin1, false},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetNthPlaceIndividual(tt.args.individuals, tt.args.fitnessComparator)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetNthPlaceIndividual() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !tt.wantErr {
//				if !reflect.DeepEqual(got.AverageFitness, tt.want.AverageFitness) {
//					t.Errorf("GetNthPlaceIndividual() = %v, want %v", got.AverageFitness, tt.want.AverageFitness)
//				}
//			}
//		})
//	}
//}
//
//var g0Pro = &Generation{ID: "g0Pro", Protagonists: []*Individual{i1}}
//var g0Ant = &Generation{ID: "g0Pro", Antagonists: []*Individual{i1}}
//var g0 = &Generation{ID: "g0", Antagonists: []*Individual{i1}, Protagonists: []*Individual{i1}}
//
//var g1Pro = &Generation{ID: "g1Pro", Protagonists: []*Individual{i1, i2}}
//var g1 = &Generation{ID: "g1",
//	Protagonists: []*Individual{i2, i1},
//	Antagonists:  []*Individual{i1, iMax}}
//var g1SortedMoreBetter = &Generation{ID: "g1",
//	Protagonists: []*Individual{i2, i1},
//	Antagonists:  []*Individual{iMax, i1}}
//var g1SortedLessBetter = &Generation{ID: "g1",
//	Protagonists: []*Individual{i1, i2},
//	Antagonists:  []*Individual{i1, iMax}}
//
//var g2Pro = &Generation{ID: "g2Pro", Protagonists: []*Individual{iMax, i1}}
//var g2 = &Generation{ID: "g2Pro",
//	Protagonists: []*Individual{iMax, i1},
//	Antagonists:  []*Individual{i1, iMax}}
//var g2SortedMoreBetter = &Generation{ID: "g2",
//	Protagonists: []*Individual{iMax, i1},
//	Antagonists:  []*Individual{iMax, i1}}
//var g2SortedLessBetter = &Generation{ID: "g2",
//	Protagonists: []*Individual{i1, iMax},
//	Antagonists:  []*Individual{i1, iMax}}
//
//var g4Pro = &Generation{ID: "g4Pro", Protagonists: []*Individual{iMax, i2}}
//var g3Pro = &Generation{ID: "g3Pro", Protagonists: []*Individual{imin1, iMax}}
//
////
//func TestCalcTopIndividualAllGenerations(t *testing.T) {
//	type args struct {
//		generations       []*Generation
//		individualKind    int
//		fitnessComparator int
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    ResultTopIndividuals
//		wantErr bool
//	}{
//		{"nil", args{nil, 1, 0}, ResultTopIndividuals{}, true},
//		{"empty", args{[]*Generation{}, 1, 0}, ResultTopIndividuals{}, true},
//		{"1", args{[]*Generation{g0Pro}, 1, 0}, ResultTopIndividuals{Individual: i1, Generation: g0Pro, Tree: ""}, false},
//		{"2", args{[]*Generation{g0Pro, g1Pro}, 1, 0}, ResultTopIndividuals{Individual: i1, Generation: g1Pro, Tree: ""}, false},
//		{"3", args{[]*Generation{g0Pro, g1Pro, g2Pro}, 1, 0}, ResultTopIndividuals{Individual: i1, Generation: g2Pro, Tree: ""}, false},
//		{"3", args{[]*Generation{g0Pro, g1Pro, g2Pro, g3Pro}, 1, 0}, ResultTopIndividuals{Individual: imin1, Generation: g3Pro, Tree: ""}, false},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := CalcNthPlaceIndividualAllGenerations(tt.args.generations, tt.args.individualKind, tt.args.fitnessComparator)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CalcNthPlaceIndividualAllGenerations() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CalcNthPlaceIndividualAllGenerations() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestSortIndividuals(t *testing.T) {
//	type args struct {
//		individuals       []*Individual
//		fitnessComparator int
//	}
//	tests := []struct {
//		name string
//		args args
//		want []*Individual
//	}{
//		{"", args{[]*Individual{i1, i2, iMax, imin1}, 0}, []*Individual{imin1, i1, i2, iMax}},
//		{"", args{[]*Individual{i1, i2, iMax, i1, imin1}, 0}, []*Individual{imin1, i1, i1, i2, iMax}},
//		{"", args{[]*Individual{i1, i2, iMax, i1, imin1, i1}, 0}, []*Individual{imin1, i1, i1, i1, i2, iMax}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := SortIndividuals(tt.args.individuals, tt.args.fitnessComparator)
//			if len(got) != len(tt.want) {
//				t.Errorf("not same lengthgot = %v, want %v", got, tt.want)
//			}
//			for i := range got {
//				if got[i].AverageFitness != tt.want[i].AverageFitness {
//					t.Errorf("got = %v, want %v", got[i].AverageFitness, tt.want[i].AverageFitness)
//					return
//				}
//			}
//		})
//	}
//}
