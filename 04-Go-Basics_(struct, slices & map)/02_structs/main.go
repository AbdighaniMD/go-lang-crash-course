package main

import (
	"fmt"
	"strconv"
)

// A struct is a collection of fields:- Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Geetring method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i'm " + strconv.Itoa(p.age)
}

func main() {
	person1 := Person{"Abdighani", "MD", "London", "M", 23}

	fmt.Println(person1.greet())

	// Struct fields are accessed using a dot
	fmt.Println("Name is " + person1.firstName + " accessed using a dot")

	// Struct fields can be accessed through a struct pointer.
	p := &person1
	p.firstName = "Luffy"

	fmt.Println("Name is changed to ", person1.firstName)

	fmt.Println(person1.greet())

}
