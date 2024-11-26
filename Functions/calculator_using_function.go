package main

import "fmt"

func main() {
	var numb1, numb2, reuslt float32
	var option string

	i := true

	for i == true {
		fmt.Printf("Enter numb1 = ")
		fmt.Scan(&numb1)

		fmt.Printf("Enter numb2 = ")
		fmt.Scan(&numb2)

		fmt.Printf("Choose option (+, -, *, /) = ")
		fmt.Scan(&option)

		switch option {
		case "+":
			reuslt = add(numb1, numb2)
		case "-":
			reuslt = sub(numb1, numb2)
		case "*":
			reuslt = mul(numb1, numb2)
		case "/":
			reuslt = div(numb1, numb2)
		default:
			fmt.Println("Invalid option")
			continue
		}
		fmt.Printf("Result = %v\n", reuslt)
	}

}

func add(x, b float32) float32 {
	return x + b
}

func sub(x, b float32) float32 {
	return x - b
}

func mul(x, b float32) float32 {
	return x * b
}

func div(x, b float32) float32 {
	return x / b
}
