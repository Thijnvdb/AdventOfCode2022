package puzzle2

import (
	. "day2/common"
	"errors"
)

// move to play
func get_move(opponent_move string, outcome string) (string, error) {
	// make same move to draw
	switch outcome {
	case Draw:
		return opponent_move, nil
	case Win:
		return Get_weakness(opponent_move), nil
	case Lose:
		return Get_resist(opponent_move), nil
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
		opponent_move := Get_weapon(round[0])

		// desired outcome
		outcome := Get_outcome(round[1])

		my_move, err := get_move(opponent_move, outcome)
		if err != nil {
			return 0, errors.New("error while getting move")
		}

		score += Calculate_score(my_move, opponent_move)
	}

	return score, nil
}

func Run(input_file string) {
	rounds := Read_input(input_file)
	result, err := simulate_game(rounds)
	if err != nil {
		print(err)
	} else {
		print(result)
	}
}
