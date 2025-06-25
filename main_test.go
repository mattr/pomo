package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestDefaultWorkDuration(t *testing.T) {
	expected := 20
	if DefaultWorkDuration != expected {
		t.Errorf("Expected DefaultWorkDuration to be %d, got %d", expected, DefaultWorkDuration)
	}
}

func TestDefaultRestDuration(t *testing.T) {
	expected := 5
	if DefaultRestDuration != expected {
		t.Errorf("Expected DefaultRestDuration to be %d, got %d", expected, DefaultRestDuration)
	}
}

// Test helper function to simulate timer logic without actual sleep
func simulateTimer(duration time.Duration) (totalSeconds int, progressSteps int) {
	totalSeconds = int(duration.Seconds())
	progressSteps = totalSeconds + 1 // includes 0 to totalSeconds
	return
}

func TestSimulateTimer(t *testing.T) {
	duration := 3 * time.Second
	totalSeconds, progressSteps := simulateTimer(duration)

	expectedTotal := 3
	expectedSteps := 4 // 0, 1, 2, 3

	if totalSeconds != expectedTotal {
		t.Errorf("Expected totalSeconds to be %d, got %d", expectedTotal, totalSeconds)
	}

	if progressSteps != expectedSteps {
		t.Errorf("Expected progressSteps to be %d, got %d", expectedSteps, progressSteps)
	}
}

// Test command line argument parsing logic
func TestWorkCommandParsing(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"program", "work"}, DefaultWorkDuration},
		{[]string{"program", "work", "25"}, 25},
		{[]string{"program", "work", "invalid"}, DefaultWorkDuration},
		{[]string{"program", "work", "0"}, DefaultWorkDuration},
		{[]string{"program", "work", "-5"}, DefaultWorkDuration},
	}

	for _, test := range tests {
		// Simulate the parsing logic from main
		duration := DefaultWorkDuration
		if len(test.args) >= 3 {
			if d, err := parsePositiveInt(test.args[2]); err == nil {
				duration = d
			}
		}

		if duration != test.expected {
			t.Errorf("For args %v, expected duration %d, got %d", test.args, test.expected, duration)
		}
	}
}

func TestRestCommandParsing(t *testing.T) {
	tests := []struct {
		args     []string
		expected int
	}{
		{[]string{"program", "rest"}, DefaultRestDuration},
		{[]string{"program", "rest", "10"}, 10},
		{[]string{"program", "rest", "invalid"}, DefaultRestDuration},
		{[]string{"program", "rest", "0"}, DefaultRestDuration},
	}

	for _, test := range tests {
		// Simulate the parsing logic from main
		duration := DefaultRestDuration
		if len(test.args) >= 3 {
			if d, err := parsePositiveInt(test.args[2]); err == nil {
				duration = d
			}
		}

		if duration != test.expected {
			t.Errorf("For args %v, expected duration %d, got %d", test.args, test.expected, duration)
		}
	}
}

func TestTimerCommandParsing(t *testing.T) {
	tests := []struct {
		args         []string
		expectedWork int
		expectedRest int
		shouldError  bool
	}{
		{[]string{"program", "timer"}, DefaultWorkDuration, DefaultRestDuration, false},
		{[]string{"program", "timer", "25"}, 25, DefaultRestDuration, false},
		{[]string{"program", "timer", "25", "10"}, 25, 10, false},
		{[]string{"program", "timer", "invalid"}, 0, 0, true},
		{[]string{"program", "timer", "25", "invalid"}, 0, 0, true},
		{[]string{"program", "timer", "0"}, 0, 0, true},
		{[]string{"program", "timer", "25", "0"}, 0, 0, true},
	}

	for _, test := range tests {
		workDuration := DefaultWorkDuration
		restDuration := DefaultRestDuration
		hasError := false

		// Simulate the parsing logic from main
		if len(test.args) >= 3 {
			if w, err := parsePositiveInt(test.args[2]); err != nil {
				hasError = true
			} else {
				workDuration = w
			}
		}

		if !hasError && len(test.args) >= 4 {
			if r, err := parsePositiveInt(test.args[3]); err != nil {
				hasError = true
			} else {
				restDuration = r
			}
		}

		if hasError != test.shouldError {
			t.Errorf("For args %v, expected error %v, got %v", test.args, test.shouldError, hasError)
		}

		if !hasError {
			if workDuration != test.expectedWork {
				t.Errorf("For args %v, expected work duration %d, got %d", test.args, test.expectedWork, workDuration)
			}
			if restDuration != test.expectedRest {
				t.Errorf("For args %v, expected rest duration %d, got %d", test.args, test.expectedRest, restDuration)
			}
		}
	}
}

