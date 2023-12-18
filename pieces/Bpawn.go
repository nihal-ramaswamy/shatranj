package pieces

type Bpawn struct {
	notation    string
	weightBoard [8][8]int
	weight      int
	movesMatrix [][]int // Need to handle capture en passant
	isWhite     bool
}

func (p Bpawn) GetWeightBoard() [8][8]int {
	return p.weightBoard
}

func (p Bpawn) GetNotation() string {
	return p.notation
}

func (p Bpawn) GetWeight() int {
	return p.weight
}

func (p Bpawn) GetMovesMatrix() [][]int {
	return p.movesMatrix
}

func (p Bpawn) IsWhite() bool {
	return p.isWhite
}

func InitBpawn() Bpawn {
	p := Bpawn{}
	p.isWhite = false
	p.notation = "p"
	p.weight = 10
	p.weightBoard = [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{5, 10, 10, -20, -20, 10, 10, 5},
		{5, -5, -10, 0, 0, -10, -5, 5},
		{0, 0, 0, 20, 20, 0, 0, 0},
		{5, 5, 10, 25, 25, 10, 5, 5},
		{10, 10, 20, 30, 30, 20, 10, 10},
		{50, 50, 50, 50, 50, 50, 50, 50},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	p.movesMatrix = [][]int{
		{1, 0},
	}

	return p
}
