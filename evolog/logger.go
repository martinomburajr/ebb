package evolog

import (
	"fmt"
	"log"
	"time"
)

const (
	LoggerEvolution  = 0
	LoggerGeneration = 1
	LoggerEpoch      = 2
	LoggerAnalysis   = 3
	LoggerPlot       = 4
	LoggerSimulation = 5
)

type Logger struct {
	Type           int       `json:"message",csv:"message"`
	Message        string    `json:"message",csv:"message"`
	IsProgress     bool      `json:"isProgress"`
	Progress       int       `json:"progress"`
	CompleteNumber int       `json:"completeNumber"`
	Timestamp      time.Time `json:"time",csv:"time"`
	RunNumber      int       `json:"runNumber"`
}

func (l *Logger) NewLog(Run, Type int, message string) *Logger {
	l.Type = Type
	l.Message = message
	l.Timestamp = time.Now()
	l.RunNumber = Run

	return l
}

func (l *Logger) DisplayMessage() {
	loggerType := ""

	switch l.Type {
	case LoggerEvolution:
		loggerType = "evo"
	case LoggerGeneration:
		loggerType = "gen"
	case LoggerAnalysis:
		loggerType = "stat"
	case LoggerPlot:
		loggerType = "plot"
	case LoggerSimulation:
		loggerType = "simulation"
	default:
		loggerType = "xxx"
	}

	logg := fmt.Sprintf("[%d]|[%s]-%s", l.RunNumber, loggerType, l.Message)

	log.Println(logg)
}
