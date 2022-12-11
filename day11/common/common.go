package common

import (
	"aoc/shared"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
)

type Item struct {
	WorryLevel int
}

func NewItem(worryLevel int) *Item {
	item := new(Item)
	item.WorryLevel = worryLevel
	return item
}

type Monkey struct {
	Number                  int
	InspectCount            int
	operationFunctionString string
	OperationFunction       func(itemWorryLevel int) int
	TestValue               int
	TestFunction            func(itemWorryLevel int) int
	testFunctionString      string
	Items                   *linkedlistqueue.Queue
	ifTrue                  int
	ifFalse                 int
}

func NewMonkey(number int, operationFunction func(itemWorryLevel int) int, operationFunctionString string, testValue int, testFunction func(itemWorryLevel int) int, testFunctionString string, items *linkedlistqueue.Queue, ifTrue int, ifFalse int) *Monkey {
	monkey := new(Monkey)
	monkey.InspectCount = 0
	monkey.Number = number
	monkey.testFunctionString = testFunctionString
	monkey.operationFunctionString = operationFunctionString
	monkey.OperationFunction = operationFunction
	monkey.TestFunction = testFunction
	monkey.ifFalse = ifFalse
	monkey.ifTrue = ifTrue
	monkey.Items = items
	monkey.TestValue = testValue
	return monkey
}

func (monkey *Monkey) Inspect(monkeys []*Monkey, relief bool, collectiveProduct int) {
	itemRaw, ok := monkey.Items.Dequeue()
	if !ok {
		return
	}

	monkey.InspectCount++
	item := itemRaw.(*Item)

	newWorryLevel := monkey.OperationFunction(item.WorryLevel)
	for newWorryLevel >= collectiveProduct {
		newWorryLevel -= collectiveProduct
	}

	item.WorryLevel = newWorryLevel

	if relief {
		item.WorryLevel = int(math.Floor(float64(item.WorryLevel) / 3.0))
	}

	newMonkeyIndex := monkey.TestFunction(item.WorryLevel)

	monkeys[newMonkeyIndex].Items.Enqueue(item)
}

func GetMonkeyBusiness(monkeys []*Monkey) int {
	inspectCounts := []int{}

	for _, monkey := range monkeys {
		inspectCounts = append(inspectCounts, monkey.InspectCount)
	}

	fmt.Print("\n")
	for i, count := range inspectCounts {
		fmt.Printf("Monkey %v: inspected %v\n", i, count)
	}

	sort.Ints(inspectCounts)

	return inspectCounts[len(inspectCounts)-2:][0] * inspectCounts[len(inspectCounts)-1:][0]
}

func PrintMonkeys(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		fmt.Printf("\nMonkey: %v\n", monkey.Number)
		fmt.Print("Items: ")
		for _, item := range monkey.Items.Values() {
			fmt.Printf("%v,", item.(*Item).WorryLevel)
		}
		fmt.Print("\n")
		fmt.Printf("InspectCount: %v\n", monkey.Number)
		fmt.Printf("Operation: %v\n", monkey.operationFunctionString)
		fmt.Printf("Test: %v\n  If True: Throw to monkey: %v\n  If False: Throw to monkey: %v\n\n", monkey.testFunctionString, monkey.ifTrue, monkey.ifFalse)
	}
}

func ParseInput(inputFile string) ([]*Monkey, error) {
	str, err := shared.ReadFileAsString(inputFile)
	if err != nil {
		return nil, err
	}

	result := []*Monkey{}

	monkeyStrings := strings.Split(str, "\n\n")
	for i, monkeyString := range monkeyStrings {
		lines := strings.Split(monkeyString, "\n")
		startingItemsStrings := strings.Split(strings.ReplaceAll(strings.ReplaceAll(lines[1], " ", ""), "Startingitems:", ""), ",")
		operationString := strings.ReplaceAll(strings.ReplaceAll(lines[2], "  ", ""), "Operation: ", "")
		TestString := strings.ReplaceAll(strings.ReplaceAll(lines[3], " ", ""), "Test:", "")
		ifTrueString := strings.ReplaceAll(strings.ReplaceAll(lines[4], " ", ""), "Iftrue:throwtomonkey", "")
		ifFalseString := strings.ReplaceAll(strings.ReplaceAll(lines[5], " ", ""), "Iffalse:throwtomonkey", "")

		ifTrue, err := strconv.ParseInt(ifTrueString, 10, 0)
		if err != nil {
			return nil, err
		}

		ifFalse, err := strconv.ParseInt(ifFalseString, 10, 0)
		if err != nil {
			return nil, err
		}

		items := linkedlistqueue.New()
		for _, startingItemString := range startingItemsStrings {
			integer, err := strconv.ParseInt(startingItemString, 10, 0)
			if err != nil {
				return nil, err
			}
			items.Enqueue(NewItem(int(integer)))
		}

		testValue, testFunc, err := ParseTestFunction(TestString, int(ifTrue), int(ifFalse))
		if err != nil {
			return nil, err
		}

		operationFunc, err := ParseOperationFunction(operationString)
		if err != nil {
			return nil, err
		}

		result = append(result, NewMonkey(i, operationFunc, operationString, testValue, testFunc, TestString, items, int(ifTrue), int(ifFalse)))
	}

	return result, nil
}

func ParseTestFunction(testFunctionString string, ifTrue int, ifFalse int) (int, func(itemWorryLevel int) int, error) {
	// should only be divisible by x
	divisibleBy, err := strconv.ParseInt(strings.ReplaceAll(testFunctionString, "divisibleby", ""), 10, 0)
	if err != nil {
		return 0, nil, err
	}

	var test = func(itemWorryLevel int) int {
		if itemWorryLevel%int(divisibleBy) == 0 {
			return ifTrue
		} else {
			return ifFalse
		}
	}

	return int(divisibleBy), test, nil
}

func ParseOperationFunction(operationStrings string) (func(itemWorryLevel int) int, error) {
	mathFunc := strings.ReplaceAll(operationStrings, "new = ", "")
	split := strings.Split(mathFunc, " ")

	if split[2] == "old" {
		switch split[1] {
		case "+":
			return func(itemWorryLevel int) int {
				return itemWorryLevel + itemWorryLevel
			}, nil
		case "*":
			return func(itemWorryLevel int) int {
				return itemWorryLevel * itemWorryLevel
			}, nil
		}
	} else {
		right, err := strconv.ParseInt(split[2], 10, 0)
		if err != nil {
			return nil, err
		}
		switch split[1] {
		case "+":
			return func(itemWorryLevel int) int {
				return itemWorryLevel + int(right)
			}, nil
		case "*":
			return func(itemWorryLevel int) int {
				return itemWorryLevel * int(right)
			}, nil
		}
	}

	return nil, errors.New("did not find operator as an option")
}
