package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%v\n", p)
}

func main() {
	brandon := person{
		firstName: "Brandon",
		lastName:  "Lenz",
		contactInfo: contactInfo{
			email: "brandonalenz@gmail.com",
			zip:   30341,
		},
	}

	brandon.print()
	brandon.updateName("Dwight")
	brandon.print()
}
