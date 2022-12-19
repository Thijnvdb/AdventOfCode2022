package puzzle1

import (
	"aoc/day15/common"
	"fmt"
)

func Run(inputFile string) error {
	sensors, err := common.ParseInput(inputFile)
	if err != nil {
		return err
	}

	fmt.Println(sensors)
	return nil
}

// determine places where a beacon cannot be present in a given row
// returns X values of the positions which cannot be present
func getBeaconNotPresentForRow(sensors []*common.Sensor) []int {
	return nil
}
