package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"sudoku_solver/internal/types"
)

func Load(path string) (types.Sudoku, error) {
	var decoded types.Sudoku

	var fileContent, err = os.ReadFile(path)
	if err != nil {
		return decoded, fmt.Errorf("could not load json file data with message: \"%s\"", err)
	}

	err = json.Unmarshal(fileContent, &decoded)
	if err != nil {
		return decoded, fmt.Errorf("could not unmrashall json with message: \"%s\"", err)
	}

	return decoded, nil
}
