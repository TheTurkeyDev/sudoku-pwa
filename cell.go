package main

type Cell struct {
	row int
	col int
}

func CellsForRow(row int) []*Cell {
	cells := make([]*Cell, 9)
	for c := 0; c < 9; c++ {
		cells[c] = &Cell{row: row, col: c}
	}
	return cells
}

func CellsForCol(col int) []*Cell {
	cells := make([]*Cell, 9)
	for r := 0; r < 9; r++ {
		cells[r] = &Cell{row: r, col: col}
	}
	return cells
}

func CellsForBox(boxRow int, boxCol int) []*Cell {
	cells := make([]*Cell, 9)
	for i := 0; i < 9; i++ {
		cells[i] = &Cell{row: (boxRow * 3) + (i / 3), col: (boxCol * 3) + (i % 3)}
	}
	return cells
}
