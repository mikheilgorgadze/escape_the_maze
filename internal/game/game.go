package game

import (
	"github.com/eiannone/keyboard"
	"github.com/mikheilgorgadze/maze/internal/config"
	"github.com/mikheilgorgadze/maze/internal/grid"
	"github.com/mikheilgorgadze/maze/internal/player"
)


type State struct {
    Grid           *grid.Grid
    Status         Status
    Player         *player.Player
    PathCount      int
    Score          int
    LivesRemaining int
    Level          int
    KeyPressed     keyboard.Key
}

func NewGameState() *State{
    g := grid.GenerateNewGrid(config.GRID_WIDTH, config.GRID_HEIGHT, config.ENEMY_COUNT)
    player := player.New(1, 1, config.STARTING_HEALTH)
    g.PlacePlayer(player.Position)
    return &State{
        Grid: g,
        Status: StatusMenu,
        Player: player,
        PathCount: 0,
        Score: 0,
        LivesRemaining: 3,
        Level: 1,
    }
}
