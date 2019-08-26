package main

import "fmt"

// type card struct {
// 	Suit  string
// 	Value string
// }

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}

func main() {
	john := person{
		firstName: "john",
		lastName:  "doe",
		contactInfo: contactInfo{
			email:   "johndoe@gmail.com",
			zipCode: 95444,
		},
	}

	(&john).updateName("foo")
	john.print()
}
