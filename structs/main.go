package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p := person{
		firstName: "Jim",
		lastName:  "Frank",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	p.print()

	p.updateName("Tommy")

	p.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
