package app

import (
	"encoding/json"
	"fmt"
	"github.com/martinomburajr/ebb/evolog"
	"github.com/martinomburajr/ebb/evolution"
	"log"
	"os"
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
	Stats StatsInfo `json:"stats_1"`

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
	OutputDir       string `json:"outputDir"`
	OutputFile      string `json:"outputFile"`
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
		fmt.Println("output path cannot be created - defaulting to stats_1.csv")
	}

	s.Config.GeneratedStatsOutputPath = fmt.Sprintf("%s/stats_1.csv", newStatsDir)
	s.Config.GeneratedStatsStratOutputPath = fmt.Sprintf("%s/stats_1-strategy.csv", newStatsDir)

	//krtPath := fmt.Sprintf("%s/stats_1-krt.csv", newStatsDir)
	//setPath := fmt.Sprintf("%s/stats_1-set.csv", newStatsDir)

	csvData, statsCSV := simulationResult.GenerateCSVData(s.Config.Params)
	//csvDataKRT, _ := simulationResult.GenerateTopologySpecificCSVData(s.Config.Params, "KRT")
	//csvDataSET, _ := simulationResult.GenerateTopologySpecificCSVData(s.Config.Params, "SET")

	//var file *os.File
	//
	//file, err = os.Create(s.Config.GeneratedStatsOutputPath)
	//if err != nil {
	//	fmt.Println("output path cannot be created - defaulting to stats_1.csv")
	//
	//	file, err = os.Create("stats_1.csv")
	//	if err != nil {
	//		log.Fatalf("failed to write to fallback (stats_1.csv), %v", err)
	//	}
	//
	//}

	//stratFile, err := os.Create(s.Config.GeneratedStatsStratOutputPath)
	//if err != nil {
	//	fmt.Println("output path cannot be created - defaulting to stats_1-strategy.csv")
	//
	//	file, err = os.Create("stats_1-strategy.csv")
	//	if err != nil {
	//		log.Fatalf("failed to write to fallback (stats_1-strategy.csv), %v", err)
	//	}
	//
	//}

	//err = simulationResult.WriteCSV(file, csvData)
	//err = simulationResult.WriteStratCSV(stratFile, statsCSV)

	err = s.writeStatsToFile(s.Config.GeneratedStatsOutputPath, simulationResult, csvData)
	if err != nil {
		log.Fatalf("failed to write to fallback (stats_1.csv), %v", err)
	}

	err = s.writeStatsToFile(s.Config.GeneratedStatsStratOutputPath, simulationResult, statsCSV)
	if err != nil {
		log.Fatalf("failed to write to fallback (stats_1.csv), %v", err)
	}

	s.LogMessage("completed stats_1", evolog.LoggerAnalysis)

	return err
}

func (s *Simulation) writeStatsToFile(path string, simulationResult evolution.SimulationResult, csvData interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("output path cannot be created - defaulting to stats_1.csv")

		file, err = os.Create("stats_1.csv")
		if err != nil {
			log.Fatalf("failed to write to fallback (stats_1.csv), %v", err)
		}

	}

	err = simulationResult.WriteCSV(file, csvData)
	return err
}

func (s *Simulation) outputPlot(randFolder string, config ApplicationConfig, simulationResult evolution.SimulationResult) error {
	newStatsDir := fmt.Sprintf("%s/%d/%s", config.Plots.OutputDir, config.Complexity, randFolder)

	err := os.MkdirAll(newStatsDir, 0777)
	if err != nil {
		fmt.Println("output path cannot be created - defaulting to stats_1.csv")
	}

	s.Config.GeneratedPlotsOutputPath = fmt.Sprintf("%s", newStatsDir)

	s.Plot(simulationResult)

	return err
}
