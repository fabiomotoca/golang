package main

import "fmt"

func main() {
	// initialize an int slice
	numbers := []int{}
	// populate it
	for k := range [11]int{} {
		numbers = append(numbers, k)
	}
	// check if is odd or even
	for _, v := range numbers {
		if (v % 2) == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
