package parser

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/marcelotoledo5000/game-log/pkg/game"
	"github.com/marcelotoledo5000/game-log/pkg/player"
)

type LogParser struct {
	*game.GameParser
}

func (lp *LogParser) ParseLog(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.Contains(line, "InitGame"):
			currentGame := &game.Game{
				Kills:        make(map[string]int),
				KillsByMeans: make(map[string]int),
			}
			lp.Games = append(lp.Games, currentGame)

		case strings.Contains(line, "ClientConnect"):
			currentGame, err := lp.CurrentGame()
			if err != nil {
				return err
			}

			id, err := player.ExtractPlayerID(line)
			if err != nil {
				return err
			}

			game.FindOrCreatePlayer(currentGame, id)

		case strings.Contains(line, "ClientUserinfoChanged"):
			id, err := player.ExtractPlayerID(line)
			if err != nil {
				return err
			}

			name := player.ExtractPlayerName(line)
			currentGame, err := lp.CurrentGame()
			if err != nil {
				return err
			}

			player := game.FindOrCreatePlayer(currentGame, id)
			player.Name = name

		case strings.Contains(line, "Kill"):
			currentGame, err := lp.CurrentGame()
			if err != nil {
				return err
			}

			killer, victim, method := extractKillData(line)

			currentGame.TotalKills++
			currentGame.KillsByMeans[method]++

			if killer == "<world>" {
				if _, ok := currentGame.Kills[victim]; ok {
					currentGame.Kills[victim]--
				} else {
					currentGame.Kills[victim] = -1
				}
			} else {
				if killer != victim {
					currentGame.Kills[killer]++
				}
			}

		case strings.Contains(line, "ShutdownGame"):
			game, err := lp.CurrentGame()
			if err != nil {
				return err
			}

			// Remove players without names from the list
			updatedPlayers := make([]*player.Player, 0)
			for _, player := range game.Players {
				if player.Name != "" {
					updatedPlayers = append(updatedPlayers, player)
				}
			}
			game.Players = updatedPlayers
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func extractKillData(line string) (string, string, string) {
	regex := regexp.MustCompile(`Kill: \d+ \d+ \d+: (.*?) killed (.*?) by\s(.+)`)
	match := regex.FindStringSubmatch(line)

	if len(match) > 2 {
		// 1: killer, 2: victim, 3: method
		return match[1], match[2], match[3]
	}

	return "", "", ""
}
