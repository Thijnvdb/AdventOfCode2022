package common

import (
	"aoc/shared"
	"errors"
	"strconv"
	"strings"
)

type Task struct {
	Sections []int
}

func ParseInput(inputFile string) ([][]Task, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	pairs := [][]Task{}

	for _, line := range lines {
		split := strings.Split(line, ",")

		tasks := []Task{}

		for _, task := range split {
			fromTo := strings.Split(task, "-")

			left, leftErr := strconv.ParseInt(fromTo[0], 10, 0)
			if leftErr != nil {
				return nil, leftErr
			}

			right, rightErr := strconv.ParseInt(fromTo[1], 10, 0)
			if rightErr != nil {
				return nil, rightErr
			}

			if left > right {
				return nil, errors.New("Section notation was not correct")
			}

			sections := []int{}
			for i := left; i <= right; i++ {
				sections = append(sections, int(i))
			}

			tasks = append(tasks, Task{Sections: sections})
		}

		pairs = append(pairs, tasks)
	}

	return pairs, nil
}
