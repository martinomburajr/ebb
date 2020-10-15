package evolution

import (
	"math"
	"math/rand"
)

const (
	SurvivorSelectionParentVsChild = "ParentVsChild"
	SurvivorSelectionHalfAndHalf   = "HalfAndHalf"
)

// ParentVsChildSurvivorSelection ensures children also competeAntagonists to gain fitness. The best from each pool are selected.
// This technique can double the time taken to complete a generation as children now have to competeAntagonists in order to be selected.
func ParentVsChildSurvivorSelection(selectedParents, selectedChildren []Individual,
	params EvolutionParams) ([]Individual, error) {
	survivors := make([]Individual, params.EachPopulationSize)

	//parentPopulationSize := int(params.Selection.Survivor.SurvivorPercentage * float64(params.EachPopulationSize))
	//childPopulationSize := params.EachPopulationSize - parentPopulationSize
	//
	//sortedParents, err := SortIndividuals(selectedParents)
	//if err != nil {
	//	return nil, err
	//}
	//
	//sortedChildren, err := SortIndividuals(selectedChildren)
	//if err != nil {
	//	return nil, err
	//}

	// Tough to directly implement

	return survivors, nil
}

// HalfAndHalfSurvivorSelection takes the top half best parents and randomly selects children for the other half
func HalfAndHalfSurvivorSelection(selectedParents, selectedChildren []Individual, survivorPercentage float64,
	populationSize int) ([]Individual, error) {

	return applyRatiodSurvivorSelection(selectedParents, selectedChildren, survivorPercentage, populationSize)
}

// HalfAndHalfSurvivorSelection takes the top half best parents and randomly selects children for the other half
func applyRatiodSurvivorSelection(selectedParents, selectedChildren []Individual, survivorPercentage float64,
	populationSize int) ([]Individual, error) {
	survivors := make([]Individual, populationSize)

	sortedParents, err := SortIndividuals(selectedParents)
	if err != nil {
		return nil, err
	}

	slider := int(math.Floor(float64(populationSize) * survivorPercentage))
	childIndices := rand.Perm(slider)

	for i := 0; i < slider; i++ {
		survivors[i] = sortedParents[i]
	}
	for i := 0; i < populationSize-slider; i++ {
		survivors[i+slider] = selectedChildren[childIndices[i]]
	}

	return survivors, nil
}
