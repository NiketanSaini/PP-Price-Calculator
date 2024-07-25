package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path) // Replace with actual file reading logic

	if err != nil {
		return nil, errors.New("Failed to open file: " + err.Error())
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read file: " + err.Error())
	}

	file.Close()
	return lines, nil
}