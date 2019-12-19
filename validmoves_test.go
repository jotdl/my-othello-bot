package main

import (
	"testing"

	"github.com/jotdl/othello"
)

func TestValidMoves(t *testing.T) {
	// shortcuts for different colors
	n, b, w := othello.None, othello.Black, othello.White

	// array of test cases for the method we want to test
	var testCases = []struct {
		name string // name of test case

		board  othello.Board // board for test case
		player othello.Color // current player we wan't to calculate turns for

		expectedValidTurns []othello.Turn // turns we expect to be calculated
	}{
		{
			name:   "Initial start for black",
			player: othello.Black,
			board: othello.Board{
				Dimension: 8,
				Fields: [][]othello.Color{
					{n, n, n, n, n, n, n, n},
					{n, n, n, n, n, n, n, n},
					{n, n, n, n, n, n, n, n},
					{n, n, n, b, w, n, n, n},
					{n, n, n, w, b, n, n, n},
					{n, n, n, n, n, n, n, n},
					{n, n, n, n, n, n, n, n},
					{n, n, n, n, n, n, n, n},
				},
			},
			expectedValidTurns: []othello.Turn{
				{Row: 2, Column: 4},
				{Row: 3, Column: 5},
				{Row: 4, Column: 2},
				{Row: 5, Column: 3},
			},
		},
	}

	// run all the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// calculate possible turns
			turns := findValidTurns(&testCase.board, testCase.player)

			// check that all expected turns are within calculated turns
			for _, expectedMove := range testCase.expectedValidTurns {
				found := false
				for _, move := range turns {
					found = found || move.Column == expectedMove.Column && move.Row == expectedMove.Row
				}

				if !found {
					t.Errorf("Didn't find expected move for row %v and column %v", expectedMove.Row, expectedMove.Column)
				}
			}
		})
	}
}
