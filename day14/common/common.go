package common

import (
	"aoc/shared"
	"aoc/shared/printing"
	"strconv"
	"strings"
	"time"
)

type State int

type Structure struct {
	Coordinates [][]int
}

const (
	Air  State = 0
	Rock State = 1
	Sand State = 2
)

type Point struct {
	X     int
	Y     int
	State State
}

func NewPoint(x int, y int) *Point {
	point := new(Point)
	point.X = x
	point.Y = y
	point.State = Air
	return point
}

type Cave = [][]*Point

// returns cave, sand start point and possible error
func ParseInput(inputFile string) (Cave, *Point, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, nil, err
	}

	maxY := 0

	structures := []Structure{}
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		coordinates := [][]int{}
		for _, coord := range coords {
			xy := strings.Split(coord, ",")
			x, errx := strconv.ParseInt(xy[0], 10, 0)
			y, erry := strconv.ParseInt(xy[1], 10, 0)
			if errx != nil {
				return nil, nil, errx
			}

			if erry != nil {
				return nil, nil, erry
			}

			if int(y) > maxY {
				maxY = int(y)
			}

			coordinates = append(coordinates, []int{int(x) - 500, int(y)})
		}

		structures = append(structures, Structure{Coordinates: coordinates})
	}

	maxX := maxY * 8
	offset := maxX / 2

	cave := Cave{}

	// make cave
	for x := 0; x < maxX+1; x++ {
		col := []*Point{}
		for y := 0; y <= maxY+2; y++ {
			col = append(col, NewPoint(x, y))
		}
		cave = append(cave, col)
	}

	setStructures(cave, structures, offset)

	return cave, cave[offset][0], nil
}

func setStructures(cave Cave, structures []Structure, offset int) {
	for _, structure := range structures {
		for i, current := range structure.Coordinates {
			if i == len(structure.Coordinates)-1 {
				continue
			}

			next := structure.Coordinates[i+1]

			if next[0] == current[0] {
				// move vertical
				if next[1] > current[1] {
					// move down (higher is down)
					for y := current[1]; y <= next[1]; y++ {
						cave[offset+current[0]][y].State = Rock
					}
				} else {
					for y := current[1]; y >= next[1]; y-- {
						cave[offset+current[0]][y].State = Rock
					}
				}
			} else {
				// move horizontal
				if next[0] > current[0] {
					// right
					for x := current[0]; x <= next[0]; x++ {
						cave[offset+x][current[1]].State = Rock
					}
				} else {
					// left
					for x := current[0]; x >= next[0]; x-- {
						cave[offset+x][current[1]].State = Rock
					}
				}
			}
		}
	}
}

// returns path of the grain and wether it fell into the void
func FloodStep(cave Cave, current *Point, path []*Point, print bool) ([]*Point, bool) {
	if print {
		printing.ResetScreen()
		PrintCave(cave)
		time.Sleep(time.Millisecond * 100)
	}

	if current.Y >= len(cave[0])-1 {
		return path, true
	}

	// if we can go down, go down

	if cave[current.X][current.Y+1].State == Air {
		cave[current.X][current.Y].State = Air
		cave[current.X][current.Y+1].State = Sand
		return FloodStep(cave, cave[current.X][current.Y+1], path, print)
	}

	// bottom left...

	if current.X-1 < 0 {
		// fell in abyss
		return path, true
	}

	if cave[current.X-1][current.Y+1].State == Air {
		cave[current.X][current.Y].State = Air
		cave[current.X-1][current.Y+1].State = Sand
		return FloodStep(cave, cave[current.X-1][current.Y+1], path, print)
	}

	// bottom right...

	if current.X+1 >= len(cave) {
		// fell in abyss
		return path, true
	}

	if cave[current.X+1][current.Y+1].State == Air {
		cave[current.X][current.Y].State = Air
		cave[current.X+1][current.Y+1].State = Sand
		return FloodStep(cave, cave[current.X+1][current.Y+1], path, print)
	}

	// came to rest
	return path, false
}

func PrintCave(cave Cave) {
	for y := 0; y < len(cave[0]); y++ {
		for x := 0; x < len(cave); x++ {
			point := cave[x][y]

			switch point.State {
			case Air:
				print(".")
			case Rock:
				print("#")
			case Sand:
				print("o")
			}
		}

		println()
	}
}
