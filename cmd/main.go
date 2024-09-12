package main

import (
	"fmt"
	solver "sudoku_solver/internal"
)

func main() {
	var sudoku, err = solver.Load("/app/source/sudoku.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("source sudoku:")
	solver.Print(&sudoku)

	err = solver.Solve(&sudoku)
	if err != nil {
		panic(err)
	}

	fmt.Println("solved sudoku:")
	solver.Print(&sudoku)
}
