package evolution

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
	"strings"
	"sync"
)

// TODO AGE
// TODO Calculate fitness average for GENERATIONS (seems off!)
type Generation struct {
	Mutex        *sync.Mutex
	ID           uint32
	Protagonists []Individual //Protagonists in a given Generation
	Antagonists  []Individual //Antagonists in a given Generation

	// Averages of all Antagonists and Protagonists in Generation
	Correlation float64
	Covariance  float64

	// AntagonistAverage is an average of AntagonistAvgFitnessValuesOfEveryIndividual
	AntagonistAverage                           float64
	AntagonistStdDevOfAvgFitnessValues          float64
	AntagonistVarianceOfAvgFitnessValues        float64
	AntagonistAvgFitnessValuesOfEveryIndividual []float64
	AntagonistsAvgAgeOfEveryIndividual          []float64
	AntagonistsAvgBirthGenOfEveryIndividual     []float64

	// ProtagonistAverage is an average of ProtagonistAvgFitnessOfEveryIndividual
	ProtagonistAverage                       float64
	ProtagonistStdDevOfAvgFitnessValues      float64
	ProtagonistVarianceOfAvgFitnessValues    float64
	ProtagonistAvgFitnessOfEveryIndividual   []float64
	ProtagonistsAvgAgeOfEveryIndividual      []float64
	ProtagonistsAvgBirthGenOfEveryIndividual []float64

	engine                       *Engine
	isComplete                   bool
	hasParentSelectionHappened   bool
	hasSurvivorSelectionHappened bool
	count                        int

	// Help with efficient ID allocation
	idAllocStart            uint32
	idAllocOffset           uint32
	AntagonistsAvgAge       float64
	AntagonistsAvgBirthGen  float64
	ProtagonistsAvgAge      float64
	ProtagonistsAvgBirthGen float64
}

func (g *Generation) CopyIndividuals(kind int) []Individual {
	individuals := make([]Individual, g.engine.Parameters.EachPopulationSize)

	if kind == IndividualAntagonist {
		copy(individuals, g.Antagonists)
	} else {
		copy(individuals, g.Protagonists)
	}

	return individuals
}

func CalculateGenerationIDBatch(kind int, idAllocStart uint32) (start uint32) {
	if kind == IndividualAntagonist {
		return idAllocStart + 1000
	} else {
		return idAllocStart + 6000
	}
}

func (e *Engine) NewBatch(need uint32) IDAllocator {
	oldOffset := e.idAllocOffset

	e.idAllocStart = oldOffset
	e.idAllocOffset = oldOffset + need
	return IDAllocator{oldOffset + 1, e.idAllocOffset}
}

// initializePopulation randomly creates a set of antagonists and protagonists
func (g *Generation) InitializePopulation(params EvolutionParams, idAlloc IDAllocator) (antagonists []Individual,
	protagonists []Individual, err error) {

	antagonists, protagonists = make([]Individual, params.EachPopulationSize), make([]Individual, params.EachPopulationSize)

	wg := sync.WaitGroup{}
	wg.Add(2)

	antagonistBatch := g.engine.NewBatch(uint32(params.EachPopulationSize * 2))
	protagonistBatch := g.engine.NewBatch(uint32(params.EachPopulationSize * 2))

	go func(wg *sync.WaitGroup, params EvolutionParams, antagonists []Individual, idAlloc IDAllocator) {
		defer wg.Done()

		antagonists1, err := NewRandomIndividuals(IndividualAntagonist, params, idAlloc)
		if err != nil {
			params.ErrorChan <- err
			return
		}

		copy(antagonists, antagonists1)

	}(&wg, params, antagonists, antagonistBatch)

	go func(wg *sync.WaitGroup, params EvolutionParams, protagonists []Individual, idAlloc IDAllocator) {
		defer wg.Done()

		protagonists1, err := NewRandomIndividuals(IndividualProtagonist, params, idAlloc)
		if err != nil {
			params.ErrorChan <- err
			return
		}

		copy(protagonists, protagonists1)

	}(&wg, params, protagonists, protagonistBatch)

	wg.Wait()

	g.Antagonists = antagonists
	g.Protagonists = protagonists

	return g.Antagonists, g.Protagonists, nil
}

