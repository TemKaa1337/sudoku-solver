package internal

import (
	"fmt"
	"golang.org/x/exp/slices"
	"sudoku_solver/internal/types"
)

var rowSquareEndingIndexes = []int{2, 5, 8}

func Print(sudoku *types.Sudoku) {
	fmt.Println("_________________________")
	for row := 0; row < types.RowCount; row++ {
		fmt.Print("| ")
		for column := 0; column < types.ColumnCount; column++ {
			fmt.Print(sudoku.Board[row][column] + " ")

			if slices.Contains(rowSquareEndingIndexes, column) {
				fmt.Print("| ")
			}
		}

		fmt.Println()
		if slices.Contains(rowSquareEndingIndexes, row) {
			fmt.Println("_________________________")
		}
	}
}
