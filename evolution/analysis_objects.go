package evolution


///////////// NEW
type CSVStrat struct {

	Num int  `csv:"num"`
	KRTTopAStrat string `csv:"KRTTopAStratIS"`
	HOFTopAStrat string `csv:"HOFTopAStratIS"`
	RRTopAStrat string `csv:"RRTopAStratIS"`
	SETTopAStrat string `csv:"SETTopAStratIS"`

	KRTTopPStrat string `csv:"KRTTopPStratIS"`
	HOFTopPStrat string `csv:"HOFTopPStratIS"`
	RRTopPStrat string `csv:"RRTopPStratIS"`
	SETTopPStrat string `csv:"SETTopPStratIS"`
}


type CSVSim struct {
	SpecEquation    string `csv:"specEquation"`
	SpecEquationLen int    `csv:"specEquationLen"`
	IVarCount       int    `csv:"iVarCount"`
	PolDegree       int    `csv:"polDeg"`

	KRTTopAEquation          string `csv:"KRTTopAEquationIS"`
	KRTTopAEquationPolDegree int    `csv:"KRTTopAEquationPDIS"`
	HoFTopAEquation          string `csv:"HOFTopAEquationIS"`
	HoFTopAEquationPolDegree int    `csv:"HoFTopAEquationPDIS"`
	RRTopAEquation           string `csv:"RRTopAEquationIS"`
	RRTopAEquationPolDegree  int    `csv:"RRTopAEquationPDIS"`
	SETTopAEquation          string `csv:"SETTopAEquationIS"`
	SETTopAEquationPolDegree int    `csv:"SETTopAEquationPDIS"`

	KRTTopPEquation          string `csv:"KRTTopPEquationIS"`
	KRTTopPEquationPolDegree int    `csv:"KRTTopPEquationPDIS"`
	HoFTopPEquation          string `csv:"HOFTopPEquationIS"`
	HoFTopPEquationPolDegree int    `csv:"HoFTopPEquationPDIS"`
	RRTopPEquation           string `csv:"RRTopPEquationIS"`
	RRTopPEquationPolDegree  int    `csv:"RRTopPEquationPDIS"`
	SETTopPEquation          string `csv:"SETTopPEquationIS"`
	SETTopPEquationPolDegree int    `csv:"SETTopPEquationPDIS"`

	KRTTopAntagonistBestFitnessInSim float64 `csv:"KRTTopABestFitIS"`
	HoFTopAntagonistBestFitnessInSim float64 `csv:"HoFTopABestFitIS"`
	RRTTopAntagonistBestFitnessInSim float64 `csv:"RRTTopABestFitIS"`
	SETTopAntagonistBestFitnessInSim float64 `csv:"SETTopABestFitIS"`

	KRTTopProtagonistBestFitnessInSim float64 `csv:"KRTTopPBestFitIS"`
	HoFTopProtagonistBestFitnessInSim float64 `csv:"HoFTopPBestFitIS"`
	RRTTopProtagonistBestFitnessInSim float64 `csv:"RRTTopPBestFitIS"`
	SETTopProtagonistBestFitnessInSim float64 `csv:"SETTopPBestFitIS"`

	KRTTopAntagonistBirthGenInSim int `csv:"KRTTopABirthGenIS"`
	HoFTopAntagonistBirthGenInSim int `csv:"HoFTopABirthGenIS"`
	RRTTopAntagonistBirthGenInSim int `csv:"RRTTopABirthGenIS"`
	SETTopAntagonistBirthGenInSim int `csv:"SETTopABirthGenIS"`

	KRTTopProtagonistBirthGenInSim int `csv:"KRTTopPBirthGenIS"`
	HoFTopProtagonistBirthGenInSim int `csv:"HoFTopPBirthGenIS"`
	RRTTopProtagonistBirthGenInSim int `csv:"RRTTopPBirthGenIS"`
	SETTopProtagonistBirthGenInSim int `csv:"SETTopPBirthGenIS"`

	KRTTopAntagonistAgeInSim int `csv:"KRTTopAAgeIS"`
	HoFTopAntagonistAgeInSim int `csv:"HoFTopAAgeIS"`
	RRTTopAntagonistAgeInSim int `csv:"RRTTopAAgeIS"`
	SETTopAntagonistAgeInSim int `csv:"SETTopAAgeIS"`

	KRTTopProtagonistAgeInSim int `csv:"KRTTopPAgeIS"`
	HoFTopProtagonistAgeInSim int `csv:"HoFTopPAgeIS"`
	RRTTopProtagonistAgeInSim int `csv:"RRTTopPAgeIS"`
	SETTopProtagonistAgeInSim int `csv:"SETTopPAgeIS"`

	KRTTopAntagonistNoCompetitionsInSim int `csv:"KRTTopANoCIS"`
	HoFTopAntagonistNoCompetitionsInSim int `csv:"HoFTopANoCIS"`
	RRTTopAntagonistNoCompetitionsInSim int `csv:"RRTTopANoCIS"`
	SETTopAntagonistNoCompetitionsInSim int `csv:"SETTopANoCIS"`

	KRTTopProtagonistNoCompetitionsInSim int `csv:"KRTTopPNoCIS"`
	HoFTopProtagonistNoCompetitionsInSim int `csv:"HoFTopPNoCIS"`
	RRTTopProtagonistNoCompetitionsInSim int `csv:"RRTTopPNoCIS"`
	SETTopProtagonistNoCompetitionsInSim int `csv:"SETTopPNoCIS"`

	KRTTopAntagonistStrategyInSim string `csv:"KRTTopAStratIS"`
	HoFTopAntagonistStrategyInSim string `csv:"HoFTopAStratIS"`
	RRTTopAntagonistStrategyInSim string `csv:"RRTTopAStratIS"`
	SETTopAntagonistStrategyInSim string `csv:"SETTopAStratIS"`

	KRTTopProtagonistStrategyInSim string `csv:"KRTTopPStratIS"`
	HoFTopProtagonistStrategyInSim string `csv:"HoFTopPStratIS"`
	RRTTopProtagonistStrategyInSim string `csv:"RRTTopPStratIS"`
	SETTopProtagonistStrategyInSim string `csv:"SETTopPStratIS"`

	_ string `csv:"top-gen-spacer"`

	KRTAntagonistsAvgAgeInSim float64 `csv:"KRTAAvgAgeIS"`
	HoFAntagonistsAvgAgeInSim float64 `csv:"HOFAAvgAgeIS"`
	RRTAntagonistsAvgAgeInSim float64 `csv:"RRAAvgAgeIS"`
	SETAntagonistsAvgAgeInSim float64 `csv:"SETAAvgAgeIS"`

	KRTProtagonistsAvgAgeInSim float64 `csv:"KRTPAvgAgeIS"`
	HoFProtagonistsAvgAgeInSim float64 `csv:"HOFPAvgAgeIS"`
	RRTProtagonistsAvgAgeInSim float64 `csv:"RRPAvgAgeIS"`
	SETProtagonistsAvgAgeInSim float64 `csv:"SETPAvgAgeIS"`

	KRTAntagonistsAvgBirthGenInSim float64 `csv:"KRTAAvgBirthGenIS"`
	HoFAntagonistsAvgBirthGenInSim float64 `csv:"HOFAAvgBirthGenIS"`
	RRTAntagonistsAvgBirthGenInSim float64 `csv:"RRAAvgBirthGenIS"`
	SETAntagonistsAvgBirthGenInSim float64 `csv:"SETAAvgBirthGenIS"`

	KRTProtagonistsAvgBirthGenInSim float64 `csv:"KRTPAvgBirthGenIS"`
	HoFProtagonistsAvgBirthGenInSim float64 `csv:"HOFPAvgBirthGenIS"`
	RRTProtagonistsAvgBirthGenInSim float64 `csv:"RRPAvgBirthGenIS"`
	SETProtagonistsAvgBirthGenInSim float64 `csv:"SETPAvgBirthGenIS"`

	// Mean of all antagonists in the generation generations
	KRTAntagonistsMeanInSim float64 `csv:"KRTAMeanIS"`
	HoFAntagonistsMeanInSim float64 `csv:"HOFAMeanIS"`
	RRAntagonistsMeanInSim  float64 `csv:"RRAMeanIS"`
	SETAntagonistsMeanInSim float64 `csv:"SETAMeanIS"`

	KRTProtagonistsMeanInSim float64 `csv:"KRTPMeanIS"`
	HoFProtagonistsMeanInSim float64 `csv:"HOFPMeanIS"`
	RRProtagonistsMeanInSim  float64 `csv:"RRPMeanIS"`
	SETProtagonistsMeanInSim float64 `csv:"SETPMeanIS"`

	// TopIndividualStdDevs
	KRTAntagonistStdDevInSim float64 `csv:"KRTAStdIS"`
	HoFAntagonistStdDevInSim float64 `csv:"HOFAStdIS"`
	RRAntagonistStdDevInSim  float64 `csv:"RRAStdIS"`
	SETAntagonistStdDevInSim float64 `csv:"SETAStdIS"`

	KRTProtagonistStdDevInSim float64 `csv:"KRTPStdIS"`
	HoFProtagonistStdDevInSim float64 `csv:"HOFPStdIS"`
	RRProtagonistStdDevInSim  float64 `csv:"RRPStdIS"`
	SETProtagonistStdDevInSim float64 `csv:"SETPStdIS"`

	// Variance of Individuals
	KRTAntagonistVarInSim float64 `csv:"KRTAVarIS"`
	HoFAntagonistVarInSim float64 `csv:"HOFAVarIS"`
	RRAntagonistVarInSim  float64 `csv:"RRAVarIS"`
	SETAntagonistVarInSim float64 `csv:"SETAVarIS"`

	KRTProtagonistVarInSim float64 `csv:"KRTPVarIS"`
	HoFProtagonistVarInSim float64 `csv:"HOFPVarIS"`
	RRProtagonistVarInSim  float64 `csv:"RRPVarIS"`
	SETProtagonistVarInSim float64 `csv:"SETPVarIS"`
}

