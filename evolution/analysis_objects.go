package evolution

import (
	"github.com/martinomburajr/masters-go/evolution"
	"github.com/martinomburajr/masters-go/simulation"
)

func ScaleCCAlgorithmToOrdinal(algorithm string) int {
	switch algorithm {
	case evolution.TopologyHallOfFame:
		return 0
	case evolution.TopologyRoundRobin:
		return 1
	case evolution.TopologyKRandom:
		return 2
	case evolution.TopologySingleEliminationTournament:
		return 3
	}
	return -1
}

func ScaleCrosoverToOrdinal(crossover string) int {
	switch crossover {
	case evolution.CrossoverSinglePoint:
		return 0
	case evolution.CrossoverUniform:
		return 1
	}
	return -1
}

type CSVBestAll struct {
	FileID                  string                                `csv:"ID"`
	bestIndividualStatistic simulation.RunBestIndividualStatistic `csv:"bestIndividualStatistic"`
	params                  evolution.EvolutionParams             `csv:"evolutionaryParams"`

	//BEST INDIVIDUAL
	SpecEquation string `csv:"specEquation"`
	SpecRange    int    `csv:"range"`
	SpecSeed     int    `csv:"seed"`

	AntagonistID           string  `csv:"AID"`
	ProtagonistID          string  `csv:"PID"`
	Antagonist             float64 `csv:"AAvg"`
	Protagonist            float64 `csv:"PAvg"`
	AntagonistBestFitness  float64 `csv:"ABestFit"`
	ProtagonistBestFitness float64 `csv:"PBestFit"`
	AntagonistStdDev       float64 `csv:"AStdDev"`
	ProtagonistStdDev      float64 `csv:"PStdDev"`
	//AntagonistAverageDelta      float64 `csv:"AAvgDelta"`
	//ProtagonistAverageDelta     float64 `csv:"PAvgDelta"`
	//AntagonistBestDelta         float64 `csv:"ABestDelta"`
	//ProtagonistBestDelta        float64 `csv:"PBestDelta"`
	AntagonistEquation          string `csv:"AEquation"`
	ProtagonistEquation         string `csv:"PEquation"`
	AntagonistStrategy          string `csv:"AStrat"`
	ProtagonistStrategy         string `csv:"PStrat"`
	AntagonistDominantStrategy  string `csv:"ADomStrat"`
	ProtagonistDominantStrategy string `csv:"PDomStrat"`
	AntagonistGeneration        int    `csv:"AGen"`
	ProtagonistGeneration       int    `csv:"PGen"`
	AntagonistBirthGen          int    `csv:"ABirthGen"`
	ProtagonistBirthGen         int    `csv:"PBirthGen"`
	AntagonistRun               int    `csv:"ARun"`
	ProtagonistRun              int    `csv:"PRun"`
	AntagonistAge               int    `csv:"AAge"`
	ProtagonistAge              int    `csv:"PAge"`
	AntagonistNoOComp           int    `csv:"ANoC"`
	ProtagonistNoOComp          int    `csv:"PNoC"`

	FinalAntagonist                  float64 `csv:"finAAvg"`
	FinalProtagonist                 float64 `csv:"finPAvg"`
	FinalAntagonistBestFitness       float64 `csv:"finABestFit"`
	FinalProtagonistBestFitness      float64 `csv:"finPBestFit"`
	FinalAntagonistStdDev            float64 `csv:"finAStdDev"`
	FinalProtagonistStdDev           float64 `csv:"finPStdDev"`
	FinalAntagonistAverageDelta      float64 `csv:"finAAvgDelta"`
	FinalProtagonistAverageDelta     float64 `csv:"finPAvgDelta"`
	FinalAntagonistBestDelta         float64 `csv:"finABestDelta"`
	FinalProtagonistBestDelta        float64 `csv:"finPBestDelta"`
	FinalAntagonistEquation          string  `csv:"finAEquation"`
	FinalProtagonistEquation         string  `csv:"finPEquation"`
	FinalAntagonistStrategy          string  `csv:"finAStrat"`
	FinalProtagonistStrategy         string  `csv:"finPStrat"`
	FinalAntagonistDominantStrategy  string  `csv:"finADomStrat"`
	FinalProtagonistDominantStrategy string  `csv:"finPDomStrat"`
	FinalAntagonistBirthGen          int     `csv:"finABirthGen"`
	FinalProtagonistBirthGen         int     `csv:"finPBirthGen"`
	FinalAntagonistAge               int     `csv:"finAAge"`
	FinalProtagonistAge              int     `csv:"finPAge"`
	FinalAntagonistNoOComp           int     `csv:"finANoC"`
	FinalProtagonistNoOComp          int     `csv:"finPNoC"`

	Run int `csv:"run"`

	// PARAMS
	GenerationCount    int     `csv:"genCount"`
	EachPopulationSize int     `csv:"popCount"`
	TopologyType       string  `csv:"topology"`
	AntStratCount      int     `csv:"antStratCount"`
	ProStratCount      int     `csv:"proStratCount"`
	AntStrat           string  `csv:"antStrat"`
	ProStrat           string  `csv:"proStrat"`
	RandTreeDepth      int     `csv:"randTreeDepth"`
	AntThreshMult      float64 `csv:"antThreshMult"`
	ProThresMult       float64 `csv:"proThresMult"`
	CrossPercent       float64 `csv:"crossPercent"`
	ProbMutation       float64 `csv:"probMutation"`
	ParentSelect       string  `csv:"parentSelect"`
	TournamentSize     int     `csv:"tournamentSize"`
	SurvivorSelect     string  `csv:"survivorSelect"`
	SurvivorPercent    float64 `csv:"survivorPercent"`
	DivByZero          string  `csv:"d0"`
	DivByZeroPen       float64 `csv:"d0Pen"`

	CrossoverType  string `csv:"crossoverType"`
	CrossoverScale int    `csv:"crossoverScale"`
	TopologyScale  int    `csv:"topologyScale"`
}

