package common

import (
	"bufio"
	"log"
	"os"
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
func Calculate_score(my_move string, opponent_move string) int {
	// always add score of chosen weapon
	score := Get_value(my_move)

	switch my_move {
	case opponent_move:
		score += 3 // tie
	case Get_weakness(opponent_move):
		score += 6 // win
	}

	return score
}

// get input from file
func Read_input(file_name string) [][]string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	//close file at end of program
	defer file.Close()

	lines := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		lines = append(lines, split)
	}

	return lines
}

// weakness, given weapon
var weakness_mapping = map[string]string{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

// reversed weakness mapping for easy lookup
var resist_mapping = map[string]string{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

// weapon to value
var value_mapping = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

// letter to weapon
var letter_mapping = map[string]string{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

// desired outcome, based on input letter
var outcome_mapping = map[string]string{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

func Get_weakness(weapon string) string {
	return weakness_mapping[weapon]
}

func Get_resist(weapon string) string {
	return resist_mapping[weapon]
}

func Get_value(weapon string) int {
	return value_mapping[weapon]
}

func Get_weapon(letter string) string {
	return letter_mapping[letter]
}

func Get_outcome(letter string) string {
	return outcome_mapping[letter]
}
