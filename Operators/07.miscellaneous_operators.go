package main

import "fmt"

func main() {
	x := 10
	ptr := &x

	// "&" Address Operator
	fmt.Println(ptr)

	// "*" Deference Operator
	fmt.Println(*ptr)
}
