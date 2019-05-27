package main

import (
	"fmt"
	"strconv"
)

// Person defined in struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday Method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried Method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person := Person{firstName: "Immanuel", lastName: "Daud", city: "Jakarta", gender: "m", age: 23}
	// or
	person2 := Person{"Immanuel", "Daud", "Jakarta", "m", 23}

	fmt.Println(person)
	fmt.Println(person2)

	// Get a single field
	fmt.Println(person.firstName)

	// Person has birthday
	person.hasBirthday()

	// getMerried
	person.getMarried("Samantha")

	// Greet method
	fmt.Println(person.greet())
}
