package main

import "fmt"

func main() {

	var fullName string
	var cgpa float32 = 3.94
	country := "Bangladesh"
	city := "Manikganj"
	const LANGUAGE = "Go"
	fullName = "ASIB"

	fmt.Printf("This is %v.\n", fullName)
	fmt.Printf("%v is a beautiful country.\n", country)
	fmt.Printf("I am %v. I achieved CGPA %v.\n", fullName, cgpa)
	fmt.Printf("I Love %v programming language.\n", LANGUAGE)
	fmt.Printf("I Live in %v.", city)

}
