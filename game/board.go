package game

import (
	"fmt"
)

type Board struct {
	board [8][8]string
}

var whitePieces = map[string]bool{
	"♙": true,
}

var blackPieces = map[string]bool{
	"♟": true,
}

func NewBoard() *Board {
	return &Board{
		[8][8]string{
			{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
			{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
			{"□", "■", "□", "■", "□", "■", "□", "■"},
			{"■", "□", "■", "□", "■", "□", "■", "□"},
			{"□", "■", "□", "■", "□", "■", "□", "■"},
			{"■", "□", "■", "□", "■", "□", "■", "□"},
			{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
			{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
		},
	}
}

func (b *Board) Display() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(b.board[i][j])
		}
		fmt.Println()
	}
}

func (b *Board) ValidMove(xFrom, yFrom, xTo, yTo int, color string) bool {
	outOfBounds := func(a int) bool {
		return a < 0 || a >= 8
	}
	emptyCell := func(x, y int) bool {
		return b.board[x][y] == "■" || b.board[x][y] == "□"
	}
	cantTake := func(x, y int) bool {
		if emptyCell(x, y) {
			return true
		}
		if color == "white" {
			if _, exists := whitePieces[b.board[x][y]]; !exists {
				return true
			}
		} else {
			if _, exists := blackPieces[b.board[x][y]]; !exists {
				return true
			}
		}
		return false
	}
	cantPut := func(x, y int) bool { // only for pawn that moves one cell per move and cant attack
		if color == "white" {
			if _, exists := whitePieces[b.board[x][y]]; exists {
				return true
			}
		} else {
			if _, exists := blackPieces[b.board[x][y]]; exists {
				return true
			}
		}
		absInt := func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}
		if yFrom != y || absInt(xFrom-x) > 1 || x >= xFrom {
			return true
		}
		return false
	}

	if outOfBounds(xFrom) || outOfBounds(yFrom) || outOfBounds(xTo) || outOfBounds(yTo) {
		return false
	}
	if cantTake(xFrom, yFrom) || cantPut(xTo, yTo) {
		return false
	}

	b.board[xTo][yTo] = b.board[xFrom][yFrom]

	if (xFrom+yFrom)%2 == 0 {
		b.board[xFrom][yFrom] = "□"
	} else {
		b.board[xFrom][yFrom] = "■"
	}

	return true
}
