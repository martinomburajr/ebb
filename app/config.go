package app

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/martinomburajr/ebb/evolog"
	"github.com/martinomburajr/ebb/evolution"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type ApplicationConfig struct {
	EnableParallelism bool                      `json:"enableParallelism"`
	Params            evolution.EvolutionParams `json:"params"`
	Runs              int                       `json:"runs"`
	Complexity        int                       `json:"complexity"`
	OutputDir         string                    `json:"outputDir"`
	OutputFile        string                    `json:"outputFile"`

	Plots PlotInfo  `json:"plots"`
	Stats StatsInfo `json:"stats"`

	// auto-generated
	GeneratedStatsOutputPath string
	GeneratedPlotsOutputPath string

	// Iter is useful when running several simulations one ofter the other. It keeps track of the index in the main for loop in the main
	// function
	Iter                          int
	GeneratedStatsStratOutputPath string
}

type PlotInfo struct {
	Length int `json:"length"`
	Height int `json:"height"`

	OutputDir  string `json:"outputDir"`
	OutputFile string `json:"outputFile"`
}

type StatsInfo struct {
	OutputDir  string `json:"outputDir"`
	OutputFile string `json:"outputFile"`
	OutputFileStrat string `json:"outputFileStrat"`
}

type Application struct {
	Config ApplicationConfig
}

func NewApplication(configPath string) (Application, error) {
	file, err := os.OpenFile(configPath, os.O_RDONLY, 0777)
	if err != nil {
		return Application{}, fmt.Errorf("error opening app file - %s", err.Error())
	}

	config := ApplicationConfig{}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Application{}, fmt.Errorf("error parsing application app - %s", err.Error())
	}

	return Application{Config: config}, nil
}

