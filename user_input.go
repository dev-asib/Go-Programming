package main

import "fmt"

func main() {

	// Input: Scan(), Scanln() & Scanf()

	var devName string
	var fristName, lastName string
	var numb1, numb2 int
	var st1, st2 string

	// Scan()
	fmt.Print("Enter your name: ")
	fmt.Scan(&devName)

	fmt.Printf("Developer name is %v.\n", devName)

	fmt.Print("Enter yoru fullname: ")
	fmt.Scan(&fristName, &lastName)

	fmt.Printf("This is %v %v.\n", fristName, lastName)

	fmt.Print("Enter two numbers = ")
	fmt.Scan(&numb1, &numb2)

	fmt.Printf("My favorite numbers = %v & %v\n", numb1, numb2)

	// Scanf()
	var firstBoy, secondBoy string
	fmt.Print("Enter first & second boy names = ")
	fmt.Scanf("%v %v", &firstBoy, &secondBoy)
	fmt.Printf("First Boy name = %v\n", firstBoy)
	fmt.Printf("%v is a second body.", secondBoy)

	// Scanln()
	fmt.Print("Enter student1 name = ")
	fmt.Scanln(&st1)

	fmt.Print("Enter student2 name = ")
	fmt.Scanln(&st2)

	fmt.Printf("Student1 = %v.\n", st1)
	fmt.Printf("Student2 = %v.\n", st2)
}
