package parser

import (
	"testing"

	"github.com/marcelotoledo5000/game-log/pkg/game"
)

func TestParseLog(t *testing.T) {
	logFilePath := "../../log/log_to_test.log"

	lp := &LogParser{
		GameParser: game.NewGameParser(),
	}

	err := lp.ParseLog(logFilePath)

	if err != nil {
		t.Fatalf("Error parsing log file: %s", err)
	}

	// Check the number of games
	if len(lp.Games) != 1 {
		t.Errorf("Expected 1 game, got %d", len(lp.Games))
	}

	// Check players in the game
	if len(lp.Games[0].Players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(lp.Games[0].Players))
	}

	// Check total kills
	if lp.Games[0].TotalKills != 4 {
		t.Errorf("Expected 4 total kills, got %d", lp.Games[0].TotalKills)
	}

	// Check specific kills
	if lp.Games[0].Kills["Isgalamido"] != 1 {
		t.Errorf("Expected 1 kills for player 'Isgalamido', got %d", lp.Games[0].Kills["Isgalamido"])
	}
	if lp.Games[0].Kills["Mocinha"] != 0 {
		t.Errorf("Expected 0 kills for player 'Mocinha', got %d", lp.Games[0].Kills["Mocinha"])
	}

	// Check specific kills by means
	if lp.Games[0].KillsByMeans["MOD_TRIGGER_HURT"] != 1 {
		t.Errorf("Expected 1 kills by means 'MOD_TRIGGER_HURT', got %d", lp.Games[0].KillsByMeans["MOD_TRIGGER_HURT"])
	}
	if lp.Games[0].KillsByMeans["MOD_ROCKET_SPLASH"] != 2 {
		t.Errorf("Expected 2 kills by means 'MOD_ROCKET_SPLASH', got %d", lp.Games[0].KillsByMeans["MOD_ROCKET_SPLASH"])
	}
	if lp.Games[0].KillsByMeans["MOD_ROCKET"] != 1 {
		t.Errorf("Expected 1 kills by means 'MOD_ROCKET', got %d", lp.Games[0].KillsByMeans["MOD_ROCKET"])
	}
}

func TestExtractKillData(t *testing.T) {
	line := "Kill: 1 2 3: killer killed victim by method"
	killer, victim, method := extractKillData(line)

	if killer != "killer" || victim != "victim" || method != "method" {
		t.Errorf("extractKillData failed. Got: %s, %s, %s, Expected: killer, victim, method", killer, victim, method)
	}
}
