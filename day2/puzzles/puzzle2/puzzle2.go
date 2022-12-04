package puzzle2

import (
	"aoc/day2/common"
	"errors"
)

// move to play
func get_move(opponent_move string, outcome string) (string, error) {
	// make same move to draw
	switch outcome {
	case common.Draw:
		return opponent_move, nil
	case common.Win:
		return common.WeaknessMapping[opponent_move], nil
	case common.Lose:
		return common.ResistMapping[opponent_move], nil
	default:
		return "", errors.New("Outcome is not a valid value") // not possible
	}
}

// simulate game, return the score
func simulate_game(rounds [][]string) (int, error) {
	score := 0
	for i := 0; i < len(rounds); i++ {
		round := rounds[i]
		if len(round) != 2 {
			continue
		}

		// move of opponent
		opponent_move := common.LetterMapping[round[0]]

		// desired outcome
		outcome := common.OutcomeMapping[round[1]]

		my_move, err := get_move(opponent_move, outcome)
		if err != nil {
			return 0, errors.New("error while getting move")
		}

		score += common.CalculateScore(my_move, opponent_move)
	}

	return score, nil
}

func Run(input_file string) error {
	rounds, err := common.ReadInput(input_file)
	result, err := simulate_game(rounds)
	if err != nil {
		return err
	}
	print(result)
	return nil
}
