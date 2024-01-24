package sudokujson

import (
	"encoding/json"
)

type SudokuFmt struct {
	Data     []int `json:"data"`
	Length   int   `json:"length"`
	MaxValue int   `json:"maxValue"`
}

var sudoku SudokuFmt

func ParseJson(b []byte) (sf *SudokuFmt, err error) {
	err = json.Unmarshal(b, &sudoku)
	if err != nil {
		return nil, err
	}
	return &sudoku, nil
}

func GenJson(sf SudokuFmt) (b []byte, err error) {
	b, err = json.Marshal(sf)
	return b, err
}
