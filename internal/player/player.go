package player

type Position struct {
    Row int
    Col int
}

type Player struct {
    Position Position
    Health   int
}

func New(row, col int, initHealth int) *Player {
    return &Player{
        Position: Position{Row: row, Col: col}, 
        Health: initHealth,
    }
}

func (p *Player) Move(rowOffset, colOffset int) {
    p.Position.Row += rowOffset
    p.Position.Col += colOffset
}

func (p *Player) IsAlive() bool {
    return p.Health > 0
}

func (p *Player) GetNewPosition(rowOffset, colOffset int) Position{
    return Position{
        Row: p.Position.Row + rowOffset,
        Col: p.Position.Col + colOffset,
    }
}

func (p *Player) GetPosition() Position{
    return p.Position
}

func (p *Player) GetHealth() int{
    return p.Health
}
