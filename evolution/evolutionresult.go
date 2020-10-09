package evolution

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"math"
	"os"
	"strings"
)

const (
	ProgressCountersEvolutionResult = 7
)

// EvolutionResult describes the outcome of an entire run. It contains all the stats from each generation as well as the
// totals.
type EvolutionResult struct {
	HasBeenAnalyzed bool

	TopAntagonistInRun  Individual
	TopProtagonistInRun Individual

	FinalAntagonist  Individual
	FinalProtagonist Individual

	AntagonistMeanFit  float64
	ProtagonistMeanFit float64

	AntagonistMeanStdDev  float64
	ProtagonistMeanStdDev float64

	AntagonistMeanVar  float64
	ProtagonistMeanVar float64

	OldestAntagonist  float64
	OldestProtagonist float64
	AvgAgeAntagonist  float64
	AvgAgeProtagonist float64

	GenerationalResults []GenerationResult
}

// Combine either averages or selects the fittest fields for the elements it's attempting to combine
//func (e *EvolutionResult) Combine(other EvolutionResult) EvolutionResult {
//
//	// Always looks at average fitness
//	topAntagonistInRun := other.TopAntagonistInRun
//	topProtagonistInRun := other.TopProtagonistInRun
//	finAntagonistInRun := other.FinalAntagonist
//	finProtagonistInRun := other.FinalProtagonist
//
//	if e.TopProtagonistInRun.AverageFitness > other.TopProtagonistInRun.AverageFitness {
//		topProtagonistInRun = e.TopProtagonistInRun
//	}
//
//	finalResult := EvolutionResult{
//		HasBeenAnalyzed:       true,
//		TopAntagonistInRun:    topAntagonistInRun,
//		TopProtagonistInRun:   topProtagonistInRun,
//		FinalAntagonist:       finAntagonistInRun,
//		FinalProtagonist:      finProtagonistInRun,
//		AntagonistMeanFit:     0,
//		ProtagonistMeanFit:    0,
//		AntagonistMeanStdDev:  0,
//		ProtagonistMeanStdDev: 0,
//		AntagonistMeanVar:     0,
//		ProtagonistMeanVar:    0,
//		OldestAntagonist:      0,
//		OldestProtagonist:     0,
//		AvgAgeAntagonist:      0,
//		AvgAgeProtagonist:     0,
//		GenerationalResults:   nil,
//	}
//
//	// Always looks at average fitness
//	topAntagonistInRun := Individual{AverageFitness: math.MinInt16}
//	topProtagonistInRun := Individual{AverageFitness: math.MinInt16}
//	finAntagonistInRun := Individual{AverageFitness: math.MinInt16}
//	finProtagonistInRun := Individual{AverageFitness: math.MinInt16}
//
//	oldAnt := float64(math.MinInt16)
//	oldPro := float64(math.MinInt16)
//
//	antAgeSum := 0.0
//	proAgeSum := 0.0
//
//	antagonistStdDevSum := 0.0
//	protagonistStdDevSum := 0.0
//
//	antagonistVarSum := 0.0
//	protagonistVarSum := 0.0
//
//	antagonistAvgSum := 0.0
//	protagonistAvgSum := 0.0
//
//	runCount := len(evolutionResults)
//
//	for i := 0; i < runCount; i++ {
//		runResult := evolutionResults[i]
//
//		currAntagonist := runResult.TopAntagonistInRun
//		currProtagonist := runResult.TopProtagonistInRun
//		currFinAntagonist := runResult.FinalAntagonist
//		currFinProtagonist := runResult.FinalProtagonist
//
//		if currAntagonist.AverageFitness > topAntagonistInRun.AverageFitness {
//			topAntagonistInRun = currAntagonist
//		}
//		if currProtagonist.AverageFitness > topProtagonistInRun.AverageFitness {
//			topProtagonistInRun = currProtagonist
//		}
//
//		if currFinAntagonist.AverageFitness > finAntagonistInRun.AverageFitness {
//			finAntagonistInRun = currFinAntagonist
//		}
//		if currFinProtagonist.AverageFitness > finProtagonistInRun.AverageFitness {
//			finProtagonistInRun = currFinProtagonist
//		}
//
//		if oldAnt < runResult.OldestAntagonist {
//			oldAnt = runResult.OldestAntagonist
//		}
//		if oldPro < runResult.OldestProtagonist {
//			oldPro = runResult.OldestProtagonist
//		}
//
//		antagonistAvgSum += runResult.AntagonistMeanFit
//		protagonistAvgSum += runResult.ProtagonistMeanFit
//
//		antagonistVarSum += runResult.AntagonistMeanVar
//		protagonistVarSum += runResult.ProtagonistMeanVar
//
//		antagonistStdDevSum += runResult.AntagonistMeanStdDev
//		protagonistStdDevSum += runResult.ProtagonistMeanStdDev
//
//		antAgeSum += runResult.AvgAgeAntagonist
//		proAgeSum += runResult.AvgAgeProtagonist
//	}
//
//
//	//Average generations
//	genLength := len(evolutionResults[0].GenerationalResults)
//	genLengthFloat := float64(genLength)
//
//	clonedGenResults := make([]GenerationResult, genLength)
//
//	topAntagonistInRunInner := Individual{AverageFitness: math.MinInt16}
//	topProtagonistInRunInner := Individual{AverageFitness: math.MinInt16}
//	//finAntagonistInRunInner := Individual{AverageFitness: math.MinInt16}
//	//finProtagonistInRunInner := Individual{AverageFitness: math.MinInt16}
//
//	oldAntInner := float64(math.MinInt16)
//	oldProInner := float64(math.MinInt16)
//
//	antAgeSumInner := 0.0
//	proAgeSumInner := 0.0
//
//	antagonistStdDevSumInner := 0.0
//	protagonistStdDevSumInner := 0.0
//
//	antagonistVarSumInner := 0.0
//	protagonistVarSumInner := 0.0
//
//	antagonistAvgSumInner := 0.0
//	protagonistAvgSumInner := 0.0
//
//	for i := 0; i < genLength; i++ {
//		for j := 0; j < len(evolutionResults); j++ {
//			run := evolutionResults[j]
//			gen := run.GenerationalResults[i]
//
//			currAntagonist := gen.BestAntagonist
//			currProtagonist := gen.BestProtagonist
//			//currFinAntagonist := gen.
//			//currFinProtagonist := gen.FinalProtagonist
//
//			if currAntagonist.AverageFitness > topAntagonistInRunInner.AverageFitness {
//				topAntagonistInRunInner = currAntagonist
//			}
//			if currProtagonist.AverageFitness > topProtagonistInRunInner.AverageFitness {
//				topProtagonistInRunInner = currProtagonist
//			}
//
//			//if currFinAntagonistInner.AverageFitness > finAntagonistInRunInner.AverageFitness {
//			//	finAntagonistInRunInner = currFinAntagonist
//			//}
//			//if currFinProtagonist.AverageFitness > finProtagonistInRunInner.AverageFitness {
//			//	finProtagonistInRunInner = currFinProtagonist
//			//}
//
//			if oldAntInner < gen.AntagonistOldAge {
//				oldAntInner = gen.AntagonistOldAge
//			}
//			if oldProInner < gen.ProtagonistOldAge {
//				oldProInner = gen.ProtagonistOldAge
//			}
//
//			antagonistAvgSumInner += gen.AllAntagonistAverageFitness
//			protagonistAvgSumInner += gen.AllProtagonistAverageFitness
//
//			antagonistVarSumInner += gen.AntagonistVarianceOfAvgFitnessValues
//			protagonistVarSumInner += gen.ProtagonistVarianceOfAvgFitnessValues
//
//			antagonistStdDevSumInner += gen.AntagonistStdDevOfAvgFitnessValues
//			protagonistStdDevSumInner += gen.ProtagonistStdDevOfAvgFitnessValues
//
//			antAgeSumInner += gen.AntagonistAvgAge
//			proAgeSumInner += gen.ProtagonistAvgAge
//
//			clonedGenResults[i] = GenerationResult{
//				ID:                  uint32(i),
//				BestAntagonist:      topAntagonistInRunInner,
//				BestProtagonist:     topProtagonistInRunInner,
//				AllAntagonistAverageFitness:   antagonistAvgSumInner / genLengthFloat,
//				AntagonistStdDevOfAvgFitnessValues:    antagonistStdDevSumInner / genLengthFloat,
//				AntagonistVarianceOfAvgFitnessValues:  antagonistVarSumInner / genLengthFloat,
//				AllProtagonistAverageFitness:  protagonistAvgSumInner / genLengthFloat,
//				ProtagonistStdDevOfAvgFitnessValues:   protagonistStdDevSumInner / genLengthFloat,
//				ProtagonistVarianceOfAvgFitnessValues: protagonistVarSumInner / genLengthFloat,
//				AntagonistOldAge:    oldAntInner,
//				ProtagonistOldAge:   oldProInner,
//				AntagonistAvgAge:    antAgeSumInner / genLengthFloat,
//				ProtagonistAvgAge:   proAgeSumInner / genLengthFloat,
//			}
//		}
//	}
//
//	genLengthFloat64 := float64(runCount)
//
//	return TopologicalResult{
//		HasBeenAnalyzed:     true,
//		TopAntagonistInRun:  topAntagonistInRun.Clone(),
//		TopProtagonistInRun: topProtagonistInRun.Clone(),
//		FinalAntagonist:     finAntagonistInRun,
//		FinalProtagonist:    finProtagonistInRun,
//
//		AntagonistMeanFit: antagonistAvgSum / genLengthFloat64,
//		ProtagonistMeanFit: protagonistAvgSum/ genLengthFloat64,
//
//		AntagonistMeanStdDev: antagonistStdDevSum/ genLengthFloat64,
//		ProtagonistMeanStdDev: protagonistStdDevSum/ genLengthFloat64,
//
//		AntagonistMeanVar: antagonistVarSum/ genLengthFloat64,
//		ProtagonistMeanVar: protagonistVarSum/ genLengthFloat64,
//
//		AvgAgeAntagonist:  antAgeSum / genLengthFloat64,
//		AvgAgeProtagonist: antAgeSum / genLengthFloat64,
//
//		OldestAntagonist:  oldAnt,
//		OldestProtagonist: oldPro,
//
//		GenerationalResults: clonedGenResults,
//	}
//}

