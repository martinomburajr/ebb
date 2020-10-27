package evolution

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
	"math/rand"
	"strings"
)

const (
	IndividualAntagonist  = 0
	IndividualProtagonist = 1
)

type Individual struct {
	ID       uint32
	Program  BinaryTree // The best program generated
	Strategy []Strategy
	Kind     int

	// Statistics
	Fitness              []float64
	Deltas               []float64
	FitnessVariance      float64
	FitnessStdDev        float64
	HasAppliedStrategy   bool
	HasCalculatedFitness bool

	// BirthGen represents the generation where this individual was spawned
	BirthGen         int
	Age              int
	BestFitness      float64 // Best fitness from all epochs
	AverageFitness   float64 // Measures average fitness throughout epoch
	BestDelta        float64
	AverageDelta     float64
	NoOfCompetitions int
}

// TODO switch to float32
// Clone will perform a hard clone on every field, set id < 0 to keep the same id.
func (i Individual) Clone(id int) Individual {
	if id >= 0 {
		i.ID = uint32(id)
	}

	dstStrategy := make([]Strategy, len(i.Strategy))
	dstFitness := make([]float64, len(i.Fitness))
	dstDeltas := make([]float64, len(i.Deltas))
	dstProgram := make([]BinaryTreeNode, len(i.Program))

	copy(dstDeltas, i.Deltas)
	copy(dstStrategy, i.Strategy)
	copy(dstFitness, i.Fitness)
	copy(dstProgram, i.Program)

	return i
}

//// TODO switch to float32
//// Clone will perform a hard clone on every field
//func (i Individual) CloneWithParentID(id, parentID int) Individual {
//	clone := i.Clone(id)
//	clone.ParentID = uint32(parentID)
//
//	return clone
//}

// TODO - Figure out how to allocate IDs

// CloneCleanse removes performance based information but keeps the strategy intact.
func (i Individual) CloneCleanse(ID uint32) (Individual, error) {

	i.ID = ID
	i.Fitness = make([]float64, 0)
	i.Deltas = make([]float64, 0)
	i.AverageFitness = 0
	i.AverageDelta = 0
	i.Program = BinaryTree{}
	i.FitnessStdDev = 0
	i.FitnessVariance = 0

	dstStrategy := make([]Strategy, len(i.Strategy))

	copy(dstStrategy, i.Strategy)

	return i, nil
}

// Mutate will mutate the Strategy in a given individual
func (i *Individual) Mutate(availableStrategies []Strategy) error {
	if availableStrategies == nil {
		return fmt.Errorf("Mutate | availableStrategies param cannot be nil")
	}
	if i.Strategy == nil {
		return fmt.Errorf("Mutate | i's strategies cannot be nil")
	}
	if len(i.Strategy) < 1 {
		return fmt.Errorf("Mutate | i's strategies cannot empty")
	}

	randIndexToMutate := rand.Intn(len(i.Strategy))

	randIndexForStrategies := rand.Intn(len(availableStrategies))
	i.Strategy[randIndexToMutate] = availableStrategies[randIndexForStrategies]
	return nil
}

// ApplyAntagonistStrategy applies the AntagonistEquation strategies to program.
func (i *Individual) ApplyAntagonistStrategy(params EvolutionParams) error {
	if i.Kind == IndividualProtagonist {
		return fmt.Errorf("ApplyAntagonistStrategy | cannot apply AntagonistsMean Strategy to Protagonist")
	}

	if i.Strategy == nil {
		return fmt.Errorf("antagonist stategy cannot be nil")
	}
	if len(i.Strategy) < 1 {
		return fmt.Errorf("antagonist Strategy cannot be empty")
	}

	i.Program = params.StartIndividual.Clone()

	for _, strategy := range i.Strategy {

		err := i.ApplyStrategy(strategy,
			params.SpecParam.AvailableSymbolicExpressions.Terminals,
			params.SpecParam.AvailableSymbolicExpressions.NonTerminals,
			params.Strategies.NewTreeNTCount)

		//log.Printf("Ant-%d: \t %s \t -->  \t %s", i.ID, strategy, i.Program.ToMathematicalString())

		if err != nil {
			return err
		}
	}

	i.HasAppliedStrategy = true
	i.NoOfCompetitions++

	return nil
}

