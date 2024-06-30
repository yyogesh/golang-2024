package person

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	FirstName string
	LastName  string
	age       int
	Address
}

func (a *Address) FullAddress() string {
	return a.City + " " + a.State
}

func New(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		age:       age,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}
}

func (p *Person) FullNameWithAge() string {
	fmt.Println("FullName :: ", p.FirstName, p.LastName, p.age)
	return p.FirstName + " " + p.LastName
}

// ************ chaining ********************
func (p *Person) SetFirstName(firstName string) *Person {
	p.FirstName = firstName
	return p
}

func (p *Person) SetLastName(lastName string) *Person {
	p.LastName = lastName
	return p
}

func (p *Person) SetAge(age int) *Person {
	p.age = age
	return p
}
