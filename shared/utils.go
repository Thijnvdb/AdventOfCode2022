package shared

import (
	"errors"
	"os"
	"strconv"
)

func RunPuzzle(puzzle1 func(string) error, puzzle2 func(string) error) error {
	puzzleIn, errPuzzle := strconv.ParseInt(os.Args[1], 10, 0)
	file_path := os.Args[2]

	if errPuzzle != nil || puzzleIn > 2 || puzzleIn < 1 {
		return errors.New("given puzzle index was not valid")
	}

	puzzle := int(puzzleIn)
	if puzzle == 1 {
		return puzzle1(file_path)
	} else if puzzle == 2 {
		return puzzle2(file_path)
	}

	return errors.New("no puzzle found")
}

func Reverse[T any](list []T) []T {
	newList := []T{}
	for i := len(list) - 1; i >= 0; i-- {
		newList = append(newList, list[i])
	}

	return newList
}
