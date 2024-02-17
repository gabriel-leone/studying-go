package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	gabriel := person{
		firstName: "Gabriel",
		lastName:  "Leone",
		contact: contactInfo{
			email:   "gabriel@leone.com",
			zipCode: 12345,
		},
	}
	gabriel.updateName("Gab")
	gabriel.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
