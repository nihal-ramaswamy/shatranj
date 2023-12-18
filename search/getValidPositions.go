package search

import (
	boardDto "board/dto"
	"pieces"
)

func getAllValidPositionsForPieceAtPosition(
	board *boardDto.Board,
	positions [][][]int,
	file, rank int,
	newFile, newRank int,
) [][][]int {
	if newFile < 0 || newFile > 7 || newRank < 0 || newRank > 7 {
		return positions
	}
	pieceAtPosition := board.GetPieceAt(file, rank)
	pieceAtNewPosition := board.GetPieceAt(newFile, newRank)
	if pieceAtNewPosition != nil &&
		pieceAtNewPosition.IsWhite() == pieceAtPosition.IsWhite() {
		return positions
	}
	positions = append(positions, [][]int{{file, rank}, {newFile, newRank}})
	return positions
}

func getAllValidPositionsForPiece(
	board *boardDto.Board,
	positions [][][]int,
	file, rank int,
) [][][]int {
	pieceAtPosition := board.GetPieceAt(file, rank)
	if pieceAtPosition == nil {
		return positions
	}
	possibleMoves := pieceAtPosition.GetMovesMatrix()
	for _, move := range possibleMoves {
		// Handling case for rook: Can go any number of steps until piece is found
		if pieceAtPosition.GetNotation() == pieces.BrookPiece.GetNotation() ||
			pieceAtPosition.GetNotation() == pieces.WrookPiece.GetNotation() {
			pieceFound := false
			for i := -8; i < -1; i++ {
				if pieceFound {
					break
				}
				tempFile := file + i*move[0]
				tempRank := rank + i*move[1]
				pieceFound = board.GetPieceAt(tempFile, tempRank) != nil
				positions = getAllValidPositionsForPieceAtPosition(
					board, positions, file, rank, tempFile, tempRank)
			}
			pieceFound = false
			for i := 1; i < 8; i++ {
				if pieceFound {
					break
				}
				tempFile := file + i*move[0]
				tempRank := rank + i*move[1]
				pieceFound = board.GetPieceAt(tempFile, tempRank) != nil
				positions = getAllValidPositionsForPieceAtPosition(
					board, positions, file, rank, tempFile, tempRank)
			}

		} else {
			tempFile := file + move[0]
			tempRank := rank + move[1]
			positions = getAllValidPositionsForPieceAtPosition(
				board, positions, file, rank, tempFile, tempRank)
		}

		// Handling case for pawn: Can go diagonally if enemy piece is found
		if pieceAtPosition.GetNotation() == pieces.WpawnPiece.GetNotation() ||
			pieceAtPosition.GetNotation() == pieces.BpawnPiece.GetNotation() {
			tempFile := file + -1
			tempRank := rank + 1
			pieceAtNewPosition := board.GetPieceAt(tempFile, tempRank)
			if pieceAtNewPosition != nil {
				positions = getAllValidPositionsForPieceAtPosition(
					board, positions, file, rank, tempFile, tempRank)
			}
			tempFile = file + 1
			tempRank = rank + 1
			pieceAtNewPosition = board.GetPieceAt(tempFile, tempRank)
			if pieceAtNewPosition != nil {

				positions = getAllValidPositionsForPieceAtPosition(
					board, positions, file, rank, tempFile, tempRank)
			}

		}
	}
	return positions
}

func GetAllValidPositions(board *boardDto.Board, playerIsWhite bool) [][][]int {
	var positions [][][]int
	for i := 0; i < 64; i++ {
		file := i / 8
		rank := i % 8
		pieceAtPosition := board.GetPieceAt(file, rank)

		if pieceAtPosition == nil || pieceAtPosition.IsWhite() != playerIsWhite {
			continue
		}

		positions = getAllValidPositionsForPiece(board, positions, file, rank)

	}
	return positions
}
