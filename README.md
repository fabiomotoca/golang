# My golang playground

## Useful Links

[The golang Standard Library](golang.org/pkg)

[Ascii Table - ASCII character codes and html, octal, hex and...](https://www.asciitable.com/)

[AlgoExpert - Resource for coding and system interviews](https://www.algoexpert.io/)

## Study Notes

Update - Apr 28

- The operator := can only be used when we first declare a variable, after that we should use =
- A variable can be declared outside a func, but it can't be assigned to a value, or we will get an error trying to compile
- The range operator is a special in the for loop, tells the compiler to iterate over a slice
- Every variable declared needs to be used otherwise the code won't compile
- Only the package main will generate an executable
- The golang has two types of arrays: arrays and slices. Arrays are fixed in size whereas slices can grow and shrink
- In golang we can extend a base type and add some extra functionality to it i.e. type deck []string -- will tell go to create an slice of strings and add a bunch of functions specifically made to work with it

Update - Apr 29

- Change the variable name for the "_" if you intend to not use it, so go compiler stops complain
- Selecting from a slice `cards[startIndexIncluding:upToNotIncluding]`
- go can return multiple values from a func i.e. `func (d deck, handSize int) (deck, deck)` will return two values of type deck
- If the args of a func has the same type we can abbreviate i.e. `func (x, y int)`
- We can use `defer` in the beginning of a func to do something just before returning the values
- We can name return values i.e. `func (x, y int) (r1 int, r2 int)`
- In golang functions are also their own data types i.e. `func test(n int)` && `x := test` && `x(10)`

Update - Apr 30

- The piece of code below can be used as example for generating a random number with a new `seed` every time that is called:
```go
func (d deck) shuffle() {
	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
```

- The basics to write tests in go is to create files with `_test.go` and use `go test` in order to validate or hypothesis

Update - May 1

- If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package

Update - May 2

- A `struct` is a data type in golang that is similar to a dict in Python or a plain object in JavaScript
```go
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	smith := person{firstName: "Agent", lastName: "Smith"}
	/* alternatively
	var smith person

	*/
	fmt.Println(smith)
}
```
- Whatever we're initializing variables without values, we got the following by default:
	- string = ""
	- int = 0
	- float = 0
	- bool = false

- Structs can also have receivers configured
- The "&" is an operator where "&variable" will give the memory address of the value this variable is pointing at
	- Use "&" to point to a memory address
- The "\*" is also an operator where "\*pointer" will give the value this memory address is pointing at
	- Use "\*" to point to a memory value

```go
/*
	*person = the expression that is a pointer to a person
	(*pointerToPerson) = the actual value of that pointer to a person
*/
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson.firstName = newFirstName
}
```

Update - May 4

- Use pointers to change these things in a function
	- int, float, string, bool, structs
- Don't worry about pointers with these
	- slices, maps, channels, pointers, functions
- Everything in Go is pass by value
- Maps vs Structs
	- Map
		- All keys must be the same type
		- All values must be the same type
		- Keys are indexed - we can iterate over them -
		- Used to represent a collection of related properties, i.e. A collection of colors and hex values
		- Don't need to know all the keys at compile time
		- Reference Type!
	- Struct
		- Values can be of different type
		- Keys don't support indexing
		- You need to know all the different fields at compile time
		- Used to represent a "thing" with a lot of different properties, i.e. A customer
		- Value Type!