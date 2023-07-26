package report

import (
	"fmt"
	"sort"
	"strings"

	"github.com/marcelotoledo5000/game-log/pkg/game"
	"github.com/marcelotoledo5000/game-log/pkg/player"
)

func GenerateReport(gp game.GameParserInterface) {
	gameParser := gp.(*game.GameParser)
	for i, game := range gameParser.Games {
		fmt.Printf("\"game_%d\": {\n", i+1)
		fmt.Printf("  \"total_kills\": %d,\n", game.TotalKills)
		fmt.Printf("  \"players\": %s,\n", getPlayerListJSON(game.PlayersSortedByID()))
		fmt.Println("  \"kills\": {")
		for _, player := range game.PlayersSortedByScore() {
			kills := game.Kills[player.Name]
			if kills != 0 { // Sorted by kills and non-zero.
				fmt.Printf("    \"%s\": %d,\n", player.Name, kills)
			}
		}
		fmt.Println("  },")
		fmt.Println("  \"kills_by_means\": {")
		for _, mean := range game.KillsByMeansSortedByKills() {
			count := game.KillsByMeans[mean]
			fmt.Printf("    \"%s\": %d,\n", mean, count)
		}
		fmt.Println("  }")
		if i == len(gameParser.Games)-1 {
			fmt.Println("}")
		} else {
			fmt.Println("},")
		}
	}

	fmt.Println("Player Ranking:")
	playerRanking := make(map[string]int)
	for _, game := range gameParser.Games {
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

func getPlayerListJSON(players []*player.Player) string {
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
