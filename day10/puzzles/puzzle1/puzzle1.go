package puzzle1

import (
	"aoc/day10/common"
	"aoc/shared"
	"fmt"
)

func Run(inputFile string) error {
	instructions, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	signalStrengths := Execute(instructions)

	fmt.Printf("Sum: %v", shared.Sum(signalStrengths))

	return nil
}

func Execute(instructions []common.Instruction) []int {
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
		case common.WAIT:
			//nothing
		case common.ADD:
			reg += instr.Value
		}
	}

	return result
}