// AnalyzeResults will look at the GenerationalResults and other data to come up with an EvolutionResult
// After the analysis is complete there should be no handle to any other data i.e no memory leaks of referenced
// data that cannot be cleared by us.
func (e *Engine) AnalyzeResults() EvolutionResult {
	//e.ProgressBar.Incr()

	// Always looks at average fitness
	topAntagonistInRun := Individual{AverageFitness: math.MinInt16}
	topProtagonistInRun := Individual{AverageFitness: math.MinInt16}

	oldAnt := float64(math.MinInt16)
	oldPro := float64(math.MinInt16)

	antAgeSum := 0.0
	proAgeSum := 0.0

	antagonistStdDevSum := 0.0
	protagonistStdDevSum := 0.0

	antagonistVarSum := 0.0
	protagonistVarSum := 0.0

	antagonistAvgSum := 0.0
	protagonistAvgSum := 0.0

	genLength := len(e.GenerationResults)

	clonedGenResults := make([]GenerationResult, genLength)

	for i := 0; i < genLength; i++ {
		gen := e.GenerationResults[i]
		currAntagonist := gen.BestAntagonist
		currProtagonist := gen.BestProtagonist

		if currAntagonist.AverageFitness > topAntagonistInRun.AverageFitness {
			topAntagonistInRun = currAntagonist
		}
		if currProtagonist.AverageFitness > topProtagonistInRun.AverageFitness {
			topProtagonistInRun = currProtagonist
		}
		if oldAnt < gen.AntagonistOldAge {
			oldAnt = gen.AntagonistOldAge
		}
		if oldPro < gen.ProtagonistOldAge {
			oldPro = gen.ProtagonistOldAge
		}

		antagonistAvgSum += gen.AllAntagonistAverageFitness
		protagonistAvgSum += gen.AllProtagonistAverageFitness

		antagonistVarSum += gen.AntagonistVariance
		protagonistVarSum += gen.ProtagonistVariance

		antagonistStdDevSum += gen.AntagonistStdDev
		protagonistStdDevSum += gen.ProtagonistStdDev

		antAgeSum += gen.AntagonistAvgAge
		proAgeSum += gen.ProtagonistAvgAge

		clonedGenResults[i] = gen.Clone()
	}

	finalAntagonist := e.GenerationResults[genLength-1].BestAntagonist
	finalProtagonist := e.GenerationResults[genLength-1].BestProtagonist

	genLengthFloat64 := float64(genLength)

	evolutionResult := EvolutionResult{
		HasBeenAnalyzed:     true,
		TopAntagonistInRun:  topAntagonistInRun.Clone(-1),
		TopProtagonistInRun: topProtagonistInRun.Clone(-1),
		FinalAntagonist:     finalAntagonist.Clone(-1),
		FinalProtagonist:    finalProtagonist.Clone(-1),
		GenerationalResults: clonedGenResults,

		AntagonistMeanFit:  antagonistAvgSum / genLengthFloat64,
		ProtagonistMeanFit: protagonistAvgSum / genLengthFloat64,

		AntagonistMeanStdDev:  antagonistStdDevSum / genLengthFloat64,
		ProtagonistMeanStdDev: protagonistStdDevSum / genLengthFloat64,

		AntagonistMeanVar:  antagonistVarSum / genLengthFloat64,
		ProtagonistMeanVar: protagonistVarSum / genLengthFloat64,

		AvgAgeAntagonist:  antAgeSum / genLengthFloat64,
		AvgAgeProtagonist: antAgeSum / genLengthFloat64,

		OldestAntagonist:  oldAnt,
		OldestProtagonist: oldPro,
	}

	//e.ProgressBar.Incr()

	return evolutionResult
}

