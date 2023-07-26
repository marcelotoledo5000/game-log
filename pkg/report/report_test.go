package report

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/marcelotoledo5000/game-log/pkg/game"
	"github.com/marcelotoledo5000/game-log/pkg/player"
)

func TestGenerateReport(t *testing.T) {
	gp := game.NewGameParser()

	// Sample games.
	game1 := &game.Game{
		TotalKills: 11,
		Players: []*player.Player{
			{Name: "Player1"},
			{Name: "Player2"},
			{Name: "Player3"},
		},
		Kills: map[string]int{
			"Player1": 5,
			"Player2": 3,
			"Player3": 1,
		},
		KillsByMeans: map[string]int{
			"MOD_UNKNOWN":    5,
			"MOD_SHOTGUN":    2,
			"MOD_RAILGUN":    1,
			"MOD_MACHINEGUN": 3,
		},
	}

	game2 := &game.Game{
		TotalKills: 10,
		Players: []*player.Player{
			{Name: "Player1"},
			{Name: "Player2"},
		},
		Kills: map[string]int{
			"Player1": 3,
			"Player2": 7,
		},
		KillsByMeans: map[string]int{
			"MOD_UNKNOWN": 10,
		},
	}

	game3 := &game.Game{
		TotalKills: 1,
		Players: []*player.Player{
			{Name: "Player3"},
			{Name: "Player4"},
		},
		Kills: map[string]int{
			"Player3": 1,
		},
		KillsByMeans: map[string]int{
			"MOD_RAILGUN": 1,
		},
	}

	gp.Games = append(gp.Games, game1, game2, game3)

	expectedOutput := `
	"game_1": {
  "total_kills": 11,
  "players": ["Player1", "Player2", "Player3"],
  "kills": {
    "Player1": 5,
    "Player2": 3,
    "Player3": 1,
  },
  "kills_by_means": {
    "MOD_UNKNOWN": 5,
    "MOD_MACHINEGUN": 3,
    "MOD_SHOTGUN": 2,
    "MOD_RAILGUN": 1,
  }
},
"game_2": {
  "total_kills": 10,
  "players": ["Player1", "Player2"],
  "kills": {
    "Player2": 7,
    "Player1": 3,
  },
  "kills_by_means": {
    "MOD_UNKNOWN": 10,
  }
},
"game_3": {
  "total_kills": 1,
  "players": ["Player3", "Player4"],
  "kills": {
    "Player3": 1,
  },
  "kills_by_means": {
    "MOD_RAILGUN": 1,
  }
}
Player Ranking:
1. Player2: 10 kills
2. Player1: 8 kills
3. Player3: 2 kills
`

	// Capture console output.
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	GenerateReport(gp)

	// Restore the default console output.
	w.Close()
	os.Stdout = old

	var buffer bytes.Buffer
	io.Copy(&buffer, r)
	r.Close()

	result := buffer.String()

	// Verify that the output is correct
	if strings.TrimSpace(result) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Output mismatch.\nExpected:\n%s\nGot:\n%s", expectedOutput, result)
	}
}
