package evolution

import (
	"fmt"
	"math/rand"
)

type HallOfFame struct {
	Engine Engine

	AntagonistArchive  []Individual
	ProtagonistArchive []Individual

	GenerationIntervals int

	BestIndividualMap BestIndividualMap
}

func (h *HallOfFame) Name(otherInfo string) string {
	return fmt.Sprintf("HoF-%s", otherInfo)
}

func (h *HallOfFame) ClearMap() {
	for key := range h.BestIndividualMap {
		delete(h.BestIndividualMap, key)
	}
}

func (h *HallOfFame) Topology(currentGeneration Generation, params EvolutionParams) (curr Generation, nextGen Generation, err error) {
	roundRobin := RoundRobin{Engine: h.Engine}
	h.BestIndividualMap = NewBestIndividualMap()

	currGen, nextGeneration, err := roundRobin.Topology(currentGeneration, params)
	if err != nil {
		return Generation{}, Generation{}, err
	}

	bestAntagonist := currGen.BestAntagonist()
	bestProtagonist := currGen.BestProtagonist()

	idAllocator := h.Engine.NewBatch(4)

	newID1 := int(bestAntagonist.ID) + int(idAllocator.idStart+1)
	newID2 := int(bestProtagonist.ID) + int(idAllocator.idStart+2)

	h.AntagonistArchive = append(h.AntagonistArchive, bestAntagonist.Clone(newID1))
	h.ProtagonistArchive = append(h.ProtagonistArchive, bestProtagonist.Clone(newID2))

	return currGen, nextGeneration, nil
}

func (h *HallOfFame) Evolve() (EvolutionResult, error) {
	engine := h.Engine
	params := engine.Parameters

	gen0, _, _, err := engine.InitializeGenerations(engine.Parameters)
	if err != nil {
		return EvolutionResult{}, err
	}

	genCount := CalculateGenerationSize(engine.Parameters)

	h.GenerationIntervals = int(engine.Parameters.Topology.HoFGenerationInterval * float64(genCount))

	if h.GenerationIntervals >= int(float64(params.EachPopulationSize)*0.1) {

		for h.GenerationIntervals >= int(float64(params.EachPopulationSize)*0.1) {
			if h.GenerationIntervals < MinAllowableGenerationsToTerminate {
				h.GenerationIntervals = params.EachPopulationSize / 2
				break
			}

			h.GenerationIntervals /= 2
			if h.GenerationIntervals == 0 {
				h.GenerationIntervals = 4
			}
		}

	}

	currGen := gen0
	for i := 0; i < genCount; i++ {
		//started := time.Now()
		// 1. CLEANSE
		currGen.CleansePopulations(engine.Parameters)

		// REINSERT HALL OF FAME
		if i%h.GenerationIntervals == 0 && i != 0 {
			//Reinsert
			permAntagonist := rand.Perm(h.GenerationIntervals)
			permProtagonist := rand.Perm(h.GenerationIntervals)

			newBatch := engine.NewBatch(uint32(4 * h.GenerationIntervals))
			count := uint32(0)

			for j := 0; j < h.GenerationIntervals; j++ {
				antagonistClone := h.AntagonistArchive[permAntagonist[j]].Clone(int(newBatch.idStart + count))
				antagonistClone.Program = nil
				count++

				protagonistClone := h.ProtagonistArchive[permProtagonist[j]].Clone(int(newBatch.idStart + count))
				protagonistClone.Program = nil

				currGen.Antagonists[permAntagonist[j]] = antagonistClone
				currGen.Protagonists[permProtagonist[j]] = protagonistClone
				count++
			}
		}

		// 2. START
		currentGen, nextGeneration, err := h.Topology(currGen, params)
		if err != nil {
			return EvolutionResult{}, err
		}

		// 3. EVALUATE
		generationResult := currentGen.RunGenerationStatistics()
		engine.GenerationResults[i] = generationResult

		//if genCount == params.GenerationsCount && params.MaxGenerations < MinAllowableGenerationsToTerminate {
		//	shouldTerminateEvolution := engine.EvaluateTerminationCriteria(currGen, generationResult, engine.Parameters)
		//	if shouldTerminateEvolution {
		//		break
		//	}
		//}

		if i == engine.Parameters.MaxGenerations-1 {
			break
		}

		currGen = nextGeneration

		// Check if Antagonists or Protagonist have clashing IDs
		idMap := make(map[uint32]int)
		for _, ind := range currGen.Antagonists {
			if idMap[ind.ID] > 1 {
				panic("ID CLASH!")
			}
			idMap[ind.ID]++
		}

		for _, ind := range currGen.Protagonists {
			if idMap[ind.ID] > 1 {
				panic("ID CLASH!")
			}
			idMap[ind.ID]++
		}

	}

	evolutionResult := engine.AnalyzeResults()

	return evolutionResult, nil
}
