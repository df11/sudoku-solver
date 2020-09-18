package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func printPuzzleLine(puzzle [9][9]int) {
	for lineNumber := 0; lineNumber < 9; lineNumber++ {
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			print(puzzle[lineNumber][columnNumber])
		}
	}
	print("\n")
}

func printPuzzle(puzzle [9][9]int) {
	var lineOutput string
	var lineNumber int

	print("\033[H\033[2J")
	for lineNumber = 0; lineNumber < 9; lineNumber++ {
		if lineNumber > 0 && lineNumber%3 == 0 {
			print("=========================================\n")
		} else {
			print("-----------------------------------------\n")
		}
		lineOutput = "|"
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			cellValue := strconv.Itoa(puzzle[lineNumber][columnNumber])
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
	time.Sleep(100 * time.Millisecond)
}

func printPossibilities(possibilities [9][9][9]bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			print(fmt.Sprintf("%d, %d: ", i, j))
			for value := 0; value < 9; value++ {
				if possibilities[i][j][value] == true {
					print(fmt.Sprintf("%d", value+1))
				}
			}
			print("\n")
		}
	}
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

func clearNumber(lineNumber int, columnNumber int, value int, possibilities *[9][9][9]bool) {
	for i := 0; i < 9; i++ {
		possibilities[lineNumber][i][value-1] = false
		possibilities[i][columnNumber][value-1] = false
		possibilities[lineNumber][columnNumber][i] = false
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			possibilities[lineNumber/3*3+i][columnNumber/3*3+j][value-1] = false
			// print(lineNumber/3*3 + i)
		}
	}
}

func initialize(src string) ([9][9]int, [9][9][9]bool) {
	splitSrc := strings.Split(src, "")
	var puzzle [9][9]int
	var possibilities [9][9][9]bool

	for lineNumber := 0; lineNumber < 9; lineNumber++ {
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			value, _ := strconv.Atoi(splitSrc[lineNumber*9+columnNumber])
			puzzle[lineNumber][columnNumber] = value

			for i := 0; i < 9; i++ {
				if value == 0 {
					possibilities[lineNumber][columnNumber][i] = true
				} else {
					possibilities[lineNumber][columnNumber][i] = false
				}
			}
		}
	}
	// printPuzzle(puzzle)
	for lineNumber := 0; lineNumber < 9; lineNumber++ {
		for columnNumber := 0; columnNumber < 9; columnNumber++ {
			value := puzzle[lineNumber][columnNumber]
			if value > 0 {
				clearNumber(lineNumber, columnNumber, value, &possibilities)
			}
		}
	}
	// for lineNumber := 0; lineNumber < 9; lineNumber++ {
	// 	for columnNumber := 0; columnNumber < 9; columnNumber++ {
	// 		if puzzle[lineNumber][columnNumber] == 0 {
	// 			possibility := checkPossibility(possibilities[lineNumber][columnNumber])
	// 			// print(fmt.Sprintf("%d, %d = %d\n", lineNumber, columnNumber, possibility))
	// 			if possibility > 0 {
	// 				print(fmt.Sprintf("%d, %d = %d", lineNumber, columnNumber, possibility))
	// 				puzzle[lineNumber][columnNumber] = possibility
	// 			}
	// 		}
	// 	}
	// }
	// printPuzzle(puzzle)
	return puzzle, possibilities
}

func checkSquare(line int, column int, puzzle *[9][9]int, possibilities *[9][9][9]bool) {
	// print(fmt.Sprintf("Check square %d %d\n", line, column))
	for value := 1; value <= 9; value++ {
		occurences := 0
		occLine := 0
		occColumn := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if possibilities[line+i][column+j][value-1] == true {
					occurences++
					occLine = line + i
					occColumn = column + j
				}
			}
		}
		if occurences == 1 {
			// print(fmt.Sprintf("%d, %d = %d\n", occLine, occColumn, value))
			puzzle[occLine][occColumn] = value
			clearNumber(occLine, occColumn, value, possibilities)
			printPuzzle(*puzzle)
		}
	}
}

func checkLine(line int, puzzle *[9][9]int, possibilities *[9][9][9]bool) {
	for value := 1; value <= 9; value++ {
		occurences := 0
		occColumn := 0
		for i := 0; i < 9; i++ {
			if possibilities[line][i][value-1] == true {
				occurences++
				occColumn = i
			}
		}
		if occurences == 1 {
			puzzle[line][occColumn] = value
			clearNumber(line, occColumn, value, possibilities)
			printPuzzle(*puzzle)
		}
	}
}

func checkColumn(column int, puzzle *[9][9]int, possibilities *[9][9][9]bool) {
	for value := 1; value <= 9; value++ {
		occurences := 0
		occLine := 0
		for i := 0; i < 9; i++ {
			if possibilities[i][column][value-1] == true {
				occurences++
				occLine = i
			}
		}
		if occurences == 1 {
			puzzle[occLine][column] = value
			clearNumber(occLine, column, value, possibilities)
			printPuzzle(*puzzle)
		}
	}
}

func solve(puzzle *[9][9]int, possibilities *[9][9][9]bool) bool {
	for try := 0; try < 10000; try++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				checkSquare(i*3, j*3, puzzle, possibilities)
				checkLine(i*3+j, puzzle, possibilities)
				checkColumn(i*3+j, puzzle, possibilities)
			}
		}
	}
	return false
}

func main() {
	res, err := http.Get("https://sudoku.com/api/getLevel/expert")
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

	resJSON := apiResponse{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		log.Fatal(err)
	}

	puzzleLine := fmt.Sprintf("%v", resJSON.Desc[apiPuzzleIdx])
	puzzleSolution := fmt.Sprintf("%v", resJSON.Desc[apiSolutionIdx])
	// print(puzzleLine)
	// puzzleLine := "000740006406800507700090004030984700820613409040000300062370005005409000070061208"

	puzzle, possibilities := initialize(puzzleLine)
	printPossibilities(possibilities)
	printPuzzle(puzzle)
	solve(&puzzle, &possibilities)
	printPuzzle(puzzle)
	printPossibilities(possibilities)
	print(puzzleLine + "\n")
	printPuzzleLine(puzzle)
	print(puzzleSolution + "\n")
	// for solve(&puzzle, &possibilities) == false {
	// printPuzzle(puzzle)
	// }

	// print(possibilities)

	// for i := 0; i < 100; i++ {
	// 	print("\033[H\033[2J")
	// 	printSudoku(fmt.Sprintf("%v", resJSON.Desc[apiPuzzleIdx]))
	// 	print(i)
	// 	print("\n")
	// 	time.Sleep(1 * time.Second)
	// }
	// fmt.Print(resJSON.Desc[apiSolutionIdx])
}
