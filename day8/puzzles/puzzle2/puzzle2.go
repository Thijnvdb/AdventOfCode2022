package puzzle2

import (
	"aoc/day8/common"
	"fmt"
)

func Run(inputFile string) error {
	forest, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	max := common.CalculateScenicScores(forest)

	common.PrintForestHighlighted(forest, max)

	fmt.Printf("Best scenic score: %v", max.ScenicScore)

	return nil
}
