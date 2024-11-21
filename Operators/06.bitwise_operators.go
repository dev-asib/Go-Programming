package main

import (
	"fmt"
)

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

	// "<<=" Left Shift Assignment
	leftShiftAssignment := 5
	leftShiftAssignment <<= 2
	fmt.Println(leftShiftAssignment)

	// ">>=" Right Shift Assignment
	rightShiftAssignment := 20
	rightShiftAssignment >>= 2
	fmt.Println(rightShiftAssignment)

	// 1 0 0 1 1
	// 0 1 0 1 1

	// "&=" Bitwise AND Assignment
	bitwiseAndAssignment := 11
	bitwiseAndAssignment &= 19
	fmt.Println(bitwiseAndAssignment)

	// "&=" Bitwise OR Assignment
	bitwiseORAssignment := 11
	bitwiseORAssignment |= 19
	fmt.Println(bitwiseORAssignment)

	// "^=" Bitwise XOR Assignment
	bitwiseXORAssignment := 11
	bitwiseXORAssignment ^= 19
	fmt.Println(bitwiseXORAssignment)

}
