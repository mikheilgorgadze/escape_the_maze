package grid

import (

	"github.com/mikheilgorgadze/maze/internal/player"
)

type Grid struct {
    Cells  [][]rune
    Width  int
    Height int
}

const (
    WALL   rune = '#'
    PATH   rune = '-' 
    PLAYER rune = 'P'
    ENEMY  rune = 'X'
    EXIT   rune = 'E'
)

func (g *Grid) IsValidMove(pos player.Position) bool {
    return pos.Row >= 0 && pos.Row < g.Height &&
           pos.Col >= 0 && pos.Col < g.Width &&
           g.Cells[pos.Row][pos.Col] != WALL
}

func (g *Grid) Render() string {
    var gridOutput string
    for _, row := range g.Cells {
        for _, cell := range row {
            gridOutput += string(cell)
        }
        gridOutput += "\n"
    }
    return gridOutput
}

func (g *Grid) GetCellType(pos player.Position) rune{
    return g.Cells[pos.Row][pos.Col]
}

func (g *Grid) MovePlayer(oldPos, newPos player.Position){
    g.Cells[oldPos.Row][oldPos.Col] = PATH
    g.Cells[newPos.Row][newPos.Col] = PLAYER
}

func (g *Grid) PlacePlayer(pos player.Position) {
    g.Cells[pos.Row][pos.Col] = PLAYER
}

