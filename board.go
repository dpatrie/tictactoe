package main

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

func (s square) String() string {
	return string(s)
}

const boardSize = 3

type Board [boardSize][boardSize]square

//0,0  0,1  0,2
//1,0  1,1  1,2
//2,0  2,1  2,2

func (b Board) HasTicTacToe() bool {
	return b.horizontalTicTacToe() || b.verticalTicTacToe() || b.diagonalTicTacToe()
}

func (b Board) Render() string {
	var board []string
	for _, row := range b {
		board = append(board, b.renderRow(row))
	}
	return strings.Join(board, "\n------------\n")
}

func (b Board) renderRow(row [boardSize]square) string {
	return fmt.Sprintf(" %s | %s | %s ", row[0], row[1], row[2])
}

func (b Board) horizontalTicTacToe() bool {
	return (b[0][0] != E && b[0][0] == b[0][1] && b[0][0] == b[0][2]) ||
		(b[1][0] != E && b[1][0] == b[1][1] && b[1][0] == b[1][2]) ||
		(b[2][0] != E && b[2][0] == b[2][1] && b[2][0] == b[2][2])
}

func (b Board) verticalTicTacToe() bool {
	return (b[0][0] != E && b[0][0] == b[1][0] && b[0][0] == b[2][0]) ||
		(b[0][1] != E && b[0][1] == b[1][1] && b[0][1] == b[2][1]) ||
		(b[0][2] != E && b[0][2] == b[1][2] && b[0][2] == b[2][2])
}

func (b Board) diagonalTicTacToe() bool {
	return (b[0][0] != E && b[0][0] == b[1][1] && b[0][0] == b[2][2]) ||
		(b[0][2] != E && b[0][2] == b[1][1] && b[0][2] == b[2][0])
}
