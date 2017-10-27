package main

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
		t.Logf("Testing board #%d\n\n%s\n\n", idx, board.TestBoard.Render())
		if board.HasTicTacToe != board.TestBoard.HasTicTacToe() {
			t.Errorf("Should have returned %t\n", board.HasTicTacToe)
		} else {
			t.Log("Success!!\n")
		}
	}
}
