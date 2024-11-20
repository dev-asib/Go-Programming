package main

import "fmt"

func main() {
	var base, height, area float32

	fmt.Printf("Enter base = ")
	fmt.Scan(&base)

	fmt.Printf("Enter height = ")
	fmt.Scan(&height)

	area = 0.5 * base * height

	fmt.Printf("Area of Triangle = %v", area)

}
