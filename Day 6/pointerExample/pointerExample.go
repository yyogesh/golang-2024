package main

import "fmt"

func main() {
	// Defining variables to store values of type string, int, boo, and float64
	var username string = "Go developer"
	var age int = 12
	var isAwesome = true
	var height = 5.5

	// Declaring pointers of types :string, int, bool and float64 using * operator
	// Pointer to a string
	var usernamePointer *string = &username
	// Pointer to an integer
	var agePointer *int = &age
	// Pointer to a boolean
	var isAwesomePointer *bool = &isAwesome
	// Pointer to a floating number
	var heightPointer *float64 = &height

	fmt.Printf("usernamePointer : %v ===> username value: %v \n", usernamePointer, *usernamePointer)
	fmt.Printf("agePointer: %v  ===> age value: %v \n", agePointer, *agePointer)
	fmt.Printf("isAwesomePointer: %v  ===> isAwesome value: %v \n", isAwesomePointer, *isAwesomePointer)
	fmt.Printf("heightPointer : %v  ===> height value: %v \n", heightPointer, *heightPointer)

	var usernamePointer1 *string

	if usernamePointer1 == nil {
		fmt.Println("username pointer is nil")
		return
	}

	fmt.Printf("usernamePointer : %v ===> username value: %v \n", usernamePointer1, *usernamePointer1)

	x := new(string)
	fmt.Printf("usernamePointer : %v ===> username value: %v \n", x, *x)

	// Declare an integer variable
	var number int = 10

	// Print the original value
	fmt.Println("Original value:", number)

	// Call the increment function, passing the address of the variable
	increment(&number)

	// Print the modified value
	fmt.Println("Value after increment:", number)
}

// This function takes a pointer to an int
func increment(value *int) {
	// Dereference the pointer and increment the value
	*value++
}
