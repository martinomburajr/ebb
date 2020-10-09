package evolution

import (
	"fmt"
	"math"
	"strings"
)

// EquationPairing refers to a set dependent and independent values for a given equation.
// For example the equation x^2 + 1 has an equation pairing of {1, 0}, {2, 1}, {5,
// 2} for dependent and independent pairs respectively
type EquationPairing struct {
	Independents IndependentVariableMap
	Dependent    float64

	ProtagonistThreshold float64
	AntagonistThreshold  float64

	// AntagonistPenalization value to give the antagonist if it creates an invalid tree evaluation e.g. DivideByZero
	AntagonistPenalization float64
	// ProtagonistPenalization value to give the protagonist if it creates an invalid tree evaluation e.g. DivideByZero
	ProtagonistPenalization float64

	DivideByZeroPenalty float64
}

type IndependentVariableMap map[rune]float64

func (e *EquationPairing) ToString() string {
	return fmt.Sprintf("  %#v  \t  %.2f  \n", e.Independents, e.Dependent)
}

// SpecMulti is the underlying data structre that contains the spec as well as threshold information
type SpecMulti []EquationPairing

// NewSpec assumes a single independent variable x with an unlimited count.
func NewSpec(inputTree BinaryTree, specParam SpecParam, fitnessStrategy FitnessStrategy) (SpecMulti,
	error) {

	if specParam.Range < 1 {
		return nil, fmt.Errorf("GenerateSpec | specParam.Range cannot be less than 0")
	}

	if fitnessStrategy.AntagonistThresholdMultiplier < 1 {
		fitnessStrategy.AntagonistThresholdMultiplier = 1
	}

	if fitnessStrategy.ProtagonistThresholdMultiplier < 1 {
		fitnessStrategy.ProtagonistThresholdMultiplier = 1
	}

	spec := make([]EquationPairing, specParam.Range)

	for i := range spec {
		spec[i].Independents = map[rune]float64{}
		xVal := i + specParam.Seed
		spec[i].Independents['x'] = float64(xVal)

		dependentVariable := EvaluateMathematicalExpression(inputTree,
			float64(xVal))

		if math.IsNaN(dependentVariable) || math.IsInf(dependentVariable, 0) {
			if xVal != 0 {
				return nil, fmt.Errorf("invalid")
			}
		}

		spec[i].Dependent = dependentVariable
		spec[i].AntagonistThreshold = dependentVariable * fitnessStrategy.AntagonistThresholdMultiplier
		spec[i].ProtagonistThreshold = dependentVariable * fitnessStrategy.ProtagonistThresholdMultiplier
		spec[i].DivideByZeroPenalty = specParam.DivideByZeroPenalty
	}
	return spec, nil
}

func (spec SpecMulti) ToString() string {
	sb := strings.Builder{}
	if spec == nil {
		return sb.String()
	}

	sb.WriteString("  x  :\t  f(x)  \n")
	for i := range spec {
		s := spec[i].ToString()
		sb.WriteString(s)
	}
	return sb.String()
}
