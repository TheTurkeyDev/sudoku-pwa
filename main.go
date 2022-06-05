package main

import (
	"fmt"
)

func main() {
	// generator := &Generator{}
	// generator.Generate()
	// board := copy2Darray(generator.board)

	// printBoard(board)
	// fmt.Println()

	// board[2][2] = 0

	// printBoard(board)
	// fmt.Println()

	// board := boardFromString("007000100103905700698210530030600020500070001070001040012094678004106209005000400") // Beginner
	// board := boardFromString("100500700000740000090002010003050480900000005085090600060200030000063000004005002") // Easy Single Position, Single Candidate
	board := boardFromString("906008070500030000704950002491000508000000000603000724100089203000040007060500109") // Medium  Single Position, Single Candidate, Candidate Lines

	board.printBoard()
	solver := &Solver{
		board: board,
	}
	solver.Solve()

	fmt.Print("Board Difficulty: ")
	fmt.Println(solver.difficulty)
	board.printBoard()
}

func boardFromString(str string) *Board {
	board := &Board{}
	board.InitEmpty()

	for i := 0; i < len(str); i++ {
		r := i / 9
		c := i % 9
		board.board[r][c] = int(str[i] - '0')
	}

	return board
}

func Copy2Darray(src [][]int) [][]int {
	cpy := make([][]int, len(src))
	for i := range src {
		cpy[i] = make([]int, len(src[i]))
		copy(cpy[i], src[i])
	}

	return cpy
}

func Copy3Darray(src [][][]int) [][][]int {
	cpy := make([][][]int, len(src))
	for i := range src {
		cpy[i] = make([][]int, len(src[i]))
		for j := range src[i] {
			cpy[i][j] = make([]int, len(src[i][j]))
			copy(cpy[i][j], src[i][j])
		}
	}

	return cpy
}

func RemoveValues(s []int, val ...int) []int {
	cpy := make([]int, len(s))
	copy(cpy, s)
	for v := range val {
		cpy = RemoveValue(cpy, val[v])
	}
	return cpy
}

func RemoveValue(s []int, val int) []int {
	index := -1
	for i := range s {
		if s[i] == val {
			index = i
			break
		}
	}
	return RemoveIndex(s, index)
}

func FindValueIndex(s []int, val int) int {
	index := -1
	for i := range s {
		if s[i] == val {
			index = i
			break
		}
	}
	return index
}

func RemoveIndex(s []int, index int) []int {
	if index == -1 {
		return s
	}
	if index >= len(s) {
		return s[:index]
	}
	return append(s[:index], s[index+1:]...)
}

func AppendIfMissing(s []int, vals ...int) []int {
	cpy := make([]int, len(s))
	copy(cpy, s)
	for i := range vals {
		indx := FindValueIndex(s, vals[i])
		if indx == -1 {
			cpy = append(cpy, vals[i])
		}
	}
	return cpy
}