package main

import (
	"encoding/json"
	"fmt"
	"structs_example/author"
	"structs_example/person"
)

type Employee1 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last`
	Age       int    `json:"age"`
}

type employee struct {
	firstName string `json:"first_name"`
	lastName  string `json:"last`
	age       int    `json:"age"`
}

func main() {
	var abc = new(Employee1)
	var abc1 = Employee1{}
	abc.FirstName = "John"
	abc1.FirstName = "John"
	fmt.Println("Start Struct tagging********************************")
	var kumar = Employee1{
		"Vijay",
		"Kumar",
		30,
	}
	fmt.Println(kumar)
	response, err := json.Marshal(kumar)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(response))

	jsonData := `{"first_name":"Vijay","LastName":"Kumar","age":30}`

	var ravi1 Employee1
	json.Unmarshal([]byte(jsonData), &ravi1)

	fmt.Println("struct", ravi1)

	fmt.Println("END ********************************")
	author.Equality()

	newCar := struct {
		make    string
		model   string
		mileage int
	}{
		make:    "Ford",
		model:   "Taurus",
		mileage: 200000,
	}
	fmt.Println(newCar)

	var raj1 = person.New("Raj", "asdf", 19)
	raj1.FullNameWithAge()

	// person.Person.

	raj1.SetFirstName("Raj1").SetLastName("asdf")

	fmt.Println(raj1.Address.City)

	// fmt.Println(errors.New("Raj"))

	fmt.Println("PERSON ********************************")
	var firstName = "first"
	var lastName = "last"
	var age = 25

	var ravi employee
	var raj = employee{}

	displayStructName(ravi)
	fmt.Println("ravi ********************************")

	displayStructName(raj)
	fmt.Println("raj ********************************")

	var ajay = employee{
		firstName: "Ajay",
		age:       30,
	}

	displayStructName(ajay)
	fmt.Println("Ajay ********************************")

	var ajay1 = employee{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}

	displayStructName(ajay1)
	fmt.Println(" ajay1 ********************************")

	var vijay = employee{
		firstName,
		lastName,
		age,
	}

	displayStructName(vijay)
	fmt.Println("vijay ********************************")

	var vijay1 = employee{
		"adf",
		"123",
		22,
	}
	displayStructPointerName(&vijay1)
	vijay1.updateAge(26)
	vijay.fullName()
	displayStructPointerName(&vijay1)
	fmt.Println("vijay1 ********************************")

	displayName(firstName, lastName, age)
	fmt.Println("vijay2 ********************************")
	var vijay2 = &employee{
		"Vijay",
		"Kumar",
		30,
	}

	fmt.Println("FirstName", vijay2.firstName, (*vijay2).firstName)

}

func displayName(firstName string, lastName string, age int) {
	fmt.Printf("My name is %s %s and I am %d years old.\n", firstName, lastName, age)
}

func displayStructName(e employee) {
	fmt.Printf("My name is %s %s and I am %d years old.\n", e.firstName, e.lastName, e.age)
}

func displayStructPointerName(e *employee) {
	//*e.age = e.age
	// (*.e).age
	fmt.Printf("My name is %s %s and I am %d years old.\n", e.firstName, e.lastName, e.age)
}

func (e employee) fullName() {
	var name = fmt.Sprintf("%s %s", e.firstName, e.lastName)
	fmt.Println(name)
}

func (e *employee) updateAge(newAge int) {
	e.age = newAge
}

func newEmployee(firstName string, lastName string, age int) *employee {
	return &employee{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}
