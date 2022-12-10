package puzzle2

import (
	"aoc/day10/common"
)

func Run(inputFile string) error {
	instructions, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	crtOut := common.ExecuteOnCRT(instructions)

	common.PrintCRT(crtOut)

	return nil
}
