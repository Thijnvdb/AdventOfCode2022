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
	mappy := map[rune]int{} // rune should make insert more gooder, since no hash conflict resolution but idk

	for _, v := range base {
		if mappy[v] == 0 {
			mappy[v] = 1
		}
	}

	for _, v := range compare {
		if mappy[v] == 1 {
			return string(v), nil
		}
	}

	return "", errors.New("No matching characters found")
}

// get matching characters in group of 3
func GetMatchInGroup(base string, compare string, compareAlso string) (string, error) {
	mappy := map[rune]int{} // rune should make insert more gooder, since no hash conflict resolution but idk

	for _, v := range base {
		if mappy[v] == 0 {
			mappy[v] = 1
		}
	}

	for _, v := range compare {
		// only care about what was already present, so do not set any other value
		if mappy[v] == 1 {
			mappy[v] = 2
		}
	}

	for _, v := range compareAlso {
		if mappy[v] == 2 {
			return string(v), nil
		}
	}

	return "", errors.New("No matching characters found")
}
