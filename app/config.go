package app

import (
	"encoding/json"
	"fmt"
	"github.com/martinomburajr/ebb/evolution"
	"log"
	"os"
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
	simulation, err := Init(a.Config)
	if err != nil {
		return err
	}

	simulationResult, err := simulation.Start()
	if err != nil {
		return err
	}

	evoParams := simulation.Config.Params
	println(simulationResult.Summary(evoParams))

	csvData := simulationResult.GenerateCSVData(evoParams)

	outputPath := fmt.Sprintf("%s/%s-%s-%d-%s.csv", a.Config.OutputDir, a.Config.OutputFile, strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", ""), a.Config.Complexity, evolution.RandString(7))

	var file *os.File

	file, err = os.Create(outputPath)
	if err != nil {
		fmt.Println("output path cannot be created - defaulting to stats.csv")

		file, err = os.Create("stats.csv")
		if err != nil {
			log.Fatalf("failed to write to fallback (stats.csv), %v", err)
		}
	}

	return simulationResult.WriteCSV(file, csvData)
}

// ParseAll begins the compaction of all data. It begins by combining all the run-level information into complexity-level
// information, which finally accumulates into the simulation-level statistics that encompasses compacted data of every
// single run performed.
func (a *Application) ParseAll() {

}
