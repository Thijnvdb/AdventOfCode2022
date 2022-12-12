package puzzle2

import (
	"aoc/day12/common"
	"aoc/shared"
	"aoc/shared/types"
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
)

func Run(inputFile string) error {
	start, end, grid, err := ParseInput(inputFile)
	if err != nil {
		return err
	}

	min := 99999999999999
	minPath := arraylist.New()
	minStart := common.NewNode(0, types.ZeroVector(), "")
	for _, startPoint := range start.Values() {
		point := startPoint.(*common.Node)
		traveled, err := common.AStar(grid, point, end, false)
		if err != nil {
			continue
		}

		if traveled.Size() < min {
			min = traveled.Size()
			minPath = traveled
			minStart = point
		}
	}

	fmt.Printf("\nTraveled path length: %v", minPath.Size()-1)
	common.PrintHighlighted(grid, minStart, end, minPath)

	return nil
}

func ParseInput(inputFile string) (*arraylist.List, *common.Node, [][]*common.Node, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, nil, nil, err
	}

	grid := [][]*common.Node{}
	// create cols
	for x := 0; x < len(lines[0]); x++ {
		grid = append(grid, []*common.Node{})
	}

	start := arraylist.New()
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
				start.Add(node)
			}

			if node.Elevation == 0 {
				start.Add(node)
			}

			grid[x] = append(grid[x], node)
		}
	}

	return start, end, grid, nil
}
