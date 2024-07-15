package main

import "fmt"

func main() {
	fmt.Println("Controller Structure")

	number := 10

	if number > 15 {
		fmt.Println("Maior a 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// Declara a variável e a inicializa "otherNumber"
	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Maior que 0")
	} else {
		fmt.Println("Menor que 0")
	}

}
