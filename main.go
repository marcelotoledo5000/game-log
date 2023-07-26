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
	"github.com/marcelotoledo5000/game-log/pkg/parser"

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

	generateReport(gameParser)
}
