package puzzle2

import (
	"aoc/day1/common"
	"fmt"
)

func Run(inputFile string) error {
	elfs, err := common.GetElfsSortedByCalories(inputFile)
	if err != nil {
		return err
	}

	total := 0
	for _, elf := range elfs[:3] {
		total += elf.Calories
	}

	fmt.Printf("Total calories of first 3 elves: %d", total)

	return nil
}
