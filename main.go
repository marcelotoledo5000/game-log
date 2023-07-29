package main

import (
	"log"
	"os"

	"github.com/marcelotoledo5000/game-log/internal/game"
	"github.com/marcelotoledo5000/game-log/internal/parser"
	"github.com/marcelotoledo5000/game-log/internal/report"
)

func main() {
	var logFilePath string

	if len(os.Args) > 1 {
		logFilePath = os.Args[1]
	} else {
		logFilePath = "log/qgames.log"
	}
	err := run(logFilePath)
	if err != nil {
		log.Fatalf("Error running program: %s", err)
	}
}

func run(logFilePath string) error {
	gameParser := game.NewGameParser()
	logParser := &parser.LogParser{GameParser: gameParser}
	err := logParser.ParseLog(logFilePath)
	if err != nil {
		return err
	}

	report.GenerateReport(gameParser)
	return nil
}
