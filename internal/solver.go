package internal

import (
	"errors"
	"strconv"
	"sudoku_solver/internal/types"
)

func Solve(sudoku *types.Sudoku) error {
	solve(sudoku, 0, 0)

	if !isSolved(sudoku) {
		return errors.New("could not solve sudoku as it is not valid")
	}

	return nil
}

func isSolved(sudoku *types.Sudoku) bool {
	for row := 0; row < types.RowCount; row++ {
		for column := 0; column < types.ColumnCount; column++ {
			if sudoku.Board[row][column] == "." {
				return false
			}
		}
	}

	return true
}

func solve(sudoku *types.Sudoku, row int, column int) {
	if isSolved(sudoku) {
		return
	}

	if sudoku.Board[row][column] == "." {
		for _, value := range types.AllowedValues {
			if !canBePlaced(sudoku, row, column, value) {
				continue
			}

			sudoku.Board[row][column] = strconv.Itoa(value)

			newRow, newColumn := pickNextCoordinate(row, column)
			solve(sudoku, newRow, newColumn)

			if isSolved(sudoku) {
				return
			}

			sudoku.Board[row][column] = "."
		}

		return
	}

	row, column = pickNextCoordinate(row, column)
	solve(sudoku, row, column)
}

func canBePlaced(sudoku *types.Sudoku, row int, column int, value int) bool {
	// check duplicates in row
	for tempColumn := 0; tempColumn < types.ColumnCount; tempColumn++ {
		if sudoku.Board[row][tempColumn] == strconv.Itoa(value) {
			return false
		}
	}

	// check duplicates in column
	for tempRow := 0; tempRow < types.ColumnCount; tempRow++ {
		if sudoku.Board[tempRow][column] == strconv.Itoa(value) {
			return false
		}
	}

	// check duplicates in square
	squareCoordinates := getSquareCoordinates(row, column)
	for _, coordinates := range squareCoordinates {
		tempRow, tempColumn := coordinates[0], coordinates[1]

		if sudoku.Board[tempRow][tempColumn] == strconv.Itoa(value) {
			return false
		}
	}

	return true
}

func getSquareCoordinates(row int, column int) [9][2]int {
	startRow := int(row/3) * 3
	startColumn := int(column/3) * 3

	var coordinates [9][2]int

	index := 0
	for row = startRow; row < startRow+3; row++ {
		for column = startColumn; column < startColumn+3; column++ {
			coordinates[index] = [2]int{row, column}
			index++
		}
	}

	return coordinates
}

func pickNextCoordinate(row int, column int) (int, int) {
	if row == types.RowCount-1 && column == types.ColumnCount-1 {
		return 0, 0
	}

	if column == types.ColumnCount-1 {
		return row + 1, 0
	}

	return row, column + 1
}
