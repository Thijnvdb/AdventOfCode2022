package puzzle1

import (
	"aoc/day7/common"
	"fmt"
)

func Run(inputFile string) error {
	commands, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	root, err := common.ConstructFileSystem(commands)
	if err != nil {
		return err
	}

	fmt.Printf("\nSize: %v\n\n", GetSumOfDirectoriesBelowSize(100000, root))

	common.PrintFileSystem(root)

	return nil
}

func GetSumOfDirectoriesBelowSize(size int, root *common.Directory) int {
	root.RecalculateSizes() // calculate all sizes
	return recurse(size, root, 0)
}

func recurse(maxSize int, dir *common.Directory, total int) int {
	if dir.Size <= maxSize {
		fmt.Printf("Adding dir: %v\n", dir.Name)
		total += dir.Size
	}

	for _, subdir := range dir.SubDirectories {
		total = recurse(maxSize, subdir, total)
	}

	fmt.Printf("total: %v\n", total)

	return total
}
