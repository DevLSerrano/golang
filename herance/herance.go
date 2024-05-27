package main

import "fmt"

type People struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}
type Student struct {
	People
	school string
	course string
}

func main() {
	student := Student{
		People{
			"leo",
			"serrano",
			20,
			80,
		},
		"engineer",
		"software",
	}

	fmt.Println(student.name)
	fmt.Println(student.age)
	fmt.Println(student.course)
	fmt.Println(student.school)
}
