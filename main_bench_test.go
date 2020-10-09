package main

import (
	"github.com/martinomburajr/ebb/app"
	"log"
	"math/rand"
	"testing"
	"time"
)


var err error

func BenchmarkMainSimulation(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	configPath := "bench_config.json"

	application, err := app.NewApplication(configPath)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	//debug.SetGCPercent(-1)

	for i := 0; i < b.N; i++{
		err = application.Begin()

		if err != nil {
			log.Fatal(err)
		}
	}
}
