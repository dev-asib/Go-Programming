package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Printf("Enter two integer numbers = ")
	fmt.Scan(&number1, &number2)

	if number1 > number2 {
		fmt.Printf("Largest number = %v", number1)
	} else if number2 > number1 {
		fmt.Printf("Largest number = %v", number2)
	} else {
		fmt.Printf("Numbers are equal.")
	}

}
