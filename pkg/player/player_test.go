package player

import (
	"testing"
)

func TestExtractPlayerID(t *testing.T) {
	t.Run("ValidClientUserinfoChanged", func(t *testing.T) {
		input := "ClientUserinfoChanged: 456"
		expectedID := "456"

		id, err := ExtractPlayerID(input)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if id != expectedID {
			t.Errorf("Expected ID '%s', got: '%s'", expectedID, id)
		}
	})

	t.Run("IDNotFound", func(t *testing.T) {
		input := "InvalidString"

		id, err := ExtractPlayerID(input)

		if id != "" {
			t.Errorf("Expected empty ID, got: '%s'", id)
		}
		if err == nil || err.Error() != "ID not found in the string" {
			t.Errorf("Expected 'ID not found in the string' error, got: %v", err)
		}
	})
}

func TestExtractPlayerName(t *testing.T) {
	t.Run("ValidName", func(t *testing.T) {
		line := "n\\Player1\\"
		expectedName := "Player1"

		name := ExtractPlayerName(line)

		if name != expectedName {
			t.Errorf("Expected name '%s', got: '%s'", expectedName, name)
		}
	})

	t.Run("NameNotFound", func(t *testing.T) {
		line := "InvalidString"

		name := ExtractPlayerName(line)

		if name != "" {
			t.Errorf("Expected empty name, got: '%s'", name)
		}
	})
}
