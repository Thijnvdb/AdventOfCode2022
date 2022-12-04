package common

import (
	"aoc/shared"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Index    int
	Meals    []int
	Calories int
}

func GetElfsSortedByCalories(inputFile string) ([]Elf, error) {
	content, err := shared.ReadFileAsString(inputFile)
	if err != nil {
		return nil, err
	}

	elfCalorieGroups := strings.Split(content, "\n\n")

	elfs := []Elf{}

	for elfNumber, elfMeals := range elfCalorieGroups {
		meals := []int{}
		calories := 0
		for _, meal := range strings.Split(elfMeals, "\n") {
			if meal == "" {
				continue
			}
			mealNum, err := strconv.ParseInt(meal, 10, 0)
			if err != nil {
				return nil, err
			}

			meals = append(meals, int(mealNum))
			calories += int(mealNum)
		}

		elf := Elf{
			Index:    elfNumber,
			Meals:    meals,
			Calories: calories,
		}

		elfs = append(elfs, elf)
	}

	sort.Slice(elfs, func(a, b int) bool {
		return elfs[a].Calories > elfs[b].Calories
	})

	return elfs, nil
}
