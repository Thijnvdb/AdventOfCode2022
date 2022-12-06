package puzzle2

import (
	"aoc/day6/common"
	"aoc/shared"
	"fmt"
)

func Run(inputFile string) error {
	stream, err := shared.ReadFileAsString(inputFile)
	if err != nil {
		return err
	}

	packetIndex, packetString, err := common.GetDistinctPacketOfSize(stream, 14)

	if err != nil {
		return err
	} else {
		fmt.Printf("Found start of packet string: found packet %v , data starting at index %v", packetString, packetIndex)
		return nil
	}
}
