package board

import "board/dto"
import "player"

func InitBoard(whitePlayer, blackPlayer player.Player) board.Board {
	b := board.Board{}

	b.BPawn = 0x000000000000FF00
	b.BKnight = 0x0000000000000042
	b.BBishop = 0x0000000000000024
	b.BRook = 0x0000000000000081
	b.BQueen = 0x0000000000000008
	b.BKing = 0x0000000000000010
	b.WPawn = 0x00FF000000000000
	b.WKnight = 0x4200000000000000
	b.WBishop = 0x2400000000000000
	b.WRook = 0x8100000000000000
	b.WQueen = 0x0800000000000000
	b.WKing = 0x1000000000000000

	b.White = whitePlayer
	b.Black = blackPlayer
	b.Turn = true // White goes first

	return b
}
