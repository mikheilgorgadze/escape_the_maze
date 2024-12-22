package main

import (
	"log"

	"github.com/mikheilgorgadze/maze/internal/game"
	"github.com/mikheilgorgadze/maze/internal/input"
	"github.com/mikheilgorgadze/maze/internal/ui"
)


func main() {

    inputHandler, err := input.NewHandler()
    if err!=nil{
        log.Fatal(err.Error())
        return
    }
    
    defer inputHandler.Close()
   
    gameState := game.NewGameState()
    ui := &ui.Renderer{}
    ui.HandleMenu(gameState)


	for {
		event, err := inputHandler.ReadKey()
		if err!= nil {
			panic(event.Err)
		}
		gameState.KeyPressed = event.Key

		switch gameState.Status {
		case game.StatusMenu:
			ui.HandleMenu(gameState)
		case game.StatusPlaying:
			ui.HandlePlaying(gameState, inputHandler)
            if gameState.Status == game.StatusGameOver || gameState.Status == game.StatusWon {
                return
            }
        case game.StatusGameOver:
            ui.HandleGameOver(gameState)
			return
		}

		if inputHandler.IsExit(event.Key) {
			break
		}
	}
}
