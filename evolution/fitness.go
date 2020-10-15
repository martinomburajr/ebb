package evolution

import (
	"math"
)

const (
	FitnessRatio = "FitnessRatio"

	DivByZeroIgnore           = "Ignore"
	DivByZeroSteadyPenalize   = "SteadyPenalization"
	DivByZeroPenalize         = "Penalize"
	DivByZeroSetSpecValueZero = "SetSpecValueZero"
)

// ThresholdedRatioFitness is a means to calculate fitness that spurs the protagonists and
// antagonists to do their best. It works by applying a threshold criteria that is based on the incoming spec.
// A mono threshold is applied by setting the protagonist and antagonist threshold values to the same value,
// this is done automatically when you select the fitness strategy at the start of the evolutionary process.
// Both individuals have to fall on their respective side and either edge closer to delta-0 for the protagonist or
// delta-infitinite for the antagonist. If they happen to fall on opposite sides they attain at most -1
// A dual threshold is used to punish both antagonist of protagonist for not performing as expected.
//// This fitness strategy works by comparing the average delta values of both protagonist and antagonist with a
//// specified threshold. Their deltas are not compared against each other as with other fitness strategies,
//// the thresholds act as markers of performance for each.
//// The porotagonist and antagonist each have their own threshold values that are embeded in the SpecMulti data
//// structure. Note this only compares the average and not the total deltas
func ThresholdedRatioFitness(spec SpecMulti, antagonist, protagonist BinaryTree) (antagonistFitness float64,
	protagonistFitness, antagonistFitnessDelta, protagonistFitnessDelta float64) {

	return thresholdedRatioFitness(spec, antagonist, protagonist)
}

// thresholdedRatioFitness performs fitness evaluation using the given antagonist and protagonist.
// It returns information regarding thresholds as well,
// they can be ignored if the function does not require information on the thresholds.
// Furthermore these values are averaged based on the length of the spec.
// A nil or empty spec will throw an error. It takes advantage of RMSE
func thresholdedRatioFitness(spec SpecMulti, antagonist, protagonist BinaryTree) (antagonistFitness, protagonistFitness, antagonistFitnessError, protagonistFitnessError float64) {
	fitnessPenalization := spec[0].DivideByZeroPenalty

	deltaProtagonist := 0.0
	deltaAntagonist := 0.0
	deltaAntagonistThreshold := 0.0
	deltaProtagonistThreshold := 0.0
	isAntagonistValid := true
	isProtagonistValid := true

	for i := range spec {
		independentXVal := spec[i].Independents['x']
		// this value gets converted to true iff the spec independent variable contains a zero value and the individuals
		// by no other reason than the spec independent variable being zero have no choice but to divice by zero.
		// this enables us to skip and ommit competitions where there was no other choice but to invalidate the fitness
		// results.
		shouldStopAndContinue := false

		if isAntagonistValid {
			dependentAntagonistVar := EvaluateMathematicalExpression(antagonist, independentXVal)

			if math.IsNaN(dependentAntagonistVar) || math.IsInf(dependentAntagonistVar, 0) {
				isAntagonistValid, shouldStopAndContinue = applyDivByZeroError(independentXVal, dependentAntagonistVar)

				if shouldStopAndContinue {
					continue
				}
			} else {
				diff := spec[i].Dependent - dependentAntagonistVar
				if math.IsNaN(diff) || math.IsInf(diff, 0) {
				} else {
					deltaAntagonist += diff * diff
				}
			}
		}

		if isProtagonistValid {
			dependentProtagonistVar := EvaluateMathematicalExpression(protagonist, independentXVal)

			if math.IsNaN(dependentProtagonistVar) || math.IsInf(dependentProtagonistVar, 0) {
				isProtagonistValid, shouldStopAndContinue = applyDivByZeroError(independentXVal, dependentProtagonistVar)

				if shouldStopAndContinue {
					continue
				}
			} else {
				diff := spec[i].Dependent - dependentProtagonistVar
				if math.IsNaN(diff) || math.IsInf(diff, 0) {
				} else {
					deltaProtagonist += diff * diff
				}
			}
		}

		antagonistThreshold := spec[i].AntagonistThreshold
		protagonistThreshold := spec[i].ProtagonistThreshold

		if !math.IsInf(antagonistThreshold, 0) && !math.IsNaN(antagonistThreshold) {
			deltaAntagonistThreshold += antagonistThreshold * antagonistThreshold
		}

		if !math.IsInf(protagonistThreshold, 0) && !math.IsNaN(protagonistThreshold) {
			deltaProtagonistThreshold += protagonistThreshold * protagonistThreshold
		}
	}

	sp := float64(len(spec))

	antagonistFitness, antagonistFitnessError = deliberateAntagonistFitness(sp, deltaAntagonist, deltaAntagonistThreshold, isAntagonistValid, fitnessPenalization)
	protagonistFitness, protagonistFitnessError = deliberateProtagonistFitness(sp, deltaProtagonist, deltaProtagonistThreshold, isProtagonistValid, fitnessPenalization)

	return antagonistFitness, protagonistFitness, antagonistFitnessError, protagonistFitnessError
}

