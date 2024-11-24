package main

import "fmt"

// 04. for Loop with break Statement

func main() {

	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
