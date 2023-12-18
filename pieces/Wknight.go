package pieces

type Wknight struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int
	isWhite     bool
}

func (p Wknight) GetNotation() string {
	return p.notation
}

func (p Wknight) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Wknight) GetWeight() int {
	return p.weight
}

func (p Wknight) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Wknight) IsWhite() bool {
	return p.isWhite
}

func InitWknight() Wknight {
	p := Wknight{}
	p.isWhite = true
	p.notation = "N"
	p.weight = 30
	p.weightBoard = [8][8]int{
		{-50, -40, -30, -30, -30, -30, -40, -50},
		{-40, -20, 0, 0, 0, 0, -20, -40},
		{-30, 0, 10, 15, 15, 10, 0, -30},
		{-30, 5, 15, 20, 20, 15, 5, -30},
		{-30, 0, 15, 20, 20, 15, 0, -30},
		{-30, 5, 10, 15, 15, 10, 5, -30},
		{-40, -20, 0, 5, 5, 0, -20, -40},
		{-50, -40, -30, -30, -30, -30, -40, -50},
	}
	p.movesMatrix = [][]int{
		{1, 2}, {2, 1}, {2, -1}, {1, -2},
		{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2},
	}
	return p
}
