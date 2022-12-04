package day3

import (
	"aoc/day3/puzzles/puzzle1"
	"aoc/day3/puzzles/puzzle2"
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