type SimulationResult struct {
	KRT TopologicalResult
	RR  TopologicalResult
	SET TopologicalResult
	HoF TopologicalResult
}

func (s *SimulationResult) Summary(params EvolutionParams) string {

	startProgram := params.StartIndividual.ToMathematicalString()

 	sb := strings.Builder{}
	topAKRT := s.KRT.TopAntagonistInRun.ToString()
	topPKRT := s.KRT.TopProtagonistInRun.ToString()

	sb.WriteString(fmt.Sprintf("---------- KRT ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: %s", topAKRT.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPKRT.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.KRT.AntagonistMeanFit))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.KRT.ProtagonistMeanFit))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	// ---------------------------------- RR --------------------------------------
	topARR := s.RR.TopAntagonistInRun.ToString()
	topPRR := s.RR.TopProtagonistInRun.ToString()
	sb.WriteString(fmt.Sprintf("---------- RR ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: \t%s", topARR.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPRR.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.RR.AntagonistMeanFit))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.RR.ProtagonistMeanFit))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	// ---------------------------------- HoF --------------------------------------
	topAHoF := s.HoF.TopAntagonistInRun.ToString()
	topPHoF:= s.HoF.TopProtagonistInRun.ToString()

	sb.WriteString(fmt.Sprintf("---------- HoF ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: \t%s", topAHoF.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPHoF.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.HoF.AntagonistMeanFit))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.HoF.ProtagonistMeanFit))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	// ---------------------------------- SET --------------------------------------
	topASET := s.SET.TopAntagonistInRun.ToString()
	topPSET:= s.SET.TopProtagonistInRun.ToString()

	sb.WriteString(fmt.Sprintf("---------- SET ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: \t%s", topASET.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPSET.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.SET.AntagonistMeanFit))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.SET.ProtagonistMeanFit))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	return sb.String()
}

