package main

import "fmt"

type Board struct {
	board   [][]int
	options [][][]int
}

func (b *Board) InitEmpty() {
	b.board = make([][]int, 9)
	b.options = make([][][]int, 9)
	for i := range b.board {
		b.board[i] = make([]int, 9)
		b.options[i] = make([][]int, 9)
		for j := range b.options[i] {
			b.board[i][j] = 0
			b.options[i][j] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		}
	}
}

func (b *Board) Copy() *Board {
	return &Board{
		board:   Copy2Darray(b.board),
		options: Copy3Darray(b.options),
	}
}

func (b *Board) PlaceNumber(row int, column int, value int) [][][]int {
	b.board[row][column] = value
	b.options[row][column] = []int{}

	for i := 0; i < 9; i++ {
		if i != row {
			b.options[i][column] = RemoveValue(b.options[i][column], value)
		}
		if i != column {
			b.options[row][i] = RemoveValue(b.options[row][i], value)
		}
	}

	innerIndex := ((row % 3) * 3) + (column % 3)

	for i := 0; i < 9; i++ {
		if i == innerIndex {
			continue
		}
		adjRow := ((row / 3) * 3) + (i / 3)
		adjColumn := ((column / 3) * 3) + (i % 3)
		b.options[adjRow][adjColumn] = RemoveValue(b.options[adjRow][adjColumn], value)
	}

	return b.options
}

func (b *Board) GenerateOptions() {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b.board[r][c] != 0 {
				b.options[r][c] = []int{}
				continue
			} else {
				b.options[r][c] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			}

			for i := 0; i < 9; i++ {
				if i != r {
					b.options[r][c] = RemoveValue(b.options[r][c], b.board[i][c])
				}
				if i != c {
					b.options[r][c] = RemoveValue(b.options[r][c], b.board[r][i])
				}
			}
			innerIndex := ((r % 3) * 3) + (c % 3)

			for i := 0; i < 9; i++ {
				if i == innerIndex {
					continue
				}
				adjRow := ((r / 3) * 3) + (i / 3)
				adjColumn := ((c / 3) * 3) + (i % 3)
				b.options[r][c] = RemoveValue(b.options[r][c], b.board[adjRow][adjColumn])
			}
		}
	}
}

func (b *Board) IsSolved() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b.board[r][c] == 0 {
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
			if b.board[r][c] != 0 {
				fmt.Print(b.board[r][c])
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