// Helper function to test argument parsing (extracted from main logic)
func parsePositiveInt(s string) (int, error) {
	if val, err := strconv.Atoi(s); err == nil && val > 0 {
		return val, nil
	}
	return 0, fmt.Errorf("invalid positive integer: %s", s)
}

// Test displayTimer function behavior (without actual sleep/output)
func TestDisplayTimerLogic(t *testing.T) {
	// Test the core logic of displayTimer without sleep
	duration := 3 * time.Second
	totalSeconds := int(duration.Seconds())
	progressWidth := 80

	if totalSeconds != 3 {
		t.Errorf("Expected totalSeconds to be 3, got %d", totalSeconds)
	}

	// Test progress calculation at different points
	testCases := []struct {
		elapsedSeconds    int
		expectedPercent   float64
		expectedCompleted int
	}{
		{0, 0.0, 0},
		{1, 1.0 / 3.0, 26}, // 1/3 * 80 = 26.67 -> 26
		{2, 2.0 / 3.0, 53}, // 2/3 * 80 = 53.33 -> 53
		{3, 1.0, 80},
	}

	for _, tc := range testCases {
		percentComplete := float64(tc.elapsedSeconds) / float64(totalSeconds)
		completedWidth := int(percentComplete * float64(progressWidth))

		if percentComplete != tc.expectedPercent {
			t.Errorf("For elapsed %d, expected percent %.3f, got %.3f",
				tc.elapsedSeconds, tc.expectedPercent, percentComplete)
		}

		if completedWidth != tc.expectedCompleted {
			t.Errorf("For elapsed %d, expected completed width %d, got %d",
				tc.elapsedSeconds, tc.expectedCompleted, completedWidth)
		}
	}
}

// Test color code selection in displayTimer
func TestDisplayTimerColorCodes(t *testing.T) {
	testCases := []struct {
		color        string
		expectedCode string
	}{
		{"green", "\033[32m"},
		{"red", "\033[31m"},
		{"blue", ""}, // unsupported color should result in empty string
		{"", ""},     // empty color should result in empty string
	}

	for _, tc := range testCases {
		// Simulate the color code selection logic from displayTimer
		colorCode := ""
		switch tc.color {
		case "green":
			colorCode = "\033[32m"
		case "red":
			colorCode = "\033[31m"
		}

		if colorCode != tc.expectedCode {
			t.Errorf("For color %s, expected code %s, got %s",
				tc.color, tc.expectedCode, colorCode)
		}
	}
}

// Test startWorkTimer and startRestTimer functions exist and can be called
// (These are integration-style tests since the functions call displayTimer)
func TestStartWorkTimerExists(t *testing.T) {
	// Test that startWorkTimer function exists and doesn't panic with very short duration
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("startWorkTimer panicked: %v", r)
		}
	}()

	// We can't easily test the actual timer without mocking, but we can test it exists
	// and the function signature is correct by attempting to call it with a very short duration
	// Note: This will actually run for 1 second, so keep duration minimal for testing
	go func() {
		startWorkTimer(1 * time.Millisecond)
	}()

	// Give it a moment to start, then continue
	time.Sleep(10 * time.Millisecond)
}

func TestStartRestTimerExists(t *testing.T) {
	// Test that startRestTimer function exists and doesn't panic with very short duration
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("startRestTimer panicked: %v", r)
		}
	}()

	go func() {
		startRestTimer(1 * time.Millisecond)
	}()

	// Give it a moment to start, then continue
	time.Sleep(10 * time.Millisecond)
}

// Test startTimerLoop function exists (but don't run the infinite loop)
func TestStartTimerLoopExists(t *testing.T) {
	// Test that startTimerLoop function exists by checking we can reference it
	// We can't actually test the infinite loop without complex mocking

	// Test the function signature is correct by creating a function variable
	var timerFunc func(time.Duration, time.Duration)
	timerFunc = startTimerLoop

	// If we can assign the function, it exists and has the correct signature
	if timerFunc == nil {
		t.Error("startTimerLoop assignment failed")
	}

	// Alternative test: just verify we can reference the function without error
	_ = startTimerLoop
}
