package main

import (
	"fmt"
	"strconv"
)

type Solver struct {
	board      *Board
	difficulty int
	cltUsed    bool
}

func (s *Solver) Solve() {
	s.difficulty = 0
	s.cltUsed = false

	s.board.GenerateOptions()

	for !s.board.IsSolved() {
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

		s.difficulty = -1
		break
	}
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
				s.board.PlaceNumber(r, c, s.board.options[r][c][0])
				logMove("Single Candidate", r, c)
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
				s.board.PlaceNumber(r, c, optsRowCpy[0])
				logMove("Single Position Row", r, c)
				return true
			}
			if len(optsColCpy) == 1 {
				s.board.PlaceNumber(r, c, optsColCpy[0])
				logMove("Single Position Col", r, c)
				return true
			}
			if len(optsBoxCpy) == 1 {
				s.board.PlaceNumber(r, c, optsBoxCpy[0])
				logMove("Single Position Box", r, c)
				return true
			}
		}
	}
	return false
}

func (s *Solver) candidateLines() bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {

			for sr := 0; sr < 3; sr++ {
				opts := []int{}
				for i := 0; i < 3; i++ {
					opts = AppendIfMissing(opts, s.board.options[(r*3)+sr][(c*3)+i]...)
				}
				for i := 0; i < 9; i++ {
					if i/3 == sr {
						continue
					}
					opts = RemoveValues(opts, s.board.options[(r*3)+(i/3)][(c*3)+(i%3)]...)
				}
				if len(opts) > 0 {
					for i := range opts {
						removed := false
						for j := 0; j < 9; j++ {
							if j/3 == c {
								continue
							}

							indx := FindValueIndex(s.board.options[(r*3)+sr][j], opts[i])
							if indx != -1 {
								s.board.options[(r*3)+sr][j] = RemoveIndex(s.board.options[(r*3)+sr][j], indx)
								removed = true
							}
						}

						if removed {
							logMove("Candidate Lines Row for "+strconv.Itoa(opts[i]), (r*3)+sr, c)
							return true
						}
					}
				}
			}

			for sc := 0; sc < 3; sc++ {
				opts := []int{}
				for i := 0; i < 3; i++ {
					opts = AppendIfMissing(opts, s.board.options[(r*3)+i][(c*3)+sc]...)
				}
				for i := 0; i < 9; i++ {
					if i%3 == sc {
						continue
					}
					opts = RemoveValues(opts, s.board.options[(r*3)+(i/3)][(c*3)+(i%3)]...)
				}
				if len(opts) > 0 {
					for i := range opts {
						removed := false
						for j := 0; j < 9; j++ {
							if j/3 == r {
								continue
							}

							indx := FindValueIndex(s.board.options[j][(c*3)+sc], opts[i])
							if indx != -1 {
								s.board.options[j][(c*3)+sc] = RemoveIndex(s.board.options[j][(c*3)+sc], indx)
								removed = true
							}
						}

						if removed {
							logMove("Candidate Lines Col for "+strconv.Itoa(opts[i]), r, (c*3)+sc)
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func logMove(msg string, r int, c int) {
	fmt.Print(msg)
	fmt.Print(" @ (")
	fmt.Print(c + 1)
	fmt.Print(",")
	fmt.Print(r + 1)
	fmt.Println(")")
}
