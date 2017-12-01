package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var p person
	p.firstName = "Alex"
	p.lastName = "Anderson"
	fmt.Printf("%+v\n", p)
}