func (s *SimulationResult) GenerateCSVData(params EvolutionParams) []CSVAvgGenerationsCombinedAcrossRuns {
	AllCSVGens := make([]CSVAvgGenerationsCombinedAcrossRuns, len(s.KRT.GenerationalResults))

	for i := 0; i < len(AllCSVGens); i++ {
		csvGen := CSVAvgGenerationsCombinedAcrossRuns{
			Generation:   i+1,
			SpecEquation: params.StartIndividual.ToMathematicalString(),
			SpecEquationLen: len(params.StartIndividual.ToMathematicalString())/2,
			IVarCount: CountVariable(params.StartIndividual.ToMathematicalString()),
			PolDegree: CountPolDegree(params.StartIndividual.ToMathematicalString()),

			KRTTopAEquation: s.KRT.TopAntagonistInRun.Program.ToMathematicalString(),
			HoFTopAEquation: s.HoF.TopAntagonistInRun.Program.ToMathematicalString(),
			RRTopAEquation:  s.RR.TopAntagonistInRun.Program.ToMathematicalString(),
			SETTopAEquation: s.SET.TopAntagonistInRun.Program.ToMathematicalString(),

			KRTTopPEquation: s.KRT.TopProtagonistInRun.Program.ToMathematicalString(),
			HoFTopPEquation: s.HoF.TopProtagonistInRun.Program.ToMathematicalString(),
			RRTopPEquation:  s.RR.TopProtagonistInRun.Program.ToMathematicalString(),
			SETTopPEquation: s.SET.TopProtagonistInRun.Program.ToMathematicalString(),

			///////////// ########################### AVERAGES #################################
			KRTAntagonistsMean: s.KRT.AntagonistMeanFit,
			HoFAntagonistsMean: s.HoF.AntagonistMeanFit,
			RRAntagonistsMean:  s.RR. AntagonistMeanFit,
			SETAntagonistsMean: s.SET.AntagonistMeanFit,

			KRTProtagonistsMean: s.KRT.ProtagonistMeanFit,
			HoFProtagonistsMean: s.HoF.ProtagonistMeanFit,
			RRProtagonistsMean:  s.RR. ProtagonistMeanFit,
			SETProtagonistsMean: s.SET.ProtagonistMeanFit,

			KRTAntagonistStdDev: s.KRT.AntagonistMeanStdDev,
			HoFAntagonistStdDev: s.HoF.AntagonistMeanStdDev,
			RRAntagonistStdDev:  s.RR. AntagonistMeanStdDev,
			SETAntagonistStdDev: s.SET.AntagonistMeanStdDev,

			KRTProtagonistStdDev: s.KRT.ProtagonistMeanStdDev,
			HoFProtagonistStdDev: s.HoF.ProtagonistMeanStdDev,
			RRProtagonistStdDev:  s.RR. ProtagonistMeanStdDev,
			SETProtagonistStdDev: s.SET.ProtagonistMeanStdDev,

			KRTAntagonistVar: s.KRT.AntagonistMeanVar,
			HoFAntagonistVar: s.HoF.AntagonistMeanVar,
			RRAntagonistVar:  s.RR. AntagonistMeanVar,
			SETAntagonistVar: s.SET.AntagonistMeanVar,

			KRTProtagonistVar: s.KRT.ProtagonistMeanVar,
			HoFProtagonistVar: s.HoF.ProtagonistMeanVar,
			RRProtagonistVar:  s.RR. ProtagonistMeanVar,
			SETProtagonistVar: s.SET.ProtagonistMeanVar,

			KRTAntagonistAverageAge: s.KRT.AvgAgeAntagonist,
			HoFAntagonistAverageAge: s.HoF.AvgAgeAntagonist,
			RRAntagonistAverageAge:  s.RR. AvgAgeAntagonist,
			SETAntagonistAverageAge: s.SET.AvgAgeAntagonist,

			KRTProtagonistAverageAge: s.KRT.AvgAgeProtagonist,
			HoFProtagonistAverageAge: s.HoF.AvgAgeProtagonist,
			RRProtagonistAverageAge:  s.RR. AvgAgeProtagonist,
			SETProtagonistAverageAge: s.SET.AvgAgeProtagonist,

			///////////// ########################### TOP INDIVIDUALS #################################

			KRTTopAntagonistsMean: s.KRT.GenerationalResults[i].BestAntagonist.AverageFitness,
			HoFTopAntagonistsMean: s.HoF.GenerationalResults[i].BestAntagonist.AverageFitness,
			RRTopAntagonistsMean:  s.RR. GenerationalResults[i].BestAntagonist.AverageFitness,
			SETTopAntagonistsMean: s.SET.GenerationalResults[i].BestAntagonist.AverageFitness,

			KRTTopProtagonistsMean: s.KRT.GenerationalResults[i].BestProtagonist.AverageFitness,
			HoFTopProtagonistsMean: s.HoF.GenerationalResults[i].BestProtagonist.AverageFitness,
			RRTopProtagonistsMean:  s.RR. GenerationalResults[i].BestProtagonist.AverageFitness,
			SETTopProtagonistsMean: s.SET.GenerationalResults[i].BestProtagonist.AverageFitness,

			KRTTopAntagonistBestFitness: s.KRT.GenerationalResults[i].BestAntagonist.BestFitness,
			HoFTopAntagonistBestFitness: s.HoF.GenerationalResults[i].BestAntagonist.BestFitness,
			RRTopAntagonistBestFitness:  s.RR. GenerationalResults[i].BestAntagonist.BestFitness,
			SETTopAntagonistBestFitness: s.SET.GenerationalResults[i].BestAntagonist.BestFitness,

			KRTTopProtagonistBestFitness: s.KRT.GenerationalResults[i].BestProtagonist.BestFitness,
			HoFTopProtagonistBestFitness: s.HoF.GenerationalResults[i].BestProtagonist.BestFitness,
			RRTopProtagonistBestFitness:  s.RR. GenerationalResults[i].BestProtagonist.BestFitness,
			SETTopProtagonistBestFitness: s.SET.GenerationalResults[i].BestProtagonist.BestFitness,

			KRTTopAntagonistStdDev: s.KRT.GenerationalResults[i].BestAntagonist.FitnessStdDev,
			HoFTopAntagonistStdDev: s.HoF.GenerationalResults[i].BestAntagonist.FitnessStdDev,
			RRTopAntagonistStdDev:  s.RR. GenerationalResults[i].BestAntagonist.FitnessStdDev,
			SETTopAntagonistStdDev: s.SET.GenerationalResults[i].BestAntagonist.FitnessStdDev,

			KRTTopProtagonistStdDev: s.KRT.GenerationalResults[i].BestProtagonist.FitnessStdDev,
			HoFTopProtagonistStdDev: s.HoF.GenerationalResults[i].BestProtagonist.FitnessStdDev,
			RRTopProtagonistStdDev:  s.RR. GenerationalResults[i].BestProtagonist.FitnessStdDev,
			SETTopProtagonistStdDev: s.SET.GenerationalResults[i].BestProtagonist.FitnessStdDev,

			KRTTopAntagonistVar: s.KRT.GenerationalResults[i].BestAntagonist.FitnessVariance,
			HoFTopAntagonistVar: s.HoF.GenerationalResults[i].BestAntagonist.FitnessVariance,
			RRTopAntagonistVar:  s.RR. GenerationalResults[i].BestAntagonist.FitnessVariance,
			SETTopAntagonistVar: s.SET.GenerationalResults[i].BestAntagonist.FitnessVariance,

			KRTTopProtagonistVar: s.KRT.GenerationalResults[i].BestProtagonist.FitnessVariance,
			HoFTopProtagonistVar: s.HoF.GenerationalResults[i].BestProtagonist.FitnessVariance,
			RRTopProtagonistVar:  s.RR. GenerationalResults[i].BestProtagonist.FitnessVariance,
			SETTopProtagonistVar: s.SET.GenerationalResults[i].BestProtagonist.FitnessVariance,

			KRTTopAntagonistAverageAge: float64(s.KRT.GenerationalResults[i].BestAntagonist.Age),
			HoFTopAntagonistAverageAge: float64(s.HoF.GenerationalResults[i].BestAntagonist.Age),
			RRTopAntagonistAverageAge:  float64(s.RR. GenerationalResults[i].BestAntagonist.Age),
			SETTopAntagonistAverageAge: float64(s.SET.GenerationalResults[i].BestAntagonist.Age),

			KRTTopProtagonistAverageAge: float64(s.KRT.GenerationalResults[i].BestProtagonist.Age),
			HoFTopProtagonistAverageAge: float64(s.HoF.GenerationalResults[i].BestProtagonist.Age),
			RRTopProtagonistAverageAge:  float64(s.RR. GenerationalResults[i].BestProtagonist.Age),
			SETTopProtagonistAverageAge: float64(s.SET.GenerationalResults[i].BestProtagonist.Age),

			KRTTopAntagonistBirthGen: float64(s.KRT.GenerationalResults[i].BestAntagonist.BirthGen),
			HoFTopAntagonistBirthGen: float64(s.HoF.GenerationalResults[i].BestAntagonist.BirthGen),
			RRTopAntagonistBirthGen:  float64(s.RR. GenerationalResults[i].BestAntagonist.BirthGen),
			SETTopAntagonistBirthGen: float64(s.SET.GenerationalResults[i].BestAntagonist.BirthGen),

			KRTTopProtagonistBirthGen: float64(s.KRT.GenerationalResults[i].BestProtagonist.BirthGen),
			HoFTopProtagonistBirthGen: float64(s.HoF.GenerationalResults[i].BestProtagonist.BirthGen),
			RRTopProtagonistBirthGen:  float64(s.RR. GenerationalResults[i].BestProtagonist.BirthGen),
			SETTopProtagonistBirthGen: float64(s.SET.GenerationalResults[i].BestProtagonist.BirthGen),

			KRTTopAntagonistStrategy: StrategiesToString(s.KRT.GenerationalResults[i].BestAntagonist),
			HoFTopAntagonistStrategy: StrategiesToString(s.HoF.GenerationalResults[i].BestAntagonist),
			RRTopAntagonistStrategy:  StrategiesToString(s.RR. GenerationalResults[i].BestAntagonist),
			SETTopAntagonistStrategy: StrategiesToString(s.SET.GenerationalResults[i].BestAntagonist),

			KRTTopProtagonistStrategy: StrategiesToString(s.KRT.GenerationalResults[i].BestProtagonist),
			HoFTopProtagonistStrategy: StrategiesToString(s.HoF.GenerationalResults[i].BestProtagonist),
			RRTopProtagonistStrategy:  StrategiesToString(s.RR. GenerationalResults[i].BestProtagonist),
			SETTopProtagonistStrategy: StrategiesToString(s.SET.GenerationalResults[i].BestProtagonist),

			KRTTopAntagonistDomStrategy: DominantStrategy(s.KRT.GenerationalResults[i].BestAntagonist),
			HoFTopAntagonistDomStrategy: DominantStrategy(s.HoF.GenerationalResults[i].BestAntagonist),
			RRTopAntagonistDomStrategy:  DominantStrategy(s.RR. GenerationalResults[i].BestAntagonist),
			SETTopAntagonistDomStrategy: DominantStrategy(s.SET.GenerationalResults[i].BestAntagonist),

			KRTTopProtagonistDomStrategy: DominantStrategy(s.KRT.GenerationalResults[i].BestProtagonist),
			HoFTopProtagonistDomStrategy: DominantStrategy(s.HoF.GenerationalResults[i].BestProtagonist),
			RRTopProtagonistDomStrategy:  DominantStrategy(s.RR. GenerationalResults[i].BestProtagonist),
			SETTopProtagonistDomStrategy: DominantStrategy(s.SET.GenerationalResults[i].BestProtagonist),
		}

		AllCSVGens[i] = csvGen
	}

	return AllCSVGens
}

