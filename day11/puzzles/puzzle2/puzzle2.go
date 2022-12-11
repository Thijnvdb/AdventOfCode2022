package puzzle2

import (
	"aoc/day11/common"
	"fmt"
)

func Run(inputFile string) error {
	monkeys, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	SimulateRounds(10000, monkeys, true)

	fmt.Printf("monkey business: %v \n", common.GetMonkeyBusiness(monkeys))
	return nil
}

func SimulateRounds(roundCount int, monkeys []*common.Monkey, printResults bool) {
	collectiveProduct := monkeys[0].TestValue
	for _, monkey := range monkeys[1:] {
		collectiveProduct *= monkey.TestValue
	}
	for round := 0; round < roundCount; round++ {
		for _, monkey := range monkeys {
			items := monkey.Items.Size()
			for i := 0; i < items; i++ {
				monkey.Inspect(monkeys, false, collectiveProduct)
			}
		}

		if printResults && (round == 0 || round == 19 || (round+1)%1000 == 0) {
			fmt.Printf("== After round %v ==\n", round+1)
			for _, monkey := range monkeys {
				fmt.Printf("Monkey %v inspected items %v times.\n", monkey.Number, monkey.InspectCount)
			}
		}
	}
}
