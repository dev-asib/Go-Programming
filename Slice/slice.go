package main

import "fmt"

func main() {
	var students []string
	var studentName string

	var totalStudent int
	fmt.Print("How many students do you want? = ")
	fmt.Scan(&totalStudent)

	for i := 0; i < totalStudent; i++ {
		fmt.Print("Enter studnet name = ")
		fmt.Scan(&studentName)
		students = append(students, studentName)
	}

	fmt.Println(students)

	var developers []string

	developers = append(developers, "Asib")
	developers = append(developers, "Samy")
	developers = append(developers, "Sabbir")

	fmt.Println(developers)
	fmt.Println(len(developers))

}
