package player

import (
	"testing"
)


func TestPlayerMovement(t *testing.T){
    testCases := []struct{
        name string
        startRow int
        startCol int
        rowOffset int
        colOffset int
        expectedRow int
        expectedCol int
    }{
        {
            name: "Move right",
            startRow: 1,
            startCol: 1,
            rowOffset: 0,
            colOffset: 1,
            expectedRow: 1,
            expectedCol: 2,
        },
        {
            name: "Move left",
            startRow: 1,
            startCol: 1,
            rowOffset: 0,
            colOffset: -1,
            expectedRow: 1,
            expectedCol: 0,
        },
        {
            name: "Move down",
            startRow: 1,
            startCol: 1,
            rowOffset: 1,
            colOffset: 0,
            expectedRow: 2,
            expectedCol: 1,
        },
        {
            name: "Move up",
            startRow: 1,
            startCol: 1,
            rowOffset: -1,
            colOffset: 0,
            expectedRow: 0,
            expectedCol: 1,
        },
    }

    for _, test := range testCases{
        t.Run(test.name, func(t *testing.T){
            player := New(test.startRow, test.startCol, 100)
            player.Move(test.rowOffset, test.colOffset)

            if player.Position.Row != test.expectedRow  || player.Position.Col != test.expectedCol {
                t.Errorf("%s: Expected position (%d,%d), got (%d,%d)",
                    test.name, test.expectedRow, test.expectedCol,
                    player.Position.Row, player.Position.Col)
            }
        })
    }
}
