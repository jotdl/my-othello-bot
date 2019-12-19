// package definition
package main

// import of dependencies
import (
	"fmt"
	"github.com/jotdl/othello"
	"time"
)

// main function
func main() {
	// create a bot from our calculation method
	//mybot := othello.FuncPlayer(calculateNextTurn)
	mybot := othello.NewBraindeadAIPlayer()

	// create a simple terminal based human player
	enemy := othello.NewTerminalPlayer()

	// create simple ai
	// enemy = othello.NewBraindeadAIPlayer()

	// create a bit advancd ai
	//enemy = othello.NewSimpleAIPlayer()

	// create an advanced ai
	// enemy = othello.NewAdvancedAIPlayer()

	// create a fully featured ai
	//enemy = othello.NewMinMaxAIPlayer()

	// create a new othello game with our bot and a terminal player
	game := othello.NewGame(mybot, enemy)

	// as long as the game is not finished repeat game execution
	for !game.Finished() {
		fmt.Println(game) // print the current board

		_, err := game.DoNextMove() // calculate next move
		if err != nil {
			panic(fmt.Errorf("error during game execution %w", err))
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println(game) // print the final board
	// print the final result
	fmt.Printf("The winner is \"Player %v\"!!!\n", game.Winner())
}

// findValidTurns needs to find all valid turns based on the given board for the current color
// Example: For the board below and player 1 (●) the moves (3,5), (4,6), (5,3) and (6,4) (Row, Column) should be valid.
//     0   1   2   3   4   5   6   7
//   ---------------------------------
// 0 |   |   |   |   |   |   |   |   |
//   ---------------------------------
// 1 |   |   |   |   |   |   |   |   |
//   ---------------------------------
// 2 |   |   |   |   |   |   |   |   |
//   ---------------------------------
// 3 |   |   |   | ● | ○ |   |   |   |
//   ---------------------------------
// 4 |   |   |   | ○ | ● |   |   |   |
//   ---------------------------------
// 5 |   |   |   |   |   |   |   |   |
//   ---------------------------------
// 6 |   |   |   |   |   |   |   |   |
//   ---------------------------------
// 7 |   |   |   |   |   |   |   |   |
//   ---------------------------------
func findValidTurns(board *othello.Board, currentPlayer othello.Color) []othello.Turn {
	return []othello.Turn{{}}
}

// calculateNextTurn calculates the next turn of our bot based on the given board
func calculateNextTurn(board *othello.Board, currentPlayer othello.Color) othello.Turn {
	// 1. find all valid turns
	moves := findValidTurns(board, currentPlayer)

	// 2. select best move
	bestMove := 0

	// 3. return a move
	return moves[bestMove]
}
