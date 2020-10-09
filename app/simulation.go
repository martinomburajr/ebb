package app

import (
	"fmt"
	"github.com/martinomburajr/ebb/evolution"
	"github.com/martinomburajr/masters-go/evolog"
	"log"
	"math/rand"
	"os"
	"sync"
)

// Simulation represents multiple evolutionary runs.
type Simulation struct {
	ParenthesizedProgram evolution.BinaryTree
	Config               ApplicationConfig

	// 	ComplexityLevel represents the number of non-terminals in the equation
	//	easy(0), intermediate(1), or complex(2)
	// Simple  |NT| =  1->4
	// Medium  |NT| =  4->8
	// Complex |NT| = 10->20
	ComplexityLevel int
	Runs            int
}

// NewStartProgram generates a new start program randomly based on the polDegree
func NewStartProgram(polDegree int, terminals, nonTerminals []rune) evolution.BinaryTree {
	if polDegree < 0 {
		polDegree = rand.Intn(15)
	}

	var validTree evolution.BinaryTree
	var err error

	if rand.Intn(100) % 2 == 0 {
		for validTree == nil || err != nil {
			validTree, err = evolution.NewRandomTreeFromPolDegreeCount(polDegree, 15, terminals, nonTerminals)
		}

		return validTree
	}

	validTree = evolution.NewRandomTreeFromIVarCount(polDegree, 15, terminals, nonTerminals)

	return validTree
}

// Init performs more work with regards to setting up the Application runtime
func Init(config ApplicationConfig) (Simulation, error) {
	// Generate Program
	startProgram := evolution.BinaryTree{}

	terminals := config.Params.SpecParam.AvailableSymbolicExpressions.Terminals
	nonTerminals := config.Params.SpecParam.AvailableSymbolicExpressions.NonTerminals
	 //_ := config.Params.SpecParam.AvailableSymbolicExpressions.Variables

	switch config.Complexity {
	case -1:
		min := 1
		max := 2
		randSize := rand.Intn(max) + min

		err := fmt.Errorf("start Program Creation Test Error")

		for err != nil {
			startProgram = NewStartProgram(randSize, terminals, nonTerminals)
			_, err = evolution.NewSpec(startProgram, config.Params.SpecParam, config.Params.FitnessStrategy)
		}

	case 0:
		min := 1
		max := 4
		randSize := rand.Intn(max) + min

		err := fmt.Errorf("start Program Creation Test Error")

		for err != nil {
			startProgram = NewStartProgram(randSize, terminals, nonTerminals)
			_, err = evolution.NewSpec(startProgram, config.Params.SpecParam, config.Params.FitnessStrategy)
		}

	case 1:
		min := 4
		max := 8
		randSize := rand.Intn(max) + min
		err := fmt.Errorf("start Program Creation Test Error")

		for err != nil {
			startProgram = NewStartProgram(randSize, terminals, nonTerminals)
			_, err = evolution.NewSpec(startProgram, config.Params.SpecParam, config.Params.FitnessStrategy)
		}

	case 2:
		min := 8
		max := 15
		randSize := rand.Intn(max) + min

		err := fmt.Errorf("start Program Creation Test Error")

		for err != nil {
			startProgram = NewStartProgram(randSize, terminals, nonTerminals)
			_, err = evolution.NewSpec(startProgram, config.Params.SpecParam, config.Params.FitnessStrategy)
		}
	default:
		panic("Simulation:Init -> Failed to init invalid ComplexityLevel [-1,2]")
	}

	errChan := make(chan error)
	doneChan := make(chan struct{})
	logChan := make(chan evolog.Logger)


	startProgram = startProgram.Sanitize()

	log.Printf("\nStarting Program: %s", startProgram.ToMathematicalString())

	// setup spec
	spec, err := evolution.NewSpec(startProgram, config.Params.SpecParam, config.Params.FitnessStrategy)
	if err != nil {
		return Simulation{}, fmt.Errorf("failed to creat spec: %v", err)
	}

	config.Params.Spec = spec
	config.Params.StartIndividual = startProgram
	config.Params.ErrorChan = errChan
	config.Params.LoggingChan = logChan
	config.Params.DoneChan = doneChan

	simulation := Simulation{
		ParenthesizedProgram: startProgram,
		Config:               config,
		ComplexityLevel:      config.Complexity,
		Runs:                 config.Runs,
	}

	// Initialize the error chan
	go func(errChan chan error, doneChan chan struct{}, logChan chan evolog.Logger) {
		log.Println("starting logger goroutine ...")

		for {
			select {
			case err := <-errChan:
				log.Fatalf("error: %v", err)
			case done := <-doneChan:
				log.Printf("\n\nSimulation Terminating...%s", done)
				os.Exit(0)
			case l := <-logChan:
				log.Printf("log: %s", l.Message)
			}
		}
	}(errChan, doneChan, logChan)

	return simulation, nil
}

