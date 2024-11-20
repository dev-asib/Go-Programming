package main

import "fmt"

func main() {
	var num1, num2 int

	num1 = 10
	num2 = 3

	// "+" Addition
	result := num1 + num2
	fmt.Printf("%v + %v = %v\n", num1, num2, result)

	// "-" Subtraction
	result = num1 - num2
	fmt.Printf("%v - %v = %v\n", num1, num2, result)

	// "*" Multiplication
	result = num1 * num2
	fmt.Printf("%v * %v = %v\n", num1, num2, result)

	// "/" Division
	result = num1 / num2
	fmt.Printf("%v / %v = %v\n", num1, num2, result)

	var result2 = float32(num1) / float32(num2)
	fmt.Printf("%v / %v = %.2f\n", num1, num2, result2)

	// "%" Modulus
	result = num1 % num2
	fmt.Printf("%v %% %v = %v\n", num1, num2, result)
}
