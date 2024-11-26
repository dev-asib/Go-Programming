package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	callByReference(&x)
	fmt.Printf("%v\n", x)
}

func callByReference(ptr *int) {
	*ptr = 30
}
