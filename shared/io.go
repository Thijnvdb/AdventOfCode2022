package shared

import (
	"bufio"
	"fmt"
	"os"
)

// Read file string array (each entry containing 1 line)
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("An error occurred while Reading the file: %v", err.Error())
	}

	//close file at end of program
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			lines = append(lines, text)
		}
	}

	return lines, nil
}

// Read file as a string
func ReadFileAsString(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("An error occurred while Reading the file: %v", err.Error())
	}

	return string(file), nil
}
