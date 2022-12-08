package puzzle1

import (
	"aoc/day8/common"
	"fmt"
)

func Run(inputFile string) error {
	forest, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	visible := common.GetVisible(forest)

	common.PrintForestVisible(forest, *visible)

	fmt.Printf("Total visible tree count: %v ", len(visible.Values()))

	return nil
}
