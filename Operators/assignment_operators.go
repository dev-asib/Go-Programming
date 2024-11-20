package main

import "fmt"

func main() {

	// "=" Assign
	var x = 10.0
	fmt.Println(x)

	// "+=" Add and Assign
	x += 5
	fmt.Print(x, "\n")

	// "-=" Subtract and Assign
	x -= 3
	fmt.Printf("%v\n", x)

	// "*=" Multiply and Assign
	x *= 2
	fmt.Printf("%v\n", x)

	// "/=" Divide and Assign
	x /= 3
	fmt.Println(x)

	// "%=" Modulus and Assign
	var z int = 24
	z %= 7
	fmt.Println(z)

}
