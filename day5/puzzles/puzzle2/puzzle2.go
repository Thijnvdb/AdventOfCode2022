package puzzle2

import (
	"aoc/day5/common"
	"aoc/shared/types"
	"errors"
	"fmt"
)

func Run(inputFile string) error {
	stateString, instructionsString, err := common.GetSplitInput(inputFile)
	if err != nil {
		return err
	}

	state, err := common.ParseInitialState(stateString)
	if err != nil {
		return err
	}

	fmt.Print("\n\nStart state:\n\n")
	common.PrintState(state)

	instructions, err := common.ParseInstructions(instructionsString)
	if err != nil {
		return err
	}

	endState, err := ExecuteMoves(state, instructions, false)
	if err != nil {
		return err
	}

	fmt.Print("\n\nEnd state:\n\n")
	common.PrintState(endState)

	result := ""
	for _, column := range endState {
		result += column.Peek()
	}

	fmt.Printf("End Result: %v ", result)

	return nil
}

func ExecuteMoves(state []*types.Stack[string], instructions []common.Instruction, printMoves bool) ([]*types.Stack[string], error) {
	moveCount := 0
	for _, instruction := range instructions {
		moveCount++
		intermediateStack := types.NewStack[string]()

		for i := 0; i < instruction.Quantity; i++ {
			boxToMove := state[instruction.From-1].Pop()
			if boxToMove == "" {
				return nil, errors.New("Box was not present!")
			}

			intermediateStack.Push(boxToMove)
		}

		for i := 0; i < instruction.Quantity; i++ {
			boxToMove := intermediateStack.Pop()
			state[instruction.To-1].Push(boxToMove)
		}

		if printMoves {
			fmt.Printf("\n\n===== Move %v =====\n\n", moveCount)
			common.PrintState(state)
		}
	}

	return state, nil
}
