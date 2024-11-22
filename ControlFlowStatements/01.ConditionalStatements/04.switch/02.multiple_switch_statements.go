package main

import "fmt"

func main() {
	var numberOfClass int
	fmt.Printf("Enter your class = ")
	fmt.Scan(&numberOfClass)

	switch numberOfClass {
	case 1, 2, 3, 4, 5:
		fmt.Println("Primary School")
	case 6, 7, 8, 9, 10:
		fmt.Println("Secondary School")
	case 11, 12:
		fmt.Println("Higher Secondary School")
	default:
		fmt.Println("Invalid class")
	}
}
