package evolution

const (
	MinAllowableGenerationsToTerminate  = 9
	TopologyRoundRobin                  = "TopologyRoundRobin"
	TopologyKRandom                     = "TopologyKRandom"
	TopologyHallOfFame                  = "TopologyHallOfFame"
	TopologySingleEliminationTournament = "TopologySET"
)

// Evolver outlines what any Evolutionary Topology should be able to do, define a Topology and Evolve
type Evolver interface {
	Topology(currentGeneration Generation, params EvolutionParams) (curr Generation, nextGen Generation, err error)
	Name(otherInfo string) string
}
