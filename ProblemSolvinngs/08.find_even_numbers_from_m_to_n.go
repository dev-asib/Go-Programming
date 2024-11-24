package main

import "fmt"

func main() {
	var m, n int

	fmt.Printf("Enter starting number m = ")
	fmt.Scan(&m)

	fmt.Printf("Enter ending number n = ")
	fmt.Scan(&n)

	for i := m; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