// TODO - Create STrategy that applies a random constant
// ApplyProtagonistStrategy applies the AntagonistEquation strategies to program.
func (i *Individual) ApplyProtagonistStrategy(antagonistTree BinaryTree, params EvolutionParams) error {
	if i.Kind == IndividualAntagonist {
		return fmt.Errorf("ApplyProtagonistStrategy | cannot apply Protagonist Strategy to AntagonistsMean")
	}
	if i.Strategy == nil {
		return fmt.Errorf("protagonist stategy cannot be nil")
	}
	if len(i.Strategy) < 1 {
		return fmt.Errorf("protagonist Strategy cannot be empty")
	}

	i.Program = antagonistTree.Clone()

	for _, strategy := range i.Strategy {

		err := i.ApplyStrategy(strategy,
			params.SpecParam.AvailableSymbolicExpressions.Terminals,
			params.SpecParam.AvailableSymbolicExpressions.NonTerminals,
			params.Strategies.NewTreeNTCount)

		//log.Printf("Pro-%d: \t%s \t -->  \t %s", i.ID, strategy,  i.Program.ToMathematicalString())

		if err != nil {
			return err
		}
	}

	i.HasAppliedStrategy = true
	i.HasAppliedStrategy = true

	i.NoOfCompetitions++

	return nil
}

func (i Individual) CloneWithTree(ID int, tree BinaryTree) Individual {
	if ID >= 0 {
		i.ID = uint32(ID)
	}

	dstStrategy := make([]Strategy, len(i.Strategy))
	dstFitness := make([]float64, len(i.Fitness))
	dstDeltas := make([]float64, len(i.Deltas))

	copy(dstDeltas, i.Deltas)
	copy(dstStrategy, i.Strategy)
	copy(dstFitness, i.Fitness)

	i.Program = tree.Clone()

	return i
}

func (i *Individual) CalculateProtagonistThresholdedFitness(params EvolutionParams) (
	protagonistFitness float64,
	delta float64, err error) {
	if !i.HasAppliedStrategy {
		return 0, 0, fmt.Errorf(" CalculateProtagonistThresholdedFitness | has not applied strategies")
	}

	if i.Kind == IndividualAntagonist {
		return 0, 0, fmt.Errorf(" CalculateProtagonistThresholdedFitness | cannot apply protagonist antagonist" +
			" fitness to" +
			" antagonist")
	}

	protagonist := i.Program

	fitnessPenalization := params.Spec[0].DivideByZeroPenalty
	deltaProtagonist := 0.0
	deltaProtagonistThreshold := 0.0
	isProtagonistValid := true

	spec := params.Spec

	for j := range spec {
		independentXVal := spec[j].Independents['x']

		shouldStopAndContinue := false

		if isProtagonistValid {
			dependentProtagonistVar := EvaluateMathematicalExpression(protagonist, independentXVal)

			if math.IsNaN(dependentProtagonistVar) || math.IsInf(dependentProtagonistVar, 0) {
				isProtagonistValid, shouldStopAndContinue = applyDivByZeroError(independentXVal, dependentProtagonistVar)

				if shouldStopAndContinue {
					continue
				}
			} else {
				diff := spec[j].Dependent - dependentProtagonistVar
				if math.IsNaN(diff) || math.IsInf(diff, 0) {
				} else {
					deltaProtagonist += math.Pow(diff, 2)
				}
			}
		} else {
			break
		}

		protagonistThreshold := spec[j].ProtagonistThreshold

		if !math.IsInf(protagonistThreshold, 0) && !math.IsNaN(protagonistThreshold) {
			deltaProtagonistThreshold += protagonistThreshold * protagonistThreshold
		}
	}

	specLen := float64(len(spec))

	protagonistFitness, protagonistFitnessError := deliberateProtagonistFitness(specLen, deltaProtagonist, deltaProtagonistThreshold, isProtagonistValid, fitnessPenalization)

	return protagonistFitness, protagonistFitnessError, nil
}