// ApplySelection applies all 3 selection methods, parent,
// reproduction and survivor to return a set of survivor antagonist and protagonists
func (g *Generation) ApplySelection() (
	antagonistSurvivors []Individual, protagonistSurvivors []Individual) {

	antReproductionBatch := g.engine.NewBatch(uint32(g.engine.Parameters.EachPopulationSize * 5))
	proReproductionBatch := g.engine.NewBatch(uint32(g.engine.Parameters.EachPopulationSize * 5))
	antChildBatch := g.engine.NewBatch(uint32(g.engine.Parameters.EachPopulationSize * 3))
	proChildBatch := g.engine.NewBatch(uint32(g.engine.Parameters.EachPopulationSize * 3))

	antagonists := g.CopyIndividuals(IndividualAntagonist)
	protagonists := g.CopyIndividuals(IndividualProtagonist)

	params := g.engine.Parameters
	populationSize := params.EachPopulationSize
	tournamentSize := params.Selection.Parent.TournamentSize
	crossoverStrategy := params.Reproduction.CrossoverStrategy
	strategies := params.Strategies.Strategies
	probMutation := params.Reproduction.ProbabilityOfMutation
	errorChan := g.engine.Parameters.ErrorChan
	survivorPercentage := g.engine.Parameters.Selection.Survivor.SurvivorPercentage

	antagonistSurvivors = make([]Individual, populationSize)
	protagonistSurvivors = make([]Individual, populationSize)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup, individuals []Individual, survivors []Individual, genCount, tournamentSize, populationSize int, idAlloc, antChildBatch IDAllocator, crossoverStrategy string, strategies []Strategy, survivorPercentage, probMutation float64, errChan chan error) {
		newIndividuals, err := applySelection(antagonists, g.count, tournamentSize, populationSize, antReproductionBatch, antChildBatch, crossoverStrategy, strategies, survivorPercentage, probMutation)
		if err != nil {
			errChan <- err
		}

		copy(antagonistSurvivors, newIndividuals)
		wg.Done()
	}(&wg, antagonists, antagonistSurvivors, g.count, tournamentSize, populationSize, antReproductionBatch, antChildBatch, crossoverStrategy, strategies, survivorPercentage, probMutation, errorChan)

	go func(wg *sync.WaitGroup, individuals []Individual, survivors []Individual, genCount, tournamentSize, populationSize int, idAlloc, proChildBatch IDAllocator, crossoverStrategy string, strategies []Strategy, survivorPercentage, probMutation float64, errChan chan error) {
		newIndividuals2, err := applySelection(protagonists, g.count, tournamentSize, populationSize, proReproductionBatch, proChildBatch, crossoverStrategy, strategies, survivorPercentage, probMutation)
		if err != nil {
			errChan <- err
		}

		copy(protagonistSurvivors, newIndividuals2)
		wg.Done()
	}(&wg, protagonists, protagonistSurvivors, g.count, tournamentSize, populationSize, proReproductionBatch, proChildBatch, crossoverStrategy, strategies, survivorPercentage, probMutation, errorChan)

	wg.Wait()

	g.hasSurvivorSelectionHappened = true
	g.hasParentSelectionHappened = true
	return antagonistSurvivors, protagonistSurvivors
}

// applySelection uses a copy of the kind of individuals to pass in
func applySelection(individuals []Individual, genCount, tournamentSize, populationSize int, idAlloc, childBatch IDAllocator, crossoverStrategy string, strategies []Strategy, survivorPercentage, probMutation float64) ([]Individual, error) {
	winnerParents, err := applyParentSelection(individuals, tournamentSize, idAlloc)
	if err != nil {
		return nil, err
	}

	antSelectedChildren, err := applyCrossover(winnerParents, genCount, childBatch, populationSize, crossoverStrategy)
	if err != nil {
		return nil, err
	}

	outgoingParents, outgoingChildren, err := applyMutation(winnerParents, antSelectedChildren, strategies, probMutation)

	survivorSelection, err := applyRatiodSurvivorSelection(outgoingParents, outgoingChildren, survivorPercentage, populationSize)

	return survivorSelection, err
}

