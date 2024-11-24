/*
1. Series: 1 + 2 + 3 + ... + N
Write a program that calculates the sum of numbers from 1 to N.

Input:
N = 10
Output:
Sum = 55

*/

package main

import "fmt"

func main() {

	var endingNumber int
	result := 0

	fmt.Printf("Enter ending number = ")
	fmt.Scan(&endingNumber)

	for i := 1; i <= endingNumber; i++ {
		result += i
	}

	fmt.Println(result)

}
