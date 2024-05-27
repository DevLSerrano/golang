package main

import "fmt"

func main() {

	var variable1 int = 10
	fmt.Println(variable1)

	var variable2 int = variable1
	fmt.Println(variable2)

	// Pointer is a memory reference
	var variable3 int
	var pointer *int

	fmt.Println(variable3, pointer)
	// 0 , nill
	variable3 = 100
	pointer = &variable3

	fmt.Println(variable3, pointer)
	// 100 , (memory path) ex 0x123213

	// unreference pointer (getting the value)
	fmt.Println(variable3, *pointer)
	// 100, 100

}
