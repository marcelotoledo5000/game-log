package player

import (
	"fmt"
	"regexp"
	"strconv"
)

type Player struct {
	Id    string
	Name  string
	Score int
}

func ExtractPlayerID(input string) (string, error) {
	regex := regexp.MustCompile(`Client(?:Connect|UserinfoChanged):\s+(\d+)`)
	match := regex.FindStringSubmatch(input)

	if len(match) < 2 {
		return "", fmt.Errorf("ID not found in the string")
	}

	id := match[1]

	// Check that the ID is a valid number
	if _, err := strconv.Atoi(id); err != nil {
		return "", fmt.Errorf("Invalid ID format")
	}

	return id, nil
}

func ExtractPlayerName(line string) string {
	regex := regexp.MustCompile(`n\\(.*?)\\`)
	match := regex.FindStringSubmatch(line)

	if len(match) > 1 {
		return match[len(match)-1]
	}

	return ""
}