func CountVariable(mathematicalString string) int {
	count := 0

	for i :=0 ; i < len(mathematicalString); i++ {
		if mathematicalString[i] == 'x' {
			count++
		}
	}

	return count
}

func CountPolDegree(mathematicalString string) int {
	count := 0

	for i :=0 ; i < len(mathematicalString)-1; i++ {
		if mathematicalString[i] == '*' && mathematicalString[i+1] == 'x' {
			count++
		}
		if mathematicalString[i] == '/' && mathematicalString[i+1] == 'x' {
			count--
		}
	}

	return count
}

// WriteCSV takes in the CSVAvg... object and writes it to the path. It DOES NOT CREATE the path if it does not exist.
func (s *SimulationResult) WriteCSV(file *os.File, data []CSVAvgGenerationsCombinedAcrossRuns) error {
	defer file.Close()

	err := gocsv.MarshalFile(data, file)
	if err != nil {
		return err
	}

	return nil
}

func NewSimulationResult(topologicalResults []TopologicalResult) SimulationResult {
	return SimulationResult{}
}

// TopologicalResult refers to the combination of multiple evolutionary runs (results) into a single result that should be further
// combined with other topological results to return a simulation result which is the ultimate and final point of analysis
type TopologicalResult struct {
	HasBeenAnalyzed bool

	Topology string

	TopAntagonistInRun  Individual
	TopProtagonistInRun Individual

	FinalAntagonist  Individual
	FinalProtagonist Individual

	AntagonistMeanFit  float64
	ProtagonistMeanFit float64

	AntagonistMeanStdDev  float64
	ProtagonistMeanStdDev float64

	AntagonistMeanVar  float64
	ProtagonistMeanVar float64

	OldestAntagonist  float64
	OldestProtagonist float64
	AvgAgeAntagonist  float64
	AvgAgeProtagonist float64

	GenerationalResults []GenerationResult
}

