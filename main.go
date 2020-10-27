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

// TODO - Age of the same individual should remain even if cloned, do not delete the age
// TODO - NoOfCompetitions

func MainSimulation(configPath string) {
	var trueConfigPath string

	parseAll, config, threads, apps := ParseFlags()

	if configPath != "" {
		trueConfigPath = configPath
	} else {
		trueConfigPath = config
	}

	application, err := app.NewApplication(trueConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	//parseAll = true
	if parseAll {
		log.Println("evo: parsing all...")

		//krtFile := "stats_1-krt.csv"
		//setFile := "stats_1-set.csv"
		//application.ParseAllTopologyAware(krtFile, "krt-all.csv")

		//allFile := "stats_1.csv"
		//application.ParseAll(allFile, "all-stats2.0.csv")

		allFileStrat := "stats_1-strategy.csv"
		application.ParseAllCSVStrat(allFileStrat, "all-stats_1-strat2.0.csv")
		return
	}

	numApplicationRuns := apps

	ThreadPool(threads, numApplicationRuns, func(iter int) {
		StartApplication(application, iter, err)
	})
	//print(threads)

	//for i := 0; i < numApplicationRuns; i++ {
	//	StartApplication(application, i, err)
	//}
}

func ParseFlags() (parseAll bool, config string, threads, apps int) {
	flag.BoolVar(&parseAll, "parseAll", false, "ParseAll begins the compaction of all data")
	flag.StringVar(&config, "config", "config.json", "defines the path of the config")
	flag.IntVar(&threads, "threads", 3, "The size of a given threadpool")
	flag.IntVar(&apps, "apps", 2000, "The number of applications to run")

	flag.Parse()

	return parseAll, config, threads, apps
}
