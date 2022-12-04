package puzzle1

import (
	common "aoc/day2/common"
)

// simulate game, return the score
func simulate_game(rounds [][]string) int {
	score := 0

	for i := 0; i < len(rounds); i++ {
		moves := rounds[i]
		if len(moves) != 2 {
			continue
		}

		my_move := common.LetterMapping[moves[1]]
		opponent_move := common.LetterMapping[moves[0]]

		score += common.CalculateScore(my_move, opponent_move)
	}

	return score
}

func Run(inputFile string) error {
	rounds, err := common.ReadInput(inputFile)
	if err != nil {
		return err
	}

	print(simulate_game(rounds))
	return nil
}
