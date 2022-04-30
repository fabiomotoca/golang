# My golang playground

## Useful Links

[The golang Standard Library](golang.org/pkg)

[Ascii Table - ASCII character codes and html, octal, hex and...](https://www.asciitable.com/)

## Study Notes

Update - Apr, 28

- The operator := can only be used when we first declare a variable, after that we should use =
- A variable can be declared outside a func, but it can't be assigned to a value or we will got an error trying to compile
- The range operator is a special in the for loop, tells the compiler to iterate over a slice
- Every variable declared needs to be used otherwise the code won't compile
- Only the package main will generate an executable
- The golang has two types of arrays: arrays and slices. Arrays are fixed in size whereas slices can grow and shrink
- In golang we can extend a base type and add some extra functionality to it i.e. type deck []string -- will tell go to create an slice of strings and add a bunch of functions specifically made to work with it

Update - Apr, 29

- Change the variable name for the "_" if you intend to not use it, so go compiler stops complain
- Selecting from a slice `cards[startIndexIncluding:upToNotIncluding]`
- go can return multiple values from a func i.e. `func (d deck, handSize int) (deck, deck)` will return two values of type deck 
- If the args of a func has the same type we can abbreviate i.e. `func (x, y int)`
- We can use `defer` in the beginning of a func to do something just before returning the values
- We can name return values i.e. `func (x, y int) (r1 int, r2 int)`
- In golang functions are also their own data types i.e. `func test(n int)` && `x := test` && `x(10)`