package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiPuzzleIdx = 0
const apiSolutionIdx = 1

type ApiResponse struct {
	Answer  string        `json:"answer"`
	Message string        `json:"message"`
	Desc    []interface{} `json:"desc"`
}

func CallAPI(difficulty string) (string, string) {
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

	resJSON := ApiResponse{}
	err = json.Unmarshal(body, &resJSON)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", resJSON.Desc[apiPuzzleIdx]), fmt.Sprintf("%v", resJSON.Desc[apiSolutionIdx])
}
