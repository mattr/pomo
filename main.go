package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

const (
    DefaultWorkDuration = 20 // default work duration in minutes
    DefaultRestDuration = 5  // default rest duration in minutes
)

func displayTimer(duration time.Duration, color string) {
    totalSeconds := int(duration.Seconds())
    progressWidth := 80

    // ANSI color codes
    colorCode := ""
    resetCode := "\033[0m"
    switch color {
    case "green":
        colorCode = "\033[32m"
    case "red":
        colorCode = "\033[31m"
    }

    for elapsedSeconds := 0; elapsedSeconds <= totalSeconds; elapsedSeconds++ {
        percentComplete := float64(elapsedSeconds) / float64(totalSeconds)
        completedWidth := int(percentComplete * float64(progressWidth))
        remainingWidth := progressWidth - completedWidth

        progressBar := colorCode + strings.Repeat("█", completedWidth) + resetCode + strings.Repeat(" ", remainingWidth)
        timeRemaining := time.Duration(totalSeconds - elapsedSeconds) * time.Second

        fmt.Printf("\r%s %s remaining", progressBar, timeRemaining)
        time.Sleep(1 * time.Second)
    }
    fmt.Println("\nTime's up!")
}

func startWorkTimer(duration time.Duration) {
    fmt.Println("Working...")
    displayTimer(duration, "green")
}

func startRestTimer(duration time.Duration) {
    fmt.Println("Resting...")
    displayTimer(duration, "red")
}

func startTimerLoop(workDuration, restDuration time.Duration) {
    for {
        startWorkTimer(workDuration)
        startRestTimer(restDuration)
    }
}

func main() {
    command := "timer" // default command
    if len(os.Args) >= 2 {
        command = os.Args[1]
    }

    switch command {
    case "work":
        duration := DefaultWorkDuration
        if len(os.Args) >= 3 {
            d, err := strconv.Atoi(os.Args[2])
            if err == nil && d > 0 {
                duration = d
            }
        }
        startWorkTimer(time.Duration(duration) * time.Minute)
    case "rest":
        duration := DefaultRestDuration
        if len(os.Args) >= 3 {
            d, err := strconv.Atoi(os.Args[2])
            if err == nil && d > 0 {
                duration = d
            }
        }
        startRestTimer(time.Duration(duration) * time.Minute)
    case "timer":
        workDuration := DefaultWorkDuration
        restDuration := DefaultRestDuration
        
        if len(os.Args) >= 3 {
            w, err := strconv.Atoi(os.Args[2])
            if err != nil || w <= 0 {
                fmt.Println("Invalid work duration provided")
                return
            }
            workDuration = w
        }
        
        if len(os.Args) >= 4 {
            r, err := strconv.Atoi(os.Args[3])
            if err != nil || r <= 0 {
                fmt.Println("Invalid rest duration provided")
                return
            }
            restDuration = r
        }
        
        startTimerLoop(time.Duration(workDuration)*time.Minute, time.Duration(restDuration)*time.Minute)
    default:
        fmt.Println("Unknown command:", command)
    }
}
