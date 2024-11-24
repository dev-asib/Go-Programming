package main

import "fmt"

// 06. Nested for Loop

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Multiplication of = ", i)
		for j := 1; j <= 10; j++ {
			mul := i * j
			fmt.Printf("%v x %v = %v\n", i, j, mul)
		}
	}

}
