package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Asahi", lastName: "Zhang", city: "Shanghai", gender: "m", age: 24}

	// Alternative
	person2 := Person{"Luna", "Sakurakouji", "Shanghai", "f", 22}

	// fmt.Println(person1.firstName, person1.age)
	// person1.age++
	// fmt.Println(person1)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())

	fmt.Println(person2.greet())
	person2.getMarried(person1.lastName)
	fmt.Println(person2.greet())
}
