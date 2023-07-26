package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	// Capture console output
	var buffer bytes.Buffer
	log.SetOutput(&buffer)

	// Call the Run function with a valid log file for the test.
	err := run("log/log_to_test.log")
	if err != nil {
		t.Errorf("Unexpected error running program: %s", err)
	}

	// Check that no errors occurred during execution
	if buffer.Len() > 0 {
		t.Errorf("Unexpected error occurred during execution:\n%s", buffer.String())
	}
}

func TestMainWithInvalidLog(t *testing.T) {
	// Capture console output (stderr)
	cmd := exec.Command("go", "run", "main.go", "invalid_log_file.log")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run the program as an external process
	err := cmd.Run()

	// Verify that an error has occurred and that the error message is as expected.
	if err == nil {
		t.Errorf("Expected error, but got nil")
	} else {
		expectedErrorMessage := "Error running program: open invalid_log_file.log: no such file or directory"
		if !strings.Contains(stderr.String(), expectedErrorMessage) {
			t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErrorMessage, stderr.String())
		}
	}
}
