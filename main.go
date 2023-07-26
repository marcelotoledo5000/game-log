package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Player struct {
	Id    string
	Name  string
	Score int
}

func (gp *GameParser) ParseLog(filePath string) error {
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
			currentGame := &Game{
				Kills:        make(map[string]int),
				KillsByMeans: make(map[string]int),
			}
			gp.games = append(gp.games, currentGame)

		case strings.Contains(line, "ClientConnect"):
			currentGame, err := gp.currentGame()
			if err != nil {
				return err
			}

			id, err := extractPlayerID(line)
			if err != nil {
				return err
			}

			findOrCreatePlayer(currentGame, id)

		case strings.Contains(line, "ClientUserinfoChanged"):
			id, err := extractPlayerID(line)
			if err != nil {
				return err
			}

			name := extractPlayerName(line)
			currentGame, err := gp.currentGame()
			if err != nil {
				return err
			}

			player := findOrCreatePlayer(currentGame, id)
			player.Name = name

		case strings.Contains(line, "Kill"):
			currentGame, err := gp.currentGame()
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
			game, err := gp.currentGame()
			if err != nil {
				return err
			}

			// Remove players without names from the list
			updatedPlayers := make([]*Player, 0)
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

func extractPlayerID(input string) (string, error) {
	regex := regexp.MustCompile(`ClientConnect:\s+(\d+)|ClientUserinfoChanged:\s+(\d+)`)
	match := regex.FindStringSubmatch(input)

	if len(match) < 2 {
		return "", fmt.Errorf("ID not found in the string")
	}

	return match[len(match)-1], nil
}

func extractPlayerName(line string) string {
	regex := regexp.MustCompile(`n\\(.*?)\\`)
	match := regex.FindStringSubmatch(line)

	if len(match) > 1 {
		return match[len(match)-1]
	}

	return ""
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

func generateReport(gp *GameParser) {
	for i, game := range gp.games {
		fmt.Printf("\"game_%d\": {\n", i+1)
		fmt.Printf("  \"total_kills\": %d,\n", game.TotalKills)
		fmt.Printf("  \"players\": %s,\n", getPlayerListJSON(game.Players))
		fmt.Println("  \"kills\": {")
		for player, kills := range game.Kills {
			fmt.Printf("    \"%s\": %d,\n", player, kills)
		}
		fmt.Println("  },")
		fmt.Println("  \"kills_by_means\": {")
		for mean, count := range game.KillsByMeans {
			fmt.Printf("    \"%s\": %d,\n", mean, count)
		}
		fmt.Println("  }")
		if i == len(gp.games)-1 {
			fmt.Println("}")
		} else {
			fmt.Println("},")
		}
	}

	fmt.Println("Player Ranking:")
	playerRanking := make(map[string]int)
	for _, game := range gp.games {
		for player, kills := range game.Kills {
			if _, ok := playerRanking[player]; ok {
				playerRanking[player] += kills
			} else {
				playerRanking[player] = kills
			}
		}
	}
	sortedRanking := sortPlayerRanking(playerRanking)
	for i, player := range sortedRanking {
		fmt.Printf("%d. %s: %d kills\n", i+1, player, playerRanking[player])
	}
}

func getPlayerListJSON(players []*Player) string {
	var names []string
	for _, player := range players {
		names = append(names, fmt.Sprintf("\"%s\"", player.Name))
	}
	return "[" + strings.Join(names, ", ") + "]"
}

func sortPlayerRanking(playerRanking map[string]int) []string {
	type playerKills struct {
		player string
		kills  int
	}
	ranking := make([]playerKills, 0)
	for player, kills := range playerRanking {
		ranking = append(ranking, playerKills{player, kills})
	}
	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].kills > ranking[j].kills
	})
	sortedRanking := make([]string, 0)
	for _, pk := range ranking {
		sortedRanking = append(sortedRanking, pk.player)
	}
	return sortedRanking
}
	"github.com/marcelotoledo5000/game-log/pkg/game"

func main() {
	var logFilePath string

	if len(os.Args) > 1 {
		logFilePath = os.Args[1]
	} else {
		logFilePath = "log/qgames.log"
	}

	gameParser := NewGameParser()
	err := gameParser.ParseLog(logFilePath)
	if err != nil {
		log.Fatalf("Error parsing log file: %s", err)
	}

	generateReport(gameParser)
}
