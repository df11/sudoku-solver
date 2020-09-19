package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const apiPuzzleIdx = 0
const apiSolutionIdx = 1

type apiResponse struct {
	Answer  string        `json:"answer"`
	Message string        `json:"message"`
	Desc    []interface{} `json:"desc"`
}

type puzzle struct {
	PuzzleString   string
	PuzzleSolution string
	Difficulty     string
	Cells          [9][9]puzzleCell
}

type puzzleCell struct {
	Value         int
	Possibilities [9]bool
}

func printSudoku(src string) {
	var lineOutput string
	var lineNumber int

	splitSrc := strings.Split(src, "")

	for lineNumber = 0; lineNumber < 9; lineNumber++ {
		if lineNumber > 0 && lineNumber%3 == 0 {
			print("=========================================\n")
		} else {
			print("-----------------------------------------\n")
		}
		lineOutput = "|"
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			cellValue := splitSrc[lineNumber*9+columnNumber]
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
	print("-----------------------------------------\n")
}

func printPuzzleLine(puzzle puzzle) {
	for lineNumber := 0; lineNumber < 9; lineNumber++ {
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			print(puzzle.Cells[lineNumber][columnNumber].Value)
		}
	}
	print("\n")
}

func printPuzzle(puzzle puzzle) {
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

// func printPossibilities(possibilities [9][9][9]bool) {
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			print(fmt.Sprintf("%d, %d: ", i, j))
// 			for value := 0; value < 9; value++ {
// 				if possibilities[i][j][value] == true {
// 					print(fmt.Sprintf("%d", value+1))
// 				}
// 			}
// 			print("\n")
// 		}
// 	}
// }

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

func clearNumber(lineNumber int, columnNumber int, value int, puzzle *puzzle) {
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
func initPuzzle(difficulty string, puzzleLine string, puzzleSolution string) puzzle {
	var puzzle puzzle

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

func checkSquare(line int, column int, puzzle *puzzle) {
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
			printPuzzle(*puzzle)
		}
	}
}

func checkLine(line int, puzzle *puzzle) {
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
			printPuzzle(*puzzle)
		}
	}
}

func checkColumn(column int, puzzle *puzzle) {
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
			printPuzzle(*puzzle)
		}
	}
}

func solve(puzzle *puzzle) bool {
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

// @TODO:
// - put functions as puzzle methods
// - add a way to know if puzzle is solved
func main() {
	var difficulty = "easy"
	var puzzleLine string
	var puzzleSolution string
	var puzzle puzzle

	if len(os.Args) == 3 {
		difficulty = "custom"
		puzzleLine = os.Args[1]
		puzzleSolution = os.Args[2]
	}
	if len(os.Args) == 2 {
		difficulty = os.Args[1]

		res, err := http.Get("https://sudoku.com/api/getLevel/" + difficulty)
		if err != nil {
			log.Fatal(err)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode != 200 {
			log.Fatal(string(body))
		}

		resJSON := apiResponse{}
		err = json.Unmarshal(body, &resJSON)
		if err != nil {
			log.Fatal(err)
		}

		puzzleLine = fmt.Sprintf("%v", resJSON.Desc[apiPuzzleIdx])
		puzzleSolution = fmt.Sprintf("%v", resJSON.Desc[apiSolutionIdx])
	}

	puzzle = initPuzzle(difficulty, puzzleLine, puzzleSolution)
	solve(&puzzle)
}
