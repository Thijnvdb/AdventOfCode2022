package puzzle1

import (
	"aoc/day1/common"
	"fmt"
)

func Run(inputFile string) error {
	elfs, err := common.GetElfsSortedByCalories(inputFile)
	if err != nil {
		return err
	}

	fmt.Printf("#1 Elf calorie count: %d", elfs[0].Calories)

	return nil
}
