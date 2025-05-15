# HackSim - Terminal Hacking Simulation Game

A terminal-based hacking simulation game built with Go, featuring an immersive text user interface (TUI) that recreates the experience of system infiltration and hacking.

![HackSim Logo](generated-icon.png)

## Overview

HackSim simulates the experience of a cyber infiltration specialist working through different hacking scenarios. Players navigate through terminal interfaces, use realistic commands, and complete mission objectives in various scenarios.

## Features

- Command-line interface using Cobra for command handling
- Interactive TUI elements powered by Bubbletea library
- Multiple hacking simulation scenarios or challenges
- Terminal-style visual feedback for actions
- Progressive difficulty with dynamic security monitoring
- File system and network navigation

## Game Scenarios

### Network Breach
Break into a corporate network and navigate through security systems to access sensitive data.

### Data Heist
Extract valuable data from a secured server while avoiding detection and covering your tracks.

### System Takeover
Gain full control of a critical infrastructure system and establish persistence.

## Commands

The game includes realistic hacker tools and commands:

- `scan` - Discover network nodes and open ports
- `connect` - Establish connections to remote systems
- `ls` - List files and directories
- `cat` - View file contents
- `crack` - Attempt to crack passwords or security mechanisms
- `exploit` - Leverage vulnerabilities to gain access
- `download` - Download files from compromised systems
- `upload` - Upload files to target systems
- `status` - Check your current mission status

## Gameplay

Players must complete mission objectives while avoiding detection. As you progress, you'll unlock more powerful tools but will face increasing security measures.

## Usage

### Basic Usage

```bash
# Start the game with the default scenario
go run main.go play

# Choose a specific scenario
go run main.go play -s data-heist

# Display help information
go run main.go help
```

### Debug Mode

For testing or development, you can use the debug mode:

```bash
go run main.go play --debug
```

### Demo Script

To see the game's functionality without the TUI interface:

```bash
go run demo.go
```

## Development

### Project Structure

- `/cmd` - Command-line interface definitions
- `/ui` - User interface components using Bubbletea
- `/game` - Core game logic and state management
- `/assets` - Game assets like ASCII art

## Design

- Colors: Primary #00FF00 (matrix green), Secondary #FF0000 (alert red), Background #000000 (terminal black), Text #33FF33 (soft green), Accent #0066FF (cyber blue)
- Design: Monospace font, classic terminal aesthetics, ASCII art elements, high contrast color scheme for readability

## Credits

HackSim is inspired by Hollywood hacking scenes and popular terminal-based tools like htop and nmap.