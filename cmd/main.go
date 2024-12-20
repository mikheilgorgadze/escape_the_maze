package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/mikheilgorgadze/maze/internal/config"
	"github.com/mikheilgorgadze/maze/internal/game"
	"github.com/mikheilgorgadze/maze/internal/grid"
)


func main() {
	keysEvents, err := keyboard.GetKeys(100)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer func() {
		_ = keyboard.Close()
	}()
   
    gameState := game.NewGameState()

	handleMenu(gameState)

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		gameState.KeyPressed = event.Key

		switch gameState.Status {
		case game.StatusMenu:
			handleMenu(gameState)
		case game.StatusPlaying:
			handlePlaying(gameState)
            if gameState.Status == game.StatusGameOver || gameState.Status == game.StatusWon {
                return
            }
        case game.StatusGameOver:
            handleGameOver(gameState)
			return
		}

		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}

func movePlayer(gameState *game.State, rowOffset, colOffset int) {
    newPos := gameState.Player.GetNewPosition(rowOffset, colOffset)
    currentDamage := 0

	if gameState.Grid.IsValidMove(newPos) && gameState.Status == game.StatusPlaying{
        cellType := gameState.Grid.GetCellType(newPos)

        switch cellType {
        case grid.EXIT:
            gameState.Status = game.StatusWon
        case grid.ENEMY:
            gameState.Player.Health -= config.DAMAGE
            currentDamage += config.DAMAGE
            if !gameState.Player.IsAlive() {
                gameState.Status = game.StatusGameOver
            }
        case grid.PATH:
            gameState.PathCount++
            gameState.Grid.MovePlayer(gameState.Player.GetPosition(), newPos)
            gameState.Player.Move(rowOffset, colOffset)
        }
        gameState.Score = gameState.PathCount*10 + currentDamage * gameState.Level
    }
}

func printGameStarted() {
	fmt.Println("GAME HAS STARTED, USE ARROWS TO MOVE PLAYER")
}

func handleMenu(gameState *game.State) {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")

	fmt.Println("Hello, welcome to the Maze")
	fmt.Println("Press ESC key to quit")
	fmt.Println("Press ENTER to start")
    fmt.Println("Your goal is to reach exit marked with 'E' by avoiding enemies (marked by 'X'). Lower score you get, the better")
    fmt.Println("Score is calculated by number of paths ('-') you take multiplied to 10 plus health points you loose multiplied by level you are playing")
	if gameState.KeyPressed == keyboard.KeyEnter {
		gameState.Status = game.StatusPlaying
        printGameStarted()
		renderGameState(gameState)
	}
}

func handlePlaying(gameState *game.State) {
	if gameState.Status != game.StatusPlaying {
		return
	}

    keyMappings := map[keyboard.Key]struct{row, col int}{
        keyboard.KeyArrowRight: {0, 1},
        keyboard.KeyArrowLeft: {0, -1},
        keyboard.KeyArrowUp: {-1, 0},
        keyboard.KeyArrowDown: {1, 0},
    }

	if move, ok := keyMappings[gameState.KeyPressed]; ok {
        movePlayer(gameState, move.row, move.col)
        if gameState.Status == game.StatusGameOver || gameState.Status == game.StatusWon {
            renderGameState(gameState) 
            handleGameOver(gameState)
            return
        }
	}
    renderGameState(gameState)
}

func handleGameOver(gameState *game.State) {
	fmt.Printf("Game Over, your total score is: %d\n", gameState.Score)
    fmt.Printf("Remaining health: %d\n", gameState.Player.Health)
    fmt.Printf("Path count: %d\n", gameState.PathCount)
}

func renderGameState(gameState *game.State){
    fmt.Print("\033[2J")
	fmt.Print("\033[H")
    fmt.Printf("Remaining health: %d\n", gameState.Player.Health)
    fmt.Printf("Path count: %d\n", gameState.PathCount)
    fmt.Printf("Current score: %d\n", gameState.Score)
    fmt.Print("\033[10;0H" + gameState.Grid.Render())
}
