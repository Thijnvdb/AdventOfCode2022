package puzzle2

import (
	"aoc/day13/puzzles/puzzle2"
	"fmt"
	"testing"
)

func TestIsLessThanScenario1(t *testing.T) {
	packets, err := puzzle2.ParseString("[[2]]\n[3]")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("packets: %v", packets)

	if !puzzle2.LessThan(packets[0], packets[1]) {
		t.Fail()
	}
}

func TestIsLessThanScenario2(t *testing.T) {
	packets, err := puzzle2.ParseString("[[]]\n[]\n[[2]]\n[3]")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("packets: %v\n", packets)

	puzzle2.PrintPackets(packets)

	if puzzle2.LessThan(packets[0], packets[1]) {
		t.Fail()
	}

	puzzle2.PrintPackets(packets)
}
