package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct {
	board      *Board
	difficulty int
	cltUsed    bool
	dpmlUsed   bool
	npUsed     bool
	silent     bool
}

func (s *Solver) Solve() {
	s.difficulty = 0
	s.cltUsed = false

	s.board.GenerateOptions()

	for !s.board.IsSolved() {
		// fmt.Println(s.board.options)
		if s.noOptions() {
			fmt.Println("No Options!")
			s.difficulty = -1
			break
		}
		if s.singleCandidate() {
			s.difficulty = s.difficulty + 100
			continue
		}
		if s.singlePosition() {
			s.difficulty = s.difficulty + 100
			continue
		}
		if s.candidateLines() {
			if s.cltUsed {
				s.difficulty = s.difficulty + 200
			} else {
				s.difficulty = s.difficulty + 350
				s.cltUsed = true
			}
			continue
		}

		if s.doublePairOrMultipleLines() {
			if s.dpmlUsed {
				s.difficulty = s.difficulty + 350
			} else {
				s.difficulty = s.difficulty + 600
				s.dpmlUsed = true
			}
			continue
		}

		if s.nakedCandidates(2) {
			if s.npUsed {
				s.difficulty = s.difficulty + 500
			} else {
				s.difficulty = s.difficulty + 750
				s.npUsed = true
			}
			continue
		}
		if s.hiddenCandidates(2) {
			if s.npUsed {
				s.difficulty = s.difficulty + 1200
			} else {
				s.difficulty = s.difficulty + 1500
				s.npUsed = true
			}
			continue
		}
		if s.nakedCandidates(3) {
			if s.npUsed {
				s.difficulty = s.difficulty + 1400
			} else {
				s.difficulty = s.difficulty + 2000
				s.npUsed = true
			}
			continue
		}
		// if s.hiddenCandidates(3) {
		// 	if s.npUsed {
		// 		s.difficulty = s.difficulty + 1600
		// 	} else {
		// 		s.difficulty = s.difficulty + 2400
		// 		s.npUsed = true
		// 	}
		// 	continue
		// }

		s.difficulty = -1
		break
	}
}

func (s *Solver) Silent() {
	s.silent = true
}

func (s *Solver) noOptions() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if s.board.board[r][c] == 0 && len(s.board.options[r][c]) == 0 {
				return true
			}
		}
	}
	return false
}

func (s *Solver) singleCandidate() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if s.board.board[r][c] != 0 {
				continue
			}
			if len(s.board.options[r][c]) == 1 {
				s.logMove("Single Candidate", strconv.Itoa(s.board.options[r][c][0]), r, c)
				s.board.PlaceNumber(r, c, s.board.options[r][c][0])
				return true
			}
		}
	}
	return false
}

func (s *Solver) singlePosition() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if s.board.board[r][c] != 0 {
				continue
			}

			opts := s.board.options[r][c]

			optsRowCpy := make([]int, len(opts))
			copy(optsRowCpy, opts)
			optsColCpy := make([]int, len(opts))
			copy(optsColCpy, opts)
			optsBoxCpy := make([]int, len(opts))
			copy(optsBoxCpy, opts)

			for i := 0; i < 9; i++ {
				if i != r {
					optsRowCpy = RemoveValues(optsRowCpy, s.board.options[i][c]...)
				}
				if i != c {
					optsColCpy = RemoveValues(optsColCpy, s.board.options[r][i]...)
				}
			}
			innerIndex := ((r % 3) * 3) + (c % 3)

			for i := 0; i < 9; i++ {
				if i == innerIndex {
					continue
				}
				adjRow := ((r / 3) * 3) + (i / 3)
				adjColumn := ((c / 3) * 3) + (i % 3)
				optsBoxCpy = RemoveValues(optsBoxCpy, s.board.options[adjRow][adjColumn]...)
			}

			if len(optsRowCpy) == 1 {
				s.logMove("Single Position Row", strconv.Itoa(optsRowCpy[0]), r, c)
				s.board.PlaceNumber(r, c, optsRowCpy[0])
				return true
			}
			if len(optsColCpy) == 1 {
				s.logMove("Single Position Col", strconv.Itoa(optsColCpy[0]), r, c)
				s.board.PlaceNumber(r, c, optsColCpy[0])
				return true
			}
			if len(optsBoxCpy) == 1 {
				s.logMove("Single Position Box", strconv.Itoa(optsBoxCpy[0]), r, c)
				s.board.PlaceNumber(r, c, optsBoxCpy[0])
				return true
			}
		}
	}
	return false
}

