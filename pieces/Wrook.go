package pieces

type Wrook struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int
	isWhite     bool
}

func (p Wrook) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Wrook) GetNotation() string {
	return p.notation
}

func (p Wrook) GetWeight() int {
	return p.weight
}

func (p Wrook) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Wrook) IsWhite() bool {
	return p.isWhite
}

func InitWrook() Wrook {
	p := Wrook{}
	p.isWhite = true
	p.notation = "R"
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
