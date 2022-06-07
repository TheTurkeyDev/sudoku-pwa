package main

type Cell struct {
	row int
	col int
}

func CellsForRow(row int) []*Cell {
	cells := make([]*Cell, 0)
	for c := 0; c < 9; c++ {
		cells = append(cells, &Cell{row: row, col: c})
	}
	return cells
}

func CellsForCol(col int) []*Cell {
	cells := make([]*Cell, 0)
	for r := 0; r < 9; r++ {
		cells = append(cells, &Cell{row: r, col: col})
	}
	return cells
}
