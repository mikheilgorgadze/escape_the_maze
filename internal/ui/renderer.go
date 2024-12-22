package ui

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/mikheilgorgadze/maze/internal/game"
	"github.com/mikheilgorgadze/maze/internal/input"
)

type Renderer struct{}


func (r *Renderer) RenderGameState(gameState *game.State){
    fmt.Print("\033[2J")
	fmt.Print("\033[H")
    fmt.Printf("Remaining health: %d\n", gameState.Player.Health)
    fmt.Printf("Path count: %d\n", gameState.PathCount)
    fmt.Printf("Current score: %d\n", gameState.Score)
    fmt.Print("\033[10;0H" + gameState.Grid.Render())
}

func (r *Renderer) HandleMenu(gameState *game.State) {
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
		r.RenderGameState(gameState)
	}
}

func (r *Renderer) HandleGameOver(gameState *game.State) {
	fmt.Printf("Game Over, your total score is: %d\n", gameState.Score)
    fmt.Printf("Remaining health: %d\n", gameState.Player.Health)
    fmt.Printf("Path count: %d\n", gameState.PathCount)
}


func (r *Renderer) HandlePlaying(gameState *game.State, inputHanlder *input.Handler) {
	if gameState.Status != game.StatusPlaying {
		return
	}

	if move, ok := inputHanlder.GetMovement(gameState.KeyPressed); ok {
        gameState.MovePlayer(move.Row, move.Col)
        if gameState.Status == game.StatusGameOver || gameState.Status == game.StatusWon {
            r.RenderGameState(gameState) 
            r.HandleGameOver(gameState)
            return
        }
	}
    r.RenderGameState(gameState)
}

func printGameStarted() {
	fmt.Println("GAME HAS STARTED, USE ARROWS TO MOVE PLAYER")
}
