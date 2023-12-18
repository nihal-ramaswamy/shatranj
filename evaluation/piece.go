package evaluation

import (
	"pieces"
)

var (
	Wpawn   = pieces.InitWpawn()
	Wknight = pieces.InitWknight()
	Wbishop = pieces.InitWbishop()
	Wrook   = pieces.InitWrook()
	Wqueen  = pieces.InitWqueen()
	Wking   = pieces.InitWking()

	Bpawn   = pieces.InitBpawn()
	Bknight = pieces.InitBknight()
	Bbishop = pieces.InitBbishop()
	Brook   = pieces.InitBrook()
	Bqueen  = pieces.InitBqueen()
	Bking   = pieces.InitBking()
)

func GetPieceValue(piece pieces.PieceInterface, file, rank int) int {
	switch piece.GetNotation() {
	case Bpawn.GetNotation():
		return -(Bpawn.GetWeightBoard()[file][rank] + Bpawn.GetWeight())
	case Bknight.GetNotation():
		return -(Bknight.GetWeightBoard()[file][rank] + Bknight.GetWeight())
	case Bbishop.GetNotation():
		return -(Bbishop.GetWeightBoard()[file][rank] + Bbishop.GetWeight())
	case Brook.GetNotation():
		return -(Brook.GetWeightBoard()[file][rank] + Brook.GetWeight())
	case Bqueen.GetNotation():
		return -(Bqueen.GetWeightBoard()[file][rank] + Bqueen.GetWeight())
	case Bking.GetNotation():
		return -(Bking.GetWeightBoard()[file][rank] + Bking.GetWeight())
	case Wpawn.GetNotation():
		return Wpawn.GetWeightBoard()[file][rank] + Wpawn.GetWeight()
	case Wknight.GetNotation():
		return Wknight.GetWeightBoard()[file][rank] + Wknight.GetWeight()
	case Wbishop.GetNotation():
		return Wbishop.GetWeightBoard()[file][rank] + Wbishop.GetWeight()
	case Wrook.GetNotation():
		return Wrook.GetWeightBoard()[file][rank] + Wrook.GetWeight()
	case Wqueen.GetNotation():
		return Wqueen.GetWeightBoard()[file][rank] + Wqueen.GetWeight()
	case Wking.GetNotation():
		return Wking.GetWeightBoard()[file][rank] + Wking.GetWeight()
	}
	return 0.0
}
