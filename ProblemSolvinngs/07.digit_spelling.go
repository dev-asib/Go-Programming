package main

import "fmt"

func main() {
	var digit int
	fmt.Printf("Enter any digit (0-9) = ")
	fmt.Scan(&digit)

	if digit == 0 {
		fmt.Println("Zero")
	} else if digit == 1 {
		fmt.Println("One")
	} else if digit == 2 {
		fmt.Println("Two")
	} else if digit == 3 {
		fmt.Println("Three")
	} else if digit == 4 {
		fmt.Println("Four")
	} else if digit == 5 {
		fmt.Println("Five")
	} else if digit == 6 {
		fmt.Println("Six")
	} else if digit == 7 {
		fmt.Println("Seven")
	} else if digit == 8 {
		fmt.Println("Eight")
	} else if digit == 9 {
		fmt.Println("Nine")
	} else {
		fmt.Println("Invalid digit")
	}
}
