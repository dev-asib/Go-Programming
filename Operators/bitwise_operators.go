package main

import "fmt"

func main() {
	x := 19
	y := 11

	// "&" AND
	and := x & y
	fmt.Printf("x & y = %v\n", and)

	// "|" OR
	or := x | y
	fmt.Printf("x | y = %v\n", or)

	// "^" XOR
	xor := x ^ y
	fmt.Printf("x ^ y = %v\n", xor)

	// "&^" AND NOT
	andNot := x &^ y
	fmt.Printf("x &^ y = %v\n", andNot)

	// "<<" Left Shift
	leftShift := 2 << 2
	fmt.Printf("2 << 2 = %v\n", leftShift)

	// ">>" Right Shift
	rightShift := 8 >> 2
	fmt.Printf("8 >> 2 = %v\n", rightShift)

}
