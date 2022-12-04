package day2

import (
	puzzle1 "aoc/day2/puzzles/puzzle1"
	puzzle2 "aoc/day2/puzzles/puzzle2"
	"errors"
)

func RunPuzzle(index int, inputFile string) error {
	if index == 1 {
		return puzzle1.Run(inputFile)
	}
	if index == 2 {
		return puzzle2.Run(inputFile)
	}
	return errors.New("puzzle index was not valid")
}
