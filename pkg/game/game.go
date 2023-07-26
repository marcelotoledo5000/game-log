package game

import (
	"fmt"

	"github.com/marcelotoledo5000/game-log/pkg/player"
)

type Game struct {
	TotalKills   int
	Players      []*player.Player
	Kills        map[string]int
	KillsByMeans map[string]int
}

type GameParserInterface interface {
	CurrentGame() (*Game, error)
}

type GameParser struct {
	Games []*Game
}

func NewGameParser() *GameParser {
	return &GameParser{
		Games: make([]*Game, 0),
	}
}

func (gp *GameParser) CurrentGame() (*Game, error) {
	if len(gp.Games) > 0 {
		current := gp.Games[len(gp.Games)-1]
		return current, nil
	}
	return nil, fmt.Errorf("No started game")
}

func FindOrCreatePlayer(game *Game, id string) *player.Player {
	for _, player := range game.Players {
		if player.Id == id {
			return player
		}
	}

	player := &player.Player{
		Id: id,
	}
	game.Players = append(game.Players, player)

	return player
}