// CombineEvolutionResults combines all evolution results (runs) into a single evolution result (runs). This is done
// by averaging values that can be averaged, and ranking those that can be ranked. The output is a total summary of all
// runs
func NewTopologicalResults(topology string, evolutionResults []EvolutionResult) TopologicalResult {
	// Always looks at average fitness
	topAntagonistInRun := Individual{AverageFitness: math.MinInt16}
	topProtagonistInRun := Individual{AverageFitness: math.MinInt16}
	finAntagonistInRun := Individual{AverageFitness: math.MinInt16}
	finProtagonistInRun := Individual{AverageFitness: math.MinInt16}

	oldAnt := float64(math.MinInt16)
	oldPro := float64(math.MinInt16)

	antAgeSum := 0.0
	proAgeSum := 0.0

	antagonistStdDevSum := 0.0
	protagonistStdDevSum := 0.0

	antagonistVarSum := 0.0
	protagonistVarSum := 0.0

	antagonistAvgSum := 0.0
	protagonistAvgSum := 0.0

	runCount := len(evolutionResults)

	for i := 0; i < runCount; i++ {
		runResult := evolutionResults[i]

		currAntagonist := runResult.TopAntagonistInRun
		currProtagonist := runResult.TopProtagonistInRun
		currFinAntagonist := runResult.FinalAntagonist
		currFinProtagonist := runResult.FinalProtagonist

		if currAntagonist.AverageFitness > topAntagonistInRun.AverageFitness {
			topAntagonistInRun = currAntagonist
		}
		if currProtagonist.AverageFitness > topProtagonistInRun.AverageFitness {
			topProtagonistInRun = currProtagonist
		}

		if currFinAntagonist.AverageFitness > finAntagonistInRun.AverageFitness {
			finAntagonistInRun = currFinAntagonist
		}
		if currFinProtagonist.AverageFitness > finProtagonistInRun.AverageFitness {
			finProtagonistInRun = currFinProtagonist
		}

		if oldAnt < runResult.OldestAntagonist {
			oldAnt = runResult.OldestAntagonist
		}
		if oldPro < runResult.OldestProtagonist {
			oldPro = runResult.OldestProtagonist
		}

		antagonistAvgSum += runResult.AntagonistMeanFit
		protagonistAvgSum += runResult.ProtagonistMeanFit

		antagonistVarSum += runResult.AntagonistMeanVar
		protagonistVarSum += runResult.ProtagonistMeanVar

		antagonistStdDevSum += runResult.AntagonistMeanStdDev
		protagonistStdDevSum += runResult.ProtagonistMeanStdDev

		antAgeSum += runResult.AvgAgeAntagonist
		proAgeSum += runResult.AvgAgeProtagonist
	}

	//Average generations
	genLength := len(evolutionResults[0].GenerationalResults)
	genLengthFloat := float64(genLength)

	clonedGenResults := make([]GenerationResult, genLength)

	topAntagonistInRunInner := Individual{AverageFitness: math.MinInt16}
	topProtagonistInRunInner := Individual{AverageFitness: math.MinInt16}
	//finAntagonistInRunInner := Individual{AverageFitness: math.MinInt16}
	//finProtagonistInRunInner := Individual{AverageFitness: math.MinInt16}

	oldAntInner := float64(math.MinInt16)
	oldProInner := float64(math.MinInt16)

	antAgeSumInner := 0.0
	proAgeSumInner := 0.0

	antagonistStdDevSumInner := 0.0
	protagonistStdDevSumInner := 0.0

	antagonistVarSumInner := 0.0
	protagonistVarSumInner := 0.0

	antagonistAvgSumInner := 0.0
	protagonistAvgSumInner := 0.0

	for i := 0; i < genLength; i++ {
		for j := 0; j < len(evolutionResults); j++ {
			run := evolutionResults[j]
			gen := run.GenerationalResults[i]

			currAntagonist := gen.BestAntagonist
			currProtagonist := gen.BestProtagonist
			//currFinAntagonist := gen.
			//currFinProtagonist := gen.FinalProtagonist

			if currAntagonist.AverageFitness > topAntagonistInRunInner.AverageFitness {
				topAntagonistInRunInner = currAntagonist
			}
			if currProtagonist.AverageFitness > topProtagonistInRunInner.AverageFitness {
				topProtagonistInRunInner = currProtagonist
			}

			//if currFinAntagonistInner.AverageFitness > finAntagonistInRunInner.AverageFitness {
			//	finAntagonistInRunInner = currFinAntagonist
			//}
			//if currFinProtagonist.AverageFitness > finProtagonistInRunInner.AverageFitness {
			//	finProtagonistInRunInner = currFinProtagonist
			//}

			if oldAntInner < gen.AntagonistOldAge {
				oldAntInner = gen.AntagonistOldAge
			}
			if oldProInner < gen.ProtagonistOldAge {
				oldProInner = gen.ProtagonistOldAge
			}

			antagonistAvgSumInner += gen.AllAntagonistAverageFitness
			protagonistAvgSumInner += gen.AllProtagonistAverageFitness

			antagonistVarSumInner += gen.AntagonistVariance
			protagonistVarSumInner += gen.ProtagonistVariance

			antagonistStdDevSumInner += gen.AntagonistStdDev
			protagonistStdDevSumInner += gen.ProtagonistStdDev

			antAgeSumInner += gen.AntagonistAvgAge
			proAgeSumInner += gen.ProtagonistAvgAge

			clonedGenResults[i] = GenerationResult{
				ID:                           uint32(i),
				BestAntagonist:               topAntagonistInRunInner,
				BestProtagonist:              topProtagonistInRunInner,
				AllAntagonistAverageFitness:  antagonistAvgSumInner / genLengthFloat,
				AntagonistStdDev:             antagonistStdDevSumInner / genLengthFloat,
				AntagonistVariance:           antagonistVarSumInner / genLengthFloat,
				AllProtagonistAverageFitness: protagonistAvgSumInner / genLengthFloat,
				ProtagonistStdDev:            protagonistStdDevSumInner / genLengthFloat,
				ProtagonistVariance:          protagonistVarSumInner / genLengthFloat,
				AntagonistOldAge:             oldAntInner,
				ProtagonistOldAge:            oldProInner,
				AntagonistAvgAge:             antAgeSumInner / genLengthFloat,
				ProtagonistAvgAge:            proAgeSumInner / genLengthFloat,
			}
		}
	}

	genLengthFloat64 := float64(runCount)

	return TopologicalResult{
		HasBeenAnalyzed:     true,
		TopAntagonistInRun:  topAntagonistInRun.Clone(-1),
		TopProtagonistInRun: topProtagonistInRun.Clone(-1),
		FinalAntagonist:     finAntagonistInRun,
		FinalProtagonist:    finProtagonistInRun,

		AntagonistMeanFit:  antagonistAvgSum / genLengthFloat64,
		ProtagonistMeanFit: protagonistAvgSum / genLengthFloat64,

		AntagonistMeanStdDev:  antagonistStdDevSum / genLengthFloat64,
		ProtagonistMeanStdDev: protagonistStdDevSum / genLengthFloat64,

		AntagonistMeanVar:  antagonistVarSum / genLengthFloat64,
		ProtagonistMeanVar: protagonistVarSum / genLengthFloat64,

		AvgAgeAntagonist:  antAgeSum / genLengthFloat64,
		AvgAgeProtagonist: antAgeSum / genLengthFloat64,

		OldestAntagonist:  oldAnt,
		OldestProtagonist: oldPro,

		GenerationalResults: clonedGenResults,
	}
}

func CalculateGenerationSize(params EvolutionParams) int {
	genCount := 0
	if params.MaxGenerations > MinAllowableGenerationsToTerminate {
		if params.FinalGeneration > 0 {
			genCount = params.FinalGeneration
		} else {
			genCount = params.MaxGenerations
		}
	} else {
		genCount = params.GenerationsCount
	}
	return genCount
}
