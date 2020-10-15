package evolution

import (
	"math"
)

// CleansePopulation removes the trees from the population and resets some of their information
func CleansePopulation(individuals []Individual) ([]Individual, error) {
	for i := range individuals {
		individual := individuals[i]

		individual.HasCalculatedFitness = false
		individual.HasAppliedStrategy = false
		individual.AverageFitness = math.MinInt32
		individual.AverageDelta = math.MinInt32
		individual.BestFitness = math.MinInt32
		individual.BestDelta = math.MinInt32
		individual.FitnessVariance = math.MinInt32
		individual.FitnessStdDev = math.MinInt32
		individual.Program = nil
		individual.Fitness = nil
		individual.Deltas = nil

		individuals[i] = individual
	}

	return individuals, nil
}
