package main

import "fmt"

type salaryCalculator interface {
	calculateSalary() float64
	report()
}

type PermanentEmployee struct {
	id          int
	basicSalary float64
	commission  float64
}

type ContractEmployee struct {
	id          int
	basicSalary float64
}

func (p PermanentEmployee) calculateSalary() float64 {
	return p.basicSalary + (p.commission/100)*p.basicSalary
}

func (c ContractEmployee) calculateSalary() float64 {
	return c.basicSalary
}

func (p PermanentEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", p.id, p.calculateSalary())
}

func (c ContractEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", c.id, c.calculateSalary())
}

// Employee<T>

func printAnything(value interface{}) {
	fmt.Println(value)
}

func printAnything1(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Printf("The value is an integer: %d\n", value.(int))
	case bool:
		fmt.Printf("The value is a boolean: %t\n", value.(bool))
	case string:
		fmt.Printf("The value is a string: %s\n", value.(string))
	default:
		fmt.Println("Unknown type")
	}
}

func printAnything2(value interface{}) {
	v1, ok := value.(int)
	if ok {
		fmt.Printf("The value is an integer: %d\n", v1)
	}

	v2, ok := value.(string)
	if ok {
		fmt.Printf("The value is an integer: %s\n", v2)
	}
}

func main() {

	printAnything1(10)
	printAnything1(true)
	printAnything1("Hello, World!")
	printAnything1(new(ContractEmployee))

	printAnything(10)
	printAnything(true)
	printAnything("Hello, World!")

	var calculator salaryCalculator
	calculator = PermanentEmployee{id: 1, basicSalary: 10000, commission: 20}
	calculator.report()
	calculator = ContractEmployee{id: 2, basicSalary: 5000}
	calculator.report()
}
