package common

import (
	"aoc/shared/colors"
	types "aoc/shared/types"
	"fmt"
	"math"
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
)

// offset to subtract from rune to get elevation
const RuneOffset = 97

func PrintGrid(grid [][]*Node, start *Node, end *Node) {
	PrintHighlighted(grid, start, end, arraylist.New())
}

func PrintHighlighted(grid [][]*Node, start *Node, end *Node, highlightedNodes *arraylist.List) {
	fmt.Println()
	for y := len(grid[0]) - 1; y >= 0; y-- {
		for x := 0; x < len(grid); x++ {
			node := grid[x][y]

			str := ColorElevation(node.Elevation, node.ElevationChar)

			if highlightedNodes.Contains(node) {
				str = colors.Yellow(node.ElevationChar)
			}

			if node == start {
				str = colors.Green(node.ElevationChar)
			}

			if node == end {
				str = colors.Red(node.ElevationChar)
			}

			fmt.Print(str)
		}
		fmt.Print("\n")
	}
}

func ColorElevation(elevation int, char string) string {
	return colors.Rgb(char, elevation*4, elevation*10, elevation*4)
}

type Node struct {
	Id            int
	ElevationChar string
	Elevation     int
	Position      types.Vector
	G             float64
	F             float64
	CameFrom      *Node
}

func (node *Node) ToString() string {
	if node.CameFrom != nil {
		return fmt.Sprintf("{ Id: %v F: %v, G: %v X: %v, Y: %v, ParentId: %v, ParentPosition: {X:%v, Y:%v} }",
			node.Id, node.F, node.G, node.Position.X, node.Position.Y, node.CameFrom.Id, node.CameFrom.Position.X, node.CameFrom.Position.Y)
	}
	return fmt.Sprintf("{ Id: %v F: %v, G: %v X: %v, Y: %v }", node.Id, node.F, node.G, node.Position.X, node.Position.Y)
}

var nodeId = 0

func NewNode(elevation int, position types.Vector, elevationChar string) *Node {
	node := new(Node)
	node.Elevation = elevation
	node.ElevationChar = elevationChar
	node.Position = position
	node.G = math.MaxInt
	node.F = math.MaxInt
	node.Id = nodeId
	node.CameFrom = nil
	nodeId++
	return node
}

func GetLowestFValue(openList *hashset.Set, goal *Node) *Node {
	var lowest *Node = nil
	min := math.Inf(1)
	for _, n := range openList.Values() {
		node := n.(*Node)
		if node.F < min {
			lowest = node
			min = node.F
		}
	}

	return lowest
}

func GetNeighbours(current *Node, grid [][]*Node) []*Node {
	neigbours := []*Node{}

	pos := current.Position

	if pos.X > 0 {
		neighbour := grid[pos.X-1][pos.Y]

		neigbours = append(neigbours, neighbour)
	}

	if pos.X < len(grid)-1 {
		neighbour := grid[pos.X+1][pos.Y]
		neigbours = append(neigbours, neighbour)
	}

	if pos.Y > 0 {
		neighbour := grid[pos.X][pos.Y-1]
		neigbours = append(neigbours, neighbour)
	}

	if pos.Y < len(grid[0])-1 {
		neighbour := grid[pos.X][pos.Y+1]
		neigbours = append(neigbours, neighbour)
	}

	return neigbours
}

func ElevationDiff(a *Node, b *Node) int {
	diff := a.Elevation - b.Elevation
	return diff
}

func ReconstructPath(current *Node) *arraylist.List {
	if current.CameFrom == nil {
		return arraylist.New(current)
	} else {
		res := ReconstructPath(current.CameFrom)
		res.Add(current)
		return res
	}
}

// assuming that the 2 nodes are neighbours
func GetWeight(current *Node, neighbour *Node) float64 {
	diff := ElevationDiff(current, neighbour)
	if diff < -1 {
		// neighbour is higher than 1 away
		return math.Inf(1) // do not cross 2 or higher
	} else {
		return 1 // this is fine
	}
}

func Heuristic(node *Node, goal *Node, grid [][]*Node) float64 {
	xdiff := math.Abs(float64(node.Position.X - goal.Position.X))
	ydiff := math.Abs(float64(node.Position.Y - goal.Position.Y))
	return xdiff + ydiff
}

func AStar(grid [][]*Node, start *Node, goal *Node, visualise bool) (*arraylist.List, error) {
	// reset all nodes
	for _, col := range grid {
		for _, node := range col {
			node.CameFrom = nil
			node.G = math.MaxInt
			node.F = math.MaxInt
		}
	}

	if visualise {
		print("\033[?25l")        // hide cursor
		print("\x1b[2J\033[0;0H") // clear console and move to 0,0
	}

	openSet := hashset.New(start)

	start.G = 0
	start.F = Heuristic(start, goal, grid)

	openListSize := openSet.Size()
	for openListSize > 0 {
		// find lowest F value
		current := GetLowestFValue(openSet, goal)
		// if goal, return traveled path
		if current == goal {
			if visualise {
				print("\033[?25h") // show cursor
			}
			return ReconstructPath(current), nil
		}

		if visualise {
			print("\033[0;0H")
			PrintHighlighted(grid, start, goal, arraylist.New(openSet.Values()...))
			time.Sleep(time.Millisecond * 20)
		}

		// remove current node from open list
		openSet.Remove(current)

		// get all neighbours
		neighbours := GetNeighbours(current, grid)
		for _, neighbour := range neighbours {
			// cost to go to this neighbour with the current path
			newG := current.G + GetWeight(current, neighbour)
			// check if new path is better
			if newG < neighbour.G {
				// found better option!
				neighbour.CameFrom = current
				neighbour.G = newG
				neighbour.F = newG + Heuristic(neighbour, goal, grid)
				openSet.Add(neighbour)
			}
		}

		openListSize = openSet.Size()
	}

	if visualise {
		print("\033[?25h") // show cursor
	}

	return nil, fmt.Errorf("could not find path")
}