type CSVAvgGenerationsCombinedAcrossRuns struct {
	//FileID string                    `csv:"ID"`
	//params evolution.EvolutionParams `csv:"evolutionaryParams"`

	//BEST INDIVIDUAL
	Generation int `csv:"gen"`

	SpecEquation    string `csv:"specEquation"`
	SpecEquationLen int    `csv:"specEquationLen"`
	IVarCount       int    `csv:"iVarCount"`
	PolDegree       int    `csv:"polDeg"`

	KRTTopAEquation          string `csv:"KRTTopAEquationIS"`
	KRTTopAEquationPolDegree int    `csv:"KRTTopAEquationPDIS"`
	HoFTopAEquation          string `csv:"HOFTopAEquationIS"`
	HoFTopAEquationPolDegree int    `csv:"HoFTopAEquationPDIS"`
	RRTopAEquation           string `csv:"RRTopAEquationIS"`
	RRTopAEquationPolDegree  int    `csv:"RRTopAEquationPDIS"`
	SETTopAEquation          string `csv:"SETTopAEquationIS"`
	SETTopAEquationPolDegree int    `csv:"SETTopAEquationPDIS"`

	KRTTopPEquation          string `csv:"KRTTopPEquationIS"`
	KRTTopPEquationPolDegree int    `csv:"KRTTopPEquationPDIS"`
	HoFTopPEquation          string `csv:"HOFTopPEquationIS"`
	HoFTopPEquationPolDegree int    `csv:"HoFTopPEquationPDIS"`
	RRTopPEquation           string `csv:"RRTopPEquationIS"`
	RRTopPEquationPolDegree  int    `csv:"RRTopPEquationPDIS"`
	SETTopPEquation          string `csv:"SETTopPEquationIS"`
	SETTopPEquationPolDegree int    `csv:"SETTopPEquationPDIS"`

	KRTTopAntagonistBestFitnessInSim float64 `csv:"KRTTopABestFitIS"`
	HoFTopAntagonistBestFitnessInSim float64 `csv:"HoFTopABestFitIS"`
	RRTTopAntagonistBestFitnessInSim float64 `csv:"RRTTopABestFitIS"`
	SETTopAntagonistBestFitnessInSim float64 `csv:"SETTopABestFitIS"`

	KRTTopProtagonistBestFitnessInSim float64 `csv:"KRTTopPBestFitIS"`
	HoFTopProtagonistBestFitnessInSim float64 `csv:"HoFTopPBestFitIS"`
	RRTTopProtagonistBestFitnessInSim float64 `csv:"RRTTopPBestFitIS"`
	SETTopProtagonistBestFitnessInSim float64 `csv:"SETTopPBestFitIS"`

	KRTTopAntagonistBirthGenInSim int `csv:"KRTTopABirthGenIS"`
	HoFTopAntagonistBirthGenInSim int `csv:"HoFTopABirthGenIS"`
	RRTTopAntagonistBirthGenInSim int `csv:"RRTTopABirthGenIS"`
	SETTopAntagonistBirthGenInSim int `csv:"SETTopABirthGenIS"`

	KRTTopProtagonistBirthGenInSim int `csv:"KRTTopPBirthGenIS"`
	HoFTopProtagonistBirthGenInSim int `csv:"HoFTopPBirthGenIS"`
	RRTTopProtagonistBirthGenInSim int `csv:"RRTTopPBirthGenIS"`
	SETTopProtagonistBirthGenInSim int `csv:"SETTopPBirthGenIS"`

	KRTTopAntagonistAgeInSim int `csv:"KRTTopAAgeIS"`
	HoFTopAntagonistAgeInSim int `csv:"HoFTopAAgeIS"`
	RRTTopAntagonistAgeInSim int `csv:"RRTTopAAgeIS"`
	SETTopAntagonistAgeInSim int `csv:"SETTopAAgeIS"`

	KRTTopProtagonistAgeInSim int `csv:"KRTTopPAgeIS"`
	HoFTopProtagonistAgeInSim int `csv:"HoFTopPAgeIS"`
	RRTTopProtagonistAgeInSim int `csv:"RRTTopPAgeIS"`
	SETTopProtagonistAgeInSim int `csv:"SETTopPAgeIS"`

	KRTTopAntagonistStrategyInSim string `csv:"KRTTopAStratIS"`
	HoFTopAntagonistStrategyInSim string `csv:"HoFTopAStratIS"`
	RRTTopAntagonistStrategyInSim string `csv:"RRTTopAStratIS"`
	SETTopAntagonistStrategyInSim string `csv:"SETTopAStratIS"`

	KRTTopProtagonistStrategyInSim string `csv:"KRTTopPStratIS"`
	HoFTopProtagonistStrategyInSim string `csv:"HoFTopPStratIS"`
	RRTTopProtagonistStrategyInSim string `csv:"RRTTopPStratIS"`
	SETTopProtagonistStrategyInSim string `csv:"SETTopPStratIS"`

	_ string `csv:"top-gen-spacer"`

	// Mean of all antagonists in the generation generations
	KRTAntagonistsMeanInSim float64 `csv:"KRTAMeanIS"`
	HoFAntagonistsMeanInSim float64 `csv:"HOFAMeanIS"`
	RRAntagonistsMeanInSim  float64 `csv:"RRAMeanIS"`
	SETAntagonistsMeanInSim float64 `csv:"SETAMeanIS"`

	KRTProtagonistsMeanInSim float64 `csv:"KRTPMeanIS"`
	HoFProtagonistsMeanInSim float64 `csv:"HOFPMeanIS"`
	RRProtagonistsMeanInSim  float64 `csv:"RRPMeanIS"`
	SETProtagonistsMeanInSim float64 `csv:"SETPMeanIS"`

	// TopIndividualStdDevs
	KRTAntagonistStdDevInSim float64 `csv:"KRTAStdIS"`
	HoFAntagonistStdDevInSim float64 `csv:"HOFAStdIS"`
	RRAntagonistStdDevInSim  float64 `csv:"RRAStdIS"`
	SETAntagonistStdDevInSim float64 `csv:"SETAStdIS"`

	KRTProtagonistStdDevInSim float64 `csv:"KRTPStdIS"`
	HoFProtagonistStdDevInSim float64 `csv:"HOFPStdIS"`
	RRProtagonistStdDevInSim  float64 `csv:"RRPStdIS"`
	SETProtagonistStdDevInSim float64 `csv:"SETPStdIS"`

	// Variance of Individuals
	KRTAntagonistVarInSim float64 `csv:"KRTAVarIS"`
	HoFAntagonistVarInSim float64 `csv:"HOFAVarIS"`
	RRAntagonistVarInSim  float64 `csv:"RRAVarIS"`
	SETAntagonistVarInSim float64 `csv:"SETAVarIS"`

	KRTProtagonistVarInSim float64 `csv:"KRTPVarIS"`
	HoFProtagonistVarInSim float64 `csv:"HOFPVarIS"`
	RRProtagonistVarInSim  float64 `csv:"RRPVarIS"`
	SETProtagonistVarInSim float64 `csv:"SETPVarIS"`

	KRTAntagonistsAvgAgeInSim float64 `csv:"KRTAAvgAgeIS"`
	HoFAntagonistsAvgAgeInSim float64 `csv:"HOFAAvgAgeIS"`
	RRTAntagonistsAvgAgeInSim float64 `csv:"RRAAvgAgeIS"`
	SETAntagonistsAvgAgeInSim float64 `csv:"SETAAvgAgeIS"`

	KRTProtagonistsAvgAgeInSim float64 `csv:"KRTPAvgAgeIS"`
	HoFProtagonistsAvgAgeInSim float64 `csv:"HOFPAvgAgeIS"`
	RRTProtagonistsAvgAgeInSim float64 `csv:"RRPAvgAgeIS"`
	SETProtagonistsAvgAgeInSim float64 `csv:"SETPAvgAgeIS"`

	KRTTopAntagonistsAvgBirthGenInSim float64 `csv:"KRTTopAAvgBirthGenIS"`
	HoFTopAntagonistsAvgBirthGenInSim float64 `csv:"HOFTopAAvgBirthGenIS"`
	RRTTopAntagonistsAvgBirthGenInSim float64 `csv:"RRTopAAvgBirthGenIS"`
	SETTopAntagonistsAvgBirthGenInSim float64 `csv:"SETTopAAvgBirthGenIS"`

	KRTTopProtagonistsAvgBirthGenInSim float64 `csv:"KRTTopPAvgBirthGenIS"`
	HoFTopProtagonistsAvgBirthGenInSim float64 `csv:"HOFTopPAvgBirthGenIS"`
	RRTopProtagonistsAvgBirthGenInSim  float64 `csv:"RRTopPAvgBirthGenIS"`
	SETTopProtagonistsAvgBirthGenInSim float64 `csv:"SETTopPAvgBirthGenIS"`

	KRTTopAntagonistNoCompetitionsInSim int `csv:"KRTTopAntagonistNoCIS"`
	HoFTopAntagonistNoCompetitionsInSim  int  `csv:"HOFTopAntagonistNoCIS"`
	RRTTopAntagonistNoCompetitionsInSim  int `csv:"RRTopAntagonistNoCIS"`
	SETTopAntagonistNoCompetitionsInSim  int `csv:"SETTopAntagonistNoCIS"`

	KRTTopProtagonistNoCompetitionsInSim int `csv:"KRTTopProtagonistNoCIS"`
	HoFTopProtagonistNoCompetitionsInSim int `csv:"HOFTopProtagonistNoCIS"`
	RRTTopProtagonistNoCompetitionsInSim int `csv:"RRTopProtagonistNoCIS"`
	SETTopProtagonistNoCompetitionsInSim int `csv:"SETTopProtagonistNoCIS"`

	_ string `csv:"gen-spacer"`


	KRTTopAntagonistStrategyARRInSim []Strategy `csv:"KRTTopStratArrInSim"`
	HoFTopAntagonistStrategyARRInSim []Strategy `csv:"KRTTopStratArrInSim"`
	RRTTopAntagonistStrategyARRInSim []Strategy `csv:"KRTTopStratArrInSim"`
	SETTopAntagonistStrategyARRInSim []Strategy `csv:"KRTTopAStratArrInSim"`

	KRTTopProtagonistStrategyARRInSim []Strategy `csv:"KRTTopPStratArrInSim"`
	HoFTopProtagonistStrategyARRInSim []Strategy `csv:"KRTTopPStratArrInSim"`
	RRTTopProtagonistStrategyARRInSim []Strategy `csv:"KRTopPStratArrInSim"`
	SETTopProtagonistStrategyARRInSim []Strategy `csv:"KRTTopPStratArrSim"`

	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL
	/////######################################################### GENERATIONAL

	// Average Antagonists Age of Individuals in that Generation
	KRTAntagonistAverageAgeInGen float64 `csv:"KRTAAvgAgeIG"`
	HoFAntagonistAverageAgeInGen float64 `csv:"HOFAAvgAgeIG"`
	RRAntagonistAverageAgeInGen  float64 `csv:"RRAAvgAgeIG"`
	SETAntagonistAverageAgeInGen float64 `csv:"SETAAvgAgeIG"`

	KRTProtagonistAverageAgeInGen float64 `csv:"KRTPAvgAgeIG"`
	HoFProtagonistAverageAgeInGen float64 `csv:"HOFPAvgAgeIG"`
	RRProtagonistAverageAgeInGen  float64 `csv:"RRPAvgAgeIG"`
	SETProtagonistAverageAgeInGen float64 `csv:"SETPAvgAgeIG"`

	// Top Individual Mean Fitness in generation
	KRTTopAntagonistsMeanInGen float64 `csv:"KRTTopAMeanIG"`
	HoFTopAntagonistsMeanInGen float64 `csv:"HOFTopAMeanIG"`
	RRTopAntagonistsMeanInGen  float64 `csv:"RRTopAMeanIG"`
	SETTopAntagonistsMeanInGen float64 `csv:"SETTopAMeanIG"`

	KRTTopProtagonistsMeanInGen float64 `csv:"KRTTopPMeanIG"`
	HoFTopProtagonistsMeanInGen float64 `csv:"HOFTopPMeanIG"`
	RRTopProtagonistsMeanInGen  float64 `csv:"RRTopPMeanIG"`
	SETTopProtagonistsMeanInGen float64 `csv:"SETTopPMeanIG"`

	// Mean of all antagonists in the generation generations
	KRTAntagonistsMeanInGen float64 `csv:"KRTAMeanIG"`
	HoFAntagonistsMeanInGen float64 `csv:"HOFAMeanIG"`
	RRAntagonistsMeanInGen  float64 `csv:"RRAMeanIG"`
	SETAntagonistsMeanInGen float64 `csv:"SETAMeanIG"`

	KRTProtagonistsMeanInGen float64 `csv:"KRTPMeanIG"`
	HoFProtagonistsMeanInGen float64 `csv:"HOFPMeanIG"`
	RRProtagonistsMeanInGen  float64 `csv:"RRPMeanIG"`
	SETProtagonistsMeanInGen float64 `csv:"SETPMeanIG"`

	// Best Fitness value of top Individual in generation
	KRTTopAntagonistBestFitnessInGen float64 `csv:"KRTTopABestIG"`
	HoFTopAntagonistBestFitnessInGen float64 `csv:"HOFTopABestIG"`
	RRTopAntagonistBestFitnessInGen  float64 `csv:"RRTopABestIG"`
	SETTopAntagonistBestFitnessInGen float64 `csv:"SETTopABestIG"`

	KRTTopProtagonistBestFitnessInGen float64 `csv:"KRTTopPBestIG"`
	HoFTopProtagonistBestFitnessInGen float64 `csv:"HOFTopPBestIG"`
	RRTopProtagonistBestFitnessInGen  float64 `csv:"RRTopPBestIG"`
	SETTopProtagonistBestFitnessInGen float64 `csv:"SETTopPBestIG"`

	// TopIndividualStdDev
	KRTTopAntagonistStdDevInGen float64 `csv:"KRTTopAStdIG"`
	HoFTopAntagonistStdDevInGen float64 `csv:"HOFTopAStdIG"`
	RRTopAntagonistStdDevInGen  float64 `csv:"RRTopAStdIG"`
	SETTopAntagonistStdDevInGen float64 `csv:"SETTopAStdIG"`

	KRTTopProtagonistStdDevInGen float64 `csv:"KRTTopPStdIG"`
	HoFTopProtagonistStdDevInGen float64 `csv:"HOFTopPStdIG"`
	RRTopProtagonistStdDevInGen  float64 `csv:"RRTopPStdIG"`
	SETTopProtagonistStdDevInGen float64 `csv:"SETTopPStdIG"`

	// Variance of TopIndividual
	KRTTopAntagonistVarInGen float64 `csv:"KRTTopAVarIG"`
	HoFTopAntagonistVarInGen float64 `csv:"HOFTopAVarIG"`
	RRTopAntagonistVarInGen  float64 `csv:"RRTopAVarIG"`
	SETTopAntagonistVarInGen float64 `csv:"SETTopAVarIG"`

	KRTTopProtagonistVarInGen float64 `csv:"KRTTopPVarIG"`
	HoFTopProtagonistVarInGen float64 `csv:"HOFTopPVarIG"`
	RRTopProtagonistVarInGen  float64 `csv:"RRTopPVarIG"`
	SETTopProtagonistVarInGen float64 `csv:"SETTopPVarIG"`

	KRTTopAntagonistAverageAgeInGen float64 `csv:"KRTATopAvgAgeIG"`
	HoFTopAntagonistAverageAgeInGen float64 `csv:"HOFATopAvgAgeIG"`
	RRTopAntagonistAverageAgeInGen  float64 `csv:"RRATopAvgAgeIG"`
	SETTopAntagonistAverageAgeInGen float64 `csv:"SETATopAvgAgeIG"`

	KRTTopProtagonistAverageAgeInGen float64 `csv:"KRTPTopAvgAgeIG"`
	HoFTopProtagonistAverageAgeInGen float64 `csv:"HOFPTopAvgAgeIG"`
	RRTopProtagonistAverageAgeInGen  float64 `csv:"RRPTopAvgAgeIG"`
	SETTopProtagonistAverageAgeInGen float64 `csv:"SETPTopAvgAgeIG"`

	KRTTopAntagonistBirthGenInGen float64 `csv:"KRTATopAvgBirthGenIG"`
	HoFTopAntagonistBirthGenInGen float64 `csv:"HOFATopAvgBirthGenIG"`
	RRTopAntagonistBirthGenInGen  float64 `csv:"RRATopAvgBirthGenIG"`
	SETTopAntagonistBirthGenInGen float64 `csv:"SETATopAvgBirthGenIG"`

	KRTTopProtagonistBirthGenInGen float64 `csv:"KRTPTopAvgBirthGenIG"`
	HoFTopProtagonistBirthGenInGen float64 `csv:"HOFPTopAvgBirthGenIG"`
	RRTopProtagonistBirthGenInGen  float64 `csv:"RRPTopAvgBirthGenIG"`
	SETTopProtagonistBirthGenInGen float64 `csv:"SETPTopAvgBirthGenIG"`

	KRTTopAntagonistStrategyInGen string `csv:"KRTATopAvgStrategyIG"`
	HoFTopAntagonistStrategyInGen string `csv:"HOFATopAvgStrategyIG"`
	RRTopAntagonistStrategyInGen  string `csv:"RRATopAvgStrategyIG"`
	SETTopAntagonistStrategyInGen string `csv:"SETATopAvgStrategyIG"`

	KRTTopProtagonistStrategyInGen string `csv:"KRTPTopAvgStrategyIG"`
	HoFTopProtagonistStrategyInGen string `csv:"HOFPTopAvgStrategyIG"`
	RRTopProtagonistStrategyInGen  string `csv:"RRPTopAvgStrategyIG"`
	SETTopProtagonistStrategyInGen string `csv:"SETPTopAvgStrategyIG"`



	KRTTopAntagonistDomStrategyInGen string `csv:"KRTATopAvgDomStrategyIG"`
	HoFTopAntagonistDomStrategyInGen string `csv:"HOFATopAvgDomStrategyIG"`
	RRTopAntagonistDomStrategyInGen  string `csv:"RRATopAvgDomStrategyIG"`
	SETTopAntagonistDomStrategyInGen string `csv:"SETATopAvgDomStrategyIG"`

	KRTTopProtagonistDomStrategyInGen    string `csv:"KRTPTopAvgDomStrategyIG"`
	HoFTopProtagonistDomStrategyInGen    string `csv:"HOFPTopAvgDomStrategyIG"`
	RRTopProtagonistDomStrategyInGen     string `csv:"RRPTopAvgDomStrategyIG"`
	SETTopProtagonistDomStrategyInGen    string `csv:"SETPTopAvgDomStrategyIG"`


}
