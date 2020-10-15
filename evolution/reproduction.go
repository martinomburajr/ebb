package evolution

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

const (
	CrossoverSinglePoint = "CrossoverSinglePoint"
	CrossoverUniform     = "CrossoverUniform"
)

// CrossoverSinglePoint performs a single-point crossover that is dictated by the crossover percentage float.
// Both parent chromosomes are split at the percentage section specified by crossoverPercentage
func SinglePointCrossover(parentA, parentB *Individual, childIDA, childIDB int) (childA Individual, childB Individual, err error) {
	// Require
	if parentA.Strategy == nil {
		return Individual{}, Individual{}, fmt.Errorf("parentA strategy cannot be nil")
	}

	if len(parentA.Strategy) < 1 {
		return Individual{}, Individual{}, fmt.Errorf("parentA strategy cannot be empty")
	}
	if parentB.Strategy == nil {
		return Individual{}, Individual{}, fmt.Errorf("parentB strategy cannot be nil")
	}
	if len(parentB.Strategy) < 1 {
		return Individual{}, Individual{}, fmt.Errorf("parentB strategy cannot be empty")
	}

	// TODO - FIX ME!!!

	// DO
	childA = parentA.Clone(childIDA)
	childA.ID = uint32(childIDA)
	childA.Fitness = nil
	childA.Deltas = nil
	childA.Program = nil
	childA.AverageFitness = math.MinInt32
	childA.FitnessStdDev = math.MinInt32
	childA.FitnessVariance = math.MinInt32
	childA.BestDelta = math.MinInt32
	childA.BestFitness = math.MinInt32
	childA.NoOfCompetitions = 0
	childA.HasAppliedStrategy = false
	childA.HasCalculatedFitness = false

	childB = parentB.Clone(childIDB)
	childB.ID = uint32(childIDB)
	childB.Fitness = nil
	childB.Program = nil
	childB.Deltas = nil
	childB.AverageFitness = math.MinInt32
	childB.FitnessStdDev = math.MinInt32
	childB.FitnessVariance = math.MinInt32
	childB.BestDelta = math.MinInt32
	childB.BestFitness = math.MinInt32
	childB.NoOfCompetitions = 0
	childB.HasAppliedStrategy = false
	childB.HasCalculatedFitness = false

	mut := sync.Mutex{}
	mut.Lock()
	if len(parentA.Strategy) >= len(parentB.Strategy) {
		prob := 0
		for prob == 0 {
			prob = rand.Intn(len(parentB.Strategy))
		}

		for i := 0; i < prob; i++ {
			childA.Strategy[i] = parentB.Strategy[i]
			childB.Strategy[i] = parentA.Strategy[i]
		}
	} else {
		prob := 0
		for prob == 0 {
			prob = rand.Intn(len(parentA.Strategy))
		}
		for i := 0; i < prob; i++ {
			childA.Strategy[i] = parentB.Strategy[i]
			childB.Strategy[i] = parentA.Strategy[i]
		}
	}
	mut.Unlock()

	return childA, childB, nil
}

// CrossoverSinglePoint performs a single-point crossover that is dictated by the crossover percentage float.
// Both parent chromosomes are split at the percentage section specified by crossoverPercentage
func UniformCrossover(parentA, parentB Individual, childIDA, childIDB int) (childA Individual,
	childB Individual,
	err error) {
	// Require
	if parentA.Strategy == nil {
		return Individual{}, Individual{}, fmt.Errorf("parentA strategy cannot be nil")
	}
	if len(parentA.Strategy) < 1 {
		return Individual{}, Individual{}, fmt.Errorf("parentA strategy cannot be empty")
	}
	if parentB.Strategy == nil {
		return Individual{}, Individual{}, fmt.Errorf("parentB strategy cannot be nil")
	}
	if len(parentB.Strategy) < 1 {
		return Individual{}, Individual{}, fmt.Errorf("parentB strategy cannot be empty")
	}

	//DO
	childA = parentA.Clone(childIDA)
	childA.ID = uint32(childIDA)
	childA.Fitness = nil
	childA.Deltas = nil
	childA.Program = nil
	childA.AverageFitness = math.MinInt32
	childA.FitnessStdDev = math.MinInt32
	childA.FitnessVariance = math.MinInt32
	childA.BestDelta = math.MinInt32
	childA.BestFitness = math.MinInt32
	childA.NoOfCompetitions = 0
	childA.HasAppliedStrategy = false
	childA.HasCalculatedFitness = false

	childB = parentB.Clone(childIDB)
	childB.ID = uint32(childIDB)
	childB.Fitness = nil
	childB.Program = nil
	childB.Deltas = nil
	childB.AverageFitness = math.MinInt32
	childB.FitnessStdDev = math.MinInt32
	childB.FitnessVariance = math.MinInt32
	childB.BestDelta = math.MinInt32
	childB.BestFitness = math.MinInt32
	childB.NoOfCompetitions = 0
	childB.HasAppliedStrategy = false
	childB.HasCalculatedFitness = false

	mut := sync.Mutex{}
	mut.Lock()
	if len(parentA.Strategy) >= len(parentB.Strategy) {
		for i := 0; i < len(parentB.Strategy); i++ {
			prob := rand.Intn(2)
			if prob == 0 {
				childA.Strategy[i] = parentA.Strategy[i]
			} else {
				childA.Strategy[i] = parentB.Strategy[i]
			}

			prob2 := rand.Intn(2)
			if prob2 == 0 {
				childB.Strategy[i] = parentA.Strategy[i]
			} else {
				childB.Strategy[i] = parentB.Strategy[i]
			}
		}
	} else {
		for i := 0; i < len(parentB.Strategy); i++ {
			prob := rand.Intn(2)
			if prob == 0 {
				childA.Strategy[i] = parentA.Strategy[i]
			} else {
				childA.Strategy[i] = parentB.Strategy[i]
			}

			prob2 := rand.Intn(2)
			if prob2 == 0 {
				childB.Strategy[i] = parentA.Strategy[i]
			} else {
				childB.Strategy[i] = parentB.Strategy[i]
			}
		}
	}
	mut.Unlock()

	return childA, childB, nil
}
