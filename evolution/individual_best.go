package evolution

type BestIndividualMap map[uint32]BestIndividual

func NewBestIndividualMap() BestIndividualMap {
	return make(map[uint32]BestIndividual)
}

type BestIndividual struct {
	BestClone Individual
}

func (b BestIndividualMap) Check(individual *Individual, fitness float64, delta float64) {
	_, ok := b[individual.ID]

	if ok {
		bi := b[individual.ID]

		initialBestFitness := bi.BestClone.BestFitness
		// Ensure that the first individual isnt tested against 0.0, rather a lower number
		if len(b[individual.ID].BestClone.Fitness) <= 0 {
			initialBestFitness = -10
		}

		// If and only if the fitness of the individual improves, then give us the program and set the bestfitness.
		if initialBestFitness < fitness {
			bi.BestClone.Program = individual.Program.Clone()

			// Store the best program, otherwise because a program is a slice, it will get modified due to the fact
			// we are giving it a reference.
			bi.BestClone.BestFitness = fitness
			bi.BestClone.BestDelta = delta
		}

		bi.BestClone.NoOfCompetitions = bi.BestClone.NoOfCompetitions + 1
		bi.BestClone.Fitness = append(bi.BestClone.Fitness, fitness)
		bi.BestClone.Deltas = append(bi.BestClone.Deltas, delta)

		b[individual.ID] = bi
	} else {
		best := BestIndividual{
			BestClone: individual.Clone(-1),
		}

		best.BestClone.BestFitness = fitness
		best.BestClone.BestDelta = delta

		best.BestClone.Fitness = append(best.BestClone.Fitness, fitness)
		best.BestClone.Deltas = append(best.BestClone.Deltas, delta)
		best.BestClone.NoOfCompetitions = individual.NoOfCompetitions
		best.BestClone.Age = individual.Age
		best.BestClone.BirthGen = individual.BirthGen

		b[individual.ID] = best
	}
}

// Deposit will populate the Antagonist and Protagonist fields of the given generation with the right Individual
// information
func (b BestIndividualMap) Deposit(eachPopulationSize int) (antagonists []Individual, protagonists []Individual) {
	antagonists = make([]Individual, eachPopulationSize)
	protagonists = make([]Individual, eachPopulationSize)

	antagonistCounter := 0
	protagonistCounter := 0

	for i := range b {
		bi := b[i]

		if bi.BestClone.Kind == IndividualAntagonist {
			antagonists[antagonistCounter] = bi.BestClone
			antagonists[antagonistCounter].BestFitness = bi.BestClone.BestFitness

			antagonists[antagonistCounter].Calculate()
			antagonists[antagonistCounter].HasCalculatedFitness = true

			antagonistCounter++
		} else if bi.BestClone.Kind == IndividualProtagonist {
			protagonists[protagonistCounter] = bi.BestClone
			protagonists[protagonistCounter].BestFitness = bi.BestClone.BestFitness

			protagonists[protagonistCounter].Calculate()
			protagonists[protagonistCounter].HasCalculatedFitness = true

			protagonistCounter++
		} else {
			panic("Invalid Individual")
		}
	}

	return antagonists, protagonists
}

// Deposit will populate the Antagonist and Protagonist fields of the given generation with the right Individual
// information
func (b BestIndividualMap) DepositSingle(eachPopulationSize int) (individuals []Individual) {
	individuals = make([]Individual, eachPopulationSize)

	counter := 0

	for i := range b {
		bi := b[i]

		individuals[counter] = bi.BestClone

		individuals[counter].Calculate()
		individuals[counter].HasCalculatedFitness = true

		counter++
	}

	return individuals
}
