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

func startWorkTimer(duration time.Duration) {
    fmt.Println("Working...")

    totalSeconds := int(duration.Seconds())
    progressWidth := 80

    for elapsedSeconds := 0; elapsedSeconds <= totalSeconds; elapsedSeconds++ {
        percentComplete := float64(elapsedSeconds) / float64(totalSeconds)
        completedWidth := int(percentComplete * float64(progressWidth))
        remainingWidth := progressWidth - completedWidth

        progressBar := "[" + strings.Repeat("=", completedWidth) + strings.Repeat(" ", remainingWidth) + "]"
        timeRemaining := time.Duration(totalSeconds - elapsedSeconds) * time.Second

        fmt.Printf("\r%s %s remaining", progressBar, timeRemaining)
        time.Sleep(1 * time.Second)
    }
    fmt.Println("\nTime's up!")
}

func startRestTimer(duration time.Duration) {
    fmt.Println("Resting...")

    totalSeconds := int(duration.Seconds())
    progressWidth := 80

    for elapsedSeconds := 0; elapsedSeconds <= totalSeconds; elapsedSeconds++ {
        percentComplete := float64(elapsedSeconds) / float64(totalSeconds)
        completedWidth := int(percentComplete * float64(progressWidth))
        remainingWidth := progressWidth - completedWidth

        progressBar := "[" + strings.Repeat("=", completedWidth) + strings.Repeat(" ", remainingWidth) + "]"
        timeRemaining := time.Duration(totalSeconds - elapsedSeconds) * time.Second

        fmt.Printf("\r%s %s remaining", progressBar, timeRemaining)
        time.Sleep(1 * time.Second)
    }
    fmt.Println("\nTime's up!")
}

func startTimerLoop(workDuration, restDuration time.Duration) {
    for {
        startWorkTimer(workDuration)
        startRestTimer(restDuration)
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command: work, rest, or timer")
        return
    }

    command := os.Args[1]

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