func (s *Simulation) Start() (evolution.SimulationResult, error) {
	KRTResults := make([]evolution.EvolutionResult, s.Runs)
	HoFResults := make([]evolution.EvolutionResult, s.Runs)
	RRResults := make([]evolution.EvolutionResult, s.Runs)
	SETResults := make([]evolution.EvolutionResult, s.Runs)

	// Setup Engine
	engine := evolution.NewEngine(s.Config.Params)

	err := engine.Validate()
	if err != nil {
		return evolution.SimulationResult{}, err
	}

	for i := 0; i < s.Runs; i++ {
		currentRun := int64(i + 1)

		krtEngine := engine.Clone()
		krtEngine.CurrentRun = currentRun

		rrEngine :=  engine.Clone()
		rrEngine.CurrentRun = currentRun

		hofEngine := engine.Clone()
		hofEngine.CurrentRun = currentRun

		setEngine := engine.Clone()
		setEngine.CurrentRun = currentRun

		// KRT
		krt := &evolution.KRandom{Engine: krtEngine}
		krtResult, err := s.startRun(&krtEngine, krt)
		if err != nil {
			return evolution.SimulationResult{}, err
		}

		KRTResults[i] = krtResult

		//HoF
		hof := evolution.HallOfFame{Engine: engine.Clone()}
		hofResult, err := s.startRun(&hofEngine, &hof)
		if err != nil {
			return evolution.SimulationResult{}, err
		}

		HoFResults[i] = hofResult

		// RR
		rr := &evolution.RoundRobin{Engine: rrEngine}
		rrResult, err := s.startRun(&rrEngine, rr)
		if err != nil {
			return evolution.SimulationResult{}, err
		}

		RRResults[i] = rrResult

		// SET
		set := evolution.SingleEliminationTournament{Engine: engine.Clone()}
		setResult, err := s.startRun(&setEngine, &set)
		if err != nil {
			return evolution.SimulationResult{}, err
		}

		SETResults[i] = setResult
	}

	// At this point all the runs are complete and we now compress all the run information into separate evolution.TopologicalResult
	// objects

	// Bring together all the different topologies into one
	var KRTTopologyResult, HoFTopologyResult, RRTopologyResult, SETTopologyResult evolution.TopologicalResult

	s.combineRuns(&KRTTopologyResult, KRTResults)

	s.combineRuns(&HoFTopologyResult, HoFResults)

	s.combineRuns(&RRTopologyResult, RRResults)

	s.combineRuns(&SETTopologyResult, SETResults)

	simulationResult := evolution.SimulationResult{
		KRT: KRTTopologyResult,
		RR:  RRTopologyResult,
		SET: SETTopologyResult,
		HoF: HoFTopologyResult,
	}

	return simulationResult, nil
}



//func (s *Simulation) StartP() (evolution.SimulationResult, error) {
//	KRTResults := make([]evolution.EvolutionResult, s.Runs)
//	HoFResults := make([]evolution.EvolutionResult, s.Runs)
//	RRResults := make([]evolution.EvolutionResult, s.Runs)
//	SETResults := make([]evolution.EvolutionResult, s.Runs)
//
//	// Setup Engine
//	engine := evolution.Engine{}
//
//	err := engine.Validate()
//	if err != nil {
//		return evolution.SimulationResult{}, err
//	}
//
//	outerWg := sync.WaitGroup{}
//
//	for i := 0; i < s.Runs; i++ {
//		i := i
//		go func() {
//			outerWg.Add(1)
//			defer outerWg.Done()
//			// Run runs in parallel for each topology
//
//			wg := sync.WaitGroup{}
//			// KRT
//				go s.startRunP(&wg, i, s.Config.Params.Clone(), engine.Clone(), KRTResults)
//
//				// HoF
//				go s.startRunP(&wg, i, s.Config.Params.Clone(), engine.Clone(), HoFResults)
//
//				// RR
//				go s.startRunP(&wg, i, s.Config.Params.Clone(), engine.Clone(), RRResults)
//
//				// SET
//				go s.startRunP(&wg, i, s.Config.Params.Clone(), engine.Clone(), SETResults)
//
//			wg.Wait()
//		}()
//	}
//
//	outerWg.Wait()
//
//	// At this point all the runs are complete and we now compress all the run information into separate evolution.TopologicalResult
//	// objects
//
//	// Bring together all the different topologies into one
//	var KRTTopologyResult, HoFTopologyResult, RRTopologyResult, SETTopologyResult evolution.TopologicalResult
//
//	wg := sync.WaitGroup{}
//
//		go s.combineRunsP(&wg, &KRTTopologyResult, KRTResults)
//
//		go s.combineRunsP(&wg, &HoFTopologyResult, HoFResults)
//
//		go s.combineRunsP(&wg, &RRTopologyResult, RRResults)
//
//		go s.combineRunsP(&wg, &SETTopologyResult, SETResults)
//
//	wg.Wait()
//
//	simulationResult := evolution.SimulationResult{
//		KRT: KRTTopologyResult,
//		RR:  HoFTopologyResult,
//		SET: RRTopologyResult,
//		HoF: SETTopologyResult,
//	}
//
//	return simulationResult, nil
//}

