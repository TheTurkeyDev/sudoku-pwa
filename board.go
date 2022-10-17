package main

import "fmt"

type Board struct {
	Solution [][]int   `json:"solution"`
	Board    [][]int   `json:"board"`
	Options  [][][]int `json:"options"`
}

func (b *Board) InitEmpty() {
	b.Board = make([][]int, 9)
	b.Solution = make([][]int, 9)
	b.Options = make([][][]int, 9)
	for i := range b.Board {
		b.Board[i] = make([]int, 9)
		b.Solution[i] = make([]int, 9)
		b.Options[i] = make([][]int, 9)
		for j := range b.Options[i] {
			b.Board[i][j] = 0
			b.Solution[i][j] = 0
			b.Options[i][j] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		}
	}
}

func (b *Board) Copy() *Board {
	return &Board{
		Board:   Copy2DArray(b.Board),
		Options: Copy3DArray(b.Options),
	}
}

func (b *Board) PlaceNumber(row int, column int, value int) [][][]int {
	b.Board[row][column] = value
	b.Options[row][column] = []int{}

	for i := 0; i < 9; i++ {
		if i != row {
			b.Options[i][column] = RemoveValue(b.Options[i][column], value)
		}
		if i != column {
			b.Options[row][i] = RemoveValue(b.Options[row][i], value)
		}
	}

	innerIndex := ((row % 3) * 3) + (column % 3)

	for i := 0; i < 9; i++ {
		if i == innerIndex {
			continue
		}
		adjRow := ((row / 3) * 3) + (i / 3)
		adjColumn := ((column / 3) * 3) + (i % 3)
		b.Options[adjRow][adjColumn] = RemoveValue(b.Options[adjRow][adjColumn], value)
	}

	return b.Options
}

func (b *Board) GenerateOptions() {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b.Board[r][c] != 0 {
				b.Options[r][c] = []int{}
				continue
			} else {
				b.Options[r][c] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}

			for i := 0; i < 9; i++ {
				if i != r {
					b.Options[r][c] = RemoveValue(b.Options[r][c], b.Board[i][c])
				}
				if i != c {
					b.Options[r][c] = RemoveValue(b.Options[r][c], b.Board[r][i])
				}
			}
			innerIndex := ((r % 3) * 3) + (c % 3)

			for i := 0; i < 9; i++ {
				if i == innerIndex {
					continue
				}
				adjRow := ((r / 3) * 3) + (i / 3)
				adjColumn := ((c / 3) * 3) + (i % 3)
				b.Options[r][c] = RemoveValue(b.Options[r][c], b.Board[adjRow][adjColumn])
			}
		}
	}
}

func (b *Board) IsSolved() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b.Board[r][c] == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) printBoard() {
	fmt.Println("+-------+-------+-------+")
	for r := 0; r < 9; r++ {
		fmt.Print("| ")
		for c := 0; c < 9; c++ {
			if b.Board[r][c] != 0 {
				fmt.Print(b.Board[r][c])
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
			if (c+1)%3 == 0 {
				fmt.Print("| ")
			}
		}
		fmt.Println()

		if (r+1)%3 == 0 {
			fmt.Println("+-------+-------+-------+")
		}
	}
}
