package main

import (
	day2 "aoc/day2"
	day3 "aoc/day3"
	"log"
	"os"
	"strconv"
)

func main() {
	dayIn, errDay := strconv.ParseInt(os.Args[1], 10, 0)
	puzzleIn, errPuzzle := strconv.ParseInt(os.Args[2], 10, 0)
	file_path := os.Args[3]

	puzzle := int(puzzleIn)
	day := int(dayIn)

	if errDay != nil || errPuzzle != nil {
		panic("Invalid input!")
	}

	var result error
	switch day {
	case 2:
		result = day2.RunPuzzle(puzzle, file_path)
	case 3:
		result = day3.RunPuzzle(puzzle, file_path)
	}

	if result != nil {
		print("oops")
		log.Fatal(result)
	}
}
