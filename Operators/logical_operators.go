package main

import "fmt"

func main() {
	var x = 10
	var y int = 30

	// "&&" AND
	z := x != y && x < y
	fmt.Println(z)

	// "||" OR
	z = x == y || y > x
	fmt.Println(z)

	// "!" NOT
	z = !(x == y)
	fmt.Println(z)
}
