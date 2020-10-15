package evolution

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type SingleEliminationTournament struct {
	Engine            Engine
	BestIndividualMap BestIndividualMap
}

func (s *SingleEliminationTournament) Topology(currentGeneration Generation,
	params EvolutionParams) (currGen Generation, nextGen Generation, err error) {

	s.BestIndividualMap = NewBestIndividualMap()

	setNoOfTournaments := params.Topology.SETNoOfTournaments

	if setNoOfTournaments == 0 {
		setNoOfTournaments = 1
	}

	fittestAntagonists := s.antagonistTopology(setNoOfTournaments, currentGeneration.Antagonists, params)
	s.protagonistTopology(setNoOfTournaments, fittestAntagonists, currentGeneration.Protagonists, params)

	consolidatedAntagonists, consolidatedProtagonists := s.ConsolidateIndividuals()

	currentGeneration.Antagonists = consolidatedAntagonists
	currentGeneration.Protagonists = consolidatedProtagonists

	// Individuals should already be set via pointer to currGen
	currentGeneration.UpdateStatisticalFields()

	antagonistSurvivors, protagonistSurvivors := currentGeneration.ApplySelection()

	newGeneration := Generation{
		ID:                           uint32(currentGeneration.count + 1),
		Protagonists:                 protagonistSurvivors,
		Antagonists:                  antagonistSurvivors,
		engine:                       currentGeneration.engine,
		isComplete:                   true,
		hasParentSelectionHappened:   true,
		hasSurvivorSelectionHappened: true,
		count:                        currentGeneration.count + 1,
	}

	nextGen.ID = uint32(currentGeneration.count + 1)
	nextGen.idAllocStart = nextGen.ID * uint32(params.IDSeparation)

	currGen = currentGeneration
	currGen.isComplete = true

	return currGen, newGeneration, nil
}

func (s *SingleEliminationTournament) createTournament(individuals []Individual) (bracket SETBracket, err error) {
	if len(individuals) < 1 {
		return nil, fmt.Errorf("createTournamentBrackets | input individuals cannot be empty")
	}

	if len(individuals) == 0 {
		return nil, fmt.Errorf("createTournamentBrackets | input individuals cannot be null")
	}

	rand.Shuffle(len(individuals), func(i, j int) {
		individuals[i], individuals[j] = individuals[j], individuals[i]
	})

	numberOfCompetitions := len(individuals) / 2
	bracket = make([]SETCompetition, numberOfCompetitions)

	for i := 0; i < numberOfCompetitions; i++ {
		bracket[i] = SETCompetition{
			individualA: &individuals[2*i],
			individualB: &individuals[(2*i)+1],
		}
	}

	return bracket, nil
}

func (s *SingleEliminationTournament) startTournamentAntagonists(tournament SETBracket, params EvolutionParams) (bestIndividual *Individual, err error) {
	if tournament == nil {
		return nil, fmt.Errorf("tournament have not been initialized | tournament is nil")
	}
	if len(tournament) < 1 {
		return nil, fmt.Errorf("tournament map is empty")
	}

	for {
		winners := make([]*Individual, 0)

		for i := 0; i < len(tournament); i++ {
			competition := tournament[i]

			// no need to clone bestAntagonist as it is already cloned by the caller
			individualAFitness, individualADelta, individualBFitness, individualBDelta, err := competition.competeAntagonists(params)
			if err != nil {
				return nil, err
			}

			s.RecordCompetitionResult(competition, individualAFitness, individualBFitness, individualADelta, individualBDelta)

			// Remove the loser from the brackets
			if individualAFitness >= individualBFitness {
				winners = append(winners, competition.individualA)
			} else {
				winners = append(winners, competition.individualB)
			}
		}

		if len(winners) == 1 {
			return winners[0], nil
		}

		tournament = NewBracket(winners)
	}
}

