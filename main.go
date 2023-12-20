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
	fmt.Println(search.GetAllValidPositions(&board, true))
}
