package pieces

type Bking struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int // Need to handle castling
	isWhite     bool
}

func (p Bking) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Bking) GetNotation() string {
	return p.notation
}

func (p Bking) GetWeight() int {
	return p.weight
}

func (p Bking) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Bking) IsWhite() bool {
	return p.isWhite
}

func InitBking() Bking {
	p := Bking{}
	p.notation = "k"
	p.isWhite = false
	p.weight = 900
	p.weightBoard = [8][8]int{
		{-30, -40, -40, -50, -50, -40, -40, -30},
		{-30, -40, -40, -50, -50, -40, -40, -30},
		{-30, -40, -40, -50, -50, -40, -40, -30},
		{-30, -40, -40, -50, -50, -40, -40, -30},
		{-20, -30, -30, -40, -40, -30, -30, -20},
		{-10, -20, -20, -20, -20, -20, -20, -10},
		{20, 20, 0, 0, 0, 0, 20, 20},
		{20, 30, 10, 0, 0, 10, 30, 20},
	}
	p.movesMatrix = [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
		{1, 1}, {-1, 1}, {-1, -1}, {1, -1},
	}

	return p
}