func (s *Solver) candidateLines() bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {

			// Find Lines in rows
			// sr is the sub row of the current box
			for sr := 0; sr < 3; sr++ {
				opts := []int{}
				// Add all the values in the row in the current box to a list of options to check for a candidate
				for i := 0; i < 3; i++ {
					opts = AppendIfMissing(opts, s.board.options[(r*3)+sr][(c*3)+i]...)
				}
				// Remove all values in the box from the options as they can't be candidates
				for i := 0; i < 9; i++ {
					// Ignore the sub row we are currently on
					if i/3 == sr {
						continue
					}
					opts = RemoveValues(opts, s.board.options[(r*3)+(i/3)][(c*3)+(i%3)]...)
				}

				// If the options list still has values, we have a candidate!
				if len(opts) > 0 {
					for i := range opts {
						removed := false
						// Go through all cells in this row and attempt to remove the current candidate from it's options
						for j := 0; j < 9; j++ {
							// Ignore the cells in our current box that the candidates were sourced from
							if j/3 == c {
								continue
							}

							indx := FindValueIndex(s.board.options[(r*3)+sr][j], opts[i])
							// If this cell contains the candidate, remove it and mark that we have removed an option
							if indx != -1 {
								s.board.options[(r*3)+sr][j] = RemoveIndex(s.board.options[(r*3)+sr][j], indx)
								removed = true
							}
						}

						// If we have removed an option, log it and return true
						if removed {
							s.logMove("Candidate Lines Row", strconv.Itoa(opts[i]), (r*3)+sr, c)
							return true
						}
					}
				}
			}

			// Find Lines in column
			// sc is the sub column of the current box
			for sc := 0; sc < 3; sc++ {
				opts := []int{}
				// Add all the values in the column in the current box to a list of options to check for a candidate
				for i := 0; i < 3; i++ {
					opts = AppendIfMissing(opts, s.board.options[(r*3)+i][(c*3)+sc]...)
				}
				// Remove all values in the box from the options as they can't be candidates
				for i := 0; i < 9; i++ {
					// Ignore the sub coulmn we are currently on
					if i%3 == sc {
						continue
					}
					opts = RemoveValues(opts, s.board.options[(r*3)+(i/3)][(c*3)+(i%3)]...)
				}

				// If the options list still has values, we have a candidate!
				if len(opts) > 0 {
					for i := range opts {
						removed := false
						// Go through all cells in this column and attempt to remove the current candidate from it's options
						for j := 0; j < 9; j++ {
							// Ignore the cells in our current box that the candidates were sourced from
							if j/3 == r {
								continue
							}

							indx := FindValueIndex(s.board.options[j][(c*3)+sc], opts[i])
							// If this cell contains the candidate, remove it and mark that we have removed an option
							if indx != -1 {
								s.board.options[j][(c*3)+sc] = RemoveIndex(s.board.options[j][(c*3)+sc], indx)
								removed = true
							}
						}

						// If we have removed an option, log it and return true
						if removed {
							s.logMove("Candidate Lines Col", strconv.Itoa(opts[i]), r, (c*3)+sc)
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func (s *Solver) doublePairOrMultipleLines() bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {

			// Find in rows
			// sr is the sub row of the current box
			for sr := 0; sr < 3; sr++ {
				opts := []int{}
				// Add all the values in the row in the current box to a list of options to check for a candidate
				for i := 0; i < 3; i++ {
					opts = AppendIfMissing(opts, s.board.options[(r*3)+sr][(c*3)+i]...)
				}
				// Remove all values in the rest of the row from the options as they can't be candidates
				for i := 0; i < 9; i++ {
					// Ignore the columns in the box we are currently in
					if i/3 == c {
						continue
					}
					opts = RemoveValues(opts, s.board.options[(r*3)+sr][i]...)
				}

				// If the options list still has values, we have a candidate!
				if len(opts) > 0 {
					for i := range opts {
						removed := false
						// Go through all cells in this box and attempt to remove the current candidate from it's options
						for j := 0; j < 9; j++ {
							// Ignore the cells in our current row that the candidates were sourced from
							if j/3 == sr {
								continue
							}

							indx := FindValueIndex(s.board.options[(r*3)+(j/3)][(c*3)+(j%3)], opts[i])
							// If this cell contains the candidate, remove it and mark that we have removed an option
							if indx != -1 {
								s.board.options[(r*3)+(j/3)][(c*3)+(j%3)] = RemoveIndex(s.board.options[(r*3)+(j/3)][(c*3)+(j%3)], indx)
								removed = true
							}
						}

						// If we have removed an option, log it and return true
						if removed {
							s.logMove("Double Pair Or Multiple Lines Row", strconv.Itoa(opts[i]), (r*3)+sr, c)
							return true
						}
					}
				}
			}

			// Find in columns
			// sc is the sub column of the current box
			for sc := 0; sc < 3; sc++ {
				opts := []int{}
				// Add all the values in the column in the current box to a list of options to check for a candidate
				for i := 0; i < 3; i++ {
					opts = AppendIfMissing(opts, s.board.options[(r*3)+i][(c*3)+sc]...)
				}
				// Remove all values in the rest of the column from the options as they can't be candidates
				for i := 0; i < 9; i++ {
					// Ignore the sub coulmn we are currently on
					if i/3 == r {
						continue
					}
					opts = RemoveValues(opts, s.board.options[i][(c*3)+sc]...)
				}

				// If the options list still has values, we have a candidate!
				if len(opts) > 0 {
					for i := range opts {
						removed := false
						// Go through all cells in this box and attempt to remove the current candidate from it's options
						for j := 0; j < 9; j++ {
							// Ignore the cells in our current column that the candidates were sourced from
							if j%3 == sc {
								continue
							}

							indx := FindValueIndex(s.board.options[(r*3)+(j/3)][(c*3)+(j%3)], opts[i])
							// If this cell contains the candidate, remove it and mark that we have removed an option
							if indx != -1 {
								s.board.options[(r*3)+(j/3)][(c*3)+(j%3)] = RemoveIndex(s.board.options[(r*3)+(j/3)][(c*3)+(j%3)], indx)
								removed = true
							}
						}

						// If we have removed an option, log it and return true
						if removed {
							s.logMove("Double Pair Or Multiple Lines Col", strconv.Itoa(opts[i]), r, (c*3)+sc)
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func (s *Solver) nakedCandidates(length int) bool {
	for r := 0; r < 9; r++ {
		if s.checkNakedCandidateCells(length, CellsForRow(r), "Row") {
			return true
		}
	}

	for c := 0; c < 9; c++ {
		if s.checkNakedCandidateCells(length, CellsForCol(c), "Col") {
			return true
		}
	}

	return false
}

func (s *Solver) checkNakedCandidateCells(length int, cells []*Cell, dir string) bool {
	occur := make(map[string][]*Cell)
	for i := range cells {
		cell := cells[i]
		if s.board.board[cell.row][cell.col] != 0 {
			continue
		}
		// Convert options list ex: [1,2,3,4] to a string ex:"1,2,3,4" for use as a key

		s, _ := json.Marshal(s.board.options[cell.row][cell.col])
		opts := strings.ReplaceAll(strings.Trim(string(s), "[]"), ",", "")
		// Increment the count of this key
		occur[opts] = append(occur[opts], &Cell{row: cell.row, col: cell.col})
	}

	// Loop over all option occurences and check that a) are the length we are looking for, and b) occur the same amount of times as the length
	for k, v := range occur {
		if len(k) == length && len(v) == length {
			removed := false
			candidates := strings.Split(k, "")
			for i := range cells {
				cell := cells[i]
				if s.board.board[cell.row][cell.col] != 0 {
					continue
				}
				// Ignore the cells with this key
				if strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s.board.options[cell.row][cell.col])), ""), "[]") == k {
					continue
				}

				// Attempt to remove each of the candidates from the rest of the cells in the row or column
				for j := range candidates {
					indx := FindValueIndex(s.board.options[cell.row][cell.col], int(candidates[j][0]-'0'))
					// If this cell contains the candidate, remove it and mark that we have removed an option
					if indx != -1 {
						s.board.options[cell.row][cell.col] = RemoveIndex(s.board.options[cell.row][cell.col], indx)
						removed = true
					}
				}
			}

			// If a candidate has been removed yay! log and return true
			if removed {
				s.logMove("Naked Candidates "+dir+" len:"+strconv.Itoa(length), k, cells[0].row, cells[0].col)
				return true
			}
		}
	}
	return false
}

func (s *Solver) hiddenCandidates(length int) bool {
	for r := 0; r < 9; r++ {
		if s.checkHiddenCandidatesCells(length, CellsForRow(r), "Row") {
			return true
		}
	}

	for c := 0; c < 9; c++ {
		if s.checkHiddenCandidatesCells(length, CellsForCol(c), "Col") {
			return true
		}
	}

	for b := 0; b < 9; b++ {
		if s.checkHiddenCandidatesCells(length, CellsForBox(b/3, b%3), "Box") {
			return true
		}
	}

	return false
}

func (s *Solver) checkHiddenCandidatesCells(length int, cells []*Cell, dir string) bool {
	occur := make(map[string]int)
	for i := range cells {
		cell := cells[i]

		for j := range s.board.options[cell.row][cell.col] {
			val := strconv.Itoa(s.board.options[cell.row][cell.col][j])
			occur[val] += 1
		}
	}

	candidates := make([]int, 0)
	for k, v := range occur {
		if v == length {
			candidates = append(candidates, int(k[0]-'0'))
		}
	}

	sort.Ints(candidates)

	combos := make([][]int, 0)
	for i := range candidates {
		for j := i + 1; j < len(candidates); j++ {
			combos = append(combos, []int{candidates[i], candidates[j]})
		}
	}

	for c := range combos {

		// Check all cells and count how many have this combo
		foundCombo := 0
		for i := range cells {
			cell := cells[i]

			// Does cell have all vals in combo
			valid := true
			for v := range combos[c] {
				if FindValueIndex(s.board.options[cell.row][cell.col], combos[c][v]) == -1 {
					valid = false
					break
				}
			}

			if !valid {
				continue
			}

			foundCombo++
		}

		// If 2 not found, this isn't a valid hidden pair
		if foundCombo != 2 {
			continue
		}

		removed := false
		for i := range cells {
			cell := cells[i]

			if FindValueIndex(s.board.options[cell.row][cell.col], combos[c][0]) != -1 && len(s.board.options[cell.row][cell.col]) != len(combos[c]) {
				s.board.options[cell.row][cell.col] = make([]int, len(combos[c]))
				copy(s.board.options[cell.row][cell.col], combos[c])
				removed = true
			}
		}

		// If a candidate has been removed yay! log and return true
		if removed {
			s.logMove("Hidden Candidates "+dir+" len:"+strconv.Itoa(length), "("+strconv.Itoa(combos[c][0])+", "+strconv.Itoa(combos[c][1])+")", cells[0].row, cells[0].col)
			return true
		}
	}

	return false
}

func (s *Solver) logMove(msg string, val string, r int, c int) {
	if !s.silent {
		fmt.Println(msg, "for", val, "@ (", c+1, ",", r+1, ")")
	}
}
