package main

import "fmt"

func main() {
	var numb1, numb2, numb3 int

	fmt.Printf("Enter three numbers = ")
	fmt.Scan(&numb1, &numb2, &numb3)

	if numb1 > numb2 && numb1 > numb3 {
		fmt.Printf("The largest number = %v", numb1)
	} else if numb2 > numb1 && numb2 > numb3 {
		fmt.Printf("The largest number = %v", numb2)
	} else if numb3 > numb1 && numb3 > numb2 {
		fmt.Printf("The largest number = %v", numb3)
	} else {
		fmt.Printf("Numbers are equal.")
	}
}
