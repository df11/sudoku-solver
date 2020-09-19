package puzzle

import (
	"strconv"
	"time"
)

func PrintPuzzle(puzzle Puzzle) {
	var lineOutput string
	var lineNumber int

	print("\033[H\033[2J")
	print("Difficuty: " + puzzle.Difficulty + "\n")
	print(puzzle.PuzzleString + "\n")
	print(puzzle.PuzzleSolution + "\n")
	for lineNumber = 0; lineNumber < 9; lineNumber++ {
		if lineNumber%3 == 0 {
			print("=========================================\n")
		} else {
			print("-----------------------------------------\n")
		}
		lineOutput = "|"
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			cellValue := strconv.Itoa(puzzle.Cells[lineNumber][columnNumber].Value)
			cellContent := " "
			if cellValue != "0" {
				cellContent = cellValue
			}
			if columnNumber%3 == 0 {
				lineOutput += "|"
			}
			lineOutput += " " + cellContent + " |"
		}
		lineOutput += "|\n"
		print(lineOutput)
	}
	print("=========================================\n")
	time.Sleep(100 * time.Millisecond)
}