type CSVCombinedGenerations struct {
	FileID string                    `csv:"ID"`
	params evolution.EvolutionParams `csv:"evolutionaryParams"`

	//BEST INDIVIDUAL
	Generation   int    `csv:"gen"`
	Topology     string `csv:"topology"`
	SpecEquation string `csv:"specEquation"`
	//SpecRange    int    `csv:"range"`
	//SpecSeed     int    `csv:"seed"`

	TopAEquation string  `csv:"topAEquation"`
	TopPEquation string  `csv:"topPEquation"`
	Correlation  float64 `csv:"correlation"`
	//Covariance   float64 `csv:"covariance"`

	Antagonist                float64 `csv:"AMean"`
	Protagonist               float64 `csv:"PMean"`
	TopAntagonistMean         float64 `csv:"topAMean"`
	TopProtagonistMean        float64 `csv:"topPMean"`
	TopAntagonistBestFitness  float64 `csv:"topABest"`
	TopProtagonistBestFitness float64 `csv:"topPBest"`
	TopAntagonistStdDev       float64 `csv:"AStd"`
	TopProtagonistStdDev      float64 `csv:"PStd"`
	TopAntagonistVar          float64 `csv:"AVar"`
	TopProtagonistVar         float64 `csv:"PVar"`
	//TopAntagonistSkew         float64 `csv:"ASkew"`
	//TopProtagonistSkew        float64 `csv:"PSkew"`
	//TopAntagonistKurtosis     float64 `csv:"AExKur"`
	//TopProtagonistKurtosis    float64 `csv:"PExKur"`

	//TopAntagonistAverageDelta      float64 `csv:"topAMeanDelta"`
	//TopProtagonistAverageDelta     float64 `csv:"topPMeanDelta"`
	//TopAntagonistBestDelta         float64 `csv:"topABestDelta"`
	//TopProtagonistBestDelta        float64 `csv:"topPBestDelta"`
	TopAntagonistStrategy          string `csv:"topAStrat"`
	TopProtagonistStrategy         string `csv:"topPStrat"`
	TopAntagonistDominantStrategy  string `csv:"topADomStrat"`
	TopProtagonistDominantStrategy string `csv:"topPDomStrat"`
	TopAntagonistGeneration        int    `csv:"topAGen"`
	TopProtagonistGeneration       int    `csv:"topPGen"`
	TopAntagonistBirthGen          int    `csv:"topABirthGen"`
	TopProtagonistBirthGen         int    `csv:"topPBirthGen"`
	TopAntagonistAge               int    `csv:"topAAge"`
	TopProtagonistAge              int    `csv:"topPAge"`

	//Run int `csv:"run"`

	// PARAMS
	GenerationCount    int     `csv:"genCount"`
	EachPopulationSize int     `csv:"popCount"`
	TopologyType       string  `csv:"topology"`
	AntStratCount      int     `csv:"antStratCount"`
	ProStratCount      int     `csv:"proStratCount"`
	AntStrat           string  `csv:"antStrat"`
	ProStrat           string  `csv:"proStrat"`
	RandTreeDepth      int     `csv:"randTreeDepth"`
	AntThreshMult      float64 `csv:"antThreshMult"`
	ProThresMult       float64 `csv:"proThresMult"`
	CrossPercent       float64 `csv:"crossPercent"`
	ProbMutation       float64 `csv:"probMutation"`
	ParentSelect       string  `csv:"parentSelect"`
	TournamentSize     int     `csv:"tournamentSize"`
	SurvivorSelect     string  `csv:"survivorSelect"`
	SurvivorPercent    float64 `csv:"survivorPercent"`
	DivByZero          string  `csv:"d0"`
	DivByZeroPen       float64 `csv:"d0Pen"`
	TopologyScale      int     `csv:"topologyScale"`
	CrossoverScale     int     `csv:"crossoverScale"`
	CrossoverType      string  `csv:"crossoverType"`
}

