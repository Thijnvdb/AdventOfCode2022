package common

import (
	"aoc/shared"
	"strings"
)

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
	Lose     = "lose"
	Win      = "win"
	Draw     = "draw"
)

// get score of round
func CalculateScore(my_move string, opponent_move string) int {
	// always add score of chosen weapon
	score := ValueMapping[my_move]

	switch my_move {
	case opponent_move:
		score += 3 // tie
	case WeaknessMapping[opponent_move]:
		score += 6 // win
	}

	return score
}

// get input from file
func ReadInput(file_name string) ([][]string, error) {
	lines, err := shared.ReadFile(file_name)
	if err != nil {
		return nil, err
	}

	letters := [][]string{}

	for _, line := range lines {
		split := strings.Split(line, " ")
		letters = append(letters, split)
	}

	return letters, nil
}

// weakness, given weapon
var WeaknessMapping = map[string]string{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

// reversed weakness mapping for easy lookup
var ResistMapping = map[string]string{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

// weapon to value
var ValueMapping = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

// letter to weapon
var LetterMapping = map[string]string{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

// desired outcome, based on input letter
var OutcomeMapping = map[string]string{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}
