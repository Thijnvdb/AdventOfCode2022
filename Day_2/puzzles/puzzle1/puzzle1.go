package puzzle1

import (
	. "day2/common"
)

// simulate game, return the score
func simulate_game(rounds [][]string) int {
	score := 0

	for i := 0; i < len(rounds); i++ {
		moves := rounds[i]
		if len(moves) != 2 {
			continue
		}

		my_move := Get_weapon(moves[1])
		opponent_move := Get_weapon(moves[0])

		score += Calculate_score(my_move, opponent_move)
	}

	return score
}

func Run(input_file string) {
	rounds := Read_input(input_file)
	print(simulate_game(rounds))
}