func (s *Simulation) combineRuns(result *evolution.TopologicalResult, evolutionResults []evolution.EvolutionResult) {
	topologicalResults := evolution.NewTopologicalResults("", evolutionResults)
	*result = topologicalResults
}

func (s *Simulation) combineRunsP(wg *sync.WaitGroup, result *evolution.TopologicalResult, evolutionResults []evolution.EvolutionResult) {
	wg.Add(1)
	defer wg.Done()

	topologicalResults := evolution.NewTopologicalResults("", evolutionResults)
	result = &topologicalResults
}

func (s *Simulation) startRun(engine *evolution.Engine, topology evolution.Evolver) (evolution.EvolutionResult, error) {
	result, err := engine.Evolve(topology)

	if err != nil {
		return evolution.EvolutionResult{}, err
	}

	return result, nil
}

//func (s *Simulation) startRunP(wg *sync.WaitGroup, index int, params evolution.EvolutionParams, engine evolution.Engine, results []evolution.EvolutionResult) {
//	wg.Add(1)
//	defer wg.Done()
//
//	competition := evolution.SingleEliminationTournament{Engine: engine}
//
//	result, err := competition.Evolve(params, &competition)
//
//	if err != nil {
//		params.ErrorChan <- err
//	}
//
//	results[index] = result
//}

type RunResult evolution.EvolutionResult

func CombineComplexityStats(baseDir string) {}

func CombineComplexityDiagrams(baseDir string) {}

type CSVOutputter interface {
	// This averages results across gen[i] across all the runs. Meaning Gen0 will be averaged across all the runs
	// Gen1 will be averaged across all the runs etc.
	OutputAveragedGenerationalStatistics(baseDir string)
	OutputEpochalStatistics(baseDir string)
}

type DiagramOutputter interface {
	// This averages results across gen[i] across all the runs. Meaning Gen0 will be averaged across all the runs
	// Gen1 will be averaged across all the runs etc.
	OutputAveragedGenerationalStatistics(baseDir string)
	OutputEpochalStatistics(baseDir string)
}

type RunCombiner interface {
	CombineStats(complexityDir string)
	CombineDiagrams(complexityDir string)
}

type ComplexityCombiner interface {
	CombineComplexityLevelStats(baseDir string)
	CombineComplexityLevelDiagrams(baseDir string)
}

type AllCombiner interface {
	CombineAllStats(baseDir string, complexityDirs []string)
	CombineRunLevelDiagrams(baseDir string, complexityDirs []string)
}

//func CombineRunResultByGeneration(topology string, runs []RunResult) error {
//	finalResult := RunResult{}
//
//
//	topAntagonistInRunSumFitAvgSum := 0.0
//	topProtagonistInRunSumFitAvgSum := 0.0
//
//
//	for i := 0; i < len(runs); i++ {
//		for j := 0; j < len(runs[i]); j++ {
//
//		}
//		topAntagonistInRunSumFitAvgSum += runs[i].result.TopAntagonistInRun.AverageFitness
//		topProtagonistInRunSumFitAvgSum += runs[i].result.TopAntagonistInRun.AverageFitness
//	}
//
//	gens := make([]analysis.CSVAvgGenerationsCombinedAcrossRuns, 0)
//
//	topAntagonistMeanSum := 0.0
//	topProtagonistMeanSum := 0.0
//	topAntagonistBestFitSum := 0.0
//	topProtagonistBestFitSum := 0.0
//	topAntagonistStdDevSumSum := 0.0
//	topProtagonistStdDevSumSum := 0.0
//	topAntagonistVarSum := 0.0
//	topProtagonistVarSum := 0.0
//	// for each run
//	for i := 0; i < len(runs); i++ {
//		// for each generation
//		for j := 0; j < len(runs[i].result.Generational.AntagonistAverageInEachGeneration); j++ {
//			gen := runs[i].result.ThoroughlySortedGenerations[j]
//
//			topAntagonistBestFitSum += runs[i].result.Generational.BestAntagonistInEachGenerationByAvgFitness[j].AverageFitness
//			topAntagonistBestFitSum += runs[i].result.Generational.BestAntagonistInEachGenerationByAvgFitness[j].AverageFitness
//		}
//	}
//
//	runs[0].result
//}
