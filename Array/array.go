package main

import "fmt"

func main() {
	var rolls [5]int
	rolls[0] = 101
	rolls[1] = 102
	rolls[2] = 103
	rolls[3] = 104
	rolls[4] = 105

	fmt.Println(rolls)
	fmt.Println(rolls[2])
	fmt.Println(rolls[3])

	var countries = [5]string{"Bangladesh", "India", "Pakistan", "Srilanka", "Afganishtan"}
	fmt.Println(countries)
	fmt.Println(countries[0])

	for i := 0; i < 5; i++ {
		fmt.Println(countries[i])
	}

	var universities = [5]string{"DU", "BUET", "DUET", "KUET", "RUET"}
	var choice int
	fmt.Printf("Choose any number between (0-4) : ")
	fmt.Scan(&choice)
	fmt.Printf("You chosen university = %v\n", universities[choice])

	var developers [5]string

	for i := 0; i < len(developers); i++ {
		fmt.Printf("Enter Developer%v name = ", i)
		fmt.Scan(&developers[i])
	}

	fmt.Println(developers)
	fmt.Println("Size of developers = ", len(developers))

}
