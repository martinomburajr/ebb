package evolution

import (
	"fmt"
	"math"
	"sort"
)

// GetTopIndividualInRun returns the best protagonist and antagonist in the entire evolutionary process
func GetTopIndividualInRun(sortedGenerations []*Generation) (topAntagonist Individual, topProtagonist Individual, err error) {
	if sortedGenerations == nil {
		return Individual{}, Individual{}, fmt.Errorf("GetGenerationalFitnessAverage | Generation cannot be nil")
	}
	if len(sortedGenerations) < 1 {
		return Individual{}, Individual{}, fmt.Errorf("GetGenerationalFitnessAverage | Generation cannot be empty")
	}

	topAntagonist = Individual{AverageFitness: math.MinInt64}
	topProtagonist = Individual{AverageFitness: math.MinInt64}

	for i := 0; i < len(sortedGenerations); i++ {
		// This ensures it picks more recent individuals
		if sortedGenerations[i].Antagonists[0].AverageFitness >= topAntagonist.AverageFitness {
			topAntagonist = sortedGenerations[i].Antagonists[0]
		}

		if sortedGenerations[i].Protagonists[0].AverageFitness >= topProtagonist.AverageFitness {
			topProtagonist = sortedGenerations[i].Protagonists[0]
		}
	}

	return topAntagonist, topProtagonist, nil
}

// SortIndividuals returns the Top N-1 individuals. In this application less is more,
// so they are sorted in ascending order, with smaller indices representing better individuals.
// It is for the user to specify the Kind of individual to pass in be it antagonist or protagonist.
func SortIndividuals(individuals []Individual) ([]Individual, error) {
	if individuals == nil {
		return nil, fmt.Errorf("SortIndividuals | individuals cannot be nil")
	}
	if len(individuals) < 1 {
		return nil, fmt.Errorf("SortIndividuals | individuals cannot be empty")
	}

	sort.Slice(individuals, func(i, j int) bool {
		return individuals[i].AverageFitness > individuals[j].AverageFitness
	})

	return individuals, nil
}

// SortIndividuals returns the Top N-1 individuals. In this application less is more,
// so they are sorted in ascending order, with smaller indices representing better individuals.
// It is for the user to specify the Kind of individual to pass in be it antagonist or protagonist.
func SortIndividualsByAvgDelta(individuals []*Individual, isMoreFitnessBetter bool) ([]*Individual, error) {
	if individuals == nil {
		return nil, fmt.Errorf("SortIndividuals | individuals cannot be nil")
	}
	if len(individuals) < 1 {
		return nil, fmt.Errorf("SortIndividuals | individuals cannot be empty")
	}

	switch isMoreFitnessBetter {
	case true:
		sort.Slice(individuals, func(i, j int) bool {
			return individuals[i].AverageDelta > individuals[j].AverageDelta
		})
	case false:
		sort.Slice(individuals, func(i, j int) bool {
			return individuals[i].AverageDelta < individuals[j].AverageDelta
		})
	default:
		// Default to More is better
		sort.Slice(individuals, func(i, j int) bool {
			return individuals[i].AverageDelta > individuals[j].AverageDelta
		})
	}
	return individuals, nil
}

// SortGenerationsThoroughly sorts each kind of individual in each generation for every generation.
// This allows for easy querying in later phases.
func SortGenerationsThoroughly(generations []*Generation) ([]*Generation, error) {
	if generations == nil {
		return nil, fmt.Errorf("SortGenerationsThoroughly | generations cannot be nil")
	}
	if len(generations) < 1 {
		return nil, fmt.Errorf("SortGenerationsThoroughly | generations cannot be empty")
	}

	sortedGenerations := make([]*Generation, len(generations))

	// TODO - Introduce Parallelism later
	for i := 0; i < len(generations); i++ {
		sortedGenerations[i] = generations[i]
		//generations[i].Mutex.Lock()

		sortedAntagonists, err := SortIndividuals(generations[i].Antagonists)
		if err != nil {
			return nil, err
		}

		sortedProtagonists, err := SortIndividuals(generations[i].Protagonists)
		if err != nil {
			return nil, err
		}

		sortedGenerations[i].Protagonists = sortedProtagonists
		sortedGenerations[i].Antagonists = sortedAntagonists

		//generations[i].Mutex.Unlock()
	}
	return sortedGenerations, nil
}
