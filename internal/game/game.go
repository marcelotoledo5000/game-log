package game

import (
	"fmt"
	"sort"

	"github.com/marcelotoledo5000/game-log/internal/player"
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

func (g *Game) PlayersSortedByID() []*player.Player {
	players := make([]*player.Player, len(g.Players))
	copy(players, g.Players)
	sort.Slice(players, func(i, j int) bool {
		return players[i].Id < players[j].Id
	})
	return players
}

func (g *Game) PlayersSortedByScore() []*player.Player {
	players := make([]*player.Player, len(g.Players))
	copy(players, g.Players)
	sort.Slice(players, func(i, j int) bool {
		return g.Kills[players[i].Name] > g.Kills[players[j].Name]
	})
	return players
}

func (g *Game) KillsByMeansSortedByKills() []string {
	means := make([]string, 0, len(g.KillsByMeans))
	for mean := range g.KillsByMeans {
		means = append(means, mean)
	}
	sort.Slice(means, func(i, j int) bool {
		return g.KillsByMeans[means[i]] > g.KillsByMeans[means[j]]
	})
	return means
}
