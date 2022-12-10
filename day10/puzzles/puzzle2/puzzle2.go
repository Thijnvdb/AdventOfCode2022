package puzzle2

import (
	"aoc/day10/common"
	"fmt"
)

func Run(inputFile string) error {
	instructions, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	crtOut := ExecuteOnCRT(instructions)

	PrintCRT(crtOut)

	return nil
}

func ExecuteOnCRT(instructions []common.Instruction) []string {
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
		case common.WAIT:
			//nothing
		case common.ADD:
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
