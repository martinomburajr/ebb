package evolution

type BasicPlot struct {
	topology string
	// data will be e.g. KRTAverages
	YAxis []float64
	// XAxis generations.
	XAxis []int
}

type TopologyPlot struct {
	Topology string

	AvgOfAllAntagonistsInGen  []float64
	AvgOfAllProtagonistsInGen []float64

	TopAntagonistsBestInGen  []float64
	TopProtagonistsBestInGen []float64

	XAxisGen []int

	SupremeAntagonistStrategies   []Strategy
	SupremeProtagonistStrategyies []Strategy
}
