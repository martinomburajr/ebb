package main

import (
	"fmt"
	"github.com/martinomburajr/ebb/app"
	"log"
	"os"
)

func StartApplication(application app.Application, i int, err error) {
	fmt.Println("__________________________START___________________________")
	application.Config.Iter = i

	err = application.Begin()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("__________________________END___________________________")
}

func ThreadPool(poolSize, queueSize int, work func(iter int)) {
	doneChan := make(chan struct{})
	workers := 0
	doneCount := 0

	if doneCount < queueSize {
		for i := 0; i < poolSize; i++ {
			go func(doneChan chan struct{}, iter int) {
				work(iter)

				doneChan <- struct{}{}
			}(doneChan, doneCount)
		}
	}

	for {
		select {
		case <-doneChan:
			doneCount++
			workers--

			log.Printf("############# - Simulation %d - Complete | Left: %d - #########", doneCount, queueSize-doneCount)

			if doneCount < queueSize {
				if workers < poolSize {
					go func(doneChan chan struct{}, iter int) {
						workers++
						work(iter)

						doneChan <- struct{}{}
					}(doneChan, doneCount)
				}
			} else {
				os.Exit(0)
			}
		default:
			//if workers < poolSize {
			//	go func(doneChan chan struct{}, iter int) {
			//		workers++
			//		work(iter)
			//
			//		doneChan <- struct{}{}
			//	}(doneChan, doneCount)
			//}
		}
	}
}
