package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
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
        duration := 20 // default duration in minutes
        if len(os.Args) >= 3 {
            d, err := strconv.Atoi(os.Args[2])
            if err == nil && d > 0 {
                duration = d
            }
        }
        startWorkTimer(time.Duration(duration) * time.Minute)
    case "rest":
        duration := 5 // default duration in minutes
        if len(os.Args) >= 3 {
            d, err := strconv.Atoi(os.Args[2])
            if err == nil && d > 0 {
                duration = d
            }
        }
        startRestTimer(time.Duration(duration) * time.Minute)
    case "timer":
        if len(os.Args) < 4 {
            fmt.Println("Please provide work and rest durations in minutes")
            return
        }
        workDuration, err1 := strconv.Atoi(os.Args[2])
        restDuration, err2 := strconv.Atoi(os.Args[3])
        if err1 != nil || err2 != nil || workDuration <= 0 || restDuration <= 0 {
            fmt.Println("Invalid durations provided")
            return
        }
        startTimerLoop(time.Duration(workDuration)*time.Minute, time.Duration(restDuration)*time.Minute)
    default:
        fmt.Println("Unknown command:", command)
    }
}
