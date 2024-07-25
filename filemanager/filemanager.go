package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath) // Replace with actual file reading logic

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

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err!= nil {
        return errors.New("Failed to create file: " + err.Error())
    }

	err = json.NewEncoder(file).Encode(data)

	if err!= nil {
        file.Close()
        return errors.New("Failed to write to file: " + err.Error())
    }

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		inputPath,
		outputPath,
	}
}