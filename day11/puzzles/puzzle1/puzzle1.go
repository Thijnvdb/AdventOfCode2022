package puzzle1

import (
	"aoc/day11/common"
	"fmt"
)

func Run(inputFile string) error {
	monkeys, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	SimulateRounds(20, monkeys, false)

	fmt.Printf("monkey business: %v \n", common.GetMonkeyBusiness(monkeys))
	return nil
}

func SimulateRounds(roundCount int, monkeys []*common.Monkey, printResults bool) {
	collectiveProduct := monkeys[0].TestValue
	for _, monkey := range monkeys[1:] {
		collectiveProduct *= monkey.TestValue
	}

	for round := 0; round < roundCount; round++ {
		if printResults {
			fmt.Printf("Round %v:\n", round)
		}
		for _, monkey := range monkeys {
			if printResults {
				fmt.Printf("  Monkey %v:\n", monkey.Number)
			}
			items := monkey.Items.Size()
			for i := 0; i < items; i++ {
				monkey.Inspect(monkeys, true, collectiveProduct)
			}
		}
	}
}
