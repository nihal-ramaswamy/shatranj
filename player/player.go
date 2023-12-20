package player

type Player struct {
	name  string
	color bool // White = true, Black = false
}

func (p Player) GetName() string {
	return p.name
}

func (p Player) GetColor() bool {
	return p.color
}

func NewPlayer(name string, color bool) Player {
	return Player{
		name:  name,
		color: color,
	}
}
