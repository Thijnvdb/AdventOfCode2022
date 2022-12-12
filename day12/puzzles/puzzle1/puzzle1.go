package puzzle1

import (
	"aoc/day12/common"
	"aoc/shared"
	"aoc/shared/types"
	"fmt"
)

func Run(inputFile string) error {
	start, end, grid, err := ParseInput(inputFile)
	if err != nil {
		return err
	}

	traveled, err := common.AStar(grid, grid[0][13], end, false)
	if err != nil {
		return err
	}

	fmt.Printf("\nTraveled path length: %v\n", traveled.Size()-1)

	common.PrintHighlighted(grid, start, end, traveled)

	return nil
}

// returns start, end and grid
func ParseInput(inputFile string) (*common.Node, *common.Node, [][]*common.Node, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, nil, nil, err
	}

	grid := [][]*common.Node{}
	// create cols
	for x := 0; x < len(lines[0]); x++ {
		grid = append(grid, []*common.Node{})
	}

	start := &common.Node{}
	end := &common.Node{}

	for y := len(lines) - 1; y >= 0; y-- {
		line := lines[y]
		for x, char := range line {
			elevation := char - common.RuneOffset
			v := types.Vector{X: x, Y: len(lines) - y - 1}
			node := common.NewNode(int(elevation), v, string(char))
			if string(char) == "E" {
				node.Elevation = 25
				end = node
			}
			if string(char) == "S" {
				node.Elevation = 0
				start = node
			}

			grid[x] = append(grid[x], node)
		}
	}

	return start, end, grid, nil
}
