package player

import (
	"fmt"
	"regexp"
)

type Player struct {
	Id    string
	Name  string
	Score int
}

func ExtractPlayerID(input string) (string, error) {
	regex := regexp.MustCompile(`ClientConnect:\s+(\d+)|ClientUserinfoChanged:\s+(\d+)`)
	match := regex.FindStringSubmatch(input)

	if len(match) < 2 {
		return "", fmt.Errorf("ID not found in the string")
	}

	return match[len(match)-1], nil
}

func ExtractPlayerName(line string) string {
	regex := regexp.MustCompile(`n\\(.*?)\\`)
	match := regex.FindStringSubmatch(line)

	if len(match) > 1 {
		return match[len(match)-1]
	}

	return ""
}
