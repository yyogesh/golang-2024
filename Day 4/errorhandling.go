package main

import (
	"errors"
	"fmt"
	"os"
)

func foo() error {
	return errors.New("Error: foo")
}

func divide(a, b int) (int, error) {
	// if b == 0 {
	// 	return 0, fmt.Errorf("canot divide by zero")
	// }
	return a / b, nil
}

func main() {
	foo()
	result, err := divide(1, 2)
	fmt.Println(result, err)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	content, err1 := os.ReadFile("abc.txt")
	// data := []byte("asdf asdf")
	// os.WriteFile("test.text", data, 0777)
	if err1 != nil {
		fmt.Println("Error: ", err1)
		return
	}

	fmt.Println(string(content))

}
