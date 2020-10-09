package evolution

import (
	"fmt"
	"math"
)

// CleansePopulation removes the trees from the population and refits them with the starter Tree.
func CleansePopulation(individuals []Individual, treeReplacer BinaryTree, idAlloc IDAllocator) ([]Individual, error) {
	for i := range individuals {
		if individuals[i].Kind == IndividualAntagonist {

			newID1 := int(idAlloc.idStart) + i

			if uint32(newID1) > idAlloc.idEnd {
				panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID1))
			}

			//tree := treeReplacer.Clone()

			individual := individuals[i]
			if individual.Program == nil {
				individual.Program = BinaryTree{}
			}

			newIndividual := individual.Clone(newID1)
			//newIndividual.Fitness = make([]float64, 0)
			//newIndividual.Deltas = make([]float64, 0)
			newIndividual.HasCalculatedFitness = false
			newIndividual.HasAppliedStrategy = false
			newIndividual.AverageFitness = math.MinInt32
			newIndividual.AverageDelta = math.MinInt32
			newIndividual.BestFitness = math.MinInt32
			newIndividual.BestDelta = math.MinInt32
			newIndividual.FitnessVariance = math.MinInt32
			newIndividual.FitnessStdDev = math.MinInt32
			newIndividual.Program = nil
			newIndividual.Strategy = individuals[i].Strategy
			individuals[i] = newIndividual

		} else {

			newID1 := int(idAlloc.idStart) + i

			if uint32(newID1) > idAlloc.idEnd {
				panic(fmt.Sprintf("Insufficient IDs allocated, hit max | curr: %d", newID1))
			}

			newIndividual := individuals[i].Clone(newID1)
			//newIndividual.Fitness = make([]float64, 0)
			//newIndividual.Deltas = make([]float64, 0)
			newIndividual.FitnessVariance = math.MinInt32
			newIndividual.FitnessStdDev = math.MinInt32
			newIndividual.HasCalculatedFitness = false
			newIndividual.HasAppliedStrategy = false
			newIndividual.AverageFitness = math.MinInt32
			newIndividual.AverageFitness = math.MinInt32
			newIndividual.AverageDelta = math.MinInt32
			newIndividual.BestFitness = math.MinInt32
			newIndividual.BestDelta = math.MinInt32
			newIndividual.Strategy = individuals[i].Strategy
			individuals[i] = newIndividual
			individuals[i].Program = BinaryTree{}
		}
	}
	return individuals, nil
}
