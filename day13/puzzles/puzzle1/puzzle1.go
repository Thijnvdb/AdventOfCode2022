package puzzle1

import (
	"aoc/day13/common"
	"aoc/shared"
	"aoc/shared/colors"
	"fmt"
	"strings"
)

func Run(inputFile string) error {
	pairs, err := parseInput(inputFile)
	if err != nil {
		return err
	}

	sum := 0
	for i, pair := range pairs {
		fmt.Printf("== Pair %v ==\n", i+1)
		fmt.Printf("- Compare %v vs %v\n", pair.Left, pair.Right)
		check, err := common.CheckPacketPair(pair.Left, pair.Right, true)
		if err != nil {
			return err
		}

		if check {
			sum += i + 1
			fmt.Printf(colors.Green("Pair %v is in the right order\n"), i+1)
		} else {
			fmt.Printf(colors.Red("Pair %v is NOT in the right order\n"), i+1)
		}
	}

	fmt.Printf(colors.Blue("\n\n Sum of indices: %v"), sum)

	return nil
}

func parseInput(inputFile string) ([]common.Pair, error) {
	str, err := shared.ReadFileAsString(inputFile)
	if err != nil {
		return nil, err
	}

	return ParseString(str)
}

func ParseString(inputString string) ([]common.Pair, error) {
	pairs := []common.Pair{}
	pairsRaw := strings.Split(inputString, "\n\n")
	for _, pairRaw := range pairsRaw {
		halves := strings.Split(pairRaw, "\n")

		pair := common.Pair{}
		for _, halve := range halves {
			packet, err := common.ParsePacket(halve)
			if err != nil {
				return nil, err
			}

			if pair.Left == nil {
				pair.Left = packet
			} else {
				pair.Right = packet
			}
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}
