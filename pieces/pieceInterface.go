package pieces

type PieceInterface interface {
	GetWeightBoard() [8][8]int
	GetNotation() string
	GetWeight() int
	GetMovesMatrix() [][]int
	IsWhite() bool
}
