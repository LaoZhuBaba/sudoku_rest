package sudokurest

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/LaoZhuBaba/sudoku_rest/autogen"
	sudokujson "github.com/LaoZhuBaba/sudoku_rest/json"
	"github.com/LaoZhuBaba/sudoku_rest/solve"
)

func init() {
	functions.HTTP("SudokuRest", SudokuRest)
	log.Printf("Cold Start detected")
}

func SudokuRest(w http.ResponseWriter, r *http.Request) {
	log.Printf("top of the function")

	// Read message body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error from io.ReadAll: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the json
	sudoku, err := sudokujson.ParseJson(bodyBytes)
	if err != nil {
		log.Printf("error from ParseJson: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if sudoku.Length != len(sudoku.Data) {
		log.Printf("value of the Length field doesn't match the length of the Data field")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Solve the puzzle
	var re *[][]int
	for _, v := range autogen.RelatedElements {
		if sudoku.Length == len(v) {
			re = &v
			break
		}
	}
	if re == nil {
		log.Printf("the size of the sudoku.Data field is unsupported")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := solve.SolveSudoku(sudoku.Data, sudoku.MaxValue, sudoku.Length-1, *re)
	if err != nil {
		log.Printf("error returned by solve.SolveSudoku: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return a response with the same JSON schema as the request, but with the Data field completed
	// Length and MaxValue don't matter in the response so just echo back the values in the request
	result := sudokujson.SudokuFmt{
		Data:     data,
		Length:   sudoku.Length,
		MaxValue: sudoku.MaxValue,
	}
	jsonResult, err := sudokujson.GenJson(result)
	if err != nil {
		log.Printf("error returned by solve.SolveSudoku: %v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", string(jsonResult))
}
