package game

import (
	"testing"

	"github.com/marcelotoledo5000/game-log/pkg/player"
)

func TestCurrentGame(t *testing.T) {
	t.Run("NoStartedGame", func(t *testing.T) {
		gp := NewGameParser()

		game, err := gp.CurrentGame()

		if game != nil {
			t.Errorf("Expected nil game when no game is started, got: %v", game)
		}
		if err == nil || err.Error() != "No started game" {
			t.Errorf("Expected 'No started game' error, got: %v", err)
		}
	})

	t.Run("StartedGame", func(t *testing.T) {
		gp := NewGameParser()
		game1 := &Game{}
		game2 := &Game{}
		gp.Games = []*Game{game1, game2}

		game, err := gp.CurrentGame()

		if game != game2 {
			t.Errorf("Expected game2 as current game, got: %v", game)
		}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func TestFindOrCreatePlayer(t *testing.T) {
	t.Run("PlayerExists", func(t *testing.T) {
		game := &Game{
			Players: []*player.Player{
				{Id: "1", Name: "Player1"},
				{Id: "2", Name: "Player2"},
			},
		}

		player := FindOrCreatePlayer(game, "1")

		if player == nil {
			t.Error("Expected player1, got nil")
		}
		if player.Name != "Player1" {
			t.Errorf("Expected player name 'Player1', got: '%s'", player.Name)
		}
	})

	t.Run("PlayerDoesNotExist", func(t *testing.T) {
		game := &Game{
			Players: []*player.Player{
				{Id: "1", Name: "Player1"},
			},
		}

		player := FindOrCreatePlayer(game, "2")

		if player == nil {
			t.Error("Expected new player, got nil")
		}
		if player.Name != "" {
			t.Errorf("Expected empty player name, got: '%s'", player.Name)
		}
	})
}
