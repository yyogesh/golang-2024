package main

import "fmt"

// Alias type
type MyInt int

type User struct {
	Name string
	Id   int
}

func (u User) getUserDetail() string {
	return fmt.Sprintf("Name: %s, ID: %d", u.Name, u.Id)
}

func (m MyInt) isEven() bool {
	return m%2 == 0
}

func main() {
	var a int = 5
	var b MyInt = 10

	fmt.Println(b.isEven(), a)

	var c = int(b)

	var d MyInt = MyInt(a)

	d.isEven()

	abc := User{
		Name: "John Doe",
		Id:   123,
	}

	abc.getUserDetail()
}
