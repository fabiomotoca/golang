# Content Fixing

Apr, 28th

- The operator := can only be used when we first declare a variable, after that we should use =
- A variable can be declared outside a func, but it can't be assigned to a value or we will got an error trying to compile
- The range operator is a special in the for loop, tells the compiler to iterate over a slice
- Every variable declared needs to be used otherwise the code won't compile
- Only the package main will generate an executable
- The golang has two types of arrays: arrays and slices. Arrays are fixed in size whereas slices can grow and shrink
- In golang we can extend a base type and add some extra functionality to it i.e. type deck []string -- will tell go to create an slice of strings and add a bunch of functions specifically made to work with it

Apr, 29th

- Change the variable name for the "_" if you intent to not use it, so go compiler stops complain
- Selecting from a slice `cards[startIndexIncluding:upToNotIncluding]`
- go can return multiple values from a func i.e. `func (d deck, handSize int) (deck, deck)` will return two values of type deck 
