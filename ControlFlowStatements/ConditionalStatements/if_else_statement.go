package main

import "fmt"

func main() {
	password := "1234Password"
	enteredPassword := "1234password"

	if enteredPassword == password {
		fmt.Println("Login Successfully")
	} else {
		fmt.Println("Password not match. Please try again!")
	}
}
