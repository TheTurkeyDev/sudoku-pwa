package main

type Difficulty int64

const (
	Beginner Difficulty = iota
	Easy
	Medium
	Hard
	Extreme // Not used yet
)

func (d Difficulty) getRange() []int {
	switch d {
	case Beginner:
		return []int{3500, 4500}
	case Easy:
		return []int{4300, 6000}
	case Medium:
		return []int{5700, 7500}
	case Hard:
		return []int{7200, 10000}
	case Extreme:
		return []int{9000, 20000}
	}

	return []int{-1, -1}
}

func (d Difficulty) Name() string {
	switch d {
	case Beginner:
		return "Beginner"
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	case Extreme:
		return "Extreme"
	}
	return "Error"
}

func (d Difficulty) getRangeMin() int {
	return d.getRange()[0]
}

func (d Difficulty) getRangeMax() int {
	return d.getRange()[1]
}
