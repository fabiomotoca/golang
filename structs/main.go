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
	smith := person{
		firstName: "Agent",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "asmith@matrix.com",
			zipCode: 94000,
		},
	}
	/* or
	var smith person
	smith.firstName = "Agent"
	smith.lastName = "Smith"
	*/

	fmt.Println(smith)
	fmt.Printf("%+v", smith)
}