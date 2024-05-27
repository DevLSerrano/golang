package main

import "fmt"

func main() {
	var variableOne string = "first"
	variableTwo := "second"

	var (
		variableThree string = "three"
		variableFour  string = "four"
	)

	variableFive, variableSix := "five", "six"

	fmt.Println(variableOne)
	fmt.Println(variableTwo)
	fmt.Println(variableThree, variableFour)
	fmt.Println(variableFive, variableSix)

	//invert values
	variableFive, variableSix = variableSix, variableFive

}
