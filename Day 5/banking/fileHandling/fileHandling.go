package fileHandling

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteValueInFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadDataFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return defaultValue, errors.New("Failed to read data from file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		fmt.Println(err)
		return defaultValue, errors.New("Failed to parse float value: " + valueText)
	}
	return value, nil
}
