package common

import (
	"aoc/shared"
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(inputFile string) ([]Command, error) {
	lines, err := shared.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	commands := []Command{}
	command := Command{}
	for _, line := range lines {
		if string(line[0]) == "$" {
			// new command
			if command.Text != nil {
				commands = append(commands, command)
			}

			command = Exec(strings.Split(line, " ")[1:], []string{})
		} else {
			// new output
			command.Output = append(command.Output, line)
		}
	}

	// add final command
	commands = append(commands, command)

	return commands, nil
}

// returns root of filesystem
func ConstructFileSystem(commands []Command) (*Directory, error) {
	root := Mkdir("/", nil)

	currentDir := root
	for _, command := range commands {
		switch command.Text[0] {
		case "cd":
			// change directory
			newDir, err := cd(currentDir, command, root)
			if err != nil {
				return root, err
			}

			currentDir = newDir
		case "ls":
			ls(currentDir, command)
		}
	}

	return root, nil
}

// also create newly discovered files/directories
func ls(currentDir *Directory, command Command) error {
	for _, line := range command.Output {
		split := strings.Split(line, " ")

		if split[0] == "dir" {
			// new directory
			newDir := Mkdir(split[1], currentDir)
			currentDir.SubDirectories[newDir.Name] = newDir
		} else {
			// new file
			size, err := strconv.ParseInt(split[0], 10, 0)
			if err != nil {
				return err
			}

			currentDir.AddFile(Touch(split[1], int(size)))
		}
	}

	return nil
}

func cd(currentDir *Directory, command Command, root *Directory) (*Directory, error) {
	cdToDir := command.Text[1]

	if len(cdToDir) > 1 && cdToDir[:2] == ".." {
		return currentDir.parent, nil
	}

	// start from root
	if string(cdToDir[0]) == "/" {
		currentDir = root
		if len(cdToDir) == 1 {
			return currentDir, nil
		} else {
			cdToDir = cdToDir[1:]
		}
	}

	value, exists := currentDir.SubDirectories[cdToDir]
	if !exists {
		return currentDir, fmt.Errorf("Directory does not exist: %v", cdToDir)
	}

	return value, nil
}

func PrintFileSystem(root *Directory) {
	printDir(root, 0)
}

func printDir(currentDir *Directory, depth int) {
	tabs := ""
	for i := 0; i < depth; i++ {
		tabs += " "
	}
	fmt.Printf("%v / %v (%v) \n", tabs, currentDir.Name, currentDir.Size)
	for _, dir := range currentDir.SubDirectories {
		printDir(dir, depth+1)
	}

	for _, file := range currentDir.Files {
		fmt.Printf("  %v - %v (%v)\n", tabs, file.Name, file.Size)
	}
}
