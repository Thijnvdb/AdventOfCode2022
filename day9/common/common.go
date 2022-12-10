package common

import (
	"aoc/shared"
	. "aoc/shared/types"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

type Move struct {
	Direction string
	Size      int
	Steps     []Vector
}

type RopePoint struct {
	Position Vector
	History  []Vector
	parent   *RopePoint
	child    *RopePoint
}

func NewRopePoint() *RopePoint {
	ropePoint := new(RopePoint)
	ropePoint.Position = Vector{X: 0, Y: 0}
	return ropePoint
}

func (rope *RopePoint) GetRopePositions() []Vector {
	if rope.child != nil {
		childPositions := append(rope.child.GetRopePositions(), rope.Position)
		return childPositions
	} else {
		return []Vector{rope.Position}
	}
}

func (rope *RopePoint) getRopeSizeInternal(size int) int {
	if rope.child == nil {
		return size
	} else {
		return rope.child.getRopeSizeInternal(size + 1)
	}
}

func (rope *RopePoint) GetRopeSize() int {
	return rope.getRopeSizeInternal(1)
}

// create rope and return head
func CreateRope(size int) *RopePoint {
	head := NewRopePoint()
	current := head
	for i := 0; i < size-1; i++ {
		child := NewRopePoint()
		current.AddChild(child)
		current = child
	}

	return head
}

func (ropePoint *RopePoint) AddChild(child *RopePoint) {
	child.parent = ropePoint
	ropePoint.child = child
}

func (ropePoint *RopePoint) GetTail() *RopePoint {
	if ropePoint.child == nil {
		return ropePoint
	} else {
		return ropePoint.child.GetTail()
	}
}

func (ropePoint *RopePoint) GetLastPosition() Vector {
	return ropePoint.History[len(ropePoint.History)-1]
}

func (knot *RopePoint) Move(step Vector) {
	knot.History = append(knot.History, knot.Position)
	knot.Position.Add(step)

	if knot.child != nil && GetDistanceBetweenKnots(knot, knot.child) > 1 {
		move := ZeroVector()
		if knot.Position.Y == knot.child.Position.Y {
			// horizontal move
			if knot.Position.X > knot.child.Position.X {
				move.X = 1
			} else {
				move.X = -1
			}
		} else if knot.Position.X == knot.child.Position.X {
			// vertical move
			if knot.Position.Y > knot.child.Position.Y {
				move.Y = 1
			} else {
				move.Y = -1
			}
		} else {
			// diagonal move
			if knot.Position.X > knot.child.Position.X {
				move.X = 1
			} else {
				move.X = -1
			}

			if knot.Position.Y > knot.child.Position.Y {
				move.Y = 1
			} else {
				move.Y = -1
			}
		}
		knot.child.Move(move)
	}
}

func GetDistanceBetweenKnots(a *RopePoint, b *RopePoint) int {
	return GetDistanceBetween(a.Position, b.Position)
}

func SimulateMoves(head *RopePoint, moves []Move, printStates bool) {
	input := bufio.NewScanner(os.Stdin)
	PrintState(head)
	for _, move := range moves {
		fmt.Printf("\n|| Move: %v %v ||\n", move.Size, move.Direction)
		for _, step := range move.Steps {
			head.Move(step)
			if printStates {
				PrintState(head)
				input.Scan()
			}
		}
	}
}

func GetUniquePositionsTraversed(ropePoint *RopePoint) *hashset.Set {
	set := hashset.New()
	for _, point := range ropePoint.History {
		set.Add(point)
	}

	set.Add(ropePoint.Position) // add final position, since it is not yet in history
	return set
}

func ParseInput(input string) ([]Move, error) {
	lines, err := shared.ReadFile(input)
	if err != nil {
		return nil, err
	}

	moves := []Move{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		size, err := strconv.ParseInt(split[1], 10, 0)
		if err != nil {
			return nil, err
		}

		move := Move{Steps: []Vector{}, Direction: split[0], Size: int(size)}
		switch split[0] {
		case "U":
			for i := 0; i < int(size); i++ {
				move.Steps = append(move.Steps, Vector{X: 0, Y: 1})
			}
		case "D":
			for i := 0; i < int(size); i++ {
				move.Steps = append(move.Steps, Vector{X: 0, Y: -1})
			}
		case "L":
			for i := 0; i < int(size); i++ {
				move.Steps = append(move.Steps, Vector{X: -1, Y: 0})
			}
		case "R":
			for i := 0; i < int(size); i++ {
				move.Steps = append(move.Steps, Vector{X: 1, Y: 0})
			}
		}

		moves = append(moves, move)
	}

	return moves, nil
}