func (g *Generation) CleansePopulations(params EvolutionParams) {
	populationSize := params.EachPopulationSize
	//antBatch := g.engine.NewBatch(uint32(populationSize * 2 * (g.count+1)))
	//proBatch := g.engine.NewBatch(uint32(populationSize * 2 * (g.count+1)))

	cleanAntagonists := make([]Individual, populationSize)
	cleanProtagonists := make([]Individual, populationSize)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup, individualsToClean, cleanIndividuals []Individual, errChan chan error) {
		defer wg.Done()

		individuals, err := CleansePopulation(individualsToClean)
		if err != nil {
			errChan <- err
		}

		copy(cleanIndividuals, individuals)
	}(&wg, g.Antagonists, cleanAntagonists, params.ErrorChan)

	go func(wg *sync.WaitGroup, individualsToClean, cleanIndividuals []Individual, errChan chan error) {
		defer wg.Done()

		individuals, err := CleansePopulation(individualsToClean)
		if err != nil {
			errChan <- err
		}

		copy(cleanIndividuals, individuals)
	}(&wg, g.Protagonists, cleanProtagonists, params.ErrorChan)

	wg.Wait()

	g.Protagonists = cleanProtagonists
	g.Antagonists = cleanAntagonists
}

// ApplyParentSelection takes in a given Generation and returns a set of individuals once the preselected parent
// selection Strategy has been applied to the Generation.
// These individuals are ready to be taken to either a new Generation or preferably through survivor selection in the
// case you do not isEqual the population to grow in size.
func (g *Generation) ApplyParentSelection(currentPopulation []Individual) ([]Individual, error) {
	idAllocator := g.engine.NewBatch(uint32(len(currentPopulation) * 2))
	return applyParentSelection(currentPopulation, g.engine.Parameters.Selection.Parent.TournamentSize, idAllocator)
}

func applyParentSelection(currentPopulation []Individual, tournamentSize int, idAlloc IDAllocator) ([]Individual, error) {
	return TournamentSelection(currentPopulation, tournamentSize, idAlloc)
}

// ApplySurvivorSelection applies the preselected survivor selection Strategy.
// It DOES NOT check to see if the parent selection has already been applied,
// as in some cases evolutionary programs may choose to run without the parent selection phase.
// The onus is on the evolutionary architect to keep this consideration in mind.
func (g *Generation) ApplyReproduction(parents []Individual) (outgoingParents []Individual, outgoingChildren []Individual, err error) {
	params := g.engine.Parameters

	idAllocator := g.engine.NewBatch(uint32(len(parents) * 2))

	children, err := applyCrossover(parents, g.count, idAllocator, params.EachPopulationSize, params.Reproduction.CrossoverStrategy)
	if err != nil {
		return nil, nil, err
	}

	return applyMutation(parents, children, params.Strategies.Strategies, params.Reproduction.ProbabilityOfMutation)
}

