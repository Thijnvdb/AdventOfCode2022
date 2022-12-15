package puzzle2

import (
	"aoc/day13/common"
	"aoc/shared"
	"fmt"
	"strings"
)

func Run(inputFile string) error {
	packets, err := parseInput(inputFile)
	if err != nil {
		return err
	}

	// manual sort, since go does not allow adding functions on *[]interface{}
	InsertionSort(packets)

	PrintPackets(packets)

	fmt.Printf("Key: %v\n", FindDecoderKey(packets))

	return nil
}

func FindDecoderKey(packets []common.Packet) int {
	key := 0
	for i, packet := range packets {
		if len(packet) > 1 {
			continue
		}

		if len(packet) == 0 {
			continue
		}

		switch t := packet[0].(type) {
		case common.Packet:
			if len(t) != 1 {
				continue
			}
			if len(t) == 1 && t[0] == float64(6) || t[0] == float64(2) {
				if key == 0 {
					key = i + 1
				} else {
					key *= i + 1
				}
			}
		}
	}

	return key
}

func PrintPackets(packets []common.Packet) {
	fmt.Println()
	for _, packet := range packets {
		fmt.Println(packet)
	}
}

func InsertionSort(packets []common.Packet) {
	for i := 0; i < len(packets); i++ {
		j := i
		for j > 0 && LessThan(packets[j], packets[j-1]) {
			temp := packets[j]
			packets[j] = packets[j-1]
			packets[j-1] = temp
			j--
		}
	}
}

// wont work and it is 1 am I quit efficienty is boring and for nerds
func QuickSort(packets []common.Packet, low int, high int) {
	if low >= high || low < 0 {
		return
	}

	pivot := partition(packets, low, high)
	QuickSort(packets, low, pivot-1)
	QuickSort(packets, pivot+1, high)
}

func partition(packets []common.Packet, low int, high int) int {
	pivot := packets[high]
	left := low
	right := high - 1

	for left < right {
		for LessThan(packets[left], pivot) {
			left++
		}
		for GreaterThan(packets[right], pivot) {
			right--
		}

		if left >= right {
			break
		}

		temp := packets[left]
		packets[left] = packets[right]
		packets[right] = temp
	}

	temp := packets[left]
	packets[left] = packets[high]
	packets[high] = temp
	return left
}

func LessThan(a common.Packet, b common.Packet) bool {
	check, err := common.GetSmallerSide(a, b, false)
	if err != nil {
		panic(fmt.Errorf("oops: %v", err.Error()))
	}
	return check == common.Left
}

func GreaterThan(a common.Packet, b common.Packet) bool {
	check, err := common.GetSmallerSide(a, b, false)
	if err != nil {
		panic(fmt.Errorf("oops: %v", err.Error()))
	}
	return check == common.Right
}

func LessThanOrEqual(a common.Packet, b common.Packet) bool {
	check, err := common.GetSmallerSide(a, b, false)
	if err != nil {
		panic(fmt.Errorf("oops: %v", err.Error()))
	}
	return check == common.Left || check == common.Same
}

func parseInput(inputFile string) ([]common.Packet, error) {
	str, err := shared.ReadFileAsString(inputFile)
	if err != nil {
		return nil, err
	}
	str += "\n[[2]]\n[[6]]"
	return ParseString(str)
}

func ParseString(inputString string) ([]common.Packet, error) {
	packets := []common.Packet{}
	noDoubleNewLine := strings.ReplaceAll(inputString, "\n\n", "\n")
	packetsRaw := strings.Split(noDoubleNewLine, "\n")
	for _, packetRaw := range packetsRaw {
		packet, err := common.ParsePacket(packetRaw)
		if err != nil {
			return nil, err
		}

		packets = append(packets, packet)
	}

	return packets, nil
}
