package main

import (
	"github.com/martinomburajr/ebb/app"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestMainSimulation(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	configPath := "bench_config.json"

	//debug.SetGCPercent(-1)

	application, err := app.NewApplication(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = application.Begin()

	if err != nil {
		log.Fatal(err)
	}
}
