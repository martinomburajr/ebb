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

	for i := 0; i < b.N; i++ {
		err = application.Begin()

		if err != nil {
			log.Fatal(err)
		}
	}
}

var (
	p = []int{}
	l = 15
)

func BenchmarkPerm(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.ResetTimer()
	b.ReportAllocs()
	//debug.SetGCPercent(-1)
	for i := 0; i < b.N; i++ {
		p = rand.Perm(l)
	}
}

func BenchmarkRandn(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	b.ResetTimer()
	b.ReportAllocs()
	//debug.SetGCPercent(-1)

	for i := 0; i < b.N; i++ {
		p = randN()
	}
}

var c = 0

func randN() []int {
	x := make([]int, l)
	for i := 0; i < l; i++ {
		x[i] = rand.Intn(l)
	}

	return x
}
