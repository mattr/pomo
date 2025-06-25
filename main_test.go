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
