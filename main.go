package main

import (
	"flag"
	"github.com/martinomburajr/ebb/app"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	MainSimulation("")
}

func MainSimulation(configPath string) {
	var trueConfigPath string

	parseAll, config := ParseFlags()

	if configPath != "" {
		trueConfigPath = configPath
	} else {
		trueConfigPath = config
	}

	if parseAll {
		log.Println("evo: parsing all...")
		return
	}

	application, err := app.NewApplication(trueConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	err = application.Begin()
	if err != nil {
		log.Fatal(err)
	}
}

func ParseFlags() (parseAll bool, config string) {
	flag.BoolVar(&parseAll, "parseAll", false, "ParseAll begins the compaction of all data")
	flag.StringVar(&config, "config", "config.json", "defines the path of the config")

	flag.Parse()

	return parseAll, config
}
