package types

import (
	"encoding/json"
)

const RowCount = 9

const ColumnCount = 9

type Sudoku struct {
	Board [][]string
}

var AllowedValues = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func (sudoku *Sudoku) UnmarshalJSON(buffer []byte) error {
	board := make([][]string, RowCount)

	var err = json.Unmarshal(buffer, &board)
	if err != nil {
		return err
	}

	sudoku.Board = board

	return nil
}
