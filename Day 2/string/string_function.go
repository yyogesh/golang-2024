package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, world!, GoLang"

	fmt.Println("Length of the string:", len(str))

	fmt.Println("Contains world ", strings.Contains(str, "world"))

	fmt.Println("Has prefix 'Hello':", strings.HasPrefix(str, "Hello"))

	fmt.Println("Has suffix 'World!':", strings.HasSuffix(str, "world!"))

	fmt.Println("Count of 'o':", strings.Count(str, "o"))

	fmt.Println("Index of 'GoLang':", strings.Index(str, "GoLang"))

	words := []string{"Hello", "GoLang", "World"}
	joined := strings.Join(words, " - ")
	fmt.Println("Joined string:", joined)

	replaced := strings.Replace(str, "world", "Universe", 1)
	fmt.Println("Replaced string:", replaced)

	fmt.Println("Replaced l:", strings.ReplaceAll(str, "l", "z"))

	split := strings.Split(str, ", ")
	fmt.Println("Split string:", split[0], split[1], split[2])

	lower := strings.ToLower(str)
	fmt.Println("Lower case string:", lower)

	upper := strings.ToUpper(str)
	fmt.Println("Upper case string:", upper)

	// user@gmail.com, test@google.com

	email := "user1234@gmail.com"
	var index = strings.Index(email, "@")
	fmt.Println(email[:index])
	fmt.Println(email[index+1:])

}
