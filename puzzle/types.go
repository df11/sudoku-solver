package puzzle

type Puzzle struct {
	PuzzleString   string
	PuzzleSolution string
	Difficulty     string
	Cells          [9][9]PuzzleCell
}

type PuzzleCell struct {
	Value         int
	Possibilities [9]bool
}
