package game

import "fmt"

type Player struct {
	name  string
	color string
}

func newPlayer(name, color string) *Player {
	return &Player{
		name:  name,
		color: color,
	}
}

func (p *Player) getMove() (int, int, int, int) {
	var xFrom, yFrom, xTo, yTo int
	fmt.Printf("%s, choose a piece you want to move (row and column): ", p.name)
	fmt.Scanf("%d %d", &xFrom, &yFrom)
	fmt.Print("\nNow choose a cell you want to move your piece to: ")
	fmt.Scanf("%d %d", &xTo, &yTo)
	return xFrom, yFrom, xTo, yTo
}
