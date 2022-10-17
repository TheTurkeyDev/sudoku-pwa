package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator struct {
	board *Board
}

func (g *Generator) GenerateAnyDifficulty() (*Board, Difficulty) {
	return g.Generate([]Difficulty{Easy, Medium, Hard})
}

func (g *Generator) Generate(rangeList []Difficulty) (*Board, Difficulty) {
	rand.Seed(time.Now().UnixNano())
	g.makeNewBoard()

	cells := make([]*Cell, 81)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cells[(row*9)+col] = &Cell{row: row, col: col}
		}
	}

	difficulty := -1
	targetDifficulty := rangeList[rand.Intn(len(rangeList))]
	board := g.board.Copy()
	fmt.Println("Generating:", targetDifficulty.Name(), "puzzle")

	for difficulty <= targetDifficulty.getRangeMin() {
		if len(cells) == 0 {
			difficulty = -1
			g.makeNewBoard()
			board = g.board.Copy()

			cells = make([]*Cell, 81)

			for row := 0; row < 9; row++ {
				for col := 0; col < 9; col++ {
					cells[(row*9)+col] = &Cell{row: row, col: col}
				}
			}
			continue
		}

		cellToRemove := cells[rand.Intn(len(cells))]
		cells = RemoveCellValue(cells, cellToRemove)
		cellCounterPart := getCellCounterPart(cellToRemove)
		cells = RemoveCellValue(cells, cellCounterPart)
		cellToRemoveOrigVal := board.Board[cellToRemove.row][cellToRemove.col]
		cellCounterPartOrigVal := board.Board[cellCounterPart.row][cellCounterPart.col]
		board.Board[cellToRemove.row][cellToRemove.col] = 0
		board.Board[cellCounterPart.row][cellCounterPart.col] = 0

		solver := &Solver{
			board: board.Copy(),
		}
		solver.Silent()
		solver.Solve()
		difficulty = solver.difficulty

		if difficulty == -1 || difficulty > targetDifficulty.getRangeMax() {
			board.Board[cellToRemove.row][cellToRemove.col] = cellToRemoveOrigVal
			board.Board[cellCounterPart.row][cellCounterPart.col] = cellCounterPartOrigVal
		}
	}
	return board, Difficulty(difficulty)
}

func (g *Generator) makeNewBoard() {
	g.board = &Board{}
	g.board.InitEmpty()
	g.fillCell(0)
	g.board.Solution = Copy2DArray(g.board.Board)
}

func RemoveCellValue(s []*Cell, val *Cell) []*Cell {
	index := -1
	for i := range s {
		if s[i].row == val.row && s[i].col == val.col {
			index = i
			break
		}
	}
	return RemoveCellIndex(s, index)
}

func RemoveCellIndex(s []*Cell, index int) []*Cell {
	if index == -1 {
		return s
	}
	if index >= len(s) {
		return s[:index]
	}
	return append(s[:index], s[index+1:]...)
}

func getCellCounterPart(c *Cell) *Cell {
	return &Cell{
		row: 8 - c.row,
		col: 8 - c.col,
	}
}

func (g *Generator) fillCell(cell int) bool {
	r := cell / 9
	c := cell % 9
	optsCpy := Copy3DArray(g.board.Options)

	complete := false
	for !complete {
		g.board.Options = Copy3DArray(optsCpy)
		opts := g.board.Options[r][c]
		if len(opts) == 0 {
			return false
		}
		chosenIndx := rand.Intn(len(opts))
		val := opts[chosenIndx]
		g.board.Options[r][c] = RemoveIndex(opts, chosenIndx)
		optsCpy[r][c] = RemoveValue(optsCpy[r][c], val)
		g.board.PlaceNumber(r, c, val)

		if cell == 80 {
			return true
		}

		complete = g.fillCell(cell + 1)
	}
	return true
}
