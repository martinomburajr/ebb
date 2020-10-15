package evolution

import (
	"math/rand"
)

func Mutate(parents []Individual, children []Individual, strategies []Strategy, paramProbOfMutation float64) (outgoingParents []Individual, outgoingChildren []Individual, err error) {
	return applyMutation(parents, children, strategies, paramProbOfMutation)
}

func applyMutation(parents, children []Individual, strategies []Strategy, paramProbOfMutation float64) (outgoingParents []Individual, outgoingChildren []Individual, err error) {
	for i := 0; i < len(parents); i++ {
		probabilityOfMutation := rand.Float64()

		if probabilityOfMutation < paramProbOfMutation {
			err := parents[i].Mutate(strategies)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	// childs
	for i := 0; i < len(children); i++ {
		probabilityOfMutation := rand.Float64()

		if probabilityOfMutation < paramProbOfMutation {
			err := children[i].Mutate(strategies)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return parents, children, nil
}
