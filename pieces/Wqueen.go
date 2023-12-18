package pieces

type Wqueen struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int
	isWhite     bool
}

func (p Wqueen) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Wqueen) GetNotation() string {
	return p.notation
}

func (p Wqueen) GetWeight() int {
	return p.weight
}

func (p Wqueen) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Wqueen) IsWhite() bool {
	return p.isWhite
}

func InitWqueen() Wqueen {
	p := Wqueen{}
	p.isWhite = true
	p.notation = "Q"
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
