package pieces

type Bqueen struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int
	isWhite     bool
}

func (p Bqueen) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Bqueen) GetNotation() string {
	return p.notation
}

func (p Bqueen) GetWeight() int {
	return p.weight
}

func (p Bqueen) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Bqueen) IsWhite() bool {
	return p.isWhite
}

func InitBqueen() Bqueen {
	p := Bqueen{}
	p.notation = "b"
	p.isWhite = false
	p.weight = 90
	p.weightBoard = [8][8]int{
		{-20, -10, -10, -5, -5, -10, -10, -20},
		{-10, 0, 0, 0, 0, 0, 0, -10},
		{-10, 0, 5, 5, 5, 5, 0, -10},
		{-5, 0, 5, 5, 5, 5, 0, -5},
		{0, 0, 5, 5, 5, 5, 0, -5},
		{-10, 5, 5, 5, 5, 5, 0, -10},
		{-10, 0, 5, 0, 0, 0, 0, -10},
		{-20, -10, -10, -5, -5, -10, -10, -20},
	}
	p.movesMatrix = [][]int{
		{1, 1}, {-1, 1}, {-1, -1}, {1, -1},
	}

	return p
}
