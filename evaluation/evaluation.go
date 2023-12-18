package evaluation

import (
	boardDto "board/dto"
)

func EvaluateBoard(board *boardDto.Board) int {
	var score int
	score = 0.0
	for i := 0; i < 64; i++ {
		score += GetPieceValue(board.GetPieceAt(i/8, i%8), i/8, i%8)
	}
	return score
}
