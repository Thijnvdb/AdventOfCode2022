package puzzle2

import (
	"aoc/day9/common"
	"fmt"
)

func Run(inputFile string) error {
	moves, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	rope := common.CreateRope(10)

	fmt.Printf("Rope size: %v", rope.GetRopeSize())

	common.SimulateMoves(rope, moves, false)

	tail := rope.GetTail()

	set := common.GetUniquePositionsTraversed(tail)

	common.PrintStateVisited(rope, set)

	fmt.Printf("\n\nUnique points visited: %v\n\n", set.Size())

	return nil
}
