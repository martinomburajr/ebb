package evolution

import (
	"fmt"
	"math"
)

// RRCompetition is defined as a coevolutionary step where protagonist and antagonist competeAntagonists.
// For example an epoch could represent a distinct interaction between two parties.
// For instance a bug mutated program (antagonist) can be challenged a variety of times (
// specified by {iterations}) by the tests (protagonist).
// The test will use up the strategies it contains and attempt to chew away at the antagonists Fitness,
// to maximize its own
type RRCompetition struct {
	id                    uint32
	generation            *Generation
	protagonist           *Individual
	antagonist            *Individual
	program               BinaryTree
	isComplete            bool
	terminalSet           []rune
	nonTerminalSet        []rune
	hasAntagonistApplied  bool
	hasProtagonistApplied bool
}

func (r *RRCompetition) compete(params EvolutionParams) (antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta float64, err error) {
	err = r.antagonist.ApplyAntagonistStrategy(params)
	inf := math.Inf(-1)

	if err != nil {
		return inf, inf, inf, inf, err
	}

	err = r.protagonist.ApplyProtagonistStrategy(r.antagonist.Program, params)
	if err != nil {
		return inf, inf, inf, inf, err
	}

	antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta = ThresholdedRatioFitness(params.Spec, r.antagonist.Program, r.protagonist.Program)

	// TODO - PUNISH DIVISIONS BY ZERO!
	if math.IsNaN(antagonistFitness) {
		print()
	}

	if math.IsNaN(protagonistFitness) {
		print()
	}

	return antagonistFitness, protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta, err
}

// AggregateFitness simply adds all the Fitness values of a given individual to come up with a total number.
// If the Fitness array is nil or empty return MaxInt8 as values such as -1 or 0 have a differnt meaning
func AggregateFitness(individual Individual) (float64, error) {
	if individual.Fitness == nil {
		return math.MaxInt8, fmt.Errorf("individuals Fitness arr cannot be nil")
	}
	if len(individual.Fitness) == 0 {
		return math.MaxInt8, fmt.Errorf("individuals Fitness arr cannot be empty")
	}

	sum := 0.0
	for i := 0; i < len(individual.Fitness); i++ {
		sum += individual.Fitness[i]
	}
	return sum, nil
}
