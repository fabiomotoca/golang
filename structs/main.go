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

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	p1 := person{
		firstName: "Agent",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "asmith@matrix.com",
			zipCode: 94000,
		},
	}
	/* or also it works like below
	var smith person
	smith.firstName = "Agent"
	smith.lastName = "Smith"
	*/

	p1.updateName("Exile")
	p1.print()
	/* or also it works like below
	p1Pointer := &p1
	p1Pointer.updateName("Exile")
	p1.print()
	*/

}
