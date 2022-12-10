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

	signalStrengths := common.Execute(instructions)

	fmt.Printf("Sum: %v", shared.Sum(signalStrengths))

	return nil
}
