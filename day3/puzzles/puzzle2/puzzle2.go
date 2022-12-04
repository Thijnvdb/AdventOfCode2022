package puzzle2

import (
	"aoc/day3/common"
	"aoc/shared"
	"fmt"
)

func Run(input_file string) error {
	lines, err := shared.ReadFile(input_file)
	if err != nil {
		return err
	}

	total := 0
	group := []string{}
	for i, line := range lines {
		group = append(group, line)

		if (i+1)%3 == 0 {
			letter, err := common.GetMatchInGroup(group[0], group[1], group[2])
			if err != nil {
				return err
			}

			priority, err := common.GetPriority(letter)
			if err != nil {
				return err
			}

			total += priority
			group = []string{}
		}
	}

	fmt.Printf("Total priority: %v", total)

	return nil
}
