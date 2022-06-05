package main

import (
	"math/rand"
	"time"
)

type Generator struct {
	board *Board
}

func (g *Generator) Generate() {
	rand.Seed(time.Now().UnixNano())
	g.board.InitEmpty()
	g.fillCell(0)
}

func (g *Generator) fillCell(cell int) bool {
	r := cell / 9
	c := cell % 9
	optsCpy := Copy3Darray(g.board.options)

	complete := false
	for !complete {
		g.board.options = Copy3Darray(optsCpy)
		opts := g.board.options[r][c]
		if len(opts) == 0 {
			return false
		}
		chosenIndx := rand.Intn(len(opts))
		val := opts[chosenIndx]
		g.board.options[r][c] = RemoveIndex(opts, chosenIndx)
		optsCpy[r][c] = RemoveValue(optsCpy[r][c], val)
		g.board.PlaceNumber(r, c, val)

		if cell == 80 {
			return true
		}

		complete = g.fillCell(cell + 1)
	}
	return true
}
