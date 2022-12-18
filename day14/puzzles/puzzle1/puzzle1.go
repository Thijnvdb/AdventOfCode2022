package puzzle1

import (
	"aoc/day14/common"
	"fmt"
)

func Run(inputFile string) error {
	cave, start, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	restCount, path := startFloodingTheEntireFookingCave(cave, start, false)

	fmt.Printf("Rest count: %v\n%v\n", restCount, len(path))

	common.PrintCave(cave)

	return nil
}

func startFloodingTheEntireFookingCave(cave common.Cave, start *common.Point, print bool) (int, []*common.Point) {
	count := 0
	for {
		start.State = common.Sand
		path, done := common.FloodStep(cave, start, []*common.Point{}, print)
		if done {
			return count, path
		}
		count++
	}
}
