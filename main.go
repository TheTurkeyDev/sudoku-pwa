package main

import "fmt"

func main() {
	generator := &Generator{}
	board, difficulty := generator.Generate()
	generator.board.printBoard()
	fmt.Print("Board Difficulty: ")
	fmt.Println(difficulty)
	board.printBoard()

	// board := boardFromString("007000100103905700698210530030600020500070001070001040012094678004106209005000400") // Beginner
	// board := boardFromString("100500700000740000090002010003050480900000005085090600060200030000063000004005002") // Easy Single Position, Single Candidate
	// board := boardFromString("906008070500030000704950002491000508000000000603000724100089203000040007060500109") // Medium  Single Position, Single Candidate, Candidate Lines
	// board := boardFromString("800204600007000001000050830900500000148060275000001009082070000700000500003809006") // Tricky Single Position, Single Candidate, Candidate Lines, Multiple Lines
	// board := boardFromString("000780420020095008007000000000360005102908307300071000000000600200630070078019000") // Tricky Single Position, Single Candidate, Candidate Lines, Double Pairs, Naked Pairs, Hidden Pairs
	// Can't solve yet. Naked triple... board := boardFromString("060500000100004609700010800040000200000842000009000010005080007902600005000005080") // Fiendish Single Position, Single Candidate, Candidate Lines, Multiple Lines, Naked Pairs, Naked Triples, Hidden Pairs

	// fmt.Println("Test")
	// board := boardFromString("200409003070000805000008000000003080120000094060800000000700000902000030400906008")

	// solver := &Solver{
	// 	board: board.Copy(),
	// }
	// solver.Silent()
	// solver.Solve()
	// fmt.Println(solver.difficulty)
}

// func boardFromString(str string) *Board {
// 	board := &Board{}
// 	board.InitEmpty()

// 	for i := 0; i < len(str); i++ {
// 		r := i / 9
// 		c := i % 9
// 		board.board[r][c] = int(str[i] - '0')
// 	}

// 	return board
// }

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
