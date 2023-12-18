package main

import (
	"board"
	"fmt"
	"search"
)

func main() {
	board := board.InitBoard()
	fmt.Println(board)
	fmt.Println(search.GetAllValidPositions(&board, true))
}