func applyCrossover(incomingParents []Individual, genCount int, idAlloc IDAllocator, populationSize int, crossoverStrategy string) (children []Individual, err error) {
	children = make([]Individual, populationSize)

	switch crossoverStrategy {
	case CrossoverSinglePoint:
		for i := 0; i < len(incomingParents); i += 2 {
			newID1 := int(idAlloc.idStart) + i
			newID2 := int(idAlloc.idStart) + i + 1

			if uint32(newID1) > idAlloc.idEnd {
				panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID1))
			}
			if uint32(newID2) > idAlloc.idEnd {
				panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID2))
			}

			child1, child2, err := SinglePointCrossover(&incomingParents[i], &incomingParents[i+1], newID1, newID2)
			if err != nil {
				return nil, err
			}

			child1.BirthGen = genCount
			child2.BirthGen = genCount
			child1.Age = 0
			child2.Age = 0

			children[i] = child1
			children[i+1] = child2
		}

	case CrossoverUniform:
		for i := 0; i < len(incomingParents); i += 2 {

			newID1 := int(idAlloc.idStart) + i
			newID2 := int(idAlloc.idStart) + i + 1

			if uint32(newID1) > idAlloc.idEnd {
				panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID1))
			}
			if uint32(newID2) > idAlloc.idEnd {
				panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID2))
			}

			child1, child2, err := UniformCrossover(incomingParents[i], incomingParents[i+1], newID1, newID2)
			if err != nil {
				return nil, err
			}

			child1.BirthGen = genCount
			child2.BirthGen = genCount
			child1.Age = 0
			child2.Age = 0

			children[i] = child1
			children[i+1] = child2
		}
	default:
		return nil, fmt.Errorf("no appropriate FixedPointCrossover operation was selected")
	}

	return children, nil
}

// ApplySurvivorSelection applies the preselected survivor selection Strategy.
// It DOES NOT check to see if the parent selection has already been applied,
// as in some cases evolutionary programs may choose to run without the parent selection phase.
// The onus is on the evolutionary architect to keep this consideration in mind.
func (g *Generation) ApplySurvivorSelection(outgoingParents []Individual,
	children []Individual) ([]Individual, error) {

	switch g.engine.Parameters.Selection.Survivor.Type {
	case SurvivorSelectionHalfAndHalf:
		return HalfAndHalfSurvivorSelection(outgoingParents, children, g.engine.Parameters.Selection.Survivor.SurvivorPercentage, g.engine.Parameters.EachPopulationSize)
	case SurvivorSelectionParentVsChild:
		return ParentVsChildSurvivorSelection(outgoingParents, children, g.engine.Parameters)
	default:
		return nil, fmt.Errorf("invalid survivor selection selected")
	}
}

// GenerateRandom creates a a random set of individuals based on the parameters passed into the
// evolution engine. To pass a tree to an individual pass it via the formal parameters and not through the evolution
// engine
// parameter section
// Antagonists are by default
// set with the StartIndividuals Program as their own
// program.
func NewRandomIndividuals(kind int, params EvolutionParams, idAlloc IDAllocator) ([]Individual, error) {
	if params.EachPopulationSize < 1 {
		return nil, fmt.Errorf("number should at least be 1")
	}

	strategyLen := len(params.Strategies.Strategies)
	if strategyLen < 1 {
		return nil, fmt.Errorf(" maxNumberOfStrategies should at least be 1")
	}

	individuals := make([]Individual, params.EachPopulationSize)

	for i := 0; i < params.EachPopulationSize; i++ {
		var randomStrategies []Strategy

		randomStrategies = GenerateRandomStrategy(params.Strategies.NumStrategiesToUse,
			params.Strategies.Strategies)

		var individual Individual

		newID := idAlloc.idStart + uint32(i)

		if newID > idAlloc.idEnd {
			panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID))
		}

		if kind == IndividualAntagonist {
			individual = Individual{
				Kind:            kind,
				ID:              newID,
				Strategy:        randomStrategies,
				Fitness:         make([]float64, 0),
				Program:         nil,
				BestFitness:     math.MinInt32,
				AverageFitness:  math.MinInt32,
				BestDelta:       math.MinInt32,
				FitnessVariance: math.MinInt32,
				FitnessStdDev:   math.MinInt32,
			}
		} else {
			individual = Individual{
				Kind:            kind,
				ID:              newID,
				Strategy:        randomStrategies,
				Fitness:         make([]float64, 0),
				Program:         nil,
				BestFitness:     math.MinInt32,
				AverageFitness:  math.MinInt32,
				BestDelta:       math.MinInt32,
				FitnessVariance: math.MinInt32,
				FitnessStdDev:   math.MinInt32,
			}
		}

		individuals[i] = individual
	}

	return individuals, nil
}

