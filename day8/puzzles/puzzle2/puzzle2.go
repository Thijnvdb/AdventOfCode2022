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

	max := CalculateScenicScores(forest)

	common.PrintForestHighlighted(forest, max)

	fmt.Printf("Best scenic score: %v", max.ScenicScore)

	return nil
}

// calculate scenic scores, return most viable tree
func CalculateScenicScores(forest []common.TreeLine) *common.Tree {
	max := new(common.Tree)
	max.ScenicScore = 0 // awful tree yuck

	for _, treeCol := range forest {
		for _, tree := range treeCol {
			score := tree.GetScenicScore(forest)
			if score > max.ScenicScore {
				max = tree
			}
		}
	}

	return max
}
