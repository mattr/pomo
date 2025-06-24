# Specification Template
> Ingest the information from this file, implement the Low-Level Tasks, and generate the code that will satisfy the High and Mid-Level Objectives.

## High-Level Objective

- Command line pomodoro timer

## Mid-Level Objective

- Build a go CLI pomodoro timer application
- Accept commands of 'work' and 'rest'
- Display the progress graphically in the terminal
- Show the remaining time in the terminal
- Accept an additional argument for the time

## Implementation Notes
- Confirm before importing any external libraries
- Write unit tests for all methods
- All functions can remain in `main.go`
- Use terminal graphics to display the progress
- Run the code using `go run . <args>`
- Carefully review each low-level task for exact code changes.

## Context

### Beginning context
- main.go
- go.mod

### Ending context  
- main.go
- go.mod

## Low-Level Tasks
> Ordered from start to finish

1. Start a work timer
```aider
Update main.go to accept a command line argument `work`: 
    add a function to start a timer for 20 minutes,
    write "Working..." to the terminal output,
    write the current progress of the timer to the terminal output so that it updates every second:
    fix the width of the display to 80 characters, and increment it as a percentage of the time elapsed,
    display the time remaining at the end of the progress bar.
```
2. Start a rest timer
```aider
Update main.go to accept a command line argument `rest`:
    add a function to start a timer for 5 minutes,
    write "Resting..." to the terminal output,
    write the current progress of the timer to the terminal output so that it updates every second:
    fix the width of the display to 80 characters, and increment it as a percentage of the time elapsed,
    display the time remaining at the end of the progress bar.
```
3. Allow a custom duration for work and rest
```aider
Update main.go:
allow a second argument which is the time to work or rest in minutes;
update the work function to use this time instead of 20 minutes,
update the rest function to use this time instead of 5 minutes.
```
4. Allow the functions to loop
```aider
Update main.go:
    add a `timer` command line argument which takes two additional parameters:
    the work time in minutes and the rest time in minutes;
    start the timer and alternate between work and rest until the user quits.
```
