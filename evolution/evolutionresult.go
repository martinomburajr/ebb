package evolution

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"math"
	"os"
	"strings"
	"time"
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

	AvgAgeAntagonist  float64
	AvgAgeProtagonist float64

	AvgBirthGenAntagonist  float64
	AvgBirthGenProtagonist float64

	GenerationalResults []GenerationResult
}

// Combine either averages or selects the fittest fields for the elements it's attempting to combine
//func (e *EvolutionResult) Combine(other EvolutionResult) EvolutionResult {
//
//	// Always looks at average fitness
//	topAntagonistInRun := other.TopAntagonistOfAllRuns
//	topProtagonistInRun := other.TopProtagonistOfAllRuns
//	finAntagonistInRun := other.FinalBestAntagonistOfAllRuns
//	finProtagonistInRun := other.FinalBestProtagonistOfAllRuns
//
//	if e.TopProtagonistOfAllRuns.AverageFitness > other.TopProtagonistOfAllRuns.AverageFitness {
//		topProtagonistInRun = e.TopProtagonistOfAllRuns
//	}
//
//	finalResult := EvolutionResult{
//		HasBeenAnalyzed:       true,
//		TopAntagonistOfAllRuns:    topAntagonistInRun,
//		TopProtagonistOfAllRuns:   topProtagonistInRun,
//		FinalBestAntagonistOfAllRuns:       finAntagonistInRun,
//		FinalBestProtagonistOfAllRuns:      finProtagonistInRun,
//		AntagonistMeanFitAcrossAllRuns:     0,
//		ProtagonistMeanFitAcrossAllRuns:    0,
//		AntagonistMeanStdDev:  0,
//		ProtagonistMeanStdDev: 0,
//		AntagonistMeanVar:     0,
//		ProtagonistMeanVar:    0,
//		OldestAntagonist:      0,
//		OldestProtagonist:     0,
//		AvgAgeAntagonist:      0,
//		AvgAgeProtagonist:     0,
//		TopGenerationalResultsAcrossAllRuns:   nil,
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
//		currAntagonist := runResult.TopAntagonistOfAllRuns
//		currProtagonist := runResult.TopProtagonistOfAllRuns
//		currFinAntagonist := runResult.FinalBestAntagonistOfAllRuns
//		currFinProtagonist := runResult.FinalBestProtagonistOfAllRuns
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
//		antagonistAvgSum += runResult.AntagonistMeanFitAcrossAllRuns
//		protagonistAvgSum += runResult.ProtagonistMeanFitAcrossAllRuns
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
//	genLength := len(evolutionResults[0].TopGenerationalResultsAcrossAllRuns)
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
//			gen := run.TopGenerationalResultsAcrossAllRuns[i]
//
//			currAntagonist := gen.BestAntagonist
//			currProtagonist := gen.BestProtagonist
//			//currFinAntagonist := gen.
//			//currFinProtagonist := gen.FinalBestProtagonistOfAllRuns
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
//		TopAntagonistOfAllRuns:  topAntagonistInRun.Clone(),
//		TopProtagonistOfAllRuns: topProtagonistInRun.Clone(),
//		FinalBestAntagonistOfAllRuns:     finAntagonistInRun,
//		FinalBestProtagonistOfAllRuns:    finProtagonistInRun,
//
//		AntagonistMeanFitAcrossAllRuns: antagonistAvgSum / genLengthFloat64,
//		ProtagonistMeanFitAcrossAllRuns: protagonistAvgSum/ genLengthFloat64,
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
//		TopGenerationalResultsAcrossAllRuns: clonedGenResults,
//	}
//}

// AnalyzeResults will look at the TopGenerationalResultsAcrossAllRuns and other data to come up with an EvolutionResult
// After the analysis is complete there should be no handle to any other data i.e no memory leaks of referenced
// data that cannot be cleared by us.
func (e *Engine) AnalyzeResults() EvolutionResult {
	//e.ProgressBar.Incr()

	// Always looks at average fitness
	topAntagonistInRun := Individual{AverageFitness: math.MinInt16}
	topProtagonistInRun := Individual{AverageFitness: math.MinInt16}

	antAgeSum := 0.0
	proAgeSum := 0.0

	antBirthGenSum := 0.0
	proBirthGenSum := 0.0

	antagonistStdDevSum := 0.0
	protagonistStdDevSum := 0.0

	antagonistVarSum := 0.0
	protagonistVarSum := 0.0

	antagonistAvgSum := 0.0
	protagonistAvgSum := 0.0

	genLength := len(e.GenerationResults)

	clonedGenResults := make([]GenerationResult, genLength)

	for i := 0; i < genLength; i++ {
		genResult := e.GenerationResults[i]
		currAntagonist := genResult.BestAntagonist
		currProtagonist := genResult.BestProtagonist

		if currAntagonist.AverageFitness > topAntagonistInRun.AverageFitness {
			topAntagonistInRun = currAntagonist
		}
		if currProtagonist.AverageFitness > topProtagonistInRun.AverageFitness {
			topProtagonistInRun = currProtagonist
		}

		antagonistAvgSum += genResult.AllAntagonistAverageFitness
		protagonistAvgSum += genResult.AllProtagonistAverageFitness

		antagonistVarSum += genResult.AntagonistVariance
		protagonistVarSum += genResult.ProtagonistVariance

		antagonistStdDevSum += genResult.AntagonistStdDev
		protagonistStdDevSum += genResult.ProtagonistStdDev

		antAgeSum += genResult.AntagonistAvgAge
		proAgeSum += genResult.ProtagonistAvgAge

		antBirthGenSum += genResult.AntagonistAvgBirthGen
		proBirthGenSum += genResult.ProtagonistAvgBirthGen

		clonedGenResults[i] = genResult.Clone()
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

		AvgBirthGenAntagonist:  antBirthGenSum / genLengthFloat64,
		AvgBirthGenProtagonist: proBirthGenSum / genLengthFloat64,
	}

	return evolutionResult
}

