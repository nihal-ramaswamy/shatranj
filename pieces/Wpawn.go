package pieces

type Wpawn struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int // Need to handle capture en passant
	isWhite     bool
}

func (p Wpawn) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Wpawn) GetNotation() string {
	return p.notation
}

func (p Wpawn) GetWeight() int {
	return p.weight
}

func (p Wpawn) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Wpawn) IsWhite() bool {
	return p.isWhite
}

func InitWpawn() Wpawn {
	p := Wpawn{}
	p.notation = "P"
	p.isWhite = true
	p.weight = 10
	p.weightBoard = [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{50, 50, 50, 50, 50, 50, 50, 50},
		{10, 10, 20, 30, 30, 20, 10, 10},
		{5, 5, 10, 25, 25, 10, 5, 5},
		{0, 0, 0, 20, 20, 0, 0, 0},
		{5, -5, -10, 0, 0, -10, -5, 5},
		{5, 10, 10, -20, -20, 10, 10, 5},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	p.movesMatrix = [][]int{
		{-1, 0},
	}

	return p
}
