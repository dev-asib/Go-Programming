package main

import "fmt"

func main() {
	country()
	country()
	proProgrammer("Asib")
	proProgrammer("Samy")
	circle(3)
	circle(5)
	addTwoNumbers(10, 20)
	addTwoNumbers(50, 50)
	subtractTwoNumbers(50, 30)
	subtractTwoNumbers(80, 40)
	fmt.Println(multiply(10, 5))
	fmt.Print(multiply(10, 20))
	fmt.Print("\n")
	textMessage := message()
	fmt.Println(textMessage)
}

func country() {
	fmt.Printf("I Love Bangladesh.\n")
}

func proProgrammer(programmer string) {
	fmt.Println("Our pro programmer = ", programmer)
}

func circle(radius float32) {
	result := 3.1416 * radius * radius
	fmt.Printf("Area of circle = %v\n", result)
}

func addTwoNumbers(numb1 int, numb2 int) {
	fmt.Printf("%v\n", numb1+numb2)
}

func subtractTwoNumbers(numb1 int, numb2 int) {
	sub := numb1 - numb2
	fmt.Printf("%v - %v = %v\n", numb1, numb2, sub)
}

func multiply(a, b int) int {
	return a * b
}

func message() string {
	return "Hi! Go Programming Language"
}
