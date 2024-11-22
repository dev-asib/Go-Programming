package main

import "fmt"

func main() {
	var number int
	fmt.Printf("Enter any integer number = ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("%v = Even Number", number)
	} else {
		fmt.Printf("%v = Odd Number", number)
	}
}
