package shared

import (
	"bufio"
	"fmt"
	"os"
)

// Read file string array (each entry containing 1 line)
func ReadFile(file_path string) ([]string, error) {
	file, err := os.Open(file_path)
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

// Read file string array (each entry containing 1 line)
func ReadFileAsString(file_path string) (string, error) {
	file, err := os.ReadFile(file_path)
	if err != nil {
		return "", fmt.Errorf("An error occurred while Reading the file: %v", err.Error())
	}

	return string(file), nil
}
