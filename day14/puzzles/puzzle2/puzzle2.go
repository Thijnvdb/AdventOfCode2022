package puzzle2

import (
	"aoc/day14/common"
	"fmt"
)

func Run(inputFile string) error {
	cave, start, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	// add floor to cave
	for i := 0; i < len(cave); i++ {
		cave[i][len(cave[0])-1].State = common.Rock
	}

	restCount, path := startFloodingTheEntireFookingCave(cave, start, false)

	common.PrintCave(cave)

	fmt.Printf("Rest count: %v\n%v\n", restCount, len(path))

	return nil
}

func startFloodingTheEntireFookingCave(cave common.Cave, start *common.Point, print bool) (int, []*common.Point) {
	count := 0
	for {
		start.State = common.Sand
		path, _ := common.FloodStep(cave, start, []*common.Point{}, print)
		count++
		if start.State == common.Sand {
			return count, path
		}
	}
}
