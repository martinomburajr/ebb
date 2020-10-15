package evolution

import (
	"fmt"
	"github.com/martinomburajr/ebb/evolog"
	"gonum.org/v1/gonum/stat"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Engine contains the most sufficient data to begin an evolutionary process. The completion of an Engine run is a single
// evolutionary run. Engines contain GenerationResults objects that hold statistical information on any given information
// engines deliberately hold no pointers to other critical data making them conceptually threadsafe. Use the Clone function
// when attempting to do a deep copy of engine elements. The ProgressBar can be used among multiple threads to update
// general progress
type Engine struct {
	// No need to keep track of generations anymore. Once they are done they can be freed as we pass on the relevant
	// data onto the GenerationResult
	//Generations []*Generation   `json:"generations"`

	Start time.Time

	GenerationResults []GenerationResult `json:"generationResults"`
	Parameters        EvolutionParams    `json:"parameters"`

	successfulGenerations                       int
	successfulGenerationsByAvg                  int
	minimumTopProtagonistThreshold              int
	minimumMeanProtagonistInGenerationThreshold int

	idAllocStart  uint32
	idAllocOffset uint32

	CurrentRun  int64

	End time.Time
}

func NewEngine(params EvolutionParams) Engine {
	return Engine{
		GenerationResults:                           make([]GenerationResult, params.GenerationsCount),
		Parameters:                                  params,
		successfulGenerations:                       0,
		successfulGenerationsByAvg:                  0,
		minimumTopProtagonistThreshold:              0,
		minimumMeanProtagonistInGenerationThreshold: 0,
	}
}

func (e *Engine) LogMessage(str string, logType int) {
	msg := fmt.Sprintf("%s", str)

	e.Parameters.LoggingChan <- evolog.Logger{Type: logType, Message: msg,
		Timestamp: time.Now()}
}

func (e *Engine) LogTime(str string) {
	msg := fmt.Sprintf("%s: %s", str, e.End.Sub(e.Start).String())

	e.Parameters.LoggingChan <- evolog.Logger{Type: evolog.LoggerEvolution, Message: msg,
		Timestamp: time.Now()}
}

func (e Engine) Clone() Engine {
	dstGenResult := make([]GenerationResult, len(e.GenerationResults))
	copy(dstGenResult, e.GenerationResults)

	e.Parameters = e.Parameters.Clone()

	return e
}

func (e *Engine) Evolve(topology Evolver) (EvolutionResult, error) {
	params := e.Parameters

	run := strconv.FormatInt(e.CurrentRun, 10)
	e.LogMessage(fmt.Sprintf("Running Topology: %s", topology.Name(run)), evolog.LoggerEvolution)

	// HoF uses a non-standard evolution process, hence it carries its own Evolve method.
	hallOfFame, isHoF := topology.(*HallOfFame)
	if isHoF {
		return hallOfFame.Evolve()
	}

	gen0, _, _, err := e.InitializeGenerations(params)
	if err != nil {
		return EvolutionResult{}, err
	}

	genCount := CalculateGenerationSize(params)

	currGen := gen0
	for i := 0; i < genCount; i++ {
		// 1. CLEANSE
		currGen.CleansePopulations(params)

		// 2. START
		currentGen, nextGeneration, err := topology.Topology(currGen, params)
		if err != nil {
			return EvolutionResult{}, err
		}

		// 3. EVALUATE
		generationResult := currentGen.RunGenerationStatistics()
		e.GenerationResults[i] = generationResult

		if i == e.Parameters.MaxGenerations-1 {
			break
		}

		currGen = nextGeneration

		// Check if Antagonists or Protagonist have clashing IDs
		//idMap := make(map[uint32]int)
		//for _, ind := range currGen.Antagonists {
		//	if idMap[ind.ID] > 1 {
		//		panic("ID CLASH!")
		//	}
		//	idMap[ind.ID]++
		//}
		//
		//for _, ind := range currGen.Protagonists {
		//	if idMap[ind.ID] > 1 {
		//		panic("ID CLASH!")
		//	}
		//	idMap[ind.ID]++
		//}
	}

	evolutionResult := e.AnalyzeResults()

	return evolutionResult, nil
}


// EvaluateTerminationCriteria looks at the current state of the Generation and checks to see if the current
// termination criteria have been achieved. If so it returns true, if not the evolution can move on to the next step
func (e *Engine) EvaluateTerminationCriteria(generation Generation, result GenerationResult,
	params EvolutionParams) (shouldTerminateEvolution bool) {

	meanPro := stat.Mean(generation.ProtagonistAvgFitnessOfEveryIndividual, nil)
	bestProtagonist := generation.BestProtagonist()

	if bestProtagonist.AverageFitness >= params.ProtagonistMinGenAvgFit {
		e.successfulGenerations++
	} else {
		e.successfulGenerations = 0
		return false
	}

	if meanPro >= params.ProtagonistMinGenAvgFit {
		e.successfulGenerationsByAvg++
	} else {
		e.successfulGenerationsByAvg = 0
		return false
	}

	// If number of successful Generations has been hit, break
	if e.successfulGenerations >= params.MinimumNumberOfSuccessfulGenerationBeforeTerminate {
		msg := fmt.Sprintf("COMPLETED CYCLE AT GENERATION: %d \n", generation.count)

		params.FinalGeneration = generation.count

		params.LoggingChan <- evolog.Logger{Type: evolog.LoggerGeneration, Message: msg,
			Timestamp: time.Now()}

		params.FinalGenerationReason = "BestClone"

		return true
	}

	if result.AllProtagonistAverageFitness >= params.MinimumGenerationMeanBeforeTerminate {
		msg := fmt.Sprintf("COMPLETED CYCLE AT GENERATION: %d \n", generation.count)

		params.FinalGeneration = generation.count

		params.LoggingChan <- evolog.Logger{Type: evolog.LoggerGeneration, Message: msg,
			Timestamp: time.Now()}

		params.FinalGenerationReason = "AvgGeneration"

		return true
	}

	return false
}

func WriteGenerationToLog(e *Engine, i int, elapsed time.Duration) {
	numGoroutine := runtime.NumGoroutine()

	msg := fmt.Sprintf("\nFile: %s\t | Spec: %s\t | Run: %d | Gen: (%d/%d) | TSz: %d | numG#: %d | Elapsed: %s",
		e.Parameters.ParamFile,
		e.Parameters.SpecParam.ExpressionParsed,
		e.Parameters.InternalCount,
		i+1,
		e.Parameters.MaxGenerations,
		e.Parameters.Strategies.NewTreeNTCount,
		numGoroutine,
		elapsed.String())

	e.Parameters.LoggingChan <- evolog.Logger{Type: evolog.LoggerGeneration, Message: msg, Timestamp: time.Now()}
}

func (e *Engine) ValidateGenerationTerminationMinimums() (minimumTopProtagonistThreshold int,
	minimumMeanProtagonistInGenerationThreshold int) {
	minimumTopProtagonistThreshold = int(e.Parameters.MinimumTopProtagonistMeanBeforeTerminate * float64(e.Parameters.
		MaxGenerations))
	minimumMeanProtagonistInGenerationThreshold = int(e.Parameters.MinimumGenerationMeanBeforeTerminate * float64(e.Parameters.
		MaxGenerations))

	if minimumTopProtagonistThreshold < MinAllowableGenerationsToTerminate {
		e.Parameters.MinimumTopProtagonistMeanBeforeTerminate = MinAllowableGenerationsToTerminate + 1
		//e.Parameters.LoggingChan <- evolog.Logger{
		//	Type:      evolog.LoggerGeneration,
		//	Message:   fmt.Sprintf("NOTE: Set MinimumTopProtagonistMeanBeforeTerminate: %f.2", e.Parameters.MinimumTopProtagonistMeanBeforeTerminate),
		//	Timestamp: time.Now(),
		//}
	}

	if minimumMeanProtagonistInGenerationThreshold < MinAllowableGenerationsToTerminate {
		e.Parameters.MinimumGenerationMeanBeforeTerminate = MinAllowableGenerationsToTerminate + 1
		//e.Parameters.LoggingChan <- evolog.Logger{
		//	Type:      evolog.LoggerGeneration,
		//	Message:   fmt.Sprintf("NOTE: Set MinimumGenerationMeanBeforeTerminate: %f.2", e.Parameters.MinimumGenerationMeanBeforeTerminate),
		//	Timestamp: time.Now(),
		//}
	}

	return minimumTopProtagonistThreshold, minimumMeanProtagonistInGenerationThreshold
}

// InitializeGenerations starts the first generation as a building block for the evolutionary process.
// It will embed the antagonists and protagonists created into its Generations slice at index [0]
func (e *Engine) InitializeGenerations(params EvolutionParams) (gen0 Generation, antagonists []Individual, protagonists []Individual, err error) {
	gen0 = Generation{
		ID:                                   1,
		Correlation:                          0,
		Covariance:                           0,
		AntagonistAverage:                    0,
		AntagonistStdDevOfAvgFitnessValues:   0,
		AntagonistVarianceOfAvgFitnessValues: 0,
		AntagonistAvgFitnessValuesOfEveryIndividual: make([]float64, 0),
		ProtagonistAverage:                          0,
		ProtagonistStdDevOfAvgFitnessValues:         0,
		ProtagonistVarianceOfAvgFitnessValues:       0,
		ProtagonistAvgFitnessOfEveryIndividual:      make([]float64, 0),
		engine:                                      e,
		isComplete:                                  false,
		hasParentSelectionHappened:                  false,
		hasSurvivorSelectionHappened:                false,
		count:                                       0,
		idAllocStart:                                0,
		idAllocOffset:                               0,
	}

	idAlloc := e.NewBatch(uint32(params.EachPopulationSize * 4))

	antagonists, protagonists, err = gen0.InitializePopulation(params, idAlloc)
	if err != nil {
		return Generation{}, nil, nil, err
	}

	gen0.ID = 1

	e.successfulGenerations = 0
	e.successfulGenerationsByAvg = 0
	e.minimumTopProtagonistThreshold, e.minimumMeanProtagonistInGenerationThreshold = e.ValidateGenerationTerminationMinimums()

	return gen0, antagonists, protagonists, err
}

// Todo Implement EvolutionProcess Validate
func (e *Engine) Validate() error {
	if e.Parameters.GenerationsCount < 1 {
		return fmt.Errorf("set number of generationCount by calling e.GenerationsCount(x)")
	}

	if e.Parameters.EachPopulationSize%4 != 0 {
		return fmt.Errorf("set number of EachPopulationSize to a number that is divisible by 2^x e.g. 8, 16, 32, 64, " +
			"128")
	}
	//if e.Parameters.SetEqualStrategyLength == true && e.Parameters.EqualStrategiesLength < 1 {
	//	return fmt.Errorf("cannot SetEqualStrategyLength to true and EqualStrategiesLength less than 1")
	//}
	if e.Parameters.StartIndividual == nil {
		return fmt.Errorf("start individual cannot have a nil Tree")
	}

	if e.Parameters.Spec == nil {
		return fmt.Errorf("spec cannot be nil")
	}

	if len(e.Parameters.Spec) < 1 {
		return fmt.Errorf("spec cannot be empty")
	}
	if e.Parameters.Selection.Survivor.SurvivorPercentage > 1 || e.Parameters.Selection.Survivor.
		SurvivorPercentage < 0 {
		return fmt.Errorf("SurvivorPercentage cannot be less than 0 or greater than 1. It is a percent value")
	}
	if e.Parameters.Selection.Parent.TournamentSize >= e.Parameters.EachPopulationSize {
		return fmt.Errorf("Tournament Size should not be greater than the population size.")
	}
	//err := e.StartIndividual.Validate()
	//if err != nil {
	//	return err
	//}

	if len(e.Parameters.Spec) < 3 {
		return fmt.Errorf("a small spec will hamper evolutionary accuracy")
	}
	return nil
}

func TruncShort(s []Strategy) string {
	sb := strings.Builder{}

	for _, str := range s {
		sb.WriteByte(str[0])
	}

	return sb.String()
}
