package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// type x = {
// 	firstName: string
// }

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate1 := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	// var appUser1 user = user{}

	appUser = user{
		firstName: firstName,
		birthdate: birthdate1,
		lastName:  lastName,
		createdAt: time.Now(),
	}

	// appUser = user{
	// 	firstName: firstName,
	// 	createdAt: time.Now(),
	// }

	// appUser = user{
	// 	firstName,
	// 	lastName,
	// 	birthdate1,
	// 	time.Now(),
	// }

	outputDetails(appUser)

}

func outputDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
