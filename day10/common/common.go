package common

import (
	"aoc/shared"
	"fmt"
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

func Execute(instructions []Instruction) []int {
	result := []int{}

	reg := 1 // register
	correction := 0
	checkInterval := 20
	for cycle := 1; cycle <= 220; cycle++ {
		instr := instructions[cycle-1]

		if (cycle-correction)%checkInterval == 0 {
			fmt.Printf("Cycle %v, Signal strength: %v\n", cycle, reg*(cycle))
			if checkInterval == 20 {
				checkInterval = 40
				correction = 20
			}
			result = append(result, reg*(cycle))
		}

		switch instr.Command {
		case WAIT:
			//nothing
		case ADD:
			reg += instr.Value
		}
	}

	return result
}

func ExecuteOnCRT(instructions []Instruction) []string {
	reg := 1 // register
	rowSize := 40

	result := []string{}
	row := ""
	for cycle := 0; cycle < len(instructions); cycle++ {
		instr := instructions[cycle]

		if reg == (cycle%rowSize) || reg-1 == (cycle%rowSize) || reg+1 == (cycle%rowSize) {
			row += "#"
		} else {
			row += "."
		}

		switch instr.Command {
		case WAIT:
			//nothing
		case ADD:
			reg += instr.Value
		}

		if cycle > 0 && (cycle+1)%rowSize == 0 {
			// end of row
			result = append(result, row)
			row = ""
		}
	}

	return result
}

func PrintCRT(output []string) {
	for _, row := range output {
		fmt.Println(row)
	}
}
