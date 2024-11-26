package main

import "fmt"

func main() {
	x := 10
	var p *int
	p = &x
	fmt.Println("x = ", x)
	fmt.Println("x = ", &x)
	fmt.Printf("p = %v\n", p)
	fmt.Printf("p = %v\n", *p)

	*p = *p - 5
	fmt.Printf("p = %v\n", *p)
	fmt.Printf("x = %v\n", x)

}
