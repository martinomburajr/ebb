package evolution

// GenerationResult is the output of a complete generation given to the Engine. The Generation itself will
// be garbage collected but the statistics stored in this GenerationResult struct will be used by the Engine
// to compute overall statistics.
type GenerationResult struct {
	ID uint32
	//Individuals
	BestAntagonist  Individual
	BestProtagonist Individual

	// Averages of all Antagonists and Protagonists in Generation
	Correlation float64
	Covariance  float64

	AllAntagonistAverageFitness float64
	AntagonistStdDev            float64
	AntagonistVariance          float64

	AllProtagonistAverageFitness float64
	ProtagonistStdDev            float64
	ProtagonistVariance          float64

	AntagonistAvgAge  float64
	ProtagonistAvgAge float64

	AntagonistAvgBirthGen  float64
	ProtagonistAvgBirthGen float64
}

func (g GenerationResult) Clone() GenerationResult {
	g.BestAntagonist = g.BestAntagonist.Clone(-1)
	g.BestProtagonist = g.BestProtagonist.Clone(-1)

	return g
}