// Runs the generational statistics
func (g *Generation) RunGenerationStatistics() (result GenerationResult) {

	correlation := stat.Correlation(g.AntagonistAvgFitnessValuesOfEveryIndividual,
		g.ProtagonistAvgFitnessOfEveryIndividual, nil)
	covariance := stat.Covariance(g.AntagonistAvgFitnessValuesOfEveryIndividual,
		g.ProtagonistAvgFitnessOfEveryIndividual, nil)

	antMean, antStd := stat.MeanStdDev(g.AntagonistAvgFitnessValuesOfEveryIndividual, nil)
	proMean, proStd := stat.MeanStdDev(g.ProtagonistAvgFitnessOfEveryIndividual, nil)

	antVar := stat.Variance(g.AntagonistAvgFitnessValuesOfEveryIndividual, nil)
	proVar := stat.Variance(g.ProtagonistAvgFitnessOfEveryIndividual, nil)

	result.AllAntagonistAverageFitness = antMean
	result.AntagonistStdDev = antStd
	result.AntagonistVariance = antVar
	result.AllProtagonistAverageFitness = proMean
	result.ProtagonistStdDev = proStd
	result.ProtagonistVariance = proVar
	result.Correlation = correlation
	result.Covariance = covariance
	result.AntagonistAvgAge = g.AntagonistsAvgAge
	result.AntagonistAvgBirthGen = g.AntagonistsAvgBirthGen
	result.ProtagonistAvgAge = g.ProtagonistsAvgAge
	result.ProtagonistAvgBirthGen = g.ProtagonistsAvgBirthGen

	result.BestAntagonist = g.BestAntagonist()
	result.BestProtagonist = g.BestProtagonist()

	//statsString := result.ToString()

	//g.Parameters.LoggingChan <- evolog.Logger{Timestamp: time.Now(), Type: evolog.LoggerGeneration, Message: statsString}

	return result
}

//UpdateStatisticalFields uses the partially populated individuals in the generation and computes their final values
//as well as populates some of the basic generational statistics. This MUST be called in every topology towards the end before CalculateGenerationalRest
func (g *Generation) UpdateStatisticalFields() {
	for i := 0; i < len(g.Protagonists); i++ {
		// Populate Antagonist Fitness Values
		antagonist := g.Antagonists[i]
		g.AntagonistAvgFitnessValuesOfEveryIndividual = append(g.AntagonistAvgFitnessValuesOfEveryIndividual, antagonist.AverageFitness)
		g.AntagonistsAvgAgeOfEveryIndividual = append(g.AntagonistsAvgAgeOfEveryIndividual, float64(antagonist.Age))
		g.AntagonistsAvgBirthGenOfEveryIndividual = append(g.AntagonistsAvgBirthGenOfEveryIndividual, float64(antagonist.BirthGen))

		// Populate Protagonists Fitness Values
		protagonist := g.Protagonists[i]
		g.ProtagonistAvgFitnessOfEveryIndividual = append(g.ProtagonistAvgFitnessOfEveryIndividual, protagonist.AverageFitness)
		g.ProtagonistsAvgAgeOfEveryIndividual = append(g.AntagonistsAvgAgeOfEveryIndividual, float64(protagonist.Age))
		g.ProtagonistsAvgBirthGenOfEveryIndividual = append(g.AntagonistsAvgBirthGenOfEveryIndividual, float64(protagonist.BirthGen))
	}

	g.AntagonistStdDevOfAvgFitnessValues = stat.StdDev(g.AntagonistAvgFitnessValuesOfEveryIndividual, nil)
	g.AntagonistVarianceOfAvgFitnessValues = stat.Variance(g.AntagonistAvgFitnessValuesOfEveryIndividual, nil)
	g.AntagonistAverage = stat.Mean(g.AntagonistAvgFitnessValuesOfEveryIndividual, nil)
	g.AntagonistsAvgAge = stat.Mean(g.AntagonistsAvgAgeOfEveryIndividual, nil)
	g.AntagonistsAvgBirthGen = stat.Mean(g.AntagonistsAvgBirthGenOfEveryIndividual, nil)

	g.ProtagonistStdDevOfAvgFitnessValues = stat.StdDev(g.ProtagonistAvgFitnessOfEveryIndividual, nil)
	g.ProtagonistVarianceOfAvgFitnessValues = stat.Variance(g.ProtagonistAvgFitnessOfEveryIndividual, nil)
	g.ProtagonistAverage = stat.Variance(g.ProtagonistAvgFitnessOfEveryIndividual, nil)
	g.ProtagonistsAvgAge = stat.Mean(g.ProtagonistsAvgAgeOfEveryIndividual, nil)
	g.ProtagonistsAvgBirthGen = stat.Mean(g.ProtagonistsAvgBirthGenOfEveryIndividual, nil)

	g.Correlation = stat.Correlation(g.AntagonistAvgFitnessValuesOfEveryIndividual, g.ProtagonistAvgFitnessOfEveryIndividual, nil)
	g.Covariance = stat.Covariance(g.AntagonistAvgFitnessValuesOfEveryIndividual, g.ProtagonistAvgFitnessOfEveryIndividual, nil)
}