type SimulationResult struct {
	KRT TopologicalResult
	RR  TopologicalResult
	SET TopologicalResult
	HoF TopologicalResult

	KRTDuration time.Duration
	SETDuration time.Duration
	HoFDuration time.Duration
	KDuration   time.Duration
}

func (s *SimulationResult) Summary(params EvolutionParams) string {

	startProgram := params.StartIndividual.ToMathematicalString()

	sb := strings.Builder{}
	topAKRT := s.KRT.TopAntagonistOfAllRuns.ToString()
	topPKRT := s.KRT.TopProtagonistOfAllRuns.ToString()

	sb.WriteString(fmt.Sprintf("---------- KRT ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: %s", topAKRT.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPKRT.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.KRT.AntagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.KRT.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	// ---------------------------------- RR --------------------------------------
	topARR := s.RR.TopAntagonistOfAllRuns.ToString()
	topPRR := s.RR.TopProtagonistOfAllRuns.ToString()
	sb.WriteString(fmt.Sprintf("---------- RR ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: \t%s", topARR.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPRR.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.RR.AntagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.RR.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	// ---------------------------------- HoF --------------------------------------
	topAHoF := s.HoF.TopAntagonistOfAllRuns.ToString()
	topPHoF := s.HoF.TopProtagonistOfAllRuns.ToString()

	sb.WriteString(fmt.Sprintf("---------- HoF ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: \t%s", topAHoF.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPHoF.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.HoF.AntagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.HoF.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	// ---------------------------------- SET --------------------------------------
	topASET := s.SET.TopAntagonistOfAllRuns.ToString()
	topPSET := s.SET.TopProtagonistOfAllRuns.ToString()

	sb.WriteString(fmt.Sprintf("---------- SET ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Start Program: %s", startProgram))
	sb.WriteRune('\n')
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Antagonist: \t%s", topASET.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Best Protagonist: \t%s", topPSET.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("Antagonists Mean Fit: %.2f", s.SET.AntagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("Protagonists Mean Fit: %.2f", s.SET.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	sb.WriteRune('\n')
	sb.WriteRune('\n')

	return sb.String()
}

func (s *SimulationResult) SimpleSummary(params EvolutionParams) string {

	sb := strings.Builder{}
	topAKRT := s.KRT.TopAntagonistOfAllRuns.ToSimpleString()
	topPKRT := s.KRT.TopProtagonistOfAllRuns.ToSimpleString()

	sb.WriteString(fmt.Sprintf("---------- KRT ------------"))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BAnt: %s", topAKRT.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BPro: %s", topPKRT.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("AllAMeanFit: %.2f \t||\t", s.KRT.AntagonistMeanFitAcrossAllRuns))
	sb.WriteString(fmt.Sprintf("AllPMeanFit: %.2f", s.KRT.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	// ---------------------------------- RR --------------------------------------
	topARR := s.RR.TopAntagonistOfAllRuns.ToSimpleString()
	topPRR := s.RR.TopProtagonistOfAllRuns.ToSimpleString()
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BAnt: %s", topARR.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BPro: %s", topPRR.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("AllAMeanFit: %.2f \t||\t", s.RR.AntagonistMeanFitAcrossAllRuns))
	sb.WriteString(fmt.Sprintf("AllPMeanFit: %.2f", s.RR.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	// ---------------------------------- HoF --------------------------------------
	topAHoF := s.HoF.TopAntagonistOfAllRuns.ToSimpleString()
	topPHoF := s.HoF.TopProtagonistOfAllRuns.ToSimpleString()
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BAnt: %s", topAHoF.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BPro: %s", topPHoF.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("AllAMeanFit: %.2f \t||\t", s.HoF.AntagonistMeanFitAcrossAllRuns))
	sb.WriteString(fmt.Sprintf("AllPMeanFit: %.2f", s.HoF.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	// ---------------------------------- SET --------------------------------------
	topASET := s.SET.TopAntagonistOfAllRuns.ToSimpleString()
	topPSET := s.SET.TopProtagonistOfAllRuns.ToSimpleString()
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BAnt: %s", topASET.String()))
	sb.WriteRune('\n')
	sb.WriteString(fmt.Sprintf("BPro: %s", topPSET.String()))
	sb.WriteRune('\n')

	sb.WriteString(fmt.Sprintf("AllAMeanFit: %.2f \t||\t", s.SET.AntagonistMeanFitAcrossAllRuns))
	sb.WriteString(fmt.Sprintf("AllPMeanFit: %.2f", s.SET.ProtagonistMeanFitAcrossAllRuns))
	sb.WriteRune('\n')

	return sb.String()
}

func (s *SimulationResult) GenerateCSVData(params EvolutionParams) ([]CSVAvgGenerationsCombinedAcrossRuns, []CSVStrat) {
	AllCSVGens := make([]CSVAvgGenerationsCombinedAcrossRuns, len(s.KRT.TopGenerationalResultsAcrossAllRuns))
	AllCSVStats := make([]CSVStrat, 0)

	for i := 0; i < len(AllCSVGens); i++ {
		csvGen := CSVAvgGenerationsCombinedAcrossRuns{
			Generation:      i + 1,
			SpecEquation:    params.StartIndividual.ToMathematicalString(),
			SpecEquationLen: len(params.StartIndividual.ToMathematicalString()) / 2,
			IVarCount:       CountVariable(params.StartIndividual.ToMathematicalString()),
			PolDegree:       CountPolDegree(params.StartIndividual.ToMathematicalString()),

			///////////// ########################### GENERATIONAL #################################
			///////////// ########################### GENERATIONAL #################################
			///////////// ########################### GENERATIONAL #################################
			///////////// ########################### GENERATIONAL #################################
			///////////// ########################### GENERATIONAL #################################
			///////////// ########################### GENERATIONAL #################################

			KRTAntagonistAverageAgeInGen: s.KRT.AvgAgeAntagonist,
			HoFAntagonistAverageAgeInGen: s.HoF.AvgAgeAntagonist,
			RRAntagonistAverageAgeInGen:  s.RR.AvgAgeAntagonist,
			SETAntagonistAverageAgeInGen: s.SET.AvgAgeAntagonist,

			KRTProtagonistAverageAgeInGen: s.KRT.AvgAgeProtagonist,
			HoFProtagonistAverageAgeInGen: s.HoF.AvgAgeProtagonist,
			RRProtagonistAverageAgeInGen:  s.RR.AvgAgeProtagonist,
			SETProtagonistAverageAgeInGen: s.SET.AvgAgeProtagonist,

			KRTAntagonistsMeanInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].AllAntagonistAverageFitness,
			HoFAntagonistsMeanInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].AllAntagonistAverageFitness,
			RRAntagonistsMeanInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].AllAntagonistAverageFitness,
			SETAntagonistsMeanInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].AllAntagonistAverageFitness,

			KRTProtagonistsMeanInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].AllProtagonistAverageFitness,
			HoFProtagonistsMeanInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].AllProtagonistAverageFitness,
			RRProtagonistsMeanInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].AllProtagonistAverageFitness,
			SETProtagonistsMeanInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].AllProtagonistAverageFitness,

			KRTTopAntagonistsMeanInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.AverageFitness,
			HoFTopAntagonistsMeanInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.AverageFitness,
			RRTopAntagonistsMeanInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.AverageFitness,
			SETTopAntagonistsMeanInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.AverageFitness,

			KRTTopProtagonistsMeanInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.AverageFitness,
			HoFTopProtagonistsMeanInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.AverageFitness,
			RRTopProtagonistsMeanInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.AverageFitness,
			SETTopProtagonistsMeanInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.AverageFitness,

			KRTTopAntagonistBestFitnessInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BestFitness,
			HoFTopAntagonistBestFitnessInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BestFitness,
			RRTopAntagonistBestFitnessInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BestFitness,
			SETTopAntagonistBestFitnessInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BestFitness,

			KRTTopProtagonistBestFitnessInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BestFitness,
			HoFTopProtagonistBestFitnessInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BestFitness,
			RRTopProtagonistBestFitnessInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BestFitness,
			SETTopProtagonistBestFitnessInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BestFitness,

			KRTTopAntagonistStdDevInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessStdDev,
			HoFTopAntagonistStdDevInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessStdDev,
			RRTopAntagonistStdDevInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessStdDev,
			SETTopAntagonistStdDevInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessStdDev,

			KRTTopProtagonistStdDevInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessStdDev,
			HoFTopProtagonistStdDevInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessStdDev,
			RRTopProtagonistStdDevInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessStdDev,
			SETTopProtagonistStdDevInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessStdDev,

			KRTTopAntagonistVarInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessVariance,
			HoFTopAntagonistVarInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessVariance,
			RRTopAntagonistVarInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessVariance,
			SETTopAntagonistVarInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.FitnessVariance,

			KRTTopProtagonistVarInGen: s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessVariance,
			HoFTopProtagonistVarInGen: s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessVariance,
			RRTopProtagonistVarInGen:  s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessVariance,
			SETTopProtagonistVarInGen: s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.FitnessVariance,

			KRTTopAntagonistAverageAgeInGen: float64(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.Age),
			HoFTopAntagonistAverageAgeInGen: float64(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.Age),
			RRTopAntagonistAverageAgeInGen:  float64(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.Age),
			SETTopAntagonistAverageAgeInGen: float64(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.Age),

			KRTTopProtagonistAverageAgeInGen: float64(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.Age),
			HoFTopProtagonistAverageAgeInGen: float64(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.Age),
			RRTopProtagonistAverageAgeInGen:  float64(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.Age),
			SETTopProtagonistAverageAgeInGen: float64(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.Age),

			KRTTopAntagonistBirthGenInGen: float64(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BirthGen),
			HoFTopAntagonistBirthGenInGen: float64(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BirthGen),
			RRTopAntagonistBirthGenInGen:  float64(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BirthGen),
			SETTopAntagonistBirthGenInGen: float64(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BirthGen),

			KRTTopProtagonistBirthGenInGen: float64(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BirthGen),
			HoFTopProtagonistBirthGenInGen: float64(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BirthGen),
			RRTopProtagonistBirthGenInGen:  float64(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BirthGen),
			SETTopProtagonistBirthGenInGen: float64(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BirthGen),

			KRTTopAntagonistNoCompetitionsInSim: s.KRT.TopAntagonistOfAllRuns.NoOfCompetitions,
			HoFTopAntagonistNoCompetitionsInSim: s.HoF.TopAntagonistOfAllRuns.NoOfCompetitions,
			RRTTopAntagonistNoCompetitionsInSim: s.RR.TopAntagonistOfAllRuns.NoOfCompetitions,
			SETTopAntagonistNoCompetitionsInSim: s.SET.TopAntagonistOfAllRuns.NoOfCompetitions,

			KRTTopProtagonistNoCompetitionsInSim: s.KRT.TopProtagonistOfAllRuns.NoOfCompetitions,
			HoFTopProtagonistNoCompetitionsInSim: s.HoF.TopProtagonistOfAllRuns.NoOfCompetitions,
			RRTTopProtagonistNoCompetitionsInSim: s.RR.TopProtagonistOfAllRuns.NoOfCompetitions,
			SETTopProtagonistNoCompetitionsInSim: s.SET.TopProtagonistOfAllRuns.NoOfCompetitions,

			KRTTopAntagonistStrategyInGen: StrategiesToString(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),
			HoFTopAntagonistStrategyInGen: StrategiesToString(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),
			RRTopAntagonistStrategyInGen:  StrategiesToString(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),
			SETTopAntagonistStrategyInGen: StrategiesToString(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),

			KRTTopProtagonistStrategyInGen: StrategiesToString(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),
			HoFTopProtagonistStrategyInGen: StrategiesToString(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),
			RRTopProtagonistStrategyInGen:  StrategiesToString(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),
			SETTopProtagonistStrategyInGen: StrategiesToString(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),

			KRTTopAntagonistDomStrategyInGen: DominantStrategy(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),
			HoFTopAntagonistDomStrategyInGen: DominantStrategy(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),
			RRTopAntagonistDomStrategyInGen:  DominantStrategy(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),
			SETTopAntagonistDomStrategyInGen: DominantStrategy(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist),

			KRTTopProtagonistDomStrategyInGen: DominantStrategy(s.KRT.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),
			HoFTopProtagonistDomStrategyInGen: DominantStrategy(s.HoF.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),
			RRTopProtagonistDomStrategyInGen:  DominantStrategy(s.RR.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),
			SETTopProtagonistDomStrategyInGen: DominantStrategy(s.SET.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist),
			

			////////////////////////////////////// TOP INDIVIDUAL /////////////////////////////////////////////////
			////////////////////////////////////// TOP INDIVIDUAL /////////////////////////////////////////////////
			////////////////////////////////////// TOP INDIVIDUAL /////////////////////////////////////////////////
			////////////////////////////////////// TOP INDIVIDUAL /////////////////////////////////////////////////
			////////////////////////////////////// TOP INDIVIDUAL /////////////////////////////////////////////////
			////////////////////////////////////// TOP INDIVIDUAL /////////////////////////////////////////////////

			KRTTopAEquation:          s.KRT.TopAntagonistOfAllRuns.Program.ToMathematicalString(),
			KRTTopAEquationPolDegree: CountPolDegree(s.KRT.TopAntagonistOfAllRuns.Program.ToMathematicalString()),
			HoFTopAEquation:          s.HoF.TopAntagonistOfAllRuns.Program.ToMathematicalString(),
			HoFTopAEquationPolDegree: CountPolDegree(s.HoF.TopAntagonistOfAllRuns.Program.ToMathematicalString()),
			RRTopAEquation:           s.RR.TopAntagonistOfAllRuns.Program.ToMathematicalString(),
			RRTopAEquationPolDegree:  CountPolDegree(s.RR.TopAntagonistOfAllRuns.Program.ToMathematicalString()),
			SETTopAEquation:          s.SET.TopAntagonistOfAllRuns.Program.ToMathematicalString(),
			SETTopAEquationPolDegree: CountPolDegree(s.SET.TopAntagonistOfAllRuns.Program.ToMathematicalString()),

			KRTTopPEquation:          s.KRT.TopProtagonistOfAllRuns.Program.ToMathematicalString(),
			KRTTopPEquationPolDegree: CountPolDegree(s.KRT.TopProtagonistOfAllRuns.Program.ToMathematicalString()),
			HoFTopPEquation:          s.HoF.TopProtagonistOfAllRuns.Program.ToMathematicalString(),
			HoFTopPEquationPolDegree: CountPolDegree(s.HoF.TopProtagonistOfAllRuns.Program.ToMathematicalString()),
			RRTopPEquation:           s.RR.TopProtagonistOfAllRuns.Program.ToMathematicalString(),
			RRTopPEquationPolDegree:  CountPolDegree(s.RR.TopProtagonistOfAllRuns.Program.ToMathematicalString()),
			SETTopPEquation:          s.SET.TopProtagonistOfAllRuns.Program.ToMathematicalString(),
			SETTopPEquationPolDegree: CountPolDegree(s.SET.TopProtagonistOfAllRuns.Program.ToMathematicalString()),

			KRTTopAntagonistBestFitnessInSim: s.KRT.TopAntagonistOfAllRuns.BestFitness,
			HoFTopAntagonistBestFitnessInSim: s.HoF.TopAntagonistOfAllRuns.BestFitness,
			RRTTopAntagonistBestFitnessInSim: s.RR.TopAntagonistOfAllRuns.BestFitness,
			SETTopAntagonistBestFitnessInSim: s.SET.TopAntagonistOfAllRuns.BestFitness,

			KRTTopProtagonistBestFitnessInSim: s.KRT.TopProtagonistOfAllRuns.BestFitness,
			HoFTopProtagonistBestFitnessInSim: s.HoF.TopProtagonistOfAllRuns.BestFitness,
			RRTTopProtagonistBestFitnessInSim: s.RR.TopProtagonistOfAllRuns.BestFitness,
			SETTopProtagonistBestFitnessInSim: s.SET.TopProtagonistOfAllRuns.BestFitness,

			KRTTopAntagonistBirthGenInSim: s.KRT.TopAntagonistOfAllRuns.BirthGen,
			HoFTopAntagonistBirthGenInSim: s.HoF.TopAntagonistOfAllRuns.BirthGen,
			RRTTopAntagonistBirthGenInSim: s.RR.TopAntagonistOfAllRuns.BirthGen,
			SETTopAntagonistBirthGenInSim: s.SET.TopAntagonistOfAllRuns.BirthGen,

			KRTTopProtagonistBirthGenInSim: s.KRT.TopProtagonistOfAllRuns.BirthGen,
			HoFTopProtagonistBirthGenInSim: s.HoF.TopProtagonistOfAllRuns.BirthGen,
			RRTTopProtagonistBirthGenInSim: s.RR.TopProtagonistOfAllRuns.BirthGen,
			SETTopProtagonistBirthGenInSim: s.SET.TopProtagonistOfAllRuns.BirthGen,

			KRTTopAntagonistAgeInSim: s.KRT.TopAntagonistOfAllRuns.Age,
			HoFTopAntagonistAgeInSim: s.HoF.TopAntagonistOfAllRuns.Age,
			RRTTopAntagonistAgeInSim: s.RR.TopAntagonistOfAllRuns.Age,
			SETTopAntagonistAgeInSim: s.SET.TopAntagonistOfAllRuns.Age,

			KRTTopProtagonistAgeInSim: s.KRT.TopProtagonistOfAllRuns.Age,
			HoFTopProtagonistAgeInSim: s.HoF.TopProtagonistOfAllRuns.Age,
			RRTTopProtagonistAgeInSim: s.RR.TopProtagonistOfAllRuns.Age,
			SETTopProtagonistAgeInSim: s.SET.TopProtagonistOfAllRuns.Age,

			KRTTopAntagonistStrategyInSim: StrategiesToString(s.KRT.TopAntagonistOfAllRuns),
			HoFTopAntagonistStrategyInSim: StrategiesToString(s.HoF.TopAntagonistOfAllRuns),
			RRTTopAntagonistStrategyInSim: StrategiesToString(s.RR.TopAntagonistOfAllRuns),
			SETTopAntagonistStrategyInSim: StrategiesToString(s.SET.TopAntagonistOfAllRuns),

			KRTTopProtagonistStrategyInSim: StrategiesToString(s.KRT.TopProtagonistOfAllRuns),
			HoFTopProtagonistStrategyInSim: StrategiesToString(s.HoF.TopProtagonistOfAllRuns),
			RRTTopProtagonistStrategyInSim: StrategiesToString(s.RR.TopProtagonistOfAllRuns),
			SETTopProtagonistStrategyInSim: StrategiesToString(s.SET.TopProtagonistOfAllRuns),

			KRTAntagonistsAvgAgeInSim: s.KRT.AvgAgeAntagonist,
			HoFAntagonistsAvgAgeInSim: s.HoF.AvgAgeAntagonist,
			RRTAntagonistsAvgAgeInSim: s.RR.AvgAgeAntagonist,
			SETAntagonistsAvgAgeInSim:  s.SET.AvgAgeAntagonist,
			KRTProtagonistsAvgAgeInSim: s.KRT.AvgAgeProtagonist,
			HoFProtagonistsAvgAgeInSim: s.HoF.AvgAgeProtagonist,
			RRTProtagonistsAvgAgeInSim: s.RR.AvgAgeProtagonist,
			SETProtagonistsAvgAgeInSim: s.SET.AvgAgeProtagonist,

			KRTTopAntagonistsAvgBirthGenInSim:  s.KRT.AvgBirthGenAntagonist,
			HoFTopAntagonistsAvgBirthGenInSim:  s.HoF.AvgBirthGenAntagonist,
			RRTTopAntagonistsAvgBirthGenInSim:  s.RR.AvgBirthGenAntagonist,
			SETTopAntagonistsAvgBirthGenInSim:  s.SET.AvgBirthGenAntagonist,
			KRTTopProtagonistsAvgBirthGenInSim: s.KRT.AvgBirthGenProtagonist,
			HoFTopProtagonistsAvgBirthGenInSim: s.HoF.AvgBirthGenProtagonist,
			RRTopProtagonistsAvgBirthGenInSim:  s.RR.AvgBirthGenProtagonist,
			SETTopProtagonistsAvgBirthGenInSim: s.SET.AvgBirthGenProtagonist,

			///////////// ########################### AVERAGES #################################
			KRTAntagonistsMeanInSim: s.KRT.AntagonistMeanFitAcrossAllRuns,
			HoFAntagonistsMeanInSim: s.HoF.AntagonistMeanFitAcrossAllRuns,
			RRAntagonistsMeanInSim:  s.RR.AntagonistMeanFitAcrossAllRuns,
			SETAntagonistsMeanInSim: s.SET.AntagonistMeanFitAcrossAllRuns,

			KRTProtagonistsMeanInSim: s.KRT.ProtagonistMeanFitAcrossAllRuns,
			HoFProtagonistsMeanInSim: s.HoF.ProtagonistMeanFitAcrossAllRuns,
			RRProtagonistsMeanInSim:  s.RR.ProtagonistMeanFitAcrossAllRuns,
			SETProtagonistsMeanInSim: s.SET.ProtagonistMeanFitAcrossAllRuns,

			KRTAntagonistStdDevInSim: s.KRT.AntagonistMeanStdDev,
			HoFAntagonistStdDevInSim: s.HoF.AntagonistMeanStdDev,
			RRAntagonistStdDevInSim:  s.RR.AntagonistMeanStdDev,
			SETAntagonistStdDevInSim: s.SET.AntagonistMeanStdDev,

			KRTProtagonistStdDevInSim: s.KRT.ProtagonistMeanStdDev,
			HoFProtagonistStdDevInSim: s.HoF.ProtagonistMeanStdDev,
			RRProtagonistStdDevInSim:  s.RR.ProtagonistMeanStdDev,
			SETProtagonistStdDevInSim: s.SET.ProtagonistMeanStdDev,

			KRTAntagonistVarInSim: s.KRT.AntagonistMeanVar,
			HoFAntagonistVarInSim: s.HoF.AntagonistMeanVar,
			RRAntagonistVarInSim:  s.RR.AntagonistMeanVar,
			SETAntagonistVarInSim: s.SET.AntagonistMeanVar,

			KRTProtagonistVarInSim: s.KRT.ProtagonistMeanVar,
			HoFProtagonistVarInSim: s.HoF.ProtagonistMeanVar,
			RRProtagonistVarInSim:  s.RR.ProtagonistMeanVar,
			SETProtagonistVarInSim: s.SET.ProtagonistMeanVar,
		}

	
		AllCSVGens[i] = csvGen
	}

	for i := 0; i < len(s.KRT.TopAntagonistOfAllRuns.Strategy); i++ {
		stat := CSVStrat{
			Num:          i,
			KRTTopAStrat: string(s.KRT.TopAntagonistOfAllRuns.Strategy[i]),
			HOFTopAStrat: string(s.HoF.TopAntagonistOfAllRuns.Strategy[i]),
			RRTopAStrat:  string(s.RR.TopAntagonistOfAllRuns.Strategy[i]),
			SETTopAStrat: string(s.SET.TopAntagonistOfAllRuns.Strategy[i]),
			KRTTopPStrat: string(s.KRT.TopProtagonistOfAllRuns.Strategy[i]),
			HOFTopPStrat: string(s.HoF.TopProtagonistOfAllRuns.Strategy[i]),
			RRTopPStrat:  string(s.RR.TopProtagonistOfAllRuns.Strategy[i]),
			SETTopPStrat: string(s.SET.TopProtagonistOfAllRuns.Strategy[i]),
		}

		AllCSVStats = append(AllCSVStats, stat)
	}

	return AllCSVGens, AllCSVStats
}

func CountVariable(mathematicalString string) int {
	count := 0

	for i := 0; i < len(mathematicalString); i++ {
		if mathematicalString[i] == 'x' {
			count++
		}
	}

	return count
}

func CountPolDegree(mathematicalString string) int {
	count := 0

	for i := 0; i < len(mathematicalString)-1; i++ {
		if mathematicalString[i] == '*' && mathematicalString[i+1] == 'x' {
			count++
		}
		if mathematicalString[i] == '*' && mathematicalString[i+1] == '0' {
			count = 0
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

// WriteStratCSV takes in the CSVAvg... object and writes it to the path. It DOES NOT CREATE the path if it does not exist.
func (s *SimulationResult) WriteStratCSV(file *os.File, data []CSVStrat) error {
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

	TopAntagonistOfAllRuns  Individual
	TopProtagonistOfAllRuns Individual

	FinalBestAntagonistOfAllRuns  Individual
	FinalBestProtagonistOfAllRuns Individual

	AntagonistMeanFitAcrossAllRuns  float64
	ProtagonistMeanFitAcrossAllRuns float64

	AntagonistMeanStdDev  float64
	ProtagonistMeanStdDev float64

	AntagonistMeanVar  float64
	ProtagonistMeanVar float64

	OldestAntagonist  float64
	OldestProtagonist float64

	AvgAgeAntagonist  float64
	AvgAgeProtagonist float64

	AvgBirthGenAntagonist  float64
	AvgBirthGenProtagonist float64

	TopGenerationalResultsAcrossAllRuns []GenerationResult
}

func (t *TopologicalResult) ToPlotterFormat(topology string) TopologyPlot {
	tp := TopologyPlot{
		Topology:                  topology,
		AvgOfAllAntagonistsInGen:  make([]float64, len(t.TopGenerationalResultsAcrossAllRuns)),
		AvgOfAllProtagonistsInGen: make([]float64, len(t.TopGenerationalResultsAcrossAllRuns)),
		TopAntagonistsBestInGen:   make([]float64, len(t.TopGenerationalResultsAcrossAllRuns)),
		TopProtagonistsBestInGen:  make([]float64, len(t.TopGenerationalResultsAcrossAllRuns)),

		XAxisGen: make([]int, len(t.TopGenerationalResultsAcrossAllRuns)),

		SupremeAntagonistStrategies:   make([]Strategy, len(t.TopProtagonistOfAllRuns.Strategy)),
		SupremeProtagonistStrategyies: make([]Strategy, len(t.TopAntagonistOfAllRuns.Strategy)),
	}

	// List out the best performance of antagonists and protagonists across all runs for each generation
	for i := range t.TopGenerationalResultsAcrossAllRuns {
		tp.AvgOfAllAntagonistsInGen[i] = t.TopGenerationalResultsAcrossAllRuns[i].AllAntagonistAverageFitness
		tp.AvgOfAllProtagonistsInGen[i] = t.TopGenerationalResultsAcrossAllRuns[i].AllProtagonistAverageFitness

		tp.TopAntagonistsBestInGen[i] = t.TopGenerationalResultsAcrossAllRuns[i].BestAntagonist.BestFitness
		tp.TopProtagonistsBestInGen[i] = t.TopGenerationalResultsAcrossAllRuns[i].BestProtagonist.BestFitness

		tp.XAxisGen[i] = i
	}

	// List out the strategies of the best antagonists and protagonists
	for i := range t.TopAntagonistOfAllRuns.Strategy {
		tp.SupremeAntagonistStrategies[i] = t.TopAntagonistOfAllRuns.Strategy[i]
		tp.SupremeProtagonistStrategyies[i] = t.TopProtagonistOfAllRuns.Strategy[i]
	}

	return tp
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

	antBirthGenSum := 0.0
	proBirthGenSum := 0.0

	//topAntNoCSum := 0.0
	//topProNoCSum := 0.0

	antagonistStdDevSum := 0.0
	protagonistStdDevSum := 0.0

	antagonistVarSum := 0.0
	protagonistVarSum := 0.0

	antagonistAvgSum := 0.0
	protagonistAvgSum := 0.0

	runCount := len(evolutionResults)
	runCountF := float64(runCount)

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

		antagonistAvgSum += runResult.AntagonistMeanFit
		protagonistAvgSum += runResult.ProtagonistMeanFit

		antagonistVarSum += runResult.AntagonistMeanVar
		protagonistVarSum += runResult.ProtagonistMeanVar

		antagonistStdDevSum += runResult.AntagonistMeanStdDev
		protagonistStdDevSum += runResult.ProtagonistMeanStdDev

		antAgeSum += runResult.AvgAgeAntagonist
		proAgeSum += runResult.AvgAgeProtagonist

		antBirthGenSum += runResult.AvgBirthGenAntagonist
		proBirthGenSum += runResult.AvgBirthGenProtagonist

		//topAntNoCSum += runResult.TopAntagonistInRun.NoOfCompetitions
		//topProNoCSum += runResult.TopProtagonistInRun.NoOfCompetitions
	}

	// Averages

	//Average generations
	genLength := len(evolutionResults[0].GenerationalResults)
	//genLengthFloat := float64(genLength)

	clonedGenResults := make([]GenerationResult, genLength)

	// For each generation
	for i := 0; i < genLength; i++ {
		topAntagonistInRunInner := Individual{AverageFitness: math.MinInt16}
		topProtagonistInRunInner := Individual{AverageFitness: math.MinInt16}
		//finAntagonistInRunInner := Individual{AverageFitness: math.MinInt16}
		//finProtagonistInRunInner := Individual{AverageFitness: math.MinInt16}

		antAgeSumInner := 0.0
		proAgeSumInner := 0.0

		antBirthGenSumInner := 0.0
		proBirthGenSumInner := 0.0

		antagonistStdDevSumInner := 0.0
		protagonistStdDevSumInner := 0.0

		antagonistVarSumInner := 0.0
		protagonistVarSumInner := 0.0

		antagonistAvgSumInner := 0.0
		protagonistAvgSumInner := 0.0

		// We shall average the value/find the best value of each run
		for j := 0; j < len(evolutionResults); j++ {
			run := evolutionResults[j]
			gen := run.GenerationalResults[i]

			currAntagonist := gen.BestAntagonist
			currProtagonist := gen.BestProtagonist
			//currFinAntagonist := gen.
			//currFinProtagonist := gen.FinalBestProtagonistOfAllRuns

			if currAntagonist.AverageFitness > topAntagonistInRunInner.AverageFitness {
				topAntagonistInRunInner = currAntagonist
			}
			if currProtagonist.AverageFitness > topProtagonistInRunInner.AverageFitness {
				topProtagonistInRunInner = currProtagonist
			}

			antagonistAvgSumInner += gen.AllAntagonistAverageFitness
			protagonistAvgSumInner += gen.AllProtagonistAverageFitness

			antagonistVarSumInner += gen.AntagonistVariance
			protagonistVarSumInner += gen.ProtagonistVariance

			antagonistStdDevSumInner += gen.AntagonistStdDev
			protagonistStdDevSumInner += gen.ProtagonistStdDev

			antAgeSumInner += gen.AntagonistAvgAge
			proAgeSumInner += gen.ProtagonistAvgAge

			antBirthGenSumInner += gen.AntagonistAvgBirthGen
			proBirthGenSumInner += gen.ProtagonistAvgBirthGen
		}

		clonedGenResults[i] = GenerationResult{
			ID:                           uint32(i),
			BestAntagonist:               topAntagonistInRunInner,
			BestProtagonist:              topProtagonistInRunInner,
			AllAntagonistAverageFitness:  antagonistAvgSumInner / runCountF, // WRONG!
			AllProtagonistAverageFitness: protagonistAvgSumInner / runCountF,
			AntagonistStdDev:             antagonistStdDevSumInner / runCountF,
			AntagonistVariance:           antagonistVarSumInner / runCountF,
			ProtagonistStdDev:            protagonistStdDevSumInner / runCountF,
			ProtagonistVariance:          protagonistVarSumInner / runCountF,
			ProtagonistAvgBirthGen:       proBirthGenSumInner / runCountF,

			AntagonistAvgBirthGen: antBirthGenSumInner / runCountF,
			AntagonistAvgAge:      antAgeSumInner / runCountF,
			ProtagonistAvgAge:     proAgeSumInner / runCountF,
		}
	}

	return TopologicalResult{
		HasBeenAnalyzed:               true,
		TopAntagonistOfAllRuns:        topAntagonistInRun.Clone(-1),
		TopProtagonistOfAllRuns:       topProtagonistInRun.Clone(-1),
		FinalBestAntagonistOfAllRuns:  finAntagonistInRun,
		FinalBestProtagonistOfAllRuns: finProtagonistInRun,

		AntagonistMeanFitAcrossAllRuns:  antagonistAvgSum / runCountF,
		ProtagonistMeanFitAcrossAllRuns: protagonistAvgSum / runCountF,

		AntagonistMeanStdDev:  antagonistStdDevSum / runCountF,
		ProtagonistMeanStdDev: protagonistStdDevSum / runCountF,

		AntagonistMeanVar:  antagonistVarSum / runCountF,
		ProtagonistMeanVar: protagonistVarSum / runCountF,

		AvgAgeAntagonist:  antAgeSum / runCountF,
		AvgAgeProtagonist: antAgeSum / runCountF,

		AvgBirthGenAntagonist: antBirthGenSum / runCountF,
		AvgBirthGenProtagonist: proBirthGenSum / runCountF,

		OldestAntagonist:  oldAnt,
		OldestProtagonist: oldPro,

		TopGenerationalResultsAcrossAllRuns: clonedGenResults,
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
