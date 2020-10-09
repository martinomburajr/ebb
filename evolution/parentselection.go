package evolution

import (
	"fmt"
	"math/rand"
)

/**
Selection is the stage of a genetic algorithm in which individual genomes are chosen from a population for later breeding (using the crossover operator).

	A generic selection procedure may be implemented as follows:

	1. The Fitness function is evaluated for each individual, providing Fitness values,
	which are then normalized. Normalization means dividing the Fitness value of each individual by the sum of all Fitness values, so that the sum of all resulting Fitness values equals 1.
	2. The population is sorted by descending Fitness values.
	3. Accumulated normalized Fitness values are computed: the accumulated Fitness value of an individual is the sum of its
	own Fitness value plus the Fitness values of all the previous individuals; the accumulated Fitness of the last individual should be 1, otherwise something went wrong in the normalization step.
	4. A random number R between 0 and 1 is chosen.
	5. The selected individual is the last one whose accumulated normalized value is greater than or equal to R.
	For a large number of individuals the above algorithm might be computationally quite demanding. A simpler and faster alternative uses the so-called stochastic acceptance.

	//https://en.wikipedia.org/wiki/Selection_(genetic_algorithm)
*/

const (
	ParentSelectionTournament = "ParentSelectionTournament" // ID for Tournament Selection
)

// TournamentSelection is a process whereby a random set of individuals from the population are selected,
// and the best in that sample succeed onto the next Generation
func TournamentSelection(population []Individual, tournamentSize int, idAlloc IDAllocator) ([]Individual, error) {
	if population == nil {
		return nil, fmt.Errorf("tournament population cannot be nil")
	}
	if len(population) < 1 {
		return nil, fmt.Errorf("tournament population cannot have len < 1")
	}
	if tournamentSize < 1 {
		return nil, fmt.Errorf("tournament size cannot be less than 1")
	}

	// do
	newPop := make([]Individual, len(population))

	for i := 0; i < len(population); i++ {
		randSelectedIndividuals := getNRandom(population, tournamentSize)

		fittest, err := tournamentSelect(randSelectedIndividuals)
		if err != nil {
			return nil, err
		}

		newID := int(idAlloc.idStart) + i

		if uint32(newID) > idAlloc.idEnd {
			panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID))
		}

		newPop[i] = fittest.Clone(newID)
	}

	return newPop, nil
}

// getNRandom selects  a random group of individiduals equivalent to the tournamentSize
func getNRandom(population []Individual, tournamentSize int) []Individual {
	newPop := make([]Individual, tournamentSize)

	for i := 0; i < tournamentSize; i++ {
		randIndex := rand.Intn(len(population))

		newPop[i] = population[randIndex]
	}

	return newPop
}

//tournamentSelect returns the fittest individual in a given tournament
func tournamentSelect(selectedIndividuals []Individual) (Individual, error) {
	fittest := selectedIndividuals[0]

	for i := range selectedIndividuals {
		if selectedIndividuals[i].AverageFitness > fittest.AverageFitness {
			individual := selectedIndividuals[i]

			fittest = individual
		}
	}

	return fittest, nil
}