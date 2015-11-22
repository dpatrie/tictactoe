package game

import (
	"fmt"
	"strings"
)

type square byte

const (
	X = 'X'
	O = 'O'
	E = ' '
)

func (s square) IsX() bool {
	return s == X
}

func (s square) IsO() bool {
	return s == O
}

func (s square) IsEmpty() bool {
	return s == E
}

func (s square) SetX() {
	s = X
}

func (s square) SetO() {
	s = O
}

func (s square) String() string {
	return string(s)
}

const boardSize = 3

type Board [boardSize][boardSize]square

func (b Board) HasTicTacToe() bool {
	return b.horizontalTicTacToe() || b.verticalTicTacToe() || b.diagonalTicTacToe()
}

func (b Board) Render() string {
	var board []string
	for _, row := range b {
		board = append(board, b.renderRow(row))
	}
	return strings.Join(board, "\n----------\n")
}

func (b Board) renderRow(row [boardSize]square) string {
	return fmt.Sprintf("%s | %s | %s", row[0], row[1], row[2])
}

//0,0  0,1  0,2
//1,0  1,1  1,2
//2,0  2,1  2,2

func (b Board) horizontalTicTacToe() bool {
	for _, row := range b {
		if row[0] == E || row[1] == E || row[2] == E {
			return false
		}
		if row[0] == row[1] && row[1] == row[2] {
			return true
		}
	}
	return false
}

func (b Board) verticalTicTacToe() bool {
	//TODO: Stef
	return false
}

func (b Board) diagonalTicTacToe() bool {
	//TODO: Stef
	return false
}
