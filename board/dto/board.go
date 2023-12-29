package board

import (
	"fmt"
	"pieces"
	"player"
)

type Board struct {
	BPawn   uint64
	BKnight uint64
	BBishop uint64
	BRook   uint64
	BQueen  uint64
	BKing   uint64

	WPawn   uint64
	WKnight uint64
	WBishop uint64
	WRook   uint64
	WQueen  uint64
	WKing   uint64

	Turn  bool
	White player.Player
	Black player.Player
}

func (b Board) GetPieceAt(row, col int) pieces.PieceInterface {
	var pieceType uint64 = 1 << uint64(row*8+col)
	if 0 != b.BPawn&pieceType {
		return pieces.BpawnPiece
	} else if 0 != b.BKnight&pieceType {
		return pieces.BknightPiece
	} else if 0 != b.BBishop&pieceType {
		return pieces.BbishopPiece
	} else if 0 != b.BRook&pieceType {
		return pieces.BrookPiece
	} else if 0 != b.BQueen&pieceType {
		return pieces.BqueenPiece
	} else if 0 != b.BKing&pieceType {
		return pieces.BkingPiece
	} else if 0 != b.WPawn&pieceType {
		return pieces.WpawnPiece
	} else if 0 != b.WKnight&pieceType {
		return pieces.WknightPiece
	} else if 0 != b.WBishop&pieceType {
		return pieces.WbishopPiece
	} else if 0 != b.WRook&pieceType {
		return pieces.WrookPiece
	} else if 0 != b.WQueen&pieceType {
		return pieces.WqueenPiece
	} else if 0 != b.WKing&pieceType {
		return pieces.WkingPiece
	} else {
		return nil
	}
}

func (b Board) CopyBoard() Board {
	return Board{
		BPawn:   b.BPawn,
		BKnight: b.BKnight,
		BBishop: b.BBishop,
		BRook:   b.BRook,
		BQueen:  b.BQueen,
		BKing:   b.BKing,
		WPawn:   b.WPawn,
		WKnight: b.WKnight,
		WBishop: b.WBishop,
		WRook:   b.WRook,
		WQueen:  b.WQueen,
		WKing:   b.WKing,
		Turn:    b.Turn,
		White:   b.White,
		Black:   b.Black,
	}
}

func (b *Board) Move(
	src []int,
	dest []int,
) {
	srcPiece := b.GetPieceAt(src[0], src[1])
	destPiece := b.GetPieceAt(dest[0], dest[1])
	if srcPiece == nil {
		return
	}

	b.Turn = !b.Turn

	var srcPiecePos uint64 = 1 << uint64(src[0]*8+src[1])
	var destPiecePos uint64 = 1 << uint64(dest[0]*8+dest[1])

	switch srcPiece.GetNotation() {
	case pieces.WpawnPiece.GetNotation():
		b.WPawn ^= srcPiecePos
		b.WPawn |= destPiecePos
	case pieces.WknightPiece.GetNotation():
		b.WKnight ^= srcPiecePos
		b.WKnight |= destPiecePos
	case pieces.WbishopPiece.GetNotation():
		b.WBishop ^= srcPiecePos
		b.WBishop |= destPiecePos
	case pieces.WrookPiece.GetNotation():
		b.WRook ^= srcPiecePos
		b.WRook |= destPiecePos
	case pieces.WqueenPiece.GetNotation():
		b.WQueen ^= srcPiecePos
		b.WQueen |= destPiecePos
	case pieces.WkingPiece.GetNotation():
		b.WKing ^= srcPiecePos
		b.WKing |= destPiecePos
	case pieces.BpawnPiece.GetNotation():
		b.BPawn ^= srcPiecePos
		b.BPawn |= destPiecePos
	case pieces.BknightPiece.GetNotation():
		b.BKnight ^= srcPiecePos
		b.BKnight |= destPiecePos
	case pieces.BbishopPiece.GetNotation():
		b.BBishop ^= srcPiecePos
		b.BBishop |= destPiecePos
	case pieces.BrookPiece.GetNotation():
		b.BRook ^= srcPiecePos
		b.BRook |= destPiecePos
	case pieces.BqueenPiece.GetNotation():
		b.BQueen ^= srcPiecePos
		b.BQueen |= destPiecePos
	case pieces.BkingPiece.GetNotation():
		b.BKing ^= srcPiecePos
		b.BKing |= destPiecePos
	}

	if nil != destPiece {

		switch destPiece.GetNotation() {
		case pieces.WpawnPiece.GetNotation():
			b.WPawn ^= destPiecePos
		case pieces.WknightPiece.GetNotation():
			b.WKnight ^= destPiecePos
		case pieces.WbishopPiece.GetNotation():
			b.WBishop ^= destPiecePos
		case pieces.WrookPiece.GetNotation():
			b.WRook ^= destPiecePos
		case pieces.WqueenPiece.GetNotation():
			b.WQueen ^= destPiecePos
		case pieces.WkingPiece.GetNotation():
			b.WKing ^= destPiecePos
		case pieces.BpawnPiece.GetNotation():
			b.BPawn ^= destPiecePos
		case pieces.BknightPiece.GetNotation():
			b.BKnight ^= destPiecePos
		case pieces.BbishopPiece.GetNotation():
			b.BBishop ^= destPiecePos
		case pieces.BrookPiece.GetNotation():
			b.BRook ^= destPiecePos
		case pieces.BqueenPiece.GetNotation():
			b.BQueen ^= destPiecePos
		case pieces.BkingPiece.GetNotation():
			b.BKing ^= destPiecePos

		}
	}
}

func (b Board) String() string {
	var boardString string
	boardString += "  +-----------------+\n"
	for i := 0; i < 8; i++ {
		boardString += fmt.Sprintf("%d | ", 8-i)
		for j := 0; j < 8; j++ {
			piece := b.GetPieceAt(i, j)
			var toPrint string
			if piece == nil {
				toPrint = " "
			} else {
				toPrint = piece.GetNotation()
			}
			boardString += fmt.Sprintf("%s ", toPrint)
		}
		boardString += "|\n"
	}
	boardString += "  +-----------------+\n"
	boardString += "    a b c d e f g h\n"
	boardString += fmt.Sprintf("Turn: %t\n", b.Turn)
	boardString += fmt.Sprintf("White: %s\n", b.White.GetName())
	boardString += fmt.Sprintf("Black: %s\n", b.Black.GetName())
	return boardString
}
