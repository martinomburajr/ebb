package evolution

import "math/rand"

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
func HalfAndHalfSurvivorSelection(selectedParents, selectedChildren []Individual,
	populationSize int) ([]Individual, error) {

	return applyHalfAndHalfSurvivorSelection(selectedParents, selectedChildren, populationSize)
}

// HalfAndHalfSurvivorSelection takes the top half best parents and randomly selects children for the other half
func applyHalfAndHalfSurvivorSelection(selectedParents, selectedChildren []Individual,
	populationSize int) ([]Individual, error) {
	survivors := make([]Individual, populationSize)

	sortedParents, err := SortIndividuals(selectedParents)
	if err != nil {
		return nil, err
	}

	halfPopulation := populationSize / 2
	childIndices := rand.Perm(halfPopulation)

	for i := 0; i < halfPopulation; i++ {
		survivors[i] = sortedParents[i]
	}
	for i := 0; i < halfPopulation; i++ {
		survivors[i+halfPopulation] = selectedChildren[childIndices[i]]
	}

	return survivors, nil
}
