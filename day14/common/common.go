package common

import (
	"aoc/shared"
	"strings"
)

type State int

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

type Cave = []*Point

func ParseInput(inputFile string) (Cave, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	rangeX := []int{9999, 0}
	rangeY := []int{9999, 0}
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		for _, coord := range coords {
			xy := strings.Split(coord, ",")
		}
	}

	return nil
}
