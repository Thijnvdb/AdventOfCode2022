package main

import (
	"aoc/day15/puzzles/puzzle1"
	"aoc/day15/puzzles/puzzle2"
	"aoc/shared"
	"fmt"
)

func main() {
	err := shared.RunPuzzle(puzzle1.Run, puzzle2.Run)
	if err != nil {
		fmt.Println("Error occured while running puzzle:", err.Error())
	}
}
