# Maze Game

A simple terminal-based maze game written in Go where players navigate through a maze while avoiding enemies.

## Current State (v0.1.0)
This game is in early development. Core features are implemented but the game is still undergoing improvements and refinements.

## Features
- Terminal-based maze navigation
- Player movement using arrow keys
- Enemy encounters
- Health system
- Score tracking
- Basic collision detection
- Exit point to win

## Prerequisites
- Go 1.19 or higher
- github.com/eiannone/keyboard package

## Installation
1. Clone the repository
```bash
git clone https://github.com/mikheilgorgadze/maze.git
```

2. Navigate to the project directory
```bash
cd maze
```

3. Install dependencies
```bash
go mod tidy
```

## How to Play
1. Run the game
```bash
go run main.go
```

2. Controls:
- Arrow keys: Move player
- Enter: Start game
- ESC: Exit game

3. Game Rules:
- Find the exit (E) while avoiding enemies (X)
- Each enemy collision reduces health
- Game ends when health reaches 0
- Lower score is better
- Score is calculated based on path taken and damage received

## Project Structure
```
maze/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   ├── game/
│   ├── grid/
│   └── player/
├── go.mod
├── go.sum
└── README.md
```

## Known Issues
- Score calculation needs refinement when player hits enemies
- No proper game ending sequence
- Limited error handling for keyboard inputs

## Contributing
This project is in active development. Feel free to submit issues and pull requests.

## Development Roadmap
- [ ] Improve scoring system
- [ ] Add difficulty levels
- [ ] Add power-ups
- [ ] Implement high score system
- [ ] Add sound effects
- [ ] Improve maze generation algorithm

## Author
Mikheil Gorgadze
