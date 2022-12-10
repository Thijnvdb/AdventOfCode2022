package common

import (
	"aoc/shared/types"
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

func PrintState(head *RopePoint) {
	positions := head.GetRopePositions()
	size := 30 //idk man

	grid := [][]string{}
	for x := -size; x < size; x++ {
		grid = append(grid, []string{})

		for y := -size; y < size; y++ {
			grid[x+size] = append(grid[x+size], "*")
		}
	}

	for i, pos := range positions {
		if pos.X+size > len(grid)-1 || pos.Y+size > len(grid)-1 || pos.X < 0 || pos.Y < 0 {
			continue // prevent error when going out of grid
		}
		grid[pos.X+size][pos.Y+size] = fmt.Sprint(i)
	}

	line := ""
	for i := 0; i < size*2; i++ {
		line += "-"
	}

	fmt.Printf("\n%v\n\n", line)
	for y := size - 1; y >= -size; y-- {
		for x := -size; x < size; x++ {
			if x+size > len(grid)-1 || y+size > len(grid)-1 || y+size < 0 || x+size < 0 {
				continue // prevent error when going out of grid
			}
			fmt.Printf("%v", grid[x+size][y+size])
		}
		fmt.Print("\n")
	}
	fmt.Printf("\n%v\n", line)
}

func PrintStateVisited(head *RopePoint, visited *hashset.Set) {
	positions := head.GetRopePositions()
	size := 50 //idk man

	grid := [][]string{}
	for x := -size; x < size; x++ {
		grid = append(grid, []string{})

		for y := -size; y < size; y++ {
			grid[x+size] = append(grid[x+size], "*")
		}
	}

	for i, pos := range positions {
		if pos.X+size > len(grid)-1 || pos.Y+size > len(grid)-1 || pos.X < 0 || pos.Y < 0 {
			continue // prevent error when going out of grid
		}
		grid[pos.X+size][pos.Y+size] = fmt.Sprint(i)
	}

	for _, point := range visited.Values() {
		if point.(types.Vector).X+size > len(grid)-1 || point.(types.Vector).Y+size > len(grid)-1 || point.(types.Vector).X < 0 || point.(types.Vector).Y < 0 {
			continue // prevent error when going out of grid
		}
		grid[point.(types.Vector).X+size][point.(types.Vector).Y+size] = "\033[0;31m#\033[0m"
	}

	for y := size - 1; y >= -size; y-- {
		for x := -size; x < size; x++ {
			if x+size > len(grid)-1 || y+size > len(grid)-1 || y+size < 0 || x+size < 0 {
				continue // prevent error when going out of grid
			}
			fmt.Printf("%v", grid[x+size][y+size])
		}
		fmt.Print("\n")
	}

}
