package main

import (
	"fmt"
)

type Address struct {
	street, city string
}

type Person struct {
	Name string
	Age int
	Addr Address
}

type EmbeddedPerson struct {
	Address
	Name string
	Age int
}

func main() {
	var p Person
	p.Name = "Bob"
	p.Addr.city = "Rennes"

	fmt.Printf("%v", p)

	var ep EmbeddedPerson
	ep.Name = "Sarah"
	ep.city = "Los Angeles"
}