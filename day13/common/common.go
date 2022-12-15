package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Pair struct {
	Left  Packet
	Right Packet
}

type Side int

const (
	Left  Side = 1
	Right Side = 2
	Same  Side = 3
)

type Packet = []interface{}

func ParsePacket(packetString string) (Packet, error) {
	var res Packet
	err := json.Unmarshal([]byte(packetString), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CheckPacketPair(left Packet, right Packet, printComparisons bool) (bool, error) {
	side, err := GetSmallerSide(left, right, printComparisons)
	return side != Right, err
}

func GetSmallerSide(left Packet, right Packet, printComparisons bool) (Side, error) {
	for i, rightEntry := range right {
		if len(left) == i {
			return Left, nil // left is smaller
		}

		leftEntry := left[i]
		smallest, err := compare(leftEntry, rightEntry, printComparisons)
		if err != nil {
			return smallest, errors.New("compare returned -1")
		}

		if smallest != Same {
			return smallest, nil
		}
	}

	// exited loop, so right is smaller or equal to left
	if len(right) < len(left) {
		return Right, nil
	}

	return Same, nil
}

// returns smaller side
func compare(leftInterface interface{}, rightInterface interface{}, printComparisons bool) (Side, error) {
	if printComparisons {
		fmt.Printf("- Compare %v vs %v\n", leftInterface, rightInterface)
	}

	switch left := leftInterface.(type) {
	case float64:
		switch right := rightInterface.(type) {
		case float64:
			if left == right {
				return Same, nil
			} else if left < right {
				return Left, nil
			} else if right < left {
				return Right, nil
			}
		case Packet:
			comp, err := compare(Packet{left}, right, printComparisons)
			return comp, err
		default:
			return -1, fmt.Errorf("could not parse interface: %v", right)
		}
	case Packet:
		switch right := rightInterface.(type) {
		case float64:
			comp, err := compare(left, Packet{right}, printComparisons)
			return comp, err
		case Packet:
			smal, err := GetSmallerSide(left, right, printComparisons)
			return smal, err
		default:
			return -1, fmt.Errorf("could not parse interface: %v", right)
		}
	default:
		return -1, fmt.Errorf("could not parse interface: %v", left)
	}

	return -1, fmt.Errorf("could not compare")
}
