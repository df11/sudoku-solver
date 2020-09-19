package puzzle

import (
	"strconv"
	"strings"
)

func Initialize(difficulty string, puzzleLine string, puzzleSolution string) Puzzle {
	var puzzle Puzzle

	puzzle.Difficulty = difficulty
	puzzle.PuzzleString = puzzleLine
	puzzle.PuzzleSolution = puzzleSolution

	splitSrc := strings.Split(puzzleLine, "")
	for lineNumber := 0; lineNumber < 9; lineNumber++ {
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			value, _ := strconv.Atoi(splitSrc[lineNumber*9+columnNumber])
			puzzle.Cells[lineNumber][columnNumber].Value = value

			for i := 0; i < 9; i++ {
				if value == 0 {
					puzzle.Cells[lineNumber][columnNumber].Possibilities[i] = true
				} else {
					puzzle.Cells[lineNumber][columnNumber].Possibilities[i] = false
				}
			}
		}
	}
	for lineNumber := 0; lineNumber < 9; lineNumber++ {
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			value := puzzle.Cells[lineNumber][columnNumber].Value
			if value > 0 {
				clearNumber(lineNumber, columnNumber, value, &puzzle)
			}
		}
	}
	return puzzle
}

func Solve(puzzle *Puzzle) bool {
	for try := 0; try < 10000; try++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				checkSquare(i*3, j*3, puzzle)
				checkLine(i*3+j, puzzle)
				checkColumn(i*3+j, puzzle)
			}
		}
	}
	return false
}

func checkPossibility(possibilities [9]bool) int {
	possibility := 0
	for i := 0; i < 9; i++ {
		if possibilities[i] == true {
			if possibility > 0 {
				return 0
			}
			possibility = i + 1
		}
	}
	return possibility
}

func clearNumber(lineNumber int, columnNumber int, value int, puzzle *Puzzle) {
	for i := 0; i < 9; i++ {
		puzzle.Cells[lineNumber][i].Possibilities[value-1] = false
		puzzle.Cells[i][columnNumber].Possibilities[value-1] = false
		puzzle.Cells[lineNumber][columnNumber].Possibilities[i] = false
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			puzzle.Cells[lineNumber/3*3+i][columnNumber/3*3+j].Possibilities[value-1] = false
		}
	}
}

func checkSquare(line int, column int, puzzle *Puzzle) {
	for value := 1; value <= 9; value++ {
		occurences := 0
		occLine := 0
		occColumn := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if puzzle.Cells[line+i][column+j].Possibilities[value-1] == true {
					occurences++
					occLine = line + i
					occColumn = column + j
				}
			}
		}
		if occurences == 1 {
			puzzle.Cells[occLine][occColumn].Value = value
			clearNumber(occLine, occColumn, value, puzzle)
			PrintPuzzle(*puzzle)
		}
	}
}

func checkLine(line int, puzzle *Puzzle) {
	for value := 1; value <= 9; value++ {
		occurences := 0
		occColumn := 0
		for i := 0; i < 9; i++ {
			if puzzle.Cells[line][i].Possibilities[value-1] == true {
				occurences++
				occColumn = i
			}
		}
		if occurences == 1 {
			puzzle.Cells[line][occColumn].Value = value
			clearNumber(line, occColumn, value, puzzle)
			PrintPuzzle(*puzzle)
		}
	}
}

func checkColumn(column int, puzzle *Puzzle) {
	for value := 1; value <= 9; value++ {
		occurences := 0
		occLine := 0
		for i := 0; i < 9; i++ {
			if puzzle.Cells[i][column].Possibilities[value-1] == true {
				occurences++
				occLine = i
			}
		}
		if occurences == 1 {
			puzzle.Cells[occLine][column].Value = value
			clearNumber(occLine, column, value, puzzle)
			PrintPuzzle(*puzzle)
		}
	}
}
