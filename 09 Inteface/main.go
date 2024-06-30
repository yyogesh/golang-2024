package main

import "fmt"

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Animal interface {
	Speak()
}

func (c Cat) Speak() {
	fmt.Println(c.Name + " cat speaking")
}

func (d Dog) Speak() {
	fmt.Println(d.Name + " dog speaking")
}

func animalSpeaking(a Animal) {
	a.Speak()
}

func main1() {
	dog := Dog{Name: "Tommy"}
	cat := Cat{Name: "Kitty"}

	animalSpeaking(dog)
	animalSpeaking(cat)
}
