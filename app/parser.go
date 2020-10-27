package app

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/martinomburajr/ebb/evolution"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ParseAll begins the compaction of all data. It begins by combining all the run-level information into complexity-level
// information, which finally accumulates into the simulation-level statistics that encompasses compacted data of every
// single run performed.
func (a *Application) ParseAll(fileName, outputFile string) {
	//runs := make([]evolution.CSVAvgGenerationsCombinedAcrossRuns, 0)
	csvSims := make([]evolution.CSVSim, 0)
	//1. Parse all CSVs into CSVAvgGenerationsCombinedAcrossRuns objects
	counter := 0
	for i := 0; i < 3; i++ {
		statsPath := fmt.Sprintf("%s/%d", a.Config.Stats.OutputDir, i)
		err := filepath.Walk(statsPath, func(path string, info os.FileInfo, err error) error {
			// If we come across a csv
			if strings.Contains(path, fileName) && !info.IsDir() {
				file, err := os.Open(path)
				defer file.Close()

				if err != nil {
					return fmt.Errorf("failed to open file at statsPath %s | err: %v", path, err)
				}

				var runs []evolution.CSVAvgGenerationsCombinedAcrossRuns
				err = gocsv.UnmarshalFile(file, &runs)
				if err != nil {
					return fmt.Errorf("failed to unmarshal | err: %v", err)
				}

				//runs = append(runs, csv)
				run := runs[0]
				csvSim := evolution.CSVSim{
					SpecEquation:                         run.SpecEquation,
					SpecEquationLen:                      run.SpecEquationLen,
					IVarCount:                            run.IVarCount,
					PolDegree:                            run.PolDegree,

					KRTTopAEquation:                      run.KRTTopAEquation,
					KRTTopAEquationPolDegree:             run.KRTTopAEquationPolDegree,

					HoFTopAEquation:                      run.HoFTopAEquation,
					HoFTopAEquationPolDegree:             run.HoFTopAEquationPolDegree,

					RRTopAEquation:                       run.RRTopAEquation,
					RRTopAEquationPolDegree:              run.RRTopAEquationPolDegree,

					SETTopAEquation:                      run.SETTopAEquation,
					SETTopAEquationPolDegree:             run.SETTopAEquationPolDegree,

					KRTTopPEquation:                      run.KRTTopPEquation,
					KRTTopPEquationPolDegree:             run.KRTTopPEquationPolDegree,

					HoFTopPEquation:                      run.HoFTopPEquation,
					HoFTopPEquationPolDegree:             run.HoFTopPEquationPolDegree,

					RRTopPEquation:                       run.RRTopPEquation,
					RRTopPEquationPolDegree:              run.RRTopPEquationPolDegree,

					SETTopPEquation:                      run.SETTopPEquation,
					SETTopPEquationPolDegree:             run.SETTopPEquationPolDegree,

					KRTTopAntagonistBestFitnessInSim:     run.KRTTopAntagonistBestFitnessInSim,
					HoFTopAntagonistBestFitnessInSim:     run.HoFTopAntagonistBestFitnessInSim,
					RRTTopAntagonistBestFitnessInSim:     run.RRTTopAntagonistBestFitnessInSim,
					SETTopAntagonistBestFitnessInSim:     run.SETTopAntagonistBestFitnessInSim,

					KRTTopProtagonistBestFitnessInSim:    run.KRTTopProtagonistBestFitnessInSim,
					HoFTopProtagonistBestFitnessInSim:    run.HoFTopProtagonistBestFitnessInSim,
					RRTTopProtagonistBestFitnessInSim:    run.RRTTopProtagonistBestFitnessInSim,
					SETTopProtagonistBestFitnessInSim:    run.SETTopProtagonistBestFitnessInSim,

					KRTTopAntagonistBirthGenInSim:        run.KRTTopAntagonistBirthGenInSim,
					HoFTopAntagonistBirthGenInSim:        run.HoFTopAntagonistBirthGenInSim,
					RRTTopAntagonistBirthGenInSim:        run.RRTTopAntagonistBirthGenInSim,
					SETTopAntagonistBirthGenInSim:        run.SETTopAntagonistBirthGenInSim,

					KRTTopProtagonistBirthGenInSim:       run.KRTTopProtagonistBirthGenInSim,
					HoFTopProtagonistBirthGenInSim:       run.HoFTopProtagonistBirthGenInSim,
					RRTTopProtagonistBirthGenInSim:       run.RRTTopProtagonistBirthGenInSim,
					SETTopProtagonistBirthGenInSim:       run.SETTopProtagonistBirthGenInSim,

					KRTTopAntagonistAgeInSim:             run.KRTTopAntagonistAgeInSim,
					HoFTopAntagonistAgeInSim:             run.HoFTopAntagonistAgeInSim,
					RRTTopAntagonistAgeInSim:             run.RRTTopAntagonistAgeInSim,
					SETTopAntagonistAgeInSim:             run.SETTopAntagonistAgeInSim,

					KRTTopProtagonistAgeInSim:            run.KRTTopProtagonistAgeInSim,
					HoFTopProtagonistAgeInSim:            run.HoFTopProtagonistAgeInSim,
					RRTTopProtagonistAgeInSim:            run.RRTTopProtagonistAgeInSim,
					SETTopProtagonistAgeInSim:            run.SETTopProtagonistAgeInSim,

					KRTAntagonistsAvgAgeInSim:            run.KRTAntagonistsAvgAgeInSim,
					HoFAntagonistsAvgAgeInSim:            run.HoFAntagonistsAvgAgeInSim,
					RRTAntagonistsAvgAgeInSim:            run.RRTAntagonistsAvgAgeInSim,
					SETAntagonistsAvgAgeInSim:            run.SETAntagonistsAvgAgeInSim,

					KRTProtagonistsAvgAgeInSim:           run.KRTProtagonistsAvgAgeInSim,
					HoFProtagonistsAvgAgeInSim:           run.HoFProtagonistsAvgAgeInSim,
					RRTProtagonistsAvgAgeInSim:           run.RRTProtagonistsAvgAgeInSim,
					SETProtagonistsAvgAgeInSim:           run.SETProtagonistsAvgAgeInSim,

					KRTAntagonistsAvgBirthGenInSim:       run.KRTTopAntagonistsAvgBirthGenInSim,
					HoFAntagonistsAvgBirthGenInSim:       run.HoFTopAntagonistsAvgBirthGenInSim,
					RRTAntagonistsAvgBirthGenInSim:       run.RRTTopAntagonistsAvgBirthGenInSim,
					SETAntagonistsAvgBirthGenInSim:       run.SETTopAntagonistsAvgBirthGenInSim,

					KRTProtagonistsAvgBirthGenInSim:      run.KRTTopProtagonistsAvgBirthGenInSim,
					HoFProtagonistsAvgBirthGenInSim:      run.HoFTopProtagonistsAvgBirthGenInSim,
					RRTProtagonistsAvgBirthGenInSim:      run.RRTopProtagonistsAvgBirthGenInSim,
					SETProtagonistsAvgBirthGenInSim:      run.SETTopProtagonistsAvgBirthGenInSim,

					KRTTopAntagonistNoCompetitionsInSim:  run.KRTTopAntagonistNoCompetitionsInSim,
					HoFTopAntagonistNoCompetitionsInSim:  run.HoFTopAntagonistNoCompetitionsInSim,
					RRTTopAntagonistNoCompetitionsInSim:  run.RRTTopAntagonistNoCompetitionsInSim,
					SETTopAntagonistNoCompetitionsInSim:  run.SETTopAntagonistNoCompetitionsInSim,

					KRTTopProtagonistNoCompetitionsInSim: run.KRTTopProtagonistNoCompetitionsInSim,
					HoFTopProtagonistNoCompetitionsInSim: run.HoFTopProtagonistNoCompetitionsInSim,
					RRTTopProtagonistNoCompetitionsInSim: run.RRTTopProtagonistNoCompetitionsInSim,
					SETTopProtagonistNoCompetitionsInSim: run.SETTopProtagonistNoCompetitionsInSim,

					KRTTopAntagonistStrategyInSim:        run.KRTTopAntagonistStrategyInSim,
					HoFTopAntagonistStrategyInSim:        run.HoFTopAntagonistStrategyInSim,
					RRTTopAntagonistStrategyInSim:        run.RRTTopAntagonistStrategyInSim,
					SETTopAntagonistStrategyInSim:        run.SETTopAntagonistStrategyInSim,

					KRTTopProtagonistStrategyInSim:       run.KRTTopProtagonistStrategyInSim,
					HoFTopProtagonistStrategyInSim:       run.HoFTopProtagonistStrategyInSim,
					RRTTopProtagonistStrategyInSim:       run.RRTTopProtagonistStrategyInSim,
					SETTopProtagonistStrategyInSim:       run.SETTopProtagonistStrategyInSim,

					KRTAntagonistsMeanInSim:              run.KRTAntagonistsMeanInSim,
					HoFAntagonistsMeanInSim:              run.HoFAntagonistsMeanInSim,
					RRAntagonistsMeanInSim:               run.RRAntagonistsMeanInSim,
					SETAntagonistsMeanInSim:              run.SETAntagonistsMeanInSim,

					KRTProtagonistsMeanInSim:             run.KRTProtagonistsMeanInSim,
					HoFProtagonistsMeanInSim:             run.HoFProtagonistsMeanInSim,
					RRProtagonistsMeanInSim:              run.RRProtagonistsMeanInSim,
					SETProtagonistsMeanInSim:             run.SETProtagonistsMeanInSim,

					KRTAntagonistStdDevInSim:             run.KRTAntagonistStdDevInSim,
					HoFAntagonistStdDevInSim:             run.HoFAntagonistStdDevInSim,
					RRAntagonistStdDevInSim:              run.RRAntagonistStdDevInSim,
					SETAntagonistStdDevInSim:             run.SETAntagonistStdDevInSim,

					KRTProtagonistStdDevInSim:            run.KRTProtagonistStdDevInSim,
					HoFProtagonistStdDevInSim:            run.HoFProtagonistStdDevInSim,
					RRProtagonistStdDevInSim:             run.RRProtagonistStdDevInSim,
					SETProtagonistStdDevInSim:            run.SETProtagonistStdDevInSim,

					KRTAntagonistVarInSim:                run.KRTAntagonistVarInSim,
					HoFAntagonistVarInSim:                run.HoFAntagonistVarInSim,
					RRAntagonistVarInSim:                 run.RRAntagonistVarInSim,
					SETAntagonistVarInSim:                run.SETAntagonistVarInSim,

					KRTProtagonistVarInSim:               run.KRTProtagonistVarInSim,
					HoFProtagonistVarInSim:               run.HoFProtagonistVarInSim,
					RRProtagonistVarInSim:                run.RRProtagonistVarInSim,
					SETProtagonistVarInSim:               run.SETProtagonistVarInSim,
				}

				csvSims = append(csvSims, csvSim)
				counter++
				log.Printf("[parser] - csv files parsed: %d ", counter)
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	//2. Grab the first row of elements and any necessary items
	create, err := os.Create(fmt.Sprintf("%s/%s", a.Config.Stats.OutputDir, outputFile))
	if err != nil {
		log.Fatal(err)
	}

	err = gocsv.MarshalFile(csvSims, create)
	if err != nil {
		log.Fatal(err)
	}
}

// ParseAll begins the compaction of all data. It begins by combining all the run-level information into complexity-level
// information, which finally accumulates into the simulation-level statistics that encompasses compacted data of every
// single run performed.
func (a *Application) ParseAllTopologyAware(fileName, outputFile string) {
	//runs := make([]evolution.CSVAvgGenerationsCombinedAcrossRuns, 0)
	csvSims := make([]evolution.CSVSimTopologyAware, 0)
	//1. Parse all CSVs into CSVAvgGenerationsCombinedAcrossRuns objects
	counter := 0
	for i := 0; i < 3; i++ {
		statsPath := fmt.Sprintf("%s/%d", a.Config.Stats.OutputDir, i)
		err := filepath.Walk(statsPath, func(path string, info os.FileInfo, err error) error {
			// If we come across a csv
			if strings.Contains(path, fileName) && !info.IsDir() {
				file, err := os.Open(path)
				defer file.Close()

				if err != nil {
					return fmt.Errorf("failed to open file at statsPath %s | err: %v", path, err)
				}

				var runs []evolution.CSVTopologySensitiveCombinedAcrossRuns
				err = gocsv.UnmarshalFile(file, &runs)
				if err != nil {
					return fmt.Errorf("failed to unmarshal | err: %v", err)
				}

				//runs = append(runs, csv)
				run := runs[0]
				csvSim := evolution.CSVSimTopologyAware{
					Topology:                          run.Topology,
					TopologyParam:                     run.ParamVal,
					SpecEquation:                      run.SpecEquation,
					SpecEquationLen:                   run.SpecEquationLen,
					IVarCount:                         run.IVarCount,
					PolDegree:                         run.PolDegree,
					TopAEquation:                   run.TopAEquation,
					TopAEquationPolDegree:          run.TopAEquationPolDegree,
					TopPEquation:                   run.TopPEquation,
					TopPEquationPolDegree:          run.TopPEquationPolDegree,
					TopAntagonistBestFitnessInSim:  run.TopAntagonistBestFitnessInSim,
					TopProtagonistBestFitnessInSim: run.TopProtagonistBestFitnessInSim,

					TopAntagonistBirthGenInSim:        run.TopAntagonistBirthGenInSim,
					TopProtagonistBirthGenInSim:       run.TopProtagonistBirthGenInSim,
					TopAntagonistAgeInSim:             run.TopAntagonistAgeInSim,
					TopProtagonistAgeInSim:            run.TopProtagonistAgeInSim,
					AntagonistsAvgAgeInSim:            run.AntagonistsAvgAgeInSim,
					ProtagonistsAvgAgeInSim:           run.ProtagonistsAvgAgeInSim,
					AntagonistsAvgBirthGenInSim:       run.TopAntagonistsAvgBirthGenInSim,
					ProtagonistsAvgBirthGenInSim:      run.TopProtagonistsAvgBirthGenInSim,
					TopAntagonistNoCompetitionsInSim:  run.TopAntagonistNoCompetitionsInSim,
					TopProtagonistNoCompetitionsInSim: run.TopProtagonistNoCompetitionsInSim,
					TopAntagonistStrategyInSim:        run.TopAntagonistStrategyInSim,
					TopProtagonistStrategyInSim:       run.TopProtagonistStrategyInSim,
					AntagonistsMeanInSim:              run.AntagonistsMeanInSim,
					ProtagonistsMeanInSim:             run.ProtagonistsMeanInSim,
					AntagonistStdDevInSim:             run.AntagonistStdDevInSim,
					ProtagonistStdDevInSim:            run.ProtagonistStdDevInSim,
					AntagonistVarInSim:                run.AntagonistVarInSim,
					ProtagonistVarInSim:               run.ProtagonistVarInSim,
				}

				csvSims = append(csvSims, csvSim)
				counter++
				log.Printf("[parser] - csv files parsed: %d ", counter)
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	//2. Grab the first row of elements and any necessary items
	create, err := os.Create(fmt.Sprintf("%s/%s", a.Config.Stats.OutputDir, outputFile))
	if err != nil {
		log.Fatal(err)
	}

	err = gocsv.MarshalFile(csvSims, create)
	if err != nil {
		log.Fatal(err)
	}
}


// ParseAll begins the compaction of all data. It begins by combining all the run-level information into complexity-level
// information, which finally accumulates into the simulation-level statistics that encompasses compacted data of every
// single run performed.
func (a *Application) ParseAllCSVStrat (fileName, outputFile string) {
	//runs := make([]evolution.CSVAvgGenerationsCombinedAcrossRuns, 0)
	csvSims := make([]evolution.CSVStratHybrid, 0)
	//1. Parse all CSVs into CSVAvgGenerationsCombinedAcrossRuns objects
	counter := 0
	for i := 0; i < 3; i++ {
		statsPath := fmt.Sprintf("%s/%d", a.Config.Stats.OutputDir, i)
		err := filepath.Walk(statsPath, func(path string, info os.FileInfo, err error) error {
			// If we come across a csv
			if strings.Contains(path, fileName) && !info.IsDir() {
				file, err := os.Open(path)
				defer file.Close()

				if err != nil {
					return fmt.Errorf("failed to open file at statsPath %s | err: %v", path, err)
				}

				var runs []evolution.CSVStrat
				err = gocsv.UnmarshalFile(file, &runs)
				if err != nil {
					return fmt.Errorf("failed to unmarshal | err: %v", err)
				}

				//pathHead := strings.SplitAfterN(path, "/",4)
				//join := strings.Join(pathHead[:len(pathHead)],"")
				//trim := strings.Trim (join, "[]")
				mainStatsPath := strings.Replace(path, "stats_1-strategy.csv", "stats_1.csv", -1)
				mainStatsFile, err := os.Open(mainStatsPath)
				defer file.Close()

				if err != nil {
					return fmt.Errorf("failed to open file at statsPath %s | err: %v", mainStatsFile, err)
				}

				var stats []evolution.CSVAvgGenerationsCombinedAcrossRuns
				err = gocsv.UnmarshalFile(mainStatsFile, &stats)
				if err != nil {
					return fmt.Errorf("failed to unmarshal | err: %v", err)
				}

				//runs = append(runs, csv)
				stat := stats[0]

				for j := 0; j < len(runs); j++ {
					run := runs[j]
					csvSim := evolution.CSVStratHybrid{
						Num:                              counter + 1,
						SpecEquation:                      stat.SpecEquation,
						SpecEquationLen:                   stat.SpecEquationLen,
						IVarCount:                         stat.IVarCount,
						PolDegree:                         stat.PolDegree,

						KRTTopAEquation:                   stat.KRTTopAEquation,
						HoFTopAEquation:                   stat.HoFTopAEquation,
						RRTopAEquation:                    stat.RRTopAEquation,
						SETTopAEquation:                   stat.SETTopAEquation,
						KRTTopPEquation:                   stat.KRTTopPEquation,
						HoFTopPEquation:                   stat.HoFTopPEquation,
						RRTopPEquation:                    stat.RRTopPEquation,
						SETTopPEquation:                   stat.SETTopPEquation,

						KRTTopAEquationPolDegree:          stat.KRTTopAEquationPolDegree,
						HoFTopAEquationPolDegree:          stat.HoFTopAEquationPolDegree,
						RRTopAEquationPolDegree:           stat.RRTopAEquationPolDegree,
						SETTopAEquationPolDegree:          stat.SETTopAEquationPolDegree,

						KRTTopPEquationPolDegree:          stat.KRTTopPEquationPolDegree,
						HoFTopPEquationPolDegree:          stat.HoFTopPEquationPolDegree,
						RRTopPEquationPolDegree:           stat.RRTopPEquationPolDegree,
						SETTopPEquationPolDegree:          stat.KRTTopPEquationPolDegree,

						KRTTopAntagonistBestFitnessInSim:  stat.KRTTopAntagonistBestFitnessInSim,
						HoFTopAntagonistBestFitnessInSim:  stat.HoFTopAntagonistBestFitnessInSim,
						RRTTopAntagonistBestFitnessInSim:  stat.RRTTopAntagonistBestFitnessInSim,
						SETTopAntagonistBestFitnessInSim:  stat.SETTopAntagonistBestFitnessInSim,
						KRTTopProtagonistBestFitnessInSim: stat.KRTTopProtagonistBestFitnessInSim,
						HoFTopProtagonistBestFitnessInSim: stat.HoFTopProtagonistBestFitnessInSim,
						RRTTopProtagonistBestFitnessInSim: stat.RRTTopProtagonistBestFitnessInSim,
						SETTopProtagonistBestFitnessInSim: stat.SETTopProtagonistBestFitnessInSim,

						KRTTopAStrat: run.KRTTopAStrat,
						HOFTopAStrat: run.HOFTopAStrat,
						RRTopAStrat:  run.RRTopAStrat,
						SETTopAStrat: run.SETTopAStrat,
						KRTTopPStrat: run.KRTTopPStrat,
						HOFTopPStrat: run.HOFTopPStrat,
						RRTopPStrat:  run.RRTopPStrat,
						SETTopPStrat: run.SETTopPStrat,

						KRTTopAStratDom:                   evolution.DominantStrategyStr(stat.KRTTopAntagonistStrategyInSim),
						HOFTopAStratDom:                   evolution.DominantStrategyStr(stat.HoFTopAntagonistStrategyInSim),
						RRTopAStratDom:                    evolution.DominantStrategyStr(stat.RRTTopAntagonistStrategyInSim),
						SETTopAStratDom:                   evolution.DominantStrategyStr(stat.SETTopAntagonistStrategyInSim),
						KRTTopPStratDom:                   evolution.DominantStrategyStr(stat.KRTTopProtagonistStrategyInSim),
						HOFTopPStratDom:                   evolution.DominantStrategyStr(stat.HoFTopProtagonistStrategyInSim),
						RRTopPStratDom:                    evolution.DominantStrategyStr(stat.RRTTopProtagonistStrategyInSim),
						SETTopPStratDom:                   evolution.DominantStrategyStr(stat.SETTopProtagonistStrategyInSim),
					}

					csvSims = append(csvSims, csvSim)
				}

				csvSims = append(csvSims, evolution.CSVStratHybrid{})


				counter++
				log.Printf("[parser] - strategy csv files parsed: %d ", counter)
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	//2. Grab the first row of elements and any necessary items
	create, err := os.Create(fmt.Sprintf("%s/%s", a.Config.Stats.OutputDir, outputFile))
	if err != nil {
		log.Fatal(err)
	}

	err = gocsv.MarshalFile(csvSims, create)
	if err != nil {
		log.Fatal(err)
	}
}