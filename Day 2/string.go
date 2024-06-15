package main

import (
	"fmt"
	"strings"
)

func main1() {
	str := "Hello, world! world! world!"
	fmt.Println(len(str))

	fmt.Println(strings.Contains(str, " World!"))
	fmt.Println(strings.HasPrefix(str, "Hello"))
	fmt.Println(strings.HasSuffix(str, "Hello"))
	fmt.Println(strings.Count(str, "world!"))
	fmt.Println(str)
	fmt.Println(strings.Index("test@gmail.com", "a"))

	email := "username@example.com"

	atIndex := strings.Index(email, "@")
	fmt.Println(email[:atIndex])
	fmt.Println(email[atIndex+1:])

	words := []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(words, " $$ "))
	fmt.Println(strings.Replace(str, "world", "Hi", 2))
	fmt.Println(strings.Split(str, "world!"))

	a := 10
	b := 20
	result := fmt.Sprintf("The sum of %d and %d is %d", a, b, a+b)
	fmt.Println(result)

	age := 25
	height := 1.75
	name := "Alice"
	info := fmt.Sprintf("%s is %d years old and %.2f meters tall.", name, age, height)
	fmt.Println(info)

	number := 123
	formattedNumber := fmt.Sprintf("%06d", number)
	fmt.Println(formattedNumber) // Output: 000123

	value := 123.456
	formattedValue := fmt.Sprintf("%8.2f", value)
	fmt.Println(formattedValue)

	// +=
	// -=

	// Arthmetics
	// Assignments
	// Comparison
	// Logical

}
