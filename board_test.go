package game

import (
	"testing"
)

func TestHasTicTacToe(t *testing.T) {
	testCase := []struct {
		HasTicTacToe bool
		TestBoard    Board
	}{
		{
			false,
			Board{
				{E, E, E},
				{E, E, E},
				{E, E, E},
			},
		},
		{
			true,
			Board{
				{X, X, X},
				{E, E, E},
				{E, E, E},
			},
		},
		{
			true,
			Board{
				{X, E, E},
				{X, E, E},
				{X, E, E},
			},
		},
		{
			true,
			Board{
				{X, E, E},
				{E, X, E},
				{E, E, X},
			},
		},
		{
			false,
			Board{
				{X, E, E},
				{E, O, E},
				{E, E, X},
			},
		},
		{
			false,
			Board{
				{X, E, E},
				{O, E, E},
				{X, E, E},
			},
		},
		{
			false,
			Board{
				{X, O, X},
				{E, E, E},
				{E, E, E},
			},
		},
	}

	for idx, board := range testCase {
		if board.HasTicTacToe != board.TestBoard.HasTicTacToe() {
			t.Errorf("Board #%d should have returned %t \n%s\n", idx, board.HasTicTacToe, board.TestBoard.Render())
		}
	}
}
