package main

import (
	"aoc/day3/puzzles/puzzle1"
	"aoc/day3/puzzles/puzzle2"
	"aoc/shared"
	"fmt"
)

func main() {
	err := shared.RunPuzzle(puzzle1.Run, puzzle2.Run)
	if err != nil {
		fmt.Println("Error occured while running puzzle:", err.Error())
	}
}
