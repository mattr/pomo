# Pomo - Pomodoro Timer

A simple command-line pomodoro timer written in Go.

## Installation

### Option 1: Download from GitHub Releases
Download the latest release for your platform from the [releases page](https://github.com/mattr/pomo/releases) and add it to your PATH.

### Option 2: Install with Go
```sh
go install github.com/mattr/pomo@latest
```

### Option 3: Build from Source
```sh
git clone https://github.com/mattr/pomo.git
cd pomo
go build . && ./pomo
```

### Option 4: Run Directly (Development)
```sh
go run .
```

## Usage

Pomo supports three main commands:

### Work Timer

Start a single work session. Default duration is 20 minutes.

```sh
# Use default work duration (20 minutes)
pomo work

# Specify custom work duration (in minutes)
pomo work 25
```

### Rest Timer

Start a single rest session. Default duration is 5 minutes.

```sh
# Use default rest duration (5 minutes)
pomo rest

# Specify custom rest duration (in minutes)
pomo rest 10
```

### Timer Loop

Start an infinite loop alternating between work and rest sessions.

```sh
# Use default durations (20 minutes work, 5 minutes rest)
pomo timer

# Specify custom work duration, default rest duration
pomo timer 25

# Specify both work and rest durations
pomo timer 25 10
```

## Examples

```sh
# Quick 5-minute work session
pomo work 5

# Standard pomodoro (25 minutes work, 5 minutes rest) loop
pomo timer 25 5

# Long rest break
pomo rest 15
```
