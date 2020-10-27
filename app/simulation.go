package app

import (
	"fmt"
	"github.com/martinomburajr/ebb/evolog"
	"github.com/martinomburajr/ebb/evolution"
	"log"
	"math/rand"
	"os"
	"time"
)

// Simulation represents multiple evolutionary runs.
type Simulation struct {
	StartTime            time.Time
	End                  time.Time
	ParenthesizedProgram evolution.BinaryTree
	Config               *ApplicationConfig

	// 	ComplexityLevel represents the number of non-terminals in the equation
	//	easy(0), intermediate(1), or complex(2)
	// Simple  |NT| =  1->4
	// Medium  |NT| =  4->8
	// Complex |NT| = 10->20
	ComplexityLevel int
	Runs            int

	LoggingChan chan evolog.Logger
	ErrorChan   chan error
}

func (s *Simulation) LogTime(str string) {
	msg := fmt.Sprintf("%s-%s", str, s.End.Sub(s.StartTime).String())

	s.LoggingChan <- evolog.Logger{Type: evolog.LoggerEvolution, Message: msg,
		Timestamp: time.Now()}
}

// NewStartProgram generates a new start program randomly based on the polDegree
func NewStartProgram(polDegree int, terminals, nonTerminals []rune) evolution.BinaryTree {
	if polDegree < 0 {
		polDegree = rand.Intn(15)
	}

	var validTree evolution.BinaryTree
	var err error

	if rand.Intn(100)%2 == 0 {
		for validTree == nil || err != nil {
			validTree, err = evolution.NewRandomTreeFromPolDegreeCount(polDegree, 18, terminals, nonTerminals)
		}

		return validTree
	}

	validTree = evolution.NewRandomTreeFromIVarCount(polDegree, 18, terminals, nonTerminals)

	return validTree
}

// Init performs more work with regards to setting up the Application runtime
func Init(config *ApplicationConfig) (Simulation, error) {
	// Generate Program
	startProgram := evolution.BinaryTree{}

	// Randomize the complexity
	config.Complexity = config.Iter % 3

	terminals := config.Params.SpecParam.AvailableSymbolicExpressions.Terminals
	nonTerminals := config.Params.SpecParam.AvailableSymbolicExpressions.NonTerminals
	//_ := config.Params.SpecParam.AvailableSymbolicExpressions.Variables

	switch config.Complexity {
	case 0:
		min := 1
		max := 6
		randSize := rand.Intn(max) + min

		err := fmt.Errorf("start Program Creation Test Error")

		for err != nil {
			startProgram = NewStartProgram(randSize, terminals, nonTerminals)
			_, err = evolution.NewSpec(startProgram, config.Params.SpecParam, config.Params.FitnessStrategy)
		}

	case 1:
		min := 6
		max := 10
		randSize := rand.Intn(max) + min

		err := fmt.Errorf("start Program Creation Test Error")

		for err != nil {
			startProgram = NewStartProgram(randSize, terminals, nonTerminals)
			_, err = evolution.NewSpec(startProgram, config.Params.SpecParam, config.Params.FitnessStrategy)
		}

	case 2:
		min := 10
		max := 18
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

	programString := startProgram.ToMathematicalString()
	startMsg := fmt.Sprintf("\nStarting Program: %s | Complexity: %d | PolDeg: %d | VarN: %d",
		programString, config.Complexity, evolution.CountPolDegree(programString), evolution.CountVariable(programString))
	paramMsg := fmt.Sprintf("Runs: %d | EachIndividual: %d | Gens: %d | StratLen: %d | KRT#: %d | SET#: %d | HoFInt: %.2f",
		config.Runs, config.Params.EachPopulationSize, config.Params.GenerationsCount, config.Params.Strategies.NumStrategiesToUse,
		config.Params.Topology.KRandomK, config.Params.Topology.SETNoOfTournaments, config.Params.Topology.HoFGenerationInterval)
	log.Println(startMsg)
	log.Println(paramMsg)
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
	simulation.LoggingChan = logChan
	simulation.ErrorChan = errChan

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
				l.DisplayMessage()
			}
		}
	}(errChan, doneChan, logChan)

	return simulation, nil
}

