package search

import (
	boardDto "board/dto"
	"evaluation"
)

const DEPTH = 3
const START_WHITE = -9999
const START_BLACK = 9999

func GetBestMove(board boardDto.Board, playerIsWhite bool) [][]int {

	allValidMoves := GetAllValidPositions(board, playerIsWhite)

	if len(allValidMoves) == 0 {
		return [][]int{}
	}

	var bestMove [][]int = allValidMoves[0]
	bestValue := START_WHITE

	if !playerIsWhite {
		bestValue = START_BLACK
	}

	for _, move := range allValidMoves {
		board.Move(move[0], move[1])
		value := minimax_alpha_beta_pruning(board, DEPTH, START_WHITE, START_BLACK, !playerIsWhite)
		board.Move(move[1], move[0])

		if playerIsWhite {
			if value > bestValue {
				bestValue = value
				bestMove = move
			}
		} else {
			if value < bestValue {
				bestValue = value
				bestMove = move
			}
		}

	}

	return bestMove
}

func minimax_alpha_beta_pruning(
	board boardDto.Board,
	depth int,
	alpha int,
	beta int,
	playerIsWhite bool,
) int {
	allValidMoves := GetAllValidPositions(board, playerIsWhite)

	if len(allValidMoves) == 0 || depth == 0 {
		return evaluation.EvaluateBoard(board)
	}

	if playerIsWhite {
		bestValue := START_WHITE
		for _, move := range allValidMoves {
			board.Move(move[0], move[1])
			value := minimax_alpha_beta_pruning(board, depth-1, alpha, beta, false)
			board.Move(move[1], move[0])
			bestValue = max(bestValue, value)
			alpha = max(alpha, bestValue)
			if beta <= alpha {
				break
			}
		}
		return bestValue
	} else {
		bestValue := START_BLACK
		for _, move := range allValidMoves {
			board.Move(move[0], move[1])
			value := minimax_alpha_beta_pruning(board, depth-1, alpha, beta, true)
			board.Move(move[1], move[0])
			bestValue = min(bestValue, value)
			beta = min(beta, bestValue)
			if beta <= alpha {
				break
			}
		}
		return bestValue
	}
}
