package pieces

type Wking struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int // Need to handle castling
	isWhite     bool
}

func (p Wking) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Wking) GetNotation() string {
	return p.notation
}

func (p Wking) GetWeight() int {
	return p.weight
}

func (p Wking) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Wking) IsWhite() bool {
	return p.isWhite
}

func InitWking() Wking {
	p := Wking{}
	p.notation = "K"
	p.isWhite = true
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
