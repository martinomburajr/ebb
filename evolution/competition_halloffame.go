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

	currGen, nextGeneration, err := roundRobin.Topology(currentGeneration, params)
	if err != nil {
		return Generation{}, Generation{}, err
	}

	bestAntagonist := currentGeneration.BestAntagonist()
	bestProtagonist := currentGeneration.BestProtagonist()

	h.AntagonistArchive = append(h.AntagonistArchive, bestAntagonist.Clone(int(bestAntagonist.ID)+200))
	h.ProtagonistArchive = append(h.ProtagonistArchive, bestProtagonist.Clone(int(bestProtagonist.ID)+200))

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

	if h.GenerationIntervals >= int(float64(params.EachPopulationSize) * 0.1) {

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
				antagonistClone:= h.AntagonistArchive[permAntagonist[j]].Clone(int(newBatch.idStart + count))
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
		currGen, nextGeneration, err := h.Topology(currGen, params)
		if err != nil {
			return EvolutionResult{}, err
		}

		// 3. EVALUATE
		generationResult := currGen.RunGenerationStatistics()
		engine.GenerationResults[i] = generationResult

		if genCount == params.GenerationsCount && params.MaxGenerations < MinAllowableGenerationsToTerminate {
			shouldTerminateEvolution := engine.EvaluateTerminationCriteria(currGen, generationResult, engine.Parameters)
			if shouldTerminateEvolution {
				break
			}
		}

		if i == engine.Parameters.MaxGenerations-1 {
			break
		}

		currGen = nextGeneration

		//engine.ProgressBar.Incr()

		// 4. LOG
		//elapsed := utils.TimeTrack(started)
		//go WriteGenerationToLog(engine, i, elapsed)
		//go WriteToDataFolders(engine.Parameters.FolderPercentages, i, engine.Parameters.GenerationsCount, engine.Parameters)
	}

	evolutionResult := engine.AnalyzeResults()

	return evolutionResult, nil
}
