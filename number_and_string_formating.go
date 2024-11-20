package main

import "fmt"

func main() {
	/// Number formating
	var decimalNumber int

	fmt.Printf("Enter any decimal number = ")
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("Binary Number of %v = %b\n", decimalNumber, decimalNumber)
	fmt.Printf("Octal Number of %v = %o\n", decimalNumber, decimalNumber)
	fmt.Printf("Hexadecimal Number of %v = %x\n", decimalNumber, decimalNumber)

	// floating point number formating
	var pi = 3.1416
	fmt.Printf("%f\n", pi)
	fmt.Printf("%.2f\n", pi)

	/// String formating
	var name = "Asib"
	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name)
}
