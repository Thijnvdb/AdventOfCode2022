package puzzle1

import (
	"aoc/day9/common"
	"fmt"
)

func Run(inputFile string) error {
	moves, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	rope := common.CreateRope(2)

	fmt.Printf("Rope size: %v", rope.GetRopeSize())

	common.SimulateMoves(rope, moves, false)

	tail := rope.GetTail()

	count := common.GetUniquePositionsTraversed(tail).Size()

	fmt.Printf("\n\nUnique points visited: %v\n\n", count)

	return nil
}