func (s *Simulation) Start() (evolution.SimulationResult, error) {
	s.LogMessage(fmt.Sprintf("Iter: %d", s.Config.Iter), evolog.LoggerSimulation)
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
		krtEngine.Topology = "KRT"

		krtEngine.CurrentRun = currentRun
		krt := &evolution.KRandom{Engine: krtEngine}
		krtResult, err := s.startRun(&krtEngine, krt)
		if err != nil {
			s.ErrorChan <- err
		}

		KRTResults[i] = krtResult
		krtEngine.LogTime("KRT Duration")
		s.printSystemStats()
		//}(&wg)

		//HoF
		//go func(wg *sync.WaitGroup) {
		//	defer wg.Done()

		hofEngine := engine.Clone()
		hofEngine.CurrentRun = currentRun

		hof := evolution.HallOfFame{Engine: engine.Clone()}
		hofResult, err := s.startRun(&hofEngine, &hof)
		if err != nil {
			s.ErrorChan <- err
		}

		HoFResults[i] = hofResult
		hofEngine.LogTime("HoF Duration")
		//s.printSystemStats()
		//}(&wg)

		// RR
		//go func(wg *sync.WaitGroup) {
		//	defer wg.Done()

		rrEngine := engine.Clone()
		rrEngine.CurrentRun = currentRun
		rr := &evolution.RoundRobin{Engine: rrEngine}
		rrResult, err := s.startRun(&rrEngine, rr)
		if err != nil {
			s.ErrorChan <- err
		}

		RRResults[i] = rrResult
		//rrEngine.LogTime("RR Duration")
		//s.printSystemStats()
		//}(&wg)

		// SET
		//go func(wg *sync.WaitGroup) {
		//	defer wg.Done()

		setEngine := engine.Clone()
		setEngine.CurrentRun = currentRun
		setEngine.Topology = "SET"

		set := evolution.SingleEliminationTournament{Engine: engine.Clone()}
		setResult, err := s.startRun(&setEngine, &set)
		if err != nil {
			s.ErrorChan <- err
		}

		SETResults[i] = setResult
		setEngine.LogTime("SET Duration")
		s.printSystemStats()
		//}(&wg)

		//wg.Wait()
	}

	// At this point all the runs are complete and we now compress all the run information into separate evolution.TopologicalResult
	// objects

	// Bring together all the different topologies into one
	var KRTTopologyResult, HoFTopologyResult, RRTopologyResult, SETTopologyResult evolution.TopologicalResult

	s.combineRuns("KRT", &KRTTopologyResult, KRTResults)

	s.combineRuns("HoF", &HoFTopologyResult, HoFResults)

	s.combineRuns("RR", &RRTopologyResult, RRResults)

	s.combineRuns("SET", &SETTopologyResult, SETResults)

	simulationResult := evolution.SimulationResult{
		KRT: KRTTopologyResult,
		RR:  RRTopologyResult,
		SET: SETTopologyResult,
		HoF: HoFTopologyResult,
	}

	return simulationResult, nil
}

func (s *Simulation) LogMessage(str string, logType int) {
	msg := fmt.Sprintf("%s", str)

	s.LoggingChan <- evolog.Logger{Type: logType, Message: msg,
		Timestamp: time.Now()}
}

func (s *Simulation) combineRuns(topology string, result *evolution.TopologicalResult, evolutionResults []evolution.EvolutionResult) {
	topologicalResults := evolution.NewTopologicalResults("", evolutionResults)
	result.Topology = topology
	*result = topologicalResults
}

func (s *Simulation) startRun(engine *evolution.Engine, topology evolution.Evolver) (evolution.EvolutionResult, error) {
	engine.Start = time.Now()
	result, err := engine.Evolve(topology)
	engine.End = time.Now()

	if err != nil {
		return evolution.EvolutionResult{}, err
	}

	return result, nil
}

type RunResult evolution.EvolutionResult
