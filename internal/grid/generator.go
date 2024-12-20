package grid

import (

	"math/rand"
    "time"

)

func GenerateNewGrid(width, height, enemyCount int) *Grid{
    grid := &Grid{
        Width: width,
        Height: height,
        Cells: make([][]rune, height),
    }

    for i := range grid.Cells {
        grid.Cells[i] = make([]rune, width)
        for j := range grid.Cells[i] {
            if i == 0 || i == height-1 || j == 0 || j == width-1{
                grid.Cells[i][j] = WALL
            } else {
                grid.Cells[i][j] = PATH
            }
        }
    }

    grid.Cells[height-2][width-2] = EXIT
    grid.GenerateEnemies(enemyCount)

    return grid
}

func (g* Grid) GenerateEnemies(enemyCount int){
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < enemyCount; {
        row := r.Intn(g.Height)
        col := r.Intn(g.Width)
        if g.Cells[row][col] == PATH {
            g.Cells[row][col] = ENEMY
            i++
        }
    }
}
