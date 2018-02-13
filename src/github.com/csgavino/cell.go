package csgavino

type Cell struct {
	X          int
	Y          int
	Alive      bool
	Neighbours int
}

func (cell *Cell) Tick() {
	if cell.Alive {
		cell.Alive = cell.Neighbours > 1 && cell.Neighbours < 4
	} else {
		cell.Alive = cell.Neighbours == 3
	}
}
