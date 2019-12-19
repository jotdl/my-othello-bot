# My Othello Bot

## Program your ultimate Othello/Reversi bot!

This is a simple starter kit to build your own othello bot and test it against various players (including yourself!) and AI's.

## Get started

1. Clone this repository

```
    git clone https://github.com/jotdl/my-othello-bot
```

2. Cd into this repository

```
    cd my-othello-bot
```

3. Run

```
    go run main.go
```

## First Task

To get your bot started you first need to know what is a valid turn and what is not.
Therefore:

1. Implement the `findValidTurns` method in `main.go`. Run `go test .` to simply check if your bot returns the valid moves for the initial board (see `validMoves_test.go` for further details). Iterate until you feel safe that you calculate all possible turns. Feel free to add more testcases to `validMoves_test.go`.
2. Once you feel safe and want to test your logic against something more advanced, replace `mybot := othello.NewBraindeadAIPlayer()` with `mybot := othello.FuncPlayer(calculateNextTurn)`. When you now run `go run .` you successfully play against a bot which always takes the first move you return from `findValidTurns`

## Second+ Task

Implement the ultimate othello bot which will win against all other bots. Therefore extend `calculateNextTurn` in `main.go` with logic you think fits best. To test your bot simply run `go run .`.

You can also adjust the bot you are playing against by changing the assignment of the `enemy`.

## Available enemies

The following enemy bots/players are currently available:

- _TerminalPlayer_: A bot which reads every turn from stdin so you can play against your bot.
- _BraindeadAIPlayer_: A bot which just does a random valid move.
- _NewSimpleAIPlayer_: Always takes the turn which leads to his highest current score.
- _AdvancedAIPlayer_: A little bit more strategic bot which will try to outplay you.
- _MinMaxAIPlayer_: A bot which follows the MinMax algorithm (See [https://en.wikipedia.org/wiki/Minimax](Minimax-Algorithm) for further details)
