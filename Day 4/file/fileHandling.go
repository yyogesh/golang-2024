package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test1.txt"
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("File not found: ", fileName)
		return
	}
	fmt.Println("File exists", fileInfo)
}
