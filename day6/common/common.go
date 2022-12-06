package common

import (
	"errors"
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

// get first index after distinct packet of a specified size and the packet
func GetDistinctPacketOfSize(stream string, size int) (int, string, error) {
	if (len(stream)) < size {
		return 0, "", errors.New("stream was too short for packet size")
	}
	for i := 0; i < len(stream)-size; i++ {
		characters := stream[i : i+size]
		set := hashset.New()
		for _, char := range characters {
			set.Add(char)
		}

		if set.Size() == size {
			return i + size, characters, nil
		}
	}

	return 0, "", fmt.Errorf("no distinct packet found of size %v ", size)
}