func deliberateAntagonistFitness(specLen float64, deltaAntagonist float64, deltaAntagonistThreshold float64, isAntagonistValid bool, fitnessPenalization float64) (fitness float64, delta float64) {
	badDeltaAntagonist := deltaAntagonistThreshold - 1

	deltaAntagonist = math.Sqrt(deltaAntagonist / specLen)
	deltaAntagonistThreshold = math.Sqrt(deltaAntagonistThreshold / specLen)

	if !isAntagonistValid {
		return fitnessPenalization, badDeltaAntagonist
	}

	//antagonists
	fitness = assignHealthyAntagonistFitness(deltaAntagonist, deltaAntagonistThreshold)

	return fitness, deltaAntagonist
}

func deliberateProtagonistFitness(specLen float64, deltaProtagonist float64, deltaProtagonistThreshold float64, isProtagonistValid bool, fitnessPenalization float64) (fitness float64, delta float64) {
	badDeltaProtagonist := deltaProtagonistThreshold - 1

	deltaProtagonist = math.Sqrt(deltaProtagonist / specLen)
	deltaProtagonistThreshold = math.Sqrt(deltaProtagonistThreshold / specLen)

	if !isProtagonistValid {
		return fitnessPenalization, badDeltaProtagonist
	}

	//antagonists
	fitness = assignHealthyProtagonistFitness(deltaProtagonist, deltaProtagonistThreshold)

	return fitness, deltaProtagonist
}


// assignHealthyAntagonistFitness assigns fitness only if the antagonist is deemed valid
func assignHealthyAntagonistFitness(deltaAntagonist float64, deltaAntagonistThreshold float64) float64 {
	antagonistFitness := 0.0

	if deltaAntagonist >= deltaAntagonistThreshold {
		if deltaAntagonist == 0 {
			antagonistFitness = -1 // This is to punish deltaAntagonist for coalescing near the spec
		} else {
			antagonistFitness = (deltaAntagonist - deltaAntagonistThreshold) / deltaAntagonist
		}
	} else {
		antagonistFitness = -1 * ((deltaAntagonistThreshold - deltaAntagonist) / deltaAntagonistThreshold)
	}

	return antagonistFitness
}

// assignHealthyProtagonistFitness assigns fitness only if the protagonist is deemed valid
func assignHealthyProtagonistFitness(deltaProtagonist float64, deltaProtagonistThreshold float64) float64 {
	protagonistFitness := 0.0

	if deltaProtagonist <= deltaProtagonistThreshold {
		if deltaProtagonist == 0 {
			protagonistFitness = 1
		} else {
			protagonistFitness = (deltaProtagonistThreshold - deltaProtagonist) / deltaProtagonistThreshold
		}
	} else {
		protagonistFitness = -1 * ((deltaProtagonist - deltaProtagonistThreshold) / deltaProtagonist)
	}

	return protagonistFitness
}

// applyDivByZeroError by default applies a steadyPenalty. This will check if the spec does not contain a zero, if it doesnt but the individual divides by zero, then maximum penalty will be given.
func applyDivByZeroError(independentXVal float64, dependentVar float64) (isIndividualValid bool, shouldStopAndContinue bool) {
	if independentXVal == 0 {
		if math.IsNaN(dependentVar) {
			isIndividualValid = true
			shouldStopAndContinue = true

			return isIndividualValid, shouldStopAndContinue
		}
		if math.IsInf(dependentVar, 0) {
			isIndividualValid = true
			shouldStopAndContinue = true

			return isIndividualValid, shouldStopAndContinue
		}

		isIndividualValid = true
		shouldStopAndContinue = false

		return isIndividualValid, shouldStopAndContinue

	} else {
		isIndividualValid = false
		shouldStopAndContinue = false

		return isIndividualValid, shouldStopAndContinue
	}
}

// calculateDelta calculates the absolute value between the truth and the supplied value
func calculateDelta(truth float64, value float64) float64 {
	return math.Abs(truth - value)
}