///////////// NEW

type CSVAvgGenerationsCombinedAcrossRuns struct {
	//FileID string                    `csv:"ID"`
	//params evolution.EvolutionParams `csv:"evolutionaryParams"`

	//BEST INDIVIDUAL
	Generation int    `csv:"gen"`
	Topology   string `csv:"topology"`

	SpecEquation    string `csv:"specEquation"`
	SpecEquationLen int    `csv:"specEquationLen"`
	IVarCount                    int `csv:"iVarCount"`
	PolDegree       int    `csv:"polDeg"`

	KRTTopAEquation string `csv:"KRTTopAEquation"`
	HoFTopAEquation string `csv:"HOFTopAEquation"`
	RRTopAEquation  string `csv:"RRTopAEquation"`
	SETTopAEquation string `csv:"SETTopAEquation"`

	KRTTopPEquation string `csv:"KRTTopPEquation"`
	HoFTopPEquation string `csv:"HOFTopPEquation"`
	RRTopPEquation  string `csv:"RRTopPEquation"`
	SETTopPEquation string `csv:"SETTopPEquation"`

	///////////// ########################### AVERAGES #################################
	///////////// ########################### AVERAGES #################################
	///////////// ########################### AVERAGES #################################
	///////////// ########################### AVERAGES #################################

	// Mean of all antagonists in the generation generations
	KRTAntagonistsMean float64 `csv:"KRTAMean"`
	HoFAntagonistsMean float64 `csv:"HOFAMean"`
	RRAntagonistsMean  float64 `csv:"RRAMean"`
	SETAntagonistsMean float64 `csv:"SETAMean"`

	KRTProtagonistsMean float64 `csv:"KRTPMean"`
	HoFProtagonistsMean float64 `csv:"HOFPMean"`
	RRProtagonistsMean  float64 `csv:"RRPMean"`
	SETProtagonistsMean float64 `csv:"SETPMean"`

	// TopIndividualStdDevs
	KRTAntagonistStdDev float64 `csv:"KRTAStd"`
	HoFAntagonistStdDev float64 `csv:"HOFAStd"`
	RRAntagonistStdDev  float64 `csv:"RRAStd"`
	SETAntagonistStdDev float64 `csv:"SETAStd"`

	KRTProtagonistStdDev float64 `csv:"KRTPStd"`
	HoFProtagonistStdDev float64 `csv:"HOFPStd"`
	RRProtagonistStdDev  float64 `csv:"RRPStd"`
	SETProtagonistStdDev float64 `csv:"SETPStd"`

	// Variance of Individuals
	KRTAntagonistVar float64 `csv:"KRTAVar"`
	HoFAntagonistVar float64 `csv:"HOFAVar"`
	RRAntagonistVar  float64 `csv:"RRAVar"`
	SETAntagonistVar float64 `csv:"SETAVar"`

	KRTProtagonistVar float64 `csv:"KRTPVar"`
	HoFProtagonistVar float64 `csv:"HOFPVar"`
	RRProtagonistVar  float64 `csv:"RRPVar"`
	SETProtagonistVar float64 `csv:"SETPVar"`

	// Average Antagonists Age of Individuals in that Generation
	KRTAntagonistAverageAge float64 `csv:"KRTAAvgAge"`
	HoFAntagonistAverageAge float64 `csv:"HOFAAvgAge"`
	RRAntagonistAverageAge  float64 `csv:"RRAAvgAge"`
	SETAntagonistAverageAge float64 `csv:"SETAAvgAge"`

	KRTProtagonistAverageAge float64 `csv:"KRTPAvgAge"`
	HoFProtagonistAverageAge float64 `csv:"HOFPAvgAge"`
	RRProtagonistAverageAge  float64 `csv:"RRPAvgAge"`
	SETProtagonistAverageAge float64 `csv:"SETPAvgAge"`

	///////////// ########################### TOP INDIVIDUALS #################################
	///////////// ########################### TOP INDIVIDUALS #################################
	///////////// ########################### TOP INDIVIDUALS #################################
	///////////// ########################### TOP INDIVIDUALS #################################
	///////////// ########################### TOP INDIVIDUALS #################################

	// Top Individual Mean Fitness
	KRTTopAntagonistsMean float64 `csv:"KRTTopAMean"`
	HoFTopAntagonistsMean float64 `csv:"HOFTopAMean"`
	RRTopAntagonistsMean  float64 `csv:"RRTopAMean"`
	SETTopAntagonistsMean float64 `csv:"SETTopAMean"`

	KRTTopProtagonistsMean float64 `csv:"KRTTopPMean"`
	HoFTopProtagonistsMean float64 `csv:"HOFTopPMean"`
	RRTopProtagonistsMean  float64 `csv:"RRTopPMean"`
	SETTopProtagonistsMean float64 `csv:"SETTopPMean"`

	// Best Fitness value of top Individual in generation
	KRTTopAntagonistBestFitness float64 `csv:"KRTTopABest"`
	HoFTopAntagonistBestFitness float64 `csv:"HOFTopABest"`
	RRTopAntagonistBestFitness  float64 `csv:"RRTopABest"`
	SETTopAntagonistBestFitness float64 `csv:"SETTopABest"`

	KRTTopProtagonistBestFitness float64 `csv:"KRTTopPBest"`
	HoFTopProtagonistBestFitness float64 `csv:"HOFTopPBest"`
	RRTopProtagonistBestFitness  float64 `csv:"RRTopPBest"`
	SETTopProtagonistBestFitness float64 `csv:"SETTopPBest"`

	// TopIndividualStdDev
	KRTTopAntagonistStdDev float64 `csv:"KRTTopAStd"`
	HoFTopAntagonistStdDev float64 `csv:"HOFTopAStd"`
	RRTopAntagonistStdDev  float64 `csv:"RRTopAStd"`
	SETTopAntagonistStdDev float64 `csv:"SETTopAStd"`

	KRTTopProtagonistStdDev float64 `csv:"KRTTopPStd"`
	HoFTopProtagonistStdDev float64 `csv:"HOFTopPStd"`
	RRTopProtagonistStdDev  float64 `csv:"RRTopPStd"`
	SETTopProtagonistStdDev float64 `csv:"SETTopPStd"`

	// Variance of TopIndividual
	KRTTopAntagonistVar float64 `csv:"KRTTopAVar"`
	HoFTopAntagonistVar float64 `csv:"HOFTopAVar"`
	RRTopAntagonistVar  float64 `csv:"RRTopAVar"`
	SETTopAntagonistVar float64 `csv:"SETTopAVar"`
	// Average Individual Age of Top Individual in that Generation
	// Average BirthGen of Top Individual
	// Top Strategy
	// Top Individual Dominant Strategy

	KRTTopProtagonistVar float64 `csv:"KRTTopPVar"`
	HoFTopProtagonistVar float64 `csv:"HOFTopPVar"`
	RRTopProtagonistVar  float64 `csv:"RRTopPVar"`
	SETTopProtagonistVar float64 `csv:"SETTopPVar"`

	KRTTopAntagonistAverageAge float64 `csv:"KRTATopAvgAge"`
	HoFTopAntagonistAverageAge float64 `csv:"HOFATopAvgAge"`
	RRTopAntagonistAverageAge  float64 `csv:"RRATopAvgAge"`
	SETTopAntagonistAverageAge float64 `csv:"SETATopAvgAge"`

	KRTTopProtagonistAverageAge float64 `csv:"KRTPTopAvgAge"`
	HoFTopProtagonistAverageAge float64 `csv:"HOFPTopAvgAge"`
	RRTopProtagonistAverageAge  float64 `csv:"RRPTopAvgAge"`
	SETTopProtagonistAverageAge float64 `csv:"SETPTopAvgAge"`

	KRTTopAntagonistBirthGen float64 `csv:"KRTATopAvgBirthGen"`
	HoFTopAntagonistBirthGen float64 `csv:"HOFATopAvgBirthGen"`
	RRTopAntagonistBirthGen  float64 `csv:"RRATopAvgBirthGen"`
	SETTopAntagonistBirthGen float64 `csv:"SETATopAvgBirthGen"`

	KRTTopProtagonistBirthGen float64 `csv:"KRTPTopAvgBirthGen"`
	HoFTopProtagonistBirthGen float64 `csv:"HOFPTopAvgBirthGen"`
	RRTopProtagonistBirthGen  float64 `csv:"RRPTopAvgBirthGen"`
	SETTopProtagonistBirthGen float64 `csv:"SETPTopAvgBirthGen"`

	KRTTopAntagonistStrategy string `csv:"KRTATopAvgStrategy"`
	HoFTopAntagonistStrategy string `csv:"HOFATopAvgStrategy"`
	RRTopAntagonistStrategy  string `csv:"RRATopAvgStrategy"`
	SETTopAntagonistStrategy string `csv:"SETATopAvgStrategy"`

	KRTTopProtagonistStrategy string `csv:"KRTPTopAvgStrategy"`
	HoFTopProtagonistStrategy string `csv:"HOFPTopAvgStrategy"`
	RRTopProtagonistStrategy  string `csv:"RRPTopAvgStrategy"`
	SETTopProtagonistStrategy string `csv:"SETPTopAvgStrategy"`

	KRTTopAntagonistDomStrategy string `csv:"KRTATopAvgDomStrategy"`
	HoFTopAntagonistDomStrategy string `csv:"HOFATopAvgDomStrategy"`
	RRTopAntagonistDomStrategy  string `csv:"RRATopAvgDomStrategy"`
	SETTopAntagonistDomStrategy string `csv:"SETATopAvgDomStrategy"`

	KRTTopProtagonistDomStrategy string `csv:"KRTPTopAvgDomStrategy"`
	HoFTopProtagonistDomStrategy string `csv:"HOFPTopAvgDomStrategy"`
	RRTopProtagonistDomStrategy  string `csv:"RRPTopAvgDomStrategy"`
	SETTopProtagonistDomStrategy string `csv:"SETPTopAvgDomStrategy"`

}
