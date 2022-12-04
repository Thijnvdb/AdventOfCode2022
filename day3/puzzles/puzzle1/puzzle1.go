package puzzle1

import (
	common "aoc/day3/common"
	shared "aoc/shared"
	"fmt"
)

func Run(input_file string) error {
	lines, err := shared.ReadFile(input_file)
	if err != nil {
		return err
	}

	total := 0
	for _, line := range lines {
		split := common.SplitBackpack(line)
		commonLetter, err := common.GetMatch(split[0], split[1])
		if err != nil {
			return err
		}

		priority, err := common.GetPriority(commonLetter)
		if err != nil {
			return err
		}

		total += priority
	}

	fmt.Printf("Total priority: %v", total)

	return nil
}
