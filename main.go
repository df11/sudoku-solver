package main

import (
	"log"
	"os"

	"github.com/df11/sudoku-solver/puzzle"
	"github.com/df11/sudoku-solver/request"
)

// @TODO:
// - put functions as puzzle methods
// - add a way to know if puzzle is solved
func main() {
	var difficulty = "easy"
	var puzzleLine string
	var puzzleSolution string

	if len(os.Args) == 3 {
		difficulty = "custom"
		puzzleLine = os.Args[1]
		puzzleSolution = os.Args[2]
		if len(puzzleLine) != 81 || len(puzzleSolution) != 81 {
			log.Fatal("Incorrect puzzle length")
		}
	}
	if len(os.Args) == 2 {
		difficulty = os.Args[1]
		puzzleLine, puzzleSolution = request.CallAPI(difficulty)
	}

	puzzleObject := puzzle.Initialize(difficulty, puzzleLine, puzzleSolution)
	puzzle.Solve(&puzzleObject)
}
