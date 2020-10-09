package evolution

import (
	"fmt"
	"math"
)

type KRandom struct {
	Engine             Engine
	BestIndividualMap BestIndividualMap
}

func (k *KRandom) Name(otherInfo string) string {
	return fmt.Sprintf("KRT-%s", otherInfo)
}

func (k *KRandom) ClearMap() {
	for key := range k.BestIndividualMap {
		delete(k.BestIndividualMap, key)
	}
}

func (k *KRandom) Topology(currentGeneration Generation, params EvolutionParams) (currGen Generation, nextGen Generation, err error) {
	k.BestIndividualMap = NewBestIndividualMap()

	tournamentLedger, err := k.createTournament(currentGeneration.Antagonists, currentGeneration.Protagonists)
	if err != nil {
		return Generation{}, Generation{}, err
	}

	bestAntagonists, bestProtagonists, err := k.startTournament(tournamentLedger, params)
	if err != nil {
		return Generation{}, Generation{}, err
	}

	// Generation Individuals will already have been set in startTournament over here.
	currentGeneration.Antagonists = bestAntagonists
	currentGeneration.Protagonists = bestProtagonists

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

// tournamentCreator will pit protagonist against parasite (competitors) based on the number of tournaments to play. It will
// return a map of protagonist(indes) -> slice of competitors to play(indices). Note that we are returning indices, so if the
// competitor count is 4, a return sample
func tournamentCreator(competitorCount int, tournaments int) map[int][]int {
	competitions := make(map[int][]int, competitorCount)

	jCounter := 0
	for i := 0; i < competitorCount; i++ {
		competitions[i] = make([]int, tournaments)

		for j := 0; j < tournaments; j++ {
			if jCounter == competitorCount {
				jCounter = 0
			}

			competitions[i][j] = jCounter
			jCounter++
		}
	}

	return competitions
}

type Compete interface {
	compete(params EvolutionParams) (antagonist Individual, protagonist Individual)
}

type KRandomCompetitions struct {
	host        *Individual
	challengers []*Individual
}

type Competition struct {
	id uint32
	protagonist *Individual
	antagonist  *Individual
}

func (k *Competition) compete(params EvolutionParams) (antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta float64, err error) {
	err = k.antagonist.ApplyAntagonistStrategy(params)
	inf := math.Inf(-1)

	if err != nil {
		return inf, inf, inf, inf, err
	}

	err = k.protagonist.ApplyProtagonistStrategy(k.antagonist.Program, params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta = ThresholdedRatioFitness(params.Spec, k.antagonist.Program, k.protagonist.Program)

	// TODO - PUNISH DIVISIONS BY ZERO!
	if math.IsNaN(antagonistFitness) {
		print()
	}

	if math.IsNaN(protagonistFitness) {
		print()
	}

	return antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta, err
}

type IDAllocator struct {
	idStart uint32
	idEnd   uint32
}

func (k *KRandom) createTournament(antagonists []Individual, protagonists []Individual) (tournamentLedger []KRandomCompetitions, err error) {
	params := k.Engine.Parameters

	tournamentSheet := tournamentCreator(params.EachPopulationSize, params.Topology.KRandomK)

	competitions := make([]KRandomCompetitions, params.EachPopulationSize)

	i := 0

	for key, v := range tournamentSheet {
		competitions[i].host = &protagonists[key]
		competitions[i].challengers = make([]*Individual, params.Topology.KRandomK)

		for j := 0; j < params.Topology.KRandomK; j++ {
			individual := antagonists[v[j]]
			competitions[i].challengers[j] = &individual
		}

		i++
	}

	return competitions, nil
}

func (k *KRandom) contains(individual Individual, individuals []Individual) (bool, int) {
	for i := 0; i < len(individuals); i++ {
		if individual.ID == individuals[i].ID {
			return true, i
		}
	}

	return false, -1
}

func (k *KRandom) RecordCompetitionResult(competition Competition, antagonistFitness, protagonistFitness, antagonistDelta, protagonistDelta float64) {
	k.BestIndividualMap.Check(competition.antagonist, antagonistFitness, antagonistDelta)
	k.BestIndividualMap.Check(competition.protagonist, protagonistFitness, protagonistDelta)
}

func (k *KRandom) ConsolidateIndividuals() (bestAntagonists, bestProtagonists []Individual) {
	eachPopulationSize := k.Engine.Parameters.EachPopulationSize
	bestAntagonists, bestProtagonists = k.BestIndividualMap.Deposit(eachPopulationSize)

	if len(bestAntagonists) != eachPopulationSize {
		panic(fmt.Sprintf("ConsolidateIndividuals: antagonists are not equal to size of eachPopulationSize %d", len(bestProtagonists)))
	}

	if len(bestProtagonists) != eachPopulationSize {
		panic(fmt.Sprintf("ConsolidateIndividuals: protagonists are not equal to size of eachPopulationSize %d", len(bestProtagonists)))
	}

	return bestAntagonists, bestProtagonists
}

func (k *KRandom) startTournament(competitions []KRandomCompetitions, params EvolutionParams) (bestAntagonists, bestProtagonists []Individual, err error) {
	if competitions == nil {
		return nil, nil, fmt.Errorf("competitions have not been initialized | competitions is nil")
	}
	if len(competitions) < 1 {
		return nil, nil, fmt.Errorf("competitions slice is empty")
	}

	for i := 0; i < len(competitions); i++ {

		for j := 0; j < len(competitions[i].challengers); j++ {
			competition := Competition{protagonist: competitions[i].host, antagonist: competitions[i].challengers[j]}

			antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta, err := competition.compete(params)
			if err != nil {
				return nil, nil, err
			}

			k.RecordCompetitionResult(competition, antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta)
		}
	}

	bestAntagonists, bestProtagonists = k.ConsolidateIndividuals()

	return bestAntagonists, bestProtagonists, nil
}
