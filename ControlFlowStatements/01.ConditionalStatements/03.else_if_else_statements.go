package main

import "fmt"

func main() {
	x := 34
	if x > 0 {
		fmt.Printf("%v is a Positive Number.", x)
	} else if x < 0 {
		fmt.Printf("%v is a Negetive Number", x)
	} else {
		fmt.Printf("%v is a Zero Number", x)
	}
}
