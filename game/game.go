package game

import "fmt"

type Game struct {
	board         *Board
	player1       *Player
	player2       *Player
	currentPlayer *Player
}

func NewGame() *Game {
	board := NewBoard()
	player1 := NewPlayer("Player 1", "white")
	player2 := NewPlayer("Player 2", "black")

	return &Game{
		board:         board,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
	}
}

func (g *Game) Start() {
	for {
		g.board.Display()
		xFrom, yFrom, xTo, yTo := g.currentPlayer.GetMove()
		fmt.Println(xFrom, yFrom, xTo, yTo)
		if !g.board.ValidMove(xFrom, yFrom, xTo, yTo, g.currentPlayer.color) {
			fmt.Println("Invalid move, try again.")
			continue
		}
		g.board.Display()
		break
	}
}

// func (g *Game) switchPlayer() {
// 	if g.currentPlayer == g.player1 {
// 		g.currentPlayer = g.player2
// 	} else {
// 		g.currentPlayer = g.player1
// 	}
// }
