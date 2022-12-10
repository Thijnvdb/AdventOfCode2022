package common

import (
	"aoc/shared"
	"strconv"
	"strings"
)

type Command string

const (
	WAIT = "wait"
	ADD  = "add"
)

type Instruction struct {
	Command Command
	Value   int
}

// Does not include noop needed for execution
func NewADDXIntstuction(value int) []Instruction {
	return []Instruction{{Command: WAIT}, {Command: ADD, Value: value}}
}

func NewNOOPInstruction() []Instruction {
	return []Instruction{{Command: WAIT}}
}

func ParseInput(inputFile string) ([]Instruction, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	instructions := []Instruction{}

	for _, line := range lines {
		split := strings.Split(line, " ")
		switch split[0] {
		case "addx":
			val, err := strconv.ParseInt(split[1], 10, 0)
			if err != nil {
				return nil, err
			}

			instructions = append(instructions, NewADDXIntstuction(int(val))...)
		case "noop":
			instructions = append(instructions, NewNOOPInstruction()...)
		}
	}

	return instructions, nil
}
