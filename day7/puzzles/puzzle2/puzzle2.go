package puzzle2

import (
	"aoc/day7/common"
	"fmt"
	"math"
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

	toDelete := FindSmallestDirToDelete(8381165, root)

	fmt.Printf("Directory to delete: %v (%v)", toDelete.Name, toDelete.Size)

	return nil
}

func FindSmallestDirToDelete(spaceToFree int, root *common.Directory) *common.Directory {
	root.RecalculateSizes() // set all sizes

	viable := recurse(spaceToFree, root, []*common.Directory{})

	minDir := new(common.Directory)
	minDir.Size = math.MaxInt // set max
	for _, dir := range viable {
		if dir.Size < minDir.Size {
			minDir = dir
		}
	}

	return minDir
}

func recurse(spaceToFree int, current *common.Directory, viable []*common.Directory) []*common.Directory {
	if current.Size >= spaceToFree {
		// this directory would be big enough to delete
		viable = append(viable, current)
	}

	for _, subdir := range current.SubDirectories {
		viable = recurse(spaceToFree, subdir, viable)
	}

	return viable
}
