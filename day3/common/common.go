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

// get matching characters (in O(n) 🥳)
func GetMatch(base string, compare string) (string, error) {
	mappy := map[string]int{}

	for _, v := range base {
		if mappy[string(v)] == 0 {
			mappy[string(v)] = 1
		}
	}

	for _, v := range compare {
		if mappy[string(v)] == 1 {
			return string(v), nil
		}
	}

	return "", errors.New("No matching characters found")
}

// get matching characters in group of 3 (in O(n) 🥳)
func GetMatchInGroup(base string, compare string, compareAlso string) (string, error) {
	mappy := map[string]int{}

	for _, v := range base {
		if mappy[string(v)] == 0 {
			mappy[string(v)] = 1
		}
	}

	for _, v := range compare {
		// only care about what was already present, so do not set any other value
		if mappy[string(v)] == 1 {
			mappy[string(v)] = 2
		}
	}

	for _, v := range compareAlso {
		if mappy[string(v)] == 2 {
			return string(v), nil
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
