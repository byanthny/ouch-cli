# ouch-cli

CLI version of ouch built with Go

## Description

Welcome to Ouch CLI, a command-line application based on //TODO OUCH LINK

## Current Progress

### Completed

- README documentation

### In Progress

- Implementing basic CLI commands structure
- Command routing and execution system

### Coming Soon

- Local state management
- Chat functionality between Quiddities
- Real-time simulation updates
- Advanced command interactions
- Multiplayer simulation support

## Features

- **Existences Management**: Create and join different Existence instances
- **XP System**: Track your Ouch level and progress
- **Offline Mode**: Interact with the simulation without an internet connection
- **Persistent State**: Your simulation state is saved between sessions
- **Cross-platform**: Works on macOS, Linux, and Windows (I hope)

## Commands

| Command       | Description                                     |
| ------------- | ----------------------------------------------- |
| `ouch create` | Create a new Existence                          |
| `ouch join`   | Join an existing Existence                      |
| `ouch theme`  | Toggle between different terminal themes        |
| `ouch bean`   | A surprise. But once you bean you can't go back |
| `ouch status` | Check the status of your Quiddity               |
| `ouch exit`   | Exit the simulation                             |

//TODO CHECK BELOW THIS LINE

## Installation

```bash
go install github.com/byanthny/ouch-cli@latest
```

Or build from source:

```bash
git clone https://github.com/byanthny/ouch-cli.git
cd ouch-cli
go build
```

## Configuration

Configuration is stored in `~/.ouch-cli.yaml` and can be edited directly or through the CLI interface.
