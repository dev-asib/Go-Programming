package main

import "fmt"

func main() {

	/*
		// Variable declaration
		var fullName, country string
		var age int
		var cgpa float32
		var isBangladeshi bool

		// Variable initialization
		fullName = "Asib"
		age = 24
		cgpa = 3.94
		isBangladeshi = true
		country = "Bangladesh"
	*/

	// // Variable declaration
	// var fullName string = "Asib"
	// var country string = "Bangladesh"
	// var age int = 20
	// var cgpa float32 = 4.00
	// var isBangladeshi bool = true

	// // Variable declaration
	// var fullName = "Asib"
	// var country = "Bangladesh"
	// var age = 20
	// var cgpa = 4.00
	// var isBangladeshi = true

	// Variable declaration
	fullName := "Asib"
	_country := "Bangladesh"
	age := 20
	CGPA := 4.00
	isBangladeshi := true

	fmt.Println("This is", fullName)
	fmt.Println("CGPA :", CGPA, ".", "It's a awesome CGPA")
	fmt.Println(fullName, "is currently", age)
	fmt.Println("Is Asib Bangladeshi? =", isBangladeshi)
	fmt.Println("Country Name:", _country)
}
