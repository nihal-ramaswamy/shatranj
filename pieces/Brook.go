package pieces

type Brook struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int
	isWhite     bool
}

func (p Brook) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Brook) GetNotation() string {
	return p.notation
}

func (p Brook) GetWeight() int {
	return p.weight
}

func (p Brook) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Brook) IsWhite() bool {
	return p.isWhite
}

func InitBrook() Brook {
	p := Brook{}
	p.notation = "r"
	p.isWhite = false
	p.weight = 50
	p.weightBoard = [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{5, 10, 10, 10, 10, 10, 10, 5},
		{-5, 0, 0, 0, 0, 0, 0, -5},
		{-5, 0, 0, 0, 0, 0, 0, -5},
		{-5, 0, 0, 0, 0, 0, 0, -5},
		{-5, 0, 0, 0, 0, 0, 0, -5},
		{-5, 0, 0, 0, 0, 0, 0, -5},
		{0, 0, 0, 5, 5, 0, 0, 0},
	}

	p.movesMatrix = [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	return p
}
