# Pomo - Pomodoro Timer

A simple command-line pomodoro timer written in Go.

## Usage

Pomo supports three main commands:

### Work Timer

Start a single work session. Default duration is 20 minutes.

```sh
# Use default work duration (20 minutes)
go run main.go work

# Specify custom work duration (in minutes)
go run main.go work 25
```

### Rest Timer

Start a single rest session. Default duration is 5 minutes.

```sh
# Use default rest duration (5 minutes)
go run main.go rest

# Specify custom rest duration (in minutes)
go run main.go rest 10
```

### Timer Loop

Start an infinite loop alternating between work and rest sessions.

```sh
# Use default durations (20 minutes work, 5 minutes rest)
go run main.go timer

# Specify custom work duration, default rest duration
go run main.go timer 25

# Specify both work and rest durations
go run main.go timer 25 10
```

## Examples

```sh
# Quick 5-minute work session
go run main.go work 5

# Standard pomodoro (25 minutes work, 5 minutes rest) loop
go run main.go timer 25 5

# Long rest break
go run main.go rest 15
```

## Building

To build the executable:

```sh
go build -o pomo main.go
```

Then run with:

```sh
./pomo work 25
```