func (i *Individual) CalculateAntagonistThresholdedFitness(params EvolutionParams) (antagonistFitness float64, delta float64, err error) {

	if !i.HasAppliedStrategy {
		return 0, 0, fmt.Errorf(" CalculateAntagonistThresholdedFitness | has not applied strategies")
	}
	if i.Kind == IndividualProtagonist {
		return 0, 0, fmt.Errorf(" CalculateAntagonistThresholdedFitness | cannot apply antagonist fitness to" +
			" protagonist")
	}

	antagonist := i.Program

	fitnessPenalization := params.Spec[0].DivideByZeroPenalty
	deltaAntagonist := 0.0
	deltaAntagonistThreshold := 0.0
	isAntagonistValid := true

	spec := params.Spec

	for j := range spec {
		independentXVal := spec[j].Independents['x']

		shouldStopAndContinue := false

		if isAntagonistValid {
			if isAntagonistValid {
				dependentAntagonistVar := EvaluateMathematicalExpression(antagonist, independentXVal)

				if math.IsNaN(dependentAntagonistVar) || math.IsInf(dependentAntagonistVar, 0) {
					isAntagonistValid, shouldStopAndContinue = applyDivByZeroError(independentXVal, dependentAntagonistVar)

					if shouldStopAndContinue {
						continue
					}
				} else {
					diff := spec[j].Dependent - dependentAntagonistVar
					if math.IsNaN(diff) || math.IsInf(diff, 0) {
					} else {
						deltaAntagonist += diff * diff
					}
				}
			}
		} else {
			break
		}

		antagonistThreshold := spec[j].AntagonistThreshold
		if !math.IsInf(antagonistThreshold, 0) && !math.IsNaN(antagonistThreshold) {
			deltaAntagonistThreshold += antagonistThreshold * antagonistThreshold
		}
	}

	specLen := float64(len(spec))

	antagonistFitness, antagonistFitnessError := deliberateAntagonistFitness(specLen, deltaAntagonist, deltaAntagonistThreshold, isAntagonistValid, fitnessPenalization)

	return antagonistFitness, antagonistFitnessError, nil
}

// ApplyStrategy takes a given Strategy and applies a transformation to the given program.
// depth defines the exact depth the treeNode can evolve to given the transformation.
// Depth of a treeNode increases exponentially. So keep depths small e.g. 1,2,3
// Ensure to place the independent variabel e.g X at the start of the SymbolicExpression terminals array.
// Otherwise there is less of a chance of having the independent variable propagate.
// The system is designed such that the first element of the terminals array will be the most prominent with regards
// to appearance.
func (i *Individual) ApplyStrategy(strategy Strategy, terminals, nonTerminals []rune, nonTerminalCount int) (err error) {
	// TODO - Remove Tree Correctness Evaluation

	switch strategy {
	case StrategyDeleteNonTerminal: // CHANGE TO DeleteNonTerminal
		i.Program = i.Program.DeleteNonTerminal()

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	case StrategyDeleteTerminal:
		i.Program = i.Program.DeleteTerminal()

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	case StrategyMutateNonTerminal:
		i.Program = i.Program.MutateNonTerminal(nonTerminals)

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	case StrategyMutateTerminal:
		i.Program = i.Program.MutateTerminal(terminals)

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	case StrategyReplaceBranch:
		tree := GenerateRandomTree(nonTerminalCount, terminals, nonTerminals)
		i.Program = i.Program.ReplaceBranch(tree)

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	case StrategyAppendRandomOperation:
		i.Program = i.Program.AppendRandomOperation(terminals, nonTerminals)

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	// DETERMINISTIC STRATEGIES
	case StrategySkip:
		// Do nothing
		break
	case StrategyFellTree:
		i.Program = i.Program.Fell()

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	case StrategyMultXD:
		i.Program = i.Program.ApplyOperatorOnTerminal('*', 'x')

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break
	case StrategyAddXD:
		i.Program = i.Program.ApplyOperatorOnTerminal('+', 'x')

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break
	case StrategySubXD:
		i.Program = i.Program.ApplyOperatorOnTerminal('-', 'x')

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break
	case StrategyDivXD:
		i.Program = i.Program.ApplyOperatorOnTerminal('/', 'x')
		//
		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	case StrategyMultCD:
		randTerminalIndex := rand.Intn(len(terminals))
		terminal := terminals[randTerminalIndex]
		i.Program = i.Program.ApplyOperatorOnTerminal('*', terminal)

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break
	case StrategyAddCD:
		randTerminalIndex := rand.Intn(len(terminals))
		terminal := terminals[randTerminalIndex]
		i.Program = i.Program.ApplyOperatorOnTerminal('+', terminal)

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break
	case StrategySubCD:
		randTerminalIndex := rand.Intn(len(terminals))
		terminal := terminals[randTerminalIndex]
		i.Program = i.Program.ApplyOperatorOnTerminal('-', terminal)

		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break
	case StrategyDivCD:
		randTerminalIndex := rand.Intn(len(terminals))
		terminal := terminals[randTerminalIndex]
		i.Program = i.Program.ApplyOperatorOnTerminal('/', terminal)
		//
		//correctness := i.Program.EvaluateTreeCorrectness()
		//if len(correctness) > 0 {
		//	log.Fatalf(correctness.ToString(i.Program))
		//}

		break

	default:
		break
	}

	return err
}