func (s *SingleEliminationTournament) startTournamentProtagonists(tournament SETBracket, bestAntagonist BinaryTree, params EvolutionParams) (bestIndividual *Individual, err error) {
	if tournament == nil {
		return nil, fmt.Errorf("tournament have not been initialized | tournament is nil")
	}
	if len(tournament) < 1 {
		return nil, fmt.Errorf("tournament map is empty")
	}

	for {
		winners := make([]*Individual, 0)

		for i := 0; i < len(tournament); i++ {
			competition := tournament[i]

			// no need to clone bestAntagonist as it is already cloned by the caller
			individualAFitness, individualADelta, individualBFitness, individualBDelta, err := competition.competeProtagonist(bestAntagonist, params)
			if err != nil {
				return nil, err
			}

			s.RecordCompetitionResult(competition, individualAFitness, individualBFitness, individualADelta, individualBDelta)

			// Remove the loser from the brackets
			if individualAFitness >= individualBFitness {
				winners = append(winners, competition.individualA)
			} else {
				winners = append(winners, competition.individualB)
			}
		}

		if len(winners) == 1 {
			return winners[0], nil
		}

		tournament = NewBracket(winners)
	}
}

func (s *SingleEliminationTournament) RecordCompetitionResult(competition SETCompetition, individualAFitness, individualBFitness, individualADelta, individualBDelta float64) {
	s.BestIndividualMap.Check(competition.individualA, individualAFitness, individualADelta)
	s.BestIndividualMap.Check(competition.individualB, individualBFitness, individualBDelta)
}

func (s *SingleEliminationTournament) ConsolidateIndividuals() (consolidatedAntagonists []Individual, consolidatedProtagonists []Individual) {
	eachPopulationSize := s.Engine.Parameters.EachPopulationSize
	consolidatedAntagonists, consolidatedProtagonists = s.BestIndividualMap.Deposit(eachPopulationSize)

	return consolidatedAntagonists, consolidatedProtagonists
}

func (s *SingleEliminationTournament) antagonistTopology(numberOfTournaments int, individuals []Individual,
	params EvolutionParams) []*Individual {

	fittestAntagonists := make([]*Individual, numberOfTournaments)

	// TODO - Watch concurrency race conditions when running multople tournaments simultaneously
	//wg := sync.WaitGroup{}

	for i := 0; i < numberOfTournaments; i++ {
		//go func(wgAntagonist *sync.WaitGroup, individuals []Individual, fittestAntagonists []*Individual, i int) {
		//	wgAntagonist.Add(1)
		//	defer wgAntagonist.Done()

		tournamentLedger, err := s.createTournament(individuals)
		if err != nil {
			//params.ErrorChan <- err
			return nil
		}

		topAntagonist, err := s.startTournamentAntagonists(tournamentLedger, params)
		if err != nil {
			//params.ErrorChan <- err
			return nil
		}

		// TODO - Check Error
		//currentGeneration.Mutex.Lock()
		fittestAntagonists[i] = topAntagonist
		//currentGeneration.Mutex.Unlock()

		//}(&wg, individuals, fittestAntagonists, i)
	}

	//wg.Wait()

	if len(fittestAntagonists) < params.EachPopulationSize {
		diff := params.EachPopulationSize - len(fittestAntagonists)
		perm := rand.Perm(diff)

		for i := 0; i < diff; i++ {
			fittestAntagonists = append(fittestAntagonists, &individuals[perm[i]])
		}
	}

	return fittestAntagonists
}

func (s *SingleEliminationTournament) protagonistTopology(numberOfTournaments int, bestAntagonists []*Individual, individuals []Individual,
	params EvolutionParams) []*Individual {

	fittestIndividuals := make([]*Individual, numberOfTournaments)

	// TODO - Watch concurrency race conditions when running multople tournaments simultaneously
	wg := sync.WaitGroup{}

	for i := 0; i < numberOfTournaments; i++ {
		bestAntagonist := bestAntagonists[i]

		//go func(wgAntagonist *sync.WaitGroup, individuals []Individual, fitesstIndividuals []*Individual, bestAntagonist *Individual, i int) {
		//	wgAntagonist.Add(1)
		//	defer wgAntagonist.Done()

		tournamentLedger, err := s.createTournament(individuals)
		if err != nil {
			params.ErrorChan <- err
			//return
		}

		topProtagonist, err := s.startTournamentProtagonists(tournamentLedger, bestAntagonist.Program.Clone(), params)
		if err != nil {
			params.ErrorChan <- err
			//return
		}

		// TODO - Check Error
		//currentGeneration.Mutex.Lock()
		fittestIndividuals[i] = topProtagonist
		//currentGeneration.Mutex.Unlock()

		//}(&wg, individuals, fittestIndividuals, bestAntagonist, i)
	}

	wg.Wait()

	if len(fittestIndividuals) < params.EachPopulationSize {
		diff := params.EachPopulationSize - len(fittestIndividuals)
		perm := rand.Perm(diff)

		for i := 0; i < diff; i++ {
			fittestIndividuals = append(fittestIndividuals, &individuals[perm[i]])
		}
	}

	return fittestIndividuals
}

