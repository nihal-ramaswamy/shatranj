package pieces

type Bbishop struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int
	isWhite     bool
}

func (p Bbishop) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Bbishop) GetNotation() string {
	return p.notation
}

func (p Bbishop) GetWeight() int {
	return p.weight
}

func (p Bbishop) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Bbishop) IsWhite() bool {
	return p.isWhite
}

func InitBbishop() Bbishop {
	p := Bbishop{}
	p.notation = "b"
	p.isWhite = false
	p.weight = 30
	p.weightBoard = [8][8]int{
		{-20, -10, -10, -10, -10, -10, -10, -20},
		{-10, 0, 0, 0, 0, 0, 0, -10},
		{-10, 0, 5, 10, 10, 5, 0, -10},
		{-10, 5, 5, 10, 10, 5, 5, -10},
		{-10, 0, 10, 10, 10, 10, 0, -10},
		{-10, 10, 10, 10, 10, 10, 10, -10},
		{-10, 5, 0, 0, 0, 0, 5, -10},
		{-20, -10, -10, -10, -10, -10, -10, -20},
	}

	p.movesMatrix = [][]int{
		{2, 0}, {0, 2}, {-2, 0}, {0, -2},
	}

	return p
}
