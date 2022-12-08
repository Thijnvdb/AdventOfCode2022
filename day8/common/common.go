package common

import (
	"aoc/shared"
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

type Tree struct {
	ScenicScore int
	Size        int
	X           int
	Y           int
}

type TreeLine []*Tree

func NewTree(X int, Y int, Size int) *Tree {
	tree := new(Tree)
	tree.X = X
	tree.Y = Y
	tree.Size = Size
	return tree
}

func (trees TreeLine) Len() int {
	return len(trees)
}

func (trees TreeLine) Less(i, j int) bool {
	return trees[i].Size < trees[j].Size
}

func (trees TreeLine) Swap(i, j int) {
	trees[i], trees[j] = trees[j], trees[i]
}

func (tree *Tree) GetScenicScore(forest []TreeLine) int {
	left, right := tree.getScenicScoreHorizontal(forest)
	top, bottom := tree.getScenicScoreVertical(forest)

	score := left * right * top * bottom
	tree.ScenicScore = score
	return score
}

func (tree *Tree) getScenicScoreHorizontal(forest []TreeLine) (int, int) {
	row := TreeLine{}
	for _, col := range forest {
		row = append(row, col[tree.Y])
	}

	// visible to the right
	scenicRight := 0
	for x := tree.X; x < len(row); x++ {
		checkTree := row[x]
		if checkTree == tree {
			continue
		}
		scenicRight++
		if row[x].Size >= tree.Size {
			break
		}
	}

	// visible to the left
	scenicLeft := 0
	for x := tree.X; x >= 0; x-- {
		checkTree := row[x]
		if checkTree == tree {
			continue
		}
		scenicLeft++
		if row[x].Size >= tree.Size {
			break
		}
	}

	return scenicLeft, scenicRight
}

func (tree *Tree) getScenicScoreVertical(forest []TreeLine) (int, int) {
	col := forest[tree.X]

	// visible to the bottom
	scenicBottom := 0
	for y := tree.Y; y < len(col); y++ {
		checkTree := col[y]
		if checkTree == tree {
			continue
		}
		scenicBottom++
		if checkTree.Size >= tree.Size {
			break
		}
	}

	// visible to the top
	scenicTop := 0
	for y := tree.Y; y >= 0; y-- {
		checkTree := col[y]
		if checkTree == tree {
			continue
		}
		scenicTop++
		if checkTree.Size >= tree.Size {
			break
		}
	}

	return scenicTop, scenicBottom
}

func ParseInput(input string) ([]TreeLine, error) {
	lines, err := shared.ReadFile(input)
	if err != nil {
		return nil, err
	}

	result := []TreeLine{}

	for i := 0; i < len(lines[0]); i++ {
		result = append(result, TreeLine{})
	}

	for y, line := range lines {
		for x, character := range line {
			size := int(character - '0')
			result[x] = append(result[x], NewTree(x, y, size))
		}
	}

	return result, nil
}

func PrintForest(forest []TreeLine) {
	fmt.Print("\n")
	for y := 0; y < len(forest[0]); y++ {
		for x := 0; x < len(forest); x++ {
			tree := forest[x][y]
			fmt.Printf(" %v ", tree.Size)
		}

		fmt.Print("\n\n")
	}
}

func PrintForestHighlighted(forest []TreeLine, highlight *Tree) {
	fmt.Print("\n")
	for y := 0; y < len(forest[0]); y++ {
		for x := 0; x < len(forest); x++ {
			tree := forest[x][y]
			if tree == highlight {
				fmt.Printf("\033[0;31m %v (%v) \033[0m", tree.Size, tree.ScenicScore)
			} else {
				fmt.Printf(" %v (%v) ", tree.Size, tree.ScenicScore)
			}
		}

		fmt.Print("\n\n")
	}
}

func PrintForestVisible(forest []TreeLine, visible hashset.Set) {
	fmt.Print("\n")
	for y := 0; y < len(forest[0]); y++ {
		for x := 0; x < len(forest); x++ {
			tree := forest[x][y]
			if visible.Contains(tree) {
				fmt.Print("\033[0;31m")
			} else {
				fmt.Print("\033[0m")
			}
			fmt.Printf(" %v ", tree.Size)
		}

		fmt.Print("\033[0m\n\n")
	}
}
