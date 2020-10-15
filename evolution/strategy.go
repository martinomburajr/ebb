package evolution

import "math/rand"

/**
	Any strategy operation below will ensure the tree remains in a valid state.
Worst case being a single terminal with value 0.
*/

type Strategy string

const (

	// ############################# NON DETERMINISTIC ############################################
	// #############################  Delete Strategies ############################################
	// All delete operations will still allow the tree to remain in a valid state.
	// Worst case scenario the resulting tree will have a root of terminal value 0.

	// StrategyDeleteNonTerminal will select any random non-terminal element from a given tree and delete it by
	// setting it to 0. This can actually fell a tree if the tree is small enough
	StrategyDeleteNonTerminal = "DeleteNonTerminalR"
	//StrategyDeleteMalicious = "DeleteMaliciousR"
	// StrategyDeleteTerminal will convert a terminal node to 0.
	StrategyDeleteTerminal = "DeleteTerminalR"
	// StrategyMutateNode randomly selects a non-terminal in a tree and changes its value to one of the available
	// nonterminals in the parameter list.
	// If the tree only contains a root that is a terminal it will ignore it.
	StrategyMutateNonTerminal = "MutateNonTerminalR"
	// StrategyMutateTerminal randomly selects a terminal in a tree and changes its value to one of the available
	// terminals in the parameter list.
	// If the tree only contains a root that is a terminal it will ignore it.
	StrategyMutateTerminal = "MutateTerminalR"
	//// StrategyReplaceBranch takes a given tree and randomly selects a branch i.
	//// e non-terminal and will swap it with a randomly generated tree of variable depth
	StrategyReplaceBranch = "ReplaceBranchR"
	//StrategyReplaceBranchX = "ReplaceBranchXR"
	//StrategyAppendRandomOperation is a generic strategy that adds a randomly generated subtree to the tail of the given tree
	//  If an add strategy encounters a 0 at the root, it will replace the 0.
	StrategyAppendRandomOperation = "AppendRandomOperationR"

	// ####################################################### DETERMINISTIC STRATEGIES #############################
	// StrategySkip performs no operations on the given subtree.
	StrategySkip = "SkipD"

	// StrategyFellTree destroys the tree and sets its root to 0 and kills it all.
	StrategyFellTree = "FellTreeD"
	StrategyMultXD   = "MultKD"
	StrategyAddXD    = "AddKD"
	StrategySubXD    = "SubKD"
	StrategyDivXD    = "DivKD"

	StrategyMultCD = "MultCD"
	StrategyAddCD  = "AddCD"
	StrategySubCD  = "SubCD"
	StrategyDivCD  = "DivCD"

	//Strategy
)

// StratToFloat converts a given strategy to an int for numeric representation.
func StratToFloat(strategy Strategy) float64 {
	switch strategy {
	case StrategyDeleteNonTerminal:
		return 0
	case StrategyDeleteTerminal:
		return 1
	case StrategyMutateNonTerminal:
		return 2
	case StrategyMutateTerminal:
		return 4
	case StrategyReplaceBranch:
		return 5
	case StrategyAppendRandomOperation:
		return 5
	case StrategySkip:
		return 6
	case StrategyFellTree:
		return 7
	case StrategyMultXD:
		return 8
	case StrategyAddXD:
		return 9
	case StrategySubXD:
		return 10
	case StrategyDivXD:
		return 11
	case StrategyMultCD:
		return 12
	case StrategyAddCD:
		return 13
	case StrategySubCD:
		return 14
	case StrategyDivCD:
		return 15
	default:
		panic("Invalid strategy!")
	}
}

var (
	AllStrategies = []Strategy{StrategyDeleteNonTerminal, StrategyDeleteTerminal, StrategyMutateNonTerminal, StrategyMutateTerminal, StrategyReplaceBranch, StrategyAppendRandomOperation, StrategySkip, StrategyFellTree, StrategyMultXD, StrategyAddXD, StrategySubXD, StrategyDivXD,
		StrategyMultCD, StrategyAddCD, StrategySubCD, StrategyDivCD}
)

// GenerateRandomStrategy creates a random Strategy list that contains some or all of the availableStrategies.
// They are randomly selected and populated.
func GenerateRandomStrategy(number int, availableStrategies []Strategy) []Strategy {
	if number < 1 {
		number = 1
	}

	if availableStrategies == nil || len(availableStrategies) < 1 {
		return []Strategy{}
	}

	strategies := make([]Strategy, number)

	for i := 0; i < number; i++ {
		strategyIndex := rand.Intn(len(availableStrategies))
		strategies[i] = availableStrategies[strategyIndex]
	}

	return strategies
}
