package main

import "fmt"

// 05. for Loop with continue Statement

func main() {

	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
