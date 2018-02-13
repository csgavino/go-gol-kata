package csgavino

type Gol struct {
	Cells []Cell
}

func NewGol(cells ...Cell) Gol {
	gol := Gol{}
	gol.Cells = cells
	return gol
}

func (gol Gol) Neighbours(x, y int) int {
	aliveCells := []Cell{}

	for _, cell := range gol.Cells {
		if cell.Alive {
			aliveCells = append(aliveCells, cell)
		}
	}

	neighbours := 0
	for _, aliveCell := range aliveCells {
		if (aliveCell.X == x-1 && aliveCell.Y == y-1) ||
			(aliveCell.X == x-1 && aliveCell.Y == y) ||
			(aliveCell.X == x-1 && aliveCell.Y == y+1) ||
			(aliveCell.X == x && aliveCell.Y == y-1) ||
			(aliveCell.X == x && aliveCell.Y == y+1) ||
			(aliveCell.X == x+1 && aliveCell.Y == y-1) ||
			(aliveCell.X == x+1 && aliveCell.Y == y) ||
			(aliveCell.X == x+1 && aliveCell.Y == y+1) {
			neighbours += 1
		}
	}

	return neighbours
}
func (gol Gol) At(x, y int) *Cell {
	for _, cell := range gol.Cells {
		if cell.X == x &&
			cell.Y == y {
			return &cell
		}
	}

	return nil
}