func DominantStrategy(individual Individual) string {
	domStrat := map[string]int{}
	for i := range individual.Strategy {
		strategy := string(individual.Strategy[i])

		stratCount := domStrat[strategy]
		domStrat[strategy] = stratCount + 1
	}

	var topStrategy string
	counter := 0
	for k, v := range domStrat {
		if v > counter {
			counter = v
			topStrategy = k
		}
	}
	return topStrategy
}

func DominantStrategyStr(str string) string {
	strategies := strings.Split(str, "|")

	domStrat := map[string]int{}
	for i := range strategies {
		strategy := string(strategies[i])
		stratCount := domStrat[strategy]
		if domStrat[strategy] > -1 {
			domStrat[strategy] = stratCount + 1
		}
	}

	var topStrategy string
	counter := 0
	for k, v := range domStrat {
		if v > counter {
			counter = v
			topStrategy = k
		}
	}
	return topStrategy
}

func StrategiesToString(individual Individual) string {
	sb := strings.Builder{}
	for _, strategy := range individual.Strategy {
		sb.WriteString(string(strategy))
		sb.WriteString("|")
	}

	final := sb.String()
	return final[:len(final)-1]
}

func StrategiesToStringArr(strategies []string) string {
	sb := strings.Builder{}
	for _, strategy := range strategies {
		sb.WriteString(string(strategy))
		sb.WriteString("|")
	}

	final := sb.String()
	if len(final) < 1 {
		return final
	}
	return final[:len(final)-1]
}

func ConvertStrategiesToString(strategies []Strategy) (stringStrategies []string) {
	for i := range strategies {
		stringStrategies = append(stringStrategies, string(strategies[i]))
	}
	return stringStrategies
}

type Antagonist Individual
type Protagonist Individual

func (i *Individual) ToString() strings.Builder {
	sb := strings.Builder{}

	var kind string
	if i.Kind == IndividualAntagonist {
		kind = "Ant"
	} else if i.Kind == IndividualProtagonist {
		kind = "Pro"
	} else {
		kind = "Unknown"
	}

	str := fmt.Sprintf("%s-%d\n"+
		"eq: %s\n"+
		"strat: %v\n"+
		"avf: %.2f | avstd: %.2f | bestf: %.2f\n"+
		"bG: %d | age: %d\n",
		kind, i.ID, i.Program.ToMathematicalString(), i.Strategy, i.AverageFitness, i.FitnessStdDev, i.BestFitness, i.BirthGen, i.Age)

	sb.WriteString(str)

	return sb
}

func (i *Individual) ToSimpleString() strings.Builder {
	sb := strings.Builder{}

	str := fmt.Sprintf("avf: %.2f | bestf: %.2f | bG: %d | age: %d",
		i.AverageFitness, i.BestFitness, i.BirthGen, i.Age)

	sb.WriteString(str)

	return sb
}

func (i *Individual) Calculate() {
	deltaMean := stat.Mean(i.Deltas, nil)
	mean, std := stat.MeanStdDev(i.Fitness, nil)
	variance := stat.Variance(i.Fitness, nil)

	i.AverageFitness = mean
	i.FitnessStdDev = std
	i.FitnessVariance = variance
	i.HasCalculatedFitness = true
	i.HasAppliedStrategy = true
	//i.Age += 1
	i.AverageDelta = deltaMean
}
