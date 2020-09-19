package main

import (
	"log"
	"os"
	"strconv"

	"github.com/df11/sudoku-solver/puzzle"
	"github.com/df11/sudoku-solver/request"
)

// @TODO:
// - put functions as puzzle methods
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
	if puzzle.Solve(&puzzleObject) {
		print("Puzzle solved (" + strconv.Itoa(puzzleObject.Iteration) + ")!\n")
	} else {
		print("Not this time (" + strconv.Itoa(puzzleObject.Iteration) + ")...\n")
	}
}
