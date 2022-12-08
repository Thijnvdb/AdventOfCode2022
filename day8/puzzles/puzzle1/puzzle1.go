package puzzle1

import (
	"aoc/day8/common"
	"aoc/shared"
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

func Run(inputFile string) error {
	forest, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	visible := GetVisible(forest)

	common.PrintForestVisible(forest, *visible)

	fmt.Printf("Total visible tree count: %v ", len(visible.Values()))

	return nil
}

func GetVisible(forest []common.TreeLine) *hashset.Set {
	set := hashset.New()
	addTreesToSet(getVisibleFromHorizontalSide(forest, true), set)
	addTreesToSet(getVisibleFromHorizontalSide(forest, false), set)
	addTreesToSet(getVisibleFromVerticalSide(forest, true), set)
	addTreesToSet(getVisibleFromVerticalSide(forest, false), set)

	return set
}

func addTreesToSet(treelines []common.TreeLine, set *hashset.Set) {
	for _, line := range treelines {
		for _, tree := range line {
			set.Add(tree)
		}
	}
}

func getVisibleFromHorizontalSide(forest []common.TreeLine, left bool) []common.TreeLine {
	visible := []common.TreeLine{}

	// iterate over vertical tree lines (top to bottom)
	for i := 0; i < len(forest[0]); i++ {
		treeRow := common.TreeLine{}
		for _, treeCol := range forest {
			treeRow = append(treeRow, treeCol[i])
		}

		if !left {
			treeRow = shared.Reverse(treeRow)
		}

		// first always visible
		visibleInRow := common.TreeLine{treeRow[0]}
		middle := treeRow[1:]

		largestPrev := treeRow[0].Size
		for _, tree := range middle {
			if tree.Size > largestPrev {
				largestPrev = tree.Size
				visibleInRow = append(visibleInRow, tree)
			}
		}

		visible = append(visible, visibleInRow)
	}

	return visible
}

func getVisibleFromVerticalSide(forest []common.TreeLine, top bool) []common.TreeLine {
	visible := []common.TreeLine{}

	// iterate over vertical tree lines (top to bottom)
	for _, treeCol := range forest {
		col := treeCol

		if !top {
			col = shared.Reverse(col)
		}

		// first always visible
		visibleInCol := common.TreeLine{col[0]}
		middle := col[1:]

		largestPrev := col[0].Size
		for _, tree := range middle {
			if tree.Size > largestPrev {
				largestPrev = tree.Size
				visibleInCol = append(visibleInCol, tree)
			}
		}

		visible = append(visible, visibleInCol)
	}

	return visible
}
