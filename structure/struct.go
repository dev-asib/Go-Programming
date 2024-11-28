package main

import "fmt"

func main() {
	student1 := Student{101, "Asib", 24}
	student2 := Student{102, "Samy", 11}
	fmt.Println(student1)
	fmt.Println(student2)

	fmt.Println("Asib Info:")
	fmt.Printf("Student1 name = %v\n", student1.name)
	fmt.Printf("Student1 id = %v\n", student1.id)

	fmt.Println("Samy Info:")
	fmt.Printf("Student2 name = %v\n", student2.name)
	fmt.Printf("Student2 id = %v\n", student2.id)

	fmt.Println("Asib Info:")
	displayInfo(student1)
	fmt.Println("Samy Info:")
	displayInfo(student2)
	fmt.Println("Sabbir Info:")
	displayInfo(Student{103, "Sabbir", 22})

	developer := Student{105, "Dev Asib", 24}
	fmt.Printf("Developer1 Info: ")
	developerInfo(developer)
	fmt.Println()
	developer.changeDeveloperName("Sabbir")
	fmt.Printf("Developer2 Info: ")
	developerInfo(developer)

}

type Student struct {
	id   int
	name string
	age  int
}

func displayInfo(student Student) {
	fmt.Printf("Student name = %v\n", student.name)
	fmt.Printf("Student id = %v\n", student.id)
	fmt.Printf("Student age = %v\n", student.age)
}

func developerInfo(dev Student) {
	fmt.Printf("name = %v, id = %v, age = %v", dev.name, dev.id, dev.age)
}

func (x *Student) changeDeveloperName(newName string) {
	x.name = newName
}
