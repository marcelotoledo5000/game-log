package main

import (
	"testing"
)

// TestParseLog tests the ParseLog function.
func TestParseLog(t *testing.T) {
	logFilePath := "log/log_to_test.log"

	gameParser := NewGameParser()
	err := gameParser.ParseLog(logFilePath)
	if err != nil {
		t.Fatalf("Error parsing log: %v", err)
	}

	// Check the number of games
	if len(gameParser.games) != 1 {
		t.Errorf("Expected 1 game, got %d", len(gameParser.games))
	}

	// Check players in the game
	if len(gameParser.games[0].Players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(gameParser.games[0].Players))
	}

	// Check total kills
	if gameParser.games[0].TotalKills != 3 {
		t.Errorf("Expected 0 total kills, got %d", gameParser.games[0].TotalKills)
	}

	// Check specific kills
	if gameParser.games[0].Kills["Isgalamido"] != 1 {
		t.Errorf("Expected 1 kills for player 'Isgalamido', got %d", gameParser.games[0].Kills["Isgalamido"])
	}

	// Check specific kills by means
	if gameParser.games[0].KillsByMeans["MOD_TRIGGER_HURT"] != 1 {
		t.Errorf("Expected 1 kills by means 'MOD_TRIGGER_HURT', got %d", gameParser.games[0].KillsByMeans["MOD_TRIGGER_HURT"])
	}
	if gameParser.games[0].KillsByMeans["MOD_ROCKET_SPLASH"] != 2 {
		t.Errorf("Expected 2 kills by means 'MOD_ROCKET_SPLASH', got %d", gameParser.games[0].KillsByMeans["MOD_ROCKET_SPLASH"])
	}
}
