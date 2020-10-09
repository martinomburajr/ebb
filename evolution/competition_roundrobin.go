package evolution

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

type RoundRobin struct {
	Engine            Engine
	CloneConsolidator BestIndividualMap
}

func (r *RoundRobin) Name(otherInfo string) string {
	return fmt.Sprintf("RR-%s", otherInfo)
}

func (r *RoundRobin) ClearMap() {
	for key := range r.CloneConsolidator {
		delete(r.CloneConsolidator, key)
	}
}

func (r *RoundRobin) Topology(currentGeneration Generation, params EvolutionParams) (currGen Generation, nextGen Generation, err error) {
	r.CloneConsolidator = NewBestIndividualMap()

	tournament, err := r.createTournament(currentGeneration.Antagonists, currentGeneration.Protagonists)
	if err != nil {
		return Generation{}, Generation{}, err
	}

	consolidatedAntagonists, consolidatedProtagonists, err := r.startTournament(tournament, params)
	if err != nil {
		return Generation{}, Generation{}, err
	}

	// Generation Individuals will already have been set in startTournament over here.
	currentGeneration.Antagonists = consolidatedAntagonists
	currentGeneration.Protagonists = consolidatedProtagonists

	currentGeneration.UpdateStatisticalFields()

	antagonistSurvivors, protagonistSurvivors := currentGeneration.ApplySelection()

	nextGen = Generation{
		ID:                           uint32(currentGeneration.count),
		Protagonists:                 protagonistSurvivors,
		Antagonists:                  antagonistSurvivors,
		engine:                       currentGeneration.engine,
		isComplete:                   false,
		hasParentSelectionHappened:   false,
		hasSurvivorSelectionHappened: false,
		count:                        currentGeneration.count + 1,
		idAllocOffset:                0,
	}

	nextGen.ID = uint32(currentGeneration.count + 1)
	nextGen.idAllocStart = nextGen.ID * uint32(params.IDSeparation)

	currGen = currentGeneration
	currGen.isComplete = true

	return currGen, nextGen, nil
}



// createTournament takes in the Generation individuals (
// protagonists and antagonists) and creates a set of uninitialized epochs
func (r *RoundRobin) createTournament(antagonists []Individual, protagonists []Individual) ([]RRCompetition, error) {
	if antagonists == nil {
		return nil, fmt.Errorf("antagonists cannot be nil in Generation")
	}
	if protagonists == nil {
		return nil, fmt.Errorf("protagonists cannot be nil in Generation")
	}

	lenAntagonists := len(antagonists)
	if lenAntagonists < 1 {
		return nil, fmt.Errorf("antagonists cannot be empty")
	}

	lenProtagonists := len(protagonists)
	if lenProtagonists < 1 {
		return nil, fmt.Errorf("protagonists cannot be empty")
	}

	competitionSize := lenAntagonists * lenProtagonists
	competitions := make([]RRCompetition, competitionSize)
	count := 0

	for i := 0; i < lenAntagonists; i++ {
		for j := 0; j < len(protagonists); j++ {

			competition := RRCompetition{
				id:          uint32(count),
				protagonist: &protagonists[j],
				antagonist:  &antagonists[i],
			}

			competitions[count] = competition

			count++
		}
	}

	return competitions, nil
}


func (r *RoundRobin) RecordCompetitionResult(competition RRCompetition, antagonistFitness, protagonistFitness, antagonistDelta, protagonistDelta float64) {
	r.CloneConsolidator.Check(competition.antagonist, antagonistFitness, antagonistDelta)
	r.CloneConsolidator.Check(competition.protagonist, protagonistFitness, protagonistDelta)
}

func (r *RoundRobin) ConsolidateIndividuals() (bestAntagonists, bestProtagonists []Individual) {
	eachPopulationSize := r.Engine.Parameters.EachPopulationSize
	bestAntagonists, bestProtagonists = r.CloneConsolidator.Deposit(eachPopulationSize)

	if len(bestAntagonists) != eachPopulationSize {
		panic(fmt.Sprintf("ConsolidateIndividuals: antagonists are not equal to size of eachPopulationSize %d", len(bestProtagonists)))
	}

	if len(bestProtagonists) != eachPopulationSize {
		panic(fmt.Sprintf("ConsolidateIndividuals: protagonists are not equal to size of eachPopulationSize %d", len(bestProtagonists)))
	}

	return bestAntagonists, bestProtagonists
}


// TODO - If performance is bad we can use pointers to generations
// runEpoch begins the run of a single epoch
func (r *RoundRobin) startTournament(competitions []RRCompetition, params EvolutionParams) (bestAntagonists, bestProtagonists []Individual, err error) {
	if competitions == nil {
		return nil, nil, fmt.Errorf("competitions have not been initialized | competitions is nil")
	}
	if len(competitions) < 1 {
		return nil, nil, fmt.Errorf("competitions slice is empty")
	}

	for i := 0; i < len(competitions); i++ {
		competition := competitions[i]

		antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta, err := competition.compete(params)
		if err != nil {
			return nil, nil, err
		}

		r.RecordCompetitionResult(competition, antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta)
	}

	bestAntagonists, bestProtagonists = r.ConsolidateIndividuals()

	return bestAntagonists, bestProtagonists, nil
}

// Compete gives protagonist and antagonists the chance to competeAntagonists against each other. A competition involves an epoch,
// that returns the Individuals of the epoch.
//func (r *RoundRobin) Compete(g *Generation) error {
//	setupEpochs, err := r.createTournament(g)
//	if err != nil {
//		return err
//	}
//
//	// Runs the epochs and returns completed epochs that contain Fitness information within each individual.
//	_, err = r.startTournament(g, setupEpochs)
//	if err != nil {
//		return err
//	}
//
//	// TODO Ensure Children of Antagonists are being created, i.e different IDs during crossover
//	// TODO use penalization when SPEc is 0
//
//	g.UpdateStatisticalFields()
//
//	return err
//}

func CoalesceFitnessStatistics(individual Individual) (fitnessToBeAppendedToGenerationAvgFitness float64) {
	deltaMean := stat.Mean(individual.Deltas, nil)
	mean, std := stat.MeanStdDev(individual.Fitness, nil)
	variance := stat.Variance(individual.Fitness, nil)

	individual.AverageFitness = mean
	individual.FitnessStdDev = std
	individual.FitnessVariance = variance
	individual.HasCalculatedFitness = true
	individual.HasAppliedStrategy = true
	individual.Age += 1
	individual.AverageDelta = deltaMean

	return mean
}
