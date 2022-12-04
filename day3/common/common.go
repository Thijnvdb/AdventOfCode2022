package common

import "errors"

const alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// split backpack in equal parts
func SplitBackpack(backpack string) []string {
	left := backpack[0:(len(backpack) / 2)]
	right := backpack[(len(backpack) / 2):]

	return []string{left, right}
}

// get priority of character
func GetPriority(letter string) (int, error) {
	if len(letter) != 1 {
		return 0, errors.New("Input was not a letter")
	}

	for i, l := range alphabet {
		if string(l) == letter {
			return i + 1, nil
		}
	}

	return 0, errors.New("No match found")
}

// get matching characters
func GetMatch(base string, compare string) (string, error) {
	for _, first := range base {
		for _, second := range compare {
			if first == second {
				return string(first), nil
			}
		}
	}

	return "", errors.New("No matching characters found")
}

// get matching characters
func GetMatchInGroup(base string, compare string, compareAlso string) (string, error) {
	for _, first := range base {
		for _, second := range compare {
			if first == second && isInString(compareAlso, string(first)) {
				return string(first), nil
			}
		}
	}

	return "", errors.New("No matching characters found")
}

// check if char is in string
func isInString(base string, letter string) bool {
	for _, v := range base {
		if letter == string(v) {
			return true
		}
	}

	return false
}
