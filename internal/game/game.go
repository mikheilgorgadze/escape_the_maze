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
        Level: 2,
    }
}

func (s *State) MovePlayer(rowOffset, colOffset int) {
    newPos := s.Player.GetNewPosition(rowOffset, colOffset)

	if s.Grid.IsValidMove(newPos) && s.Status == StatusPlaying{
        cellType := s.Grid.GetCellType(newPos)

        switch cellType {
        case grid.EXIT:
            s.Status = StatusWon
        case grid.ENEMY:
            s.takeDamage(config.DAMAGE)
            if !s.Player.IsAlive() {
                s.Status = StatusGameOver
            }
        case grid.PATH:
            s.PathCount++
            s.Grid.MovePlayerOnGrid(s.Player.GetPosition(), newPos)
            s.Player.Move(rowOffset, colOffset)
            s.Score += config.PATH_SCORE
        }
    }
}

func (s *State) takeDamage(damage int){
    s.Player.Health -= damage
    s.Score += damage
}