func (g *Generation) PrintIndividuals() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("\n\n--------------------------------GENERATION: %d\t\t------------------------------\n", g.ID))

	sb.WriteString("\nEQUATION: \t\t")
	sb.WriteString(g.engine.Parameters.StartIndividual.ToMathematicalString())
	sb.WriteRune('\n')

	sb.WriteString("ANTAGONISTS -------------")
	sb.WriteRune('\n')
	for i := 0; i < len(g.Antagonists); i++ {
		builder := g.Antagonists[i].ToString()
		sb.WriteString(builder.String())
		sb.WriteRune('\n')
	}

	sb.WriteString("\nPROTAGONISTS -------------")
	sb.WriteRune('\n')

	for i := 0; i < len(g.Protagonists); i++ {
		builder := g.Protagonists[i].ToString()
		sb.WriteString(builder.String())
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (g *Generation) BestAntagonist() Individual {
	best := Individual{AverageFitness: math.MinInt64}

	for i := 0; i < len(g.Antagonists); i++ {
		currAnt := g.Antagonists[i]

		if currAnt.AverageFitness > best.AverageFitness {
			best = currAnt
		}
	}

	return best
}

func (g *Generation) BestProtagonist() Individual {
	best := Individual{AverageFitness: math.MinInt64}

	for i := 0; i < len(g.Protagonists); i++ {
		currAnt := g.Protagonists[i]

		if currAnt.AverageFitness > best.AverageFitness {
			best = currAnt
		}
	}

	return best
}

func (g *GenerationResult) ToString() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("\n%d\n", g.ID))
	sb.WriteString(fmt.Sprintf("CorrelationInGeneration: : %.2f\n", g.Correlation))
	sb.WriteString(fmt.Sprintf("CovarianceInGeneration: : %.2f\n", g.Covariance))
	sb.WriteString("_______________________________\n")
	sb.WriteString(fmt.Sprintf("AntagonistStdDevInGeneration : %.2f\n", g.AntagonistStdDev))
	sb.WriteString(fmt.Sprintf("AntagonistAvg : %.2f\n", g.AllAntagonistAverageFitness))
	sb.WriteString(fmt.Sprintf("AntagonistStdDevInGeneration : %.2f\n", g.AntagonistStdDev))
	sb.WriteString(fmt.Sprintf("AntagonistVarianceInGeneration : %.2f\n", g.AntagonistVariance))
	sb.WriteString("<===================================>\n")
	sb.WriteString(fmt.Sprintf("ProtagonistAverageInGeneration : %.2f\n", g.AllProtagonistAverageFitness))
	sb.WriteString(fmt.Sprintf("ProtagonistStdDevInGeneration : %.2f\n", g.ProtagonistStdDev))
	sb.WriteString(fmt.Sprintf("ProtagonistVarianceInGeneration : %.2f\n", g.ProtagonistVariance))

	return sb.String()
}
