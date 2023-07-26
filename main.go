package main

import (
	"log"
	"os"

	"github.com/marcelotoledo5000/game-log/pkg/game"
	"github.com/marcelotoledo5000/game-log/pkg/parser"
	"github.com/marcelotoledo5000/game-log/pkg/report"
)

func main() {
	var logFilePath string

	if len(os.Args) > 1 {
		logFilePath = os.Args[1]
	} else {
		logFilePath = "log/qgames.log"
	}

	gameParser := game.NewGameParser()
	logParser := &parser.LogParser{GameParser: gameParser}
	err := logParser.ParseLog(logFilePath)
	if err != nil {
		log.Fatalf("Error parsing log file: %s", err)
	}

	report.GenerateReport(gameParser)
}
