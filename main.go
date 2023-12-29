package main

import (
	"board"
	"fmt"
	"player"
	"search"
)

func main() {
	whitePlayer := player.NewPlayer("white", true)
	blackPlayer := player.NewPlayer("black", false)
	board := board.InitBoard(whitePlayer, blackPlayer)

	fmt.Println(board)

	startPlayer := true
	for i := 0; i < 5; i++ {
		bestMove := search.GetBestMove(board, startPlayer)
		board.Move(bestMove[0], bestMove[1])
		startPlayer = !startPlayer
		fmt.Println(board)
	}
}