type SETCompetition struct {
	individualA *Individual
	individualB *Individual
}

type SETBracket []SETCompetition

func NewBracket(individuals []*Individual) SETBracket {
	numCompetitions := len(individuals) / 2
	s := make([]SETCompetition, numCompetitions)

	counter := 0
	for i := 0; i < len(individuals); i += 2 {
		s[counter].individualA = individuals[i]
		s[counter].individualB = individuals[i+1]
		counter++
	}

	return s
}

func (k *SETCompetition) competeAntagonists(params EvolutionParams) (antagonistFitness, antagonistFitnessDelta, protagonistFitness, protagonistFitnessDelta float64, err error) {
	inf := math.Inf(-1)
	var individualAFitness, individualADelta, individualBFitness, individualBDelta = inf, inf, inf, inf

	individualA := k.individualA
	individualB := k.individualB

	err = individualA.ApplyAntagonistStrategy(params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	err = individualB.ApplyAntagonistStrategy(params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	individualAFitness, individualADelta, err = individualA.CalculateAntagonistThresholdedFitness(
		params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	individualBFitness, individualBDelta, err = individualB.
		CalculateAntagonistThresholdedFitness(params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	//individualA.Fitness = append(individualA.Fitness, individualAFitness)
	//individualA.Deltas = append(individualA.Deltas, individualADelta)
	////individualA.Parent.Fitness = append(individualA.Parent.Fitness, individualAFitness)
	////individualA.Parent.Deltas = append(individualA.Parent.Deltas, individualADelta)
	//
	//individualB.Fitness = append(individualB.Fitness, individualBFitness)
	//individualB.Deltas = append(individualB.Deltas, individualBDelta)
	////individualB.Parent.Fitness = append(individualB.Parent.Fitness, individualBFitness)
	////individualB.Parent.Deltas = append(individualB.Parent.Deltas, individualBDelta)

	return individualAFitness, individualADelta, individualBFitness, individualBDelta, nil
}

func (k *SETCompetition) competeProtagonist(bestAntagonistTree BinaryTree, params EvolutionParams) (antagonistFitness, antagonistFitnessDelta, protagonistFitness, protagonistFitnessDelta float64, err error) {
	inf := math.Inf(-1)
	var individualAFitness, individualADelta, individualBFitness, individualBDelta = inf, inf, inf, inf

	individualA := k.individualA
	individualB := k.individualB

	err = individualA.ApplyProtagonistStrategy(bestAntagonistTree, params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	err = individualB.ApplyProtagonistStrategy(bestAntagonistTree, params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	individualAFitness, individualADelta, err = individualA.CalculateProtagonistThresholdedFitness(params)
	if err != nil {
		return inf, inf, inf, inf, err
	}
	individualBFitness, individualBDelta, err = individualB.CalculateProtagonistThresholdedFitness(params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	//individualA.Fitness = append(individualA.Fitness, individualAFitness)
	//individualA.Deltas = append(individualA.Deltas, individualADelta)
	//individualA.Parent.Fitness = append(individualA.Parent.Fitness, individualAFitness)
	//individualA.Parent.Deltas = append(individualA.Parent.Deltas, individualADelta)
	//
	//individualB.Fitness = append(individualB.Fitness, individualBFitness)
	//individualB.Deltas = append(individualB.Deltas, individualBDelta)
	//individualB.Parent.Fitness = append(individualB.Parent.Fitness, individualBFitness)
	//individualB.Parent.Deltas = append(individualB.Parent.Deltas, individualBDelta)

	return individualAFitness, individualADelta, individualBFitness, individualBDelta, nil
}

func (s *SingleEliminationTournament) Name(otherInfo string) string {
	return fmt.Sprintf("SET-%s", otherInfo)
}

func (s *SingleEliminationTournament) ClearMap() {
	for key := range s.BestIndividualMap {
		delete(s.BestIndividualMap, key)
	}
}