// Begin starts the entire evolutionary process and starts the simulations
func (a Application) Begin() error {
	simulation, err := Init(&a.Config)
	if err != nil {
		return err
	}

	simulation.StartTime = time.Now()
	simulationResult, err := simulation.Start()
	if err != nil {
		return err
	}

	// ############################# STATS ############################

	randFolder := fmt.Sprintf("%s-%s", strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", ""), evolution.RandString(4))

	err = simulation.outputStats(randFolder, a.Config, simulationResult)
	if err != nil {
		return err
	}

	// ################################# PLOTS ##############################

	//err = simulation.outputPlot(randFolder, a.Config, simulationResult)
	//if err != nil {
	//	return err
	//}

	simulation.End = time.Now()
	simulation.LogTime("Simulation Duration")

	return err
}

func (s *Simulation) printSystemStats() {
	s.LogMessage(fmt.Sprintf("NumGoroutines: %d", runtime.NumGoroutine()), evolog.LoggerSimulation)
}

func (s *Simulation) outputStats(randFolder string, config ApplicationConfig, simulationResult evolution.SimulationResult) error {
	newStatsDir := fmt.Sprintf("%s/%d/%s", config.Stats.OutputDir, config.Complexity, randFolder)

	err := os.MkdirAll(newStatsDir, 0777)
	if err != nil {
		fmt.Println("output path cannot be created - defaulting to stats.csv")
	}

	s.Config.GeneratedStatsOutputPath = fmt.Sprintf("%s/stats.csv", newStatsDir)
	s.Config.GeneratedStatsStratOutputPath = fmt.Sprintf("%s/stats-strategy.csv", newStatsDir)

	csvData, statsCSV := simulationResult.GenerateCSVData(s.Config.Params)

	var file *os.File

	file, err = os.Create(s.Config.GeneratedStatsOutputPath)
	if err != nil {
		fmt.Println("output path cannot be created - defaulting to stats.csv")

		file, err = os.Create("stats.csv")
		if err != nil {
			log.Fatalf("failed to write to fallback (stats.csv), %v", err)
		}

	}

	stratFile, err := os.Create(s.Config.GeneratedStatsStratOutputPath)
	if err != nil {
		fmt.Println("output path cannot be created - defaulting to stats-strategy.csv")

		file, err = os.Create("stats-strategy.csv")
		if err != nil {
			log.Fatalf("failed to write to fallback (stats-strategy.csv), %v", err)
		}

	}

	err = simulationResult.WriteCSV(file, csvData)
	err = simulationResult.WriteStratCSV(stratFile, statsCSV)

	s.LogMessage("completed stats", evolog.LoggerAnalysis)

	return err
}

func (s *Simulation) outputPlot(randFolder string, config ApplicationConfig, simulationResult evolution.SimulationResult) error {
	newStatsDir := fmt.Sprintf("%s/%d/%s", config.Plots.OutputDir, config.Complexity, randFolder)

	err := os.MkdirAll(newStatsDir, 0777)
	if err != nil {
		fmt.Println("output path cannot be created - defaulting to stats.csv")
	}

	s.Config.GeneratedPlotsOutputPath = fmt.Sprintf("%s", newStatsDir)

	s.Plot(simulationResult)

	return err
}

// ParseAll begins the compaction of all data. It begins by combining all the run-level information into complexity-level
// information, which finally accumulates into the simulation-level statistics that encompasses compacted data of every
// single run performed.
func (a *Application) ParseAll() {
	//runs := make([]evolution.CSVAvgGenerationsCombinedAcrossRuns, 0)
	csvSims := make([]evolution.CSVSim, 0)
	//1. Parse all CSVs into CSVAvgGenerationsCombinedAcrossRuns objects
	for i := 0; i < 3; i++ {
		err := filepath.Walk(fmt.Sprintf("%s/%d", a.Config.Stats.OutputDir, i), func(path string, info os.FileInfo, err error) error {
			// If we come across a csv
			if strings.Contains(path, "csv") && !info.IsDir() {
				file, err := os.Open(path)
				defer file.Close()

				if err != nil {
					return fmt.Errorf("failed to open file at path %s | err: %v", path, err)
				}

				var runs []evolution.CSVAvgGenerationsCombinedAcrossRuns
				err = gocsv.UnmarshalFile(file, &runs)
				if err != nil {
					return fmt.Errorf("failed to unmarshal | err: %v", err)
				}

				//runs = append(runs, csv)
				run := runs[0]
				csvSim := evolution.CSVSim{
					SpecEquation:                      run.SpecEquation,
					SpecEquationLen:                   run.SpecEquationLen,
					IVarCount:                         run.IVarCount,
					PolDegree:                         run.PolDegree,
					KRTTopAEquation:                   run.KRTTopAEquation,
					KRTTopAEquationPolDegree:          run.KRTTopAEquationPolDegree,
					HoFTopAEquation:                   run.HoFTopAEquation,
					HoFTopAEquationPolDegree:          run.KRTTopAEquationPolDegree,
					RRTopAEquation:                    run.RRTopAEquation,
					RRTopAEquationPolDegree:           run.RRTopAEquationPolDegree,
					SETTopAEquation:                   run.SETTopAEquation,
					SETTopAEquationPolDegree:          run.SETTopAEquationPolDegree,
					KRTTopPEquation:                   run.KRTTopPEquation,
					KRTTopPEquationPolDegree:          run.KRTTopPEquationPolDegree,
					HoFTopPEquation:                   run.HoFTopPEquation,
					HoFTopPEquationPolDegree:          run.KRTTopPEquationPolDegree,
					RRTopPEquation:                    run.RRTopPEquation,
					RRTopPEquationPolDegree:           run.RRTopPEquationPolDegree,
					SETTopPEquation:                   run.SETTopPEquation,
					SETTopPEquationPolDegree:          run.SETTopPEquationPolDegree,
					KRTTopAntagonistBestFitnessInSim:  run.KRTTopAntagonistBestFitnessInSim,
					HoFTopAntagonistBestFitnessInSim:  run.HoFTopAntagonistBestFitnessInSim,
					RRTTopAntagonistBestFitnessInSim:  run.RRTTopAntagonistBestFitnessInSim,
					SETTopAntagonistBestFitnessInSim:  run.SETTopAntagonistBestFitnessInSim,
					KRTTopProtagonistBestFitnessInSim: run.KRTTopProtagonistBestFitnessInSim,
					HoFTopProtagonistBestFitnessInSim: run.HoFTopProtagonistBestFitnessInSim,
					RRTTopProtagonistBestFitnessInSim: run.RRTTopProtagonistBestFitnessInSim,
					SETTopProtagonistBestFitnessInSim: run.SETTopProtagonistBestFitnessInSim,
					KRTTopAntagonistBirthGenInSim:     run.KRTTopAntagonistBirthGenInSim,
					HoFTopAntagonistBirthGenInSim:     run.HoFTopAntagonistBirthGenInSim,
					RRTTopAntagonistBirthGenInSim:     run.RRTTopAntagonistBirthGenInSim,
					SETTopAntagonistBirthGenInSim:     run.SETTopAntagonistBirthGenInSim,
					KRTTopProtagonistBirthGenInSim:    run.KRTTopProtagonistBirthGenInSim,
					HoFTopProtagonistBirthGenInSim:    run.HoFTopProtagonistBirthGenInSim,
					RRTTopProtagonistBirthGenInSim:    run.RRTTopProtagonistBirthGenInSim,
					SETTopProtagonistBirthGenInSim:    run.SETTopProtagonistBirthGenInSim,
					KRTTopAntagonistAgeInSim:          run.KRTTopAntagonistAgeInSim,
					HoFTopAntagonistAgeInSim:          run.HoFTopAntagonistAgeInSim,
					RRTTopAntagonistAgeInSim:          run.RRTTopAntagonistAgeInSim,
					SETTopAntagonistAgeInSim:          run.SETTopAntagonistAgeInSim,
					KRTTopProtagonistAgeInSim:         run.KRTTopProtagonistAgeInSim,
					HoFTopProtagonistAgeInSim:         run.HoFTopProtagonistAgeInSim,
					RRTTopProtagonistAgeInSim:         run.RRTTopProtagonistAgeInSim,
					SETTopProtagonistAgeInSim:         run.SETTopProtagonistAgeInSim,
					KRTAntagonistsAvgAgeInSim:         run.KRTAntagonistsAvgAgeInSim,
					HoFAntagonistsAvgAgeInSim:         run.HoFAntagonistsAvgAgeInSim,
					RRTAntagonistsAvgAgeInSim:         run.RRTAntagonistsAvgAgeInSim,
					SETAntagonistsAvgAgeInSim:         run.SETAntagonistsAvgAgeInSim,
					KRTProtagonistsAvgAgeInSim:        run.KRTProtagonistsAvgAgeInSim,
					HoFProtagonistsAvgAgeInSim:        run.HoFProtagonistsAvgAgeInSim,
					RRTProtagonistsAvgAgeInSim:        run.RRTProtagonistsAvgAgeInSim,
					SETProtagonistsAvgAgeInSim:        run.SETProtagonistsAvgAgeInSim,
					KRTAntagonistsAvgBirthGenInSim:    run.KRTTopAntagonistsAvgBirthGenInSim,
					HoFAntagonistsAvgBirthGenInSim:    run.HoFTopAntagonistsAvgBirthGenInSim,
					RRTAntagonistsAvgBirthGenInSim:    run.RRTTopAntagonistsAvgBirthGenInSim,
					SETAntagonistsAvgBirthGenInSim:    run.SETTopAntagonistsAvgBirthGenInSim,
					KRTProtagonistsAvgBirthGenInSim:   run.KRTTopProtagonistsAvgBirthGenInSim,
					HoFProtagonistsAvgBirthGenInSim:   run.HoFTopProtagonistsAvgBirthGenInSim,
					RRTProtagonistsAvgBirthGenInSim:   run.RRTopProtagonistsAvgBirthGenInSim,
					SETProtagonistsAvgBirthGenInSim:   run.SETTopProtagonistsAvgBirthGenInSim,
					KRTTopAntagonistNoCompetitionsInSim : run.KRTTopAntagonistNoCompetitionsInSim ,
					HoFTopAntagonistNoCompetitionsInSim: run.HoFTopAntagonistNoCompetitionsInSim ,
					RRTTopAntagonistNoCompetitionsInSim: run.RRTTopAntagonistNoCompetitionsInSim ,
					SETTopAntagonistNoCompetitionsInSim: run.SETTopAntagonistNoCompetitionsInSim ,
					KRTTopProtagonistNoCompetitionsInSim: run.KRTTopProtagonistNoCompetitionsInSim,
					HoFTopProtagonistNoCompetitionsInSim: run.HoFTopProtagonistNoCompetitionsInSim,
					RRTTopProtagonistNoCompetitionsInSim: run.RRTTopProtagonistNoCompetitionsInSim,
					SETTopProtagonistNoCompetitionsInSim: run.SETTopProtagonistNoCompetitionsInSim,
					KRTTopAntagonistStrategyInSim:     run.KRTTopAntagonistStrategyInSim,
					HoFTopAntagonistStrategyInSim:     run.HoFTopAntagonistStrategyInSim,
					RRTTopAntagonistStrategyInSim:     run.RRTTopAntagonistStrategyInSim,
					SETTopAntagonistStrategyInSim:     run.SETTopAntagonistStrategyInSim,
					KRTTopProtagonistStrategyInSim:    run.KRTTopProtagonistStrategyInSim,
					HoFTopProtagonistStrategyInSim:    run.HoFTopProtagonistStrategyInSim,
					RRTTopProtagonistStrategyInSim:    run.RRTTopProtagonistStrategyInSim,
					SETTopProtagonistStrategyInSim:    run.SETTopProtagonistStrategyInSim,
					KRTAntagonistsMeanInSim:           run.KRTAntagonistsMeanInSim,
					HoFAntagonistsMeanInSim:           run.HoFAntagonistsMeanInSim,
					RRAntagonistsMeanInSim:            run.RRAntagonistsMeanInSim,
					SETAntagonistsMeanInSim:           run.SETAntagonistsMeanInSim,
					KRTProtagonistsMeanInSim:          run.KRTProtagonistsMeanInSim,
					HoFProtagonistsMeanInSim:          run.HoFProtagonistsMeanInSim,
					RRProtagonistsMeanInSim:           run.RRProtagonistsMeanInSim,
					SETProtagonistsMeanInSim:          run.SETProtagonistsMeanInSim,
					KRTAntagonistStdDevInSim:          run.KRTAntagonistStdDevInSim,
					HoFAntagonistStdDevInSim:          run.HoFAntagonistStdDevInSim,
					RRAntagonistStdDevInSim:           run.RRAntagonistStdDevInSim,
					SETAntagonistStdDevInSim:          run.SETAntagonistStdDevInSim,
					KRTProtagonistStdDevInSim:         run.KRTProtagonistStdDevInSim,
					HoFProtagonistStdDevInSim:         run.HoFProtagonistStdDevInSim,
					RRProtagonistStdDevInSim:          run.RRProtagonistStdDevInSim,
					SETProtagonistStdDevInSim:         run.SETProtagonistStdDevInSim,
					KRTAntagonistVarInSim:             run.KRTAntagonistVarInSim,
					HoFAntagonistVarInSim:             run.HoFAntagonistVarInSim,
					RRAntagonistVarInSim:              run.RRAntagonistVarInSim,
					SETAntagonistVarInSim:             run.SETAntagonistVarInSim,
					KRTProtagonistVarInSim:            run.KRTProtagonistVarInSim,
					HoFProtagonistVarInSim:            run.HoFProtagonistVarInSim,
					RRProtagonistVarInSim:             run.RRProtagonistVarInSim,
					SETProtagonistVarInSim:            run.SETProtagonistVarInSim,
				}

				csvSims = append(csvSims, csvSim)
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	//2. Grab the first row of elements and any necessary items
	create, err := os.Create(fmt.Sprintf("%s/%s.csv", a.Config.Stats.OutputDir, "allstats"))
	if err != nil {
		log.Fatal(err)
	}

	err = gocsv.MarshalFile(csvSims, create)
	if err != nil {
		log.Fatal(err)
	}
}
