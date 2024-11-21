package main

import "fmt"

func main() {
	var x = 10
	var y = 20

	z := x == y
	fmt.Printf("%v == %v = %v\n", x, y, z)

	z = x != y
	fmt.Printf("%v != %v = %v\n", x, y, z)

	z = x > y
	fmt.Printf("%v > %v = %v\n", x, y, z)

	z = x < y
	fmt.Printf("%v < %v = %v\n", x, y, z)

	z = x <= y
	fmt.Printf("%v <= %v = %v\n", x, y, z)

	z = x >= y
	fmt.Printf("%v >= %v = %v\n", x, y, z)

}
