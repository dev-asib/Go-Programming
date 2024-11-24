/*
3. Series: 1³ + 2³ + 3³ + ... + N³
Write a program to calculate the sum of cubes of numbers from 1 to N.

Input:
N = 3
Output:
Sum of cubes = 36

*/

package main

import "fmt"

func main() {

	var n int
	result := 0

	fmt.Printf("Enter ending number n = ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		result += i * i * i
	}

	fmt.Println(result)

}
