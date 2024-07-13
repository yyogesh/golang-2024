package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, errors.New("error opening file")
	}
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return errors.New("error creating file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		file.Close()
		return errors.New("faild to convert data to json")
	}

	defer file.Close()
	return nil
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{InputFilePath: inputFilePath, OutputFilePath: outputFilePath}
}
