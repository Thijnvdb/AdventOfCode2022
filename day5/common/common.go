package common

import (
	"aoc/shared"
	"aoc/shared/types"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	Quantity int
	From     int
	To       int
}

// return 2 parts of the input seperated
func GetSplitInput(inputFile string) (string, string, error) {
	inputString, err := shared.ReadFileAsString(inputFile)
	if err != nil {
		return "", "", err
	}

	//split in 2 halves: state & instructions
	split := strings.Split(inputString, "\n\n")
	if len(split) != 2 {
		return "", "", errors.New("Input was not correct (split resulted in more than 2 halves)")
	}
	return split[0], split[1], nil
}

func ParseInitialState(initialStateString string) ([]*types.Stack[string], error) {
	lines := strings.Split(initialStateString, "\n")

	columns := []*types.Stack[string]{}

	// step to take to get to the next item
	fourth := int(math.Ceil(float64(len(lines[0])-1) / 4))

	for i := 0; i < int(math.Ceil(float64(len(lines[0])-1)/4)); i++ {
		columns = append(columns, types.NewStack[string]())
	}

	for y := len(lines) - 2; y >= 0; y-- {
		line := lines[y]
		println(line)
		for i := 0; i < fourth; i++ {
			stringIndex := i*4 + 1 // index of the character we need to read
			character := string(line[stringIndex])
			if len(strings.ReplaceAll(character, " ", "")) > 0 {
				columns[i].Push(character)
			}
		}
	}

	return columns, nil
}

func ParseInstructions(instructionsString string) ([]Instruction, error) {
	lines := strings.Split(instructionsString, "\n")

	instructions := []Instruction{}

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		line := strings.ReplaceAll(l, " ", "")   // remove spaces
		split := strings.Split(line[4:], "from") // skipping 'move'

		quantity, err := strconv.ParseInt(split[0], 10, 0)
		if err != nil {
			return nil, err
		}

		fromTo := strings.Split(split[1], "to")
		from, err := strconv.ParseInt(fromTo[0], 10, 0)
		if err != nil {
			return nil, err
		}

		to, err := strconv.ParseInt(fromTo[1], 10, 0)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, Instruction{Quantity: int(quantity), From: int(from), To: int(to)})
	}

	return instructions, nil
}

func PrintState(state []*types.Stack[string]) {
	max := 0
	items := [][]string{}
	for _, stack := range state {
		vals := stack.Values()
		if len(vals) > max {
			max = len(vals)
		}

		items = append(items, vals)
	}

	for y := max; y >= 0; y-- {
		for x := 0; x < len(items); x++ {
			list := items[x]
			if y >= len(list) {
				fmt.Print("     ")
				continue
			}
			fmt.Printf(" [%v] ", list[y])
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
	for i := range state {
		fmt.Printf(" (%v) ", i)
	}
}
