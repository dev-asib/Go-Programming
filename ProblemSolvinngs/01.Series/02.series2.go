/*
2. Series: 1² + 2² + 3² + ... + N²
Write a program to calculate the sum of squares of numbers from 1 to N.

Input:
N = 5
Output:
Sum of squares = 55

*/

package main

import "fmt"

func main() {

	var n int
	result := 0

	fmt.Printf("Enter ending number n = ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		result += i * i
	}

	fmt.Println(result)

}
