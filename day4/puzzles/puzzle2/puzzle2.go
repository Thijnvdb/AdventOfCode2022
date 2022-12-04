package puzzle2

import (
	"aoc/day4/puzzles/common"
	"fmt"
)

func Run(inputFile string) error {
	tasks, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	total := 0
	for _, task := range tasks {
		if checkIfOverlaps(task[0].Sections, task[1].Sections) {
			total++
		}
	}

	fmt.Printf("Overlapping count: %v", total)

	return nil
}

// check if the 2 sections overlap at all
func checkIfOverlaps(firstSection []int, secondSection []int) bool {
	return (firstSection[0] <= secondSection[0] && secondSection[0] <= firstSection[len(firstSection)-1]) ||
		(firstSection[0] <= secondSection[len(secondSection)-1] && secondSection[len(secondSection)-1] <= firstSection[len(firstSection)-1]) ||
		(secondSection[0] <= firstSection[0] && firstSection[0] <= secondSection[len(secondSection)-1]) ||
		(secondSection[0] <= firstSection[len(firstSection)-1] && firstSection[len(firstSection)-1] <= secondSection[len(secondSection)-1])
}
